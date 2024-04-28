package main

import (
	pb "cliente/proto"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertData(c *fiber.Ctx) error {
	fmt.Println("endpoint /insert")
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "error",
		})
	}

	rank := Data{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}

	conn, err := grpc.Dial("service-producer:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Println("Error al establecer la conexión con el servidor:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "error",
		})
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  rank.Name,
		Album: rank.Album,
		Year:  rank.Year,
		Rank:  rank.Rank,
	})
	if err != nil {
		log.Println("Error al enviar datos al servidor:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "error",
		})
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())

	return c.JSON(fiber.Map{
		"msg": "ok",
	})
}

func main() {
	fmt.Println("corriendo cliente")
	app := fiber.New()

	// Endpoint para la raíz "/"
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("endpoint /")
		return c.SendString("Hola a todos")
	})

	app.Post("/insert", func(c *fiber.Ctx) error {
		err := insertData(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "error",
			})
		}
		return c.JSON(fiber.Map{
			"msg": "ok",
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
