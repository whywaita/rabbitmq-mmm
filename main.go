package main

import (
	"fmt"
	"log"
	"os"

	"github.com/whywaita/rabbitmq-mmm/lib"

	"github.com/urfave/cli/v2"
)

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "rabbitmq-mmm"
	app.Usage = "Consume RabbitMQ msg (no ack)"
	app.Action = consumeMsg

	return app
}

func consumeMsg(ctx *cli.Context) error {
	filterStr := ctx.Args().Get(0)
	return lib.ConsumeMsg(filterStr, handleMessage)
}

func main() {
	lib.LoadConfig()

	app := makeApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMessage(msg []byte, queue lib.Queue) error {
	fmt.Printf("%s from %s", string(msg), queue.Name)

	return nil
}
