# rabbitmq-mmm (msg miru miru)

RabbitMQ Consumer CLI Library

## Usage

- write message handler

ex:

```Go
func handleMessage(msg []byte, queue lib.Queue) error {
	fmt.Printf("%s from %s", string(msg), queue.Name)

	return nil
}
```

- write your `main.go`
- execute!

```shell script
$ RMQ_USERNAME="user" RMQ_PASSWORD="password" RMQ_HOST="localhost" go run main.go <queue filter string>
```