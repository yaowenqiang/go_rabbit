package main

import (
    "bytes"
    "encoding/gob"
    "log"
    "github.com/yaowenqiang/go_rabbit/distributed/dto"
    "github.com/yaowenqiang/go_rabbit/distributed/datamanager"
    "github.com/yaowenqiang/go_rabbit/distributed/qutils"
)

const url = "amqp://guest:guest@localhost:5672"

func main() {
    conn, ch := qutils.GetChannel(url)
    defer conn.Close()
    defer ch.Close()

    msgs, err := ch.Consume(
        qutils.PersistReadingQueue,
        "",
        false,
        true,
        false,
        false,
        nil,
    )

    if err != nil {
        log.Fatalf("Failed to get access to messages ; %s", err)
    }

    for msg := range msgs {
        buf := bytes.NewReader(msg.Body)
        dec := gob.NewDecoder(buf)
        sd := &dto.SensorMessage{}
        dec.Decode(sd)

        err := datamanager.SaveReading(sd)

        if err != nil {
            log.Printf("Failed to save reading from sensor \v. error: ",sd.Name, err.Error() )
        } else {
            msg.Ack(false)
        }
    }
}
