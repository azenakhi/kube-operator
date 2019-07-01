package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

var Log = zerolog.New(output).With().Timestamp().Logger()
