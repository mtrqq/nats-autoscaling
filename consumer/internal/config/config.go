package config

import (
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	NatsUrl         = nats.DefaultURL
	EventsSubject   = "events.*"
	ConsumerName    = "EVENTS_PROCESSING"
	PollingInterval = time.Second / 2
	TimePerMessage  = time.Second
)

func getNumberEnv(key string) (float64, bool) {
	if numStr := os.Getenv(key); numStr != "" {
		number, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			log.Error().
				Err(err).
				Msgf("Failed parsing %s param, discarding", key)
		}

		return number, err == nil
	}

	return 0.0, false
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if url := os.Getenv("NATS_URL"); url != "" {
		NatsUrl = url
	}

	if timeout, success := getNumberEnv("POLLING_TIMEOUT"); success {
		PollingInterval = time.Duration(timeout * float64(time.Second))
	}

	if perSecond, success := getNumberEnv("MESSAGES_PER_SECOND"); success {
		TimePerMessage = time.Duration(1.0 / perSecond * float64(time.Second))
	}
}
