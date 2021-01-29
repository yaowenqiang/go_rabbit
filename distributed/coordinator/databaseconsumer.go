package coordinator

import (
    "time"
    "bytes"
    "encoding/gob"
    "github.com/streadway/amqp"
    "github.com/yaowenqiang/go_rabbit/distributed/dto"
    "github.com/yaowenqiang/go_rabbit/distributed/qutils"
)


const maxRate = 5 * time.Second

type DatabaseConsumer struct {
    er EventRaiser
    conn *amqp.Connection
    ch *amqp.Channel
    queue *amqp.Queue
    sources []string
}

func NewDatabaseConsumer(er EventRaiser) *DatabaseConsumer {
    dc := DatabaseConsumer {
        er: er,
    }
    dc.conn, dc.ch = qutils.GetChannel(url)
    dc.queue = qutils.GetQueue(
        qutils.PersistReadingQueue,
        dc.ch,
        false,
    )

    dc.er.AddListener("dataSourceDiscovered", func(eventdata interface{}) {
        dc.SubscribeToDataEvent(eventdata.(string))

    })
    return &dc
}

func (dc *DatabaseConsumer) SubscribeToDataEvent(eventName string) {
    for _, v := range dc.sources {
        if v == eventName {
            return
        }
    }

    dc.er.AddListener("MessageReceived" + eventName, func() func(interface{}) {
        prevTime := time.Unix(0, 0)
        buf := new(bytes.Buffer)
        return func(eventdata interface{}) {
            ed := eventdata.(EventData)
            if time.Since(prevTime) > maxRate {
                prevTime = time.Now()

                sm :=  dto.SensorMessage{
                    Name: ed.Name,
                    Value: ed.Value,
                    Timestamp: ed.Timestamp,
                }

                buf.Reset()

                enc := gob.NewEncoder(buf)

                enc.Encode(sm)

                msg := amqp.Publishing{
                    Body: buf.Bytes(),
                }

                dc.ch.Publish(
                    "",
                    qutils.PersistReadingQueue,
                    false,
                    false,
                    msg,
                )

            }
        }
    }())
}
