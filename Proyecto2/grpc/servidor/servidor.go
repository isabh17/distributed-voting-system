package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	pb "servidor/proto"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port        = ":3001"
	kafkaBroker = "my-cluster-kafka-bootstrap:9092" // Cambiar por la dirección de tu broker de Kafka
	topic       = "topic-sopes1"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	fmt.Println(data)

	// Enviar datos al servidor de Kafka
	if err := sendDataToKafka(data); err != nil {
		log.Println("Error al enviar datos a Kafka:", err)
	}

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el album"}, nil
}

func sendDataToKafka(data Data) error {
	// Configuración de Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// Serializa los datos a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Envía el mensaje a Kafka
	if err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	}); err != nil {
		return err
	}

	fmt.Println("Datos enviados correctamente a Kafka.")
	return nil
}

func main() {
	fmt.Println("Ejecutando Server")

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}

}
