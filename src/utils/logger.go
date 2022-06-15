package project_utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func InfoLogger(data string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	log.Info().Msg(data)
}


func DebugLogger(data string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	log.Debug().Msg(data)
}

func ErrorLogger(data string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	log.Error().Msg(data)
}

func PanicLogger(data string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	log.Panic().Msg(data)
}

func LogLogger(data string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	log.Log().Msg(data)
}

func WarnLogger(data string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	log.Warn().Msg(data)
}
