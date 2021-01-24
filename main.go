package main

import(
    "fmt"
    "log"
    "time"
    "github.com/streadway/amqp"
)

func main() {
    go client()
    go server()

    var a string
    fmt.Scanln(&a)
}

func client() {
    conn, ch, q := getQueue()
    defer conn.Close()
    defer ch.Close()

    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )

    faildError(err, "Failed to register a consumer")

    for msg := range msgs {
        log.Printf("Received message with message: %s", msg.Body)
    }

}
func server() {
    conn, ch, q := getQueue()
    defer conn.Close()
    defer ch.Close()
    msg := amqp.Publishing{
        ContentType: "text/plain",
        Body: []byte("Hello RabbitMQ"),
    }

    for {
        ch.Publish("", q.Name, false, false, msg)
        time.Sleep(1 * time.Second)

    }
}

func getQueue() (*amqp.Connection,*amqp.Channel, *amqp.Queue) {
    conn, err := amqp.Dial("amqp://guest@localhost:5672")
    faildError(err, "Failed to connect to RabbitMQ")
    ch, err := conn.Channel()
    faildError(err, "Failed to open a channeo")

    q, err := ch.QueueDeclare("hello",
        false,//durable
        false, //autoDelete
        false,//exclusive
        false,//noWait
        nil,//args
    )

    faildError(err, "Failed to declare a queue")
    return conn, ch, &q



}

func faildError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", err, msg)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}


