package main

import (
	"time"

	"github.com/mtrqq/nats-autoscaling/internal/config"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func main() {
	nc, err := nats.Connect(config.NatsUrl)
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

	for {
		if messages, err := sub.Fetch(1); err == nil {
			log.Debug().Msgf("Picked up a message (%s)", messages[0].Data)
			time.Sleep(config.TimePerMessage)
			messages[0].Ack()
		} else if len(messages) == 0 || err == nats.ErrTimeout {
			log.Debug().Msg("No messages available (timeout waiting)")
			time.Sleep(config.PollingInterval)
		} else {
			log.Error().Err(err).Msg("Failed fetching a message")
		}
	}
}
