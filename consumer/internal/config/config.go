package config

import (
    "os"

    "github.com/nats-io/nats.go"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

var (
    NatsUrl       = nats.DefaultURL
    EventsSubject = "events.*"
    ConsumerName  = "events-processing"
)

func init() {
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    if url := os.Getenv("NATS_URL"); url != "" {
        url = nats.DefaultURL
    }
}
