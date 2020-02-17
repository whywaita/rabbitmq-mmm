package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/streadway/amqp"
)

type Queue struct {
	Name  string `json:"name"`
	VHost string `json:"vhost"`
}

var (
	Logger *log.Logger
)

func ConsumeMsg(filter string, handler func([]byte, Queue) error) error {
	cctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	queues, err := getQueueList()
	if err != nil {
		return err
	}
	qs := filterRelatedQueues(queues, filter)

	conn, err := amqp.Dial(c.toDSN())
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	for _, queue := range qs {
		msgs, err := ch.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		go func() {
			Logger.Printf("Consume start queue: %s\n", queue.Name)
		CONSUMER:
			for {
				select {
				case <-cctx.Done():
					break CONSUMER
				case m, ok := <-msgs:
					if ok {
						err := handler(m.Body, queue)
						if err != nil {
							Logger.Println(err)
							continue
						}
					}
				}
			}
		}()
	}

	<-signals
	return nil
}

func getQueueList() ([]Queue, error) {
	// get all queue list

	manager := fmt.Sprintf("http://%s:15672/api/queues", c.Host)
	client := &http.Client{}
	req, err := http.NewRequest("GET", manager, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	v := make([]Queue, 0)
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func filterRelatedQueues(qs []Queue, queryName string) []Queue {
	// return queues include queueName in Queue.Name

	var result []Queue
	for _, q := range qs {
		if strings.Contains(q.Name, queryName) {
			result = append(result, q)
		}
	}

	return result
}
