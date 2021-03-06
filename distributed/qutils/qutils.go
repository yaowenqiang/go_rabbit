package qutils

import (
    "fmt"
    "log"
    "github.com/streadway/amqp"
)

const SensorListQueue = "SensorList"
const SensorDiscoveryExchange = "SesorDiscovery"
const PersistReadingQueue = "PersistReading"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
    conn, err := amqp.Dial(url)
    faildOnError(err, "Failed to connect to RabbitMQ")
    ch, err := conn.Channel()
    faildOnError(err, "Failed to open a channeo")
    return conn, ch
}

func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
    q, err := ch.QueueDeclare(
        name,
        false,
        autoDelete,
        false,
        false,
        nil,
    )

    faildOnError(err, "Failed to declare queue")

    return &q

}

func faildOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", err, msg)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}
