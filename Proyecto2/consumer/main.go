package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

// Configuración de MongoDB
const (
	uri      = "mongodb://root:secret@mongodb-service:27017/?authSource=admin"
	addredis = "redis-service:6379"
)

func main() {
	fmt.Println("Ejecutando Consumer v3.3 ")

	//configurando Kafka
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-bootstrap:9092",
		"group.id":          "test-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Error en crear consumer: %s", err)
		os.Exit(1)
	}

	topic := "topic-sopes1"
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("No se pudo suscribir al topic: %s", err)
		os.Exit(1)
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Configura las opciones del cliente
	opts := options.Client().ApplyURI(uri)

	// Crea un nuevo cliente y se conecta al servidor MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println("Error al conectar con MongoDB:", err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Println("Error al cerrar la conexión con MongoDB:", err)
		}
	}()

	// Ping al servidor MongoDB para confirmar la conexión exitosa
	if err := client.Ping(context.TODO(), nil); err != nil {
		fmt.Println("Error al hacer ping a MongoDB:", err)
		return
	}

	fmt.Println("Conexión exitosa con MongoDB")

	// Crea un nuevo cliente Redis
	clientRedis := redis.NewClient(&redis.Options{
		Addr: addredis, // Reemplaza con la IP o nombre del servicio de Redis
		DB:   0,        // Número de base de datos (por defecto es 0)
	})

	// Ping al servidor Redis para verificar la conexión
	pong, err := clientRedis.Ping().Result()
	if err != nil {
		fmt.Println("Error al conectar con Redis:", err)
		return
	}
	fmt.Println("Conexión exitosa con Redis:", pong)

	// Set up a WaitGroup for goroutines
	var wg sync.WaitGroup

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}

			// Procesa el mensaje
			var data Data
			err = json.Unmarshal(ev.Value, &data)
			if err != nil {
				fmt.Printf("Error al decodificar el mensaje: %s\n", err)
				continue
			}

			// Incrementa el contador de WaitGroup
			wg.Add(2)

			// Guarda en MongoDB en una goroutine
			go func(data Data) {
				defer wg.Done()
				fmt.Println("Insert Mongo")
				key := data.Name + "-" + data.Album + "-" + data.Year + "-" + data.Rank
				// Selecciona la base de datos y la colección
				db := client.Database("votos")              // Nombre de tu base de datos
				collection := db.Collection("mycollection") // Nombre de tu colección

				_, err = collection.InsertOne(context.TODO(), bson.M{"key": key, "Time": time.Now()})
				if err != nil {
					fmt.Printf("Error al guardar en MongoDB: %s\n", err)
				} else {
					fmt.Println("Se inserto en Mongo", data.Name)
				}

			}(data)

			//Guarda en Redis con gorutine
			go func(data Data) {
				defer wg.Done()
				fmt.Println("Insert Redis")
				key := data.Name + "-" + data.Album + "-" + data.Year + "-" + data.Rank
				// Inserta el JSON serializado en Redis
				err = clientRedis.HIncrBy("teams", key, 1).Err()
				if err != nil {
					fmt.Println("Error al insertar el JSON en Redis:", err)
					return
				} else {
					fmt.Println("voto registrado en redis ", key)
				}
			}(data)

			go func(data Data) {
				// Imprime el mensaje
				fmt.Printf("Mensaje recibido:\n")
				fmt.Printf("Name: %s\n", data.Name)
				fmt.Printf("Album: %s\n", data.Album)
				fmt.Printf("Year: %s\n", data.Year)
				fmt.Printf("Rank: %s\n", data.Rank)
			}(data)

		}
	}

	// Espera a que todas las goroutines terminen
	wg.Wait()
	c.Close()

	fmt.Println("Consumidor")
}
