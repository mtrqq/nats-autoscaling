package main

import (
    "fmt"
    "time"

    "github.com/mtrqq/nats-autoscaling/internal/config"
    "github.com/nats-io/nats.go"
    "github.com/rs/zerolog/log"
)

func main() {
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal().
            Err(err).
            Msg("Failed to connect to NATS")
    }

    defer nc.Drain()

    js, _ := nc.JetStream()
    sub, err := js.PullSubscribe(config.EventsSubject, config.ConsumerName)
    if err != nil {
        log.Fatal().
            Err(err).
            Msg("Failed to create a pull subscription")
    }

    // Simple Async Stream Publisher
    for i := 0; i < 500; i++ {
        js.PublishAsync("ORDERS.scratch", []byte("hello"))
    }
    select {
    case <-js.PublishAsyncComplete():
    case <-time.After(5 * time.Second):
        fmt.Println("Did not resolve in time")
    }

    // Simple Async Ephemeral Consumer
    js.Subscribe("ORDERS.*", func(m *nats.Msg) {
        fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
    })

    // Simple Sync Durable Consumer (optional SubOpts at the end)
    sub, err := js.SubscribeSync("ORDERS.*", nats.Durable("MONITOR"), nats.MaxDeliver(3))
    m, err := sub.NextMsg(timeout)

    // Simple Pull Consumer
    sub, err := js.PullSubscribe("ORDERS.*", "MONITOR")
    msgs, err := sub.Fetch(10)

    // Unsubscribe
    sub.Unsubscribe()

    // Drain
    sub.Drain()
}
