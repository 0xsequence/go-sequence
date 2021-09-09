package logadapter

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Wraps a zerolog.Logger to provide common Logger and StdLogger interfaces

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type StdLogger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

var (
	_ Logger    = &LogAdapter{}
	_ StdLogger = &LogAdapter{}
)

func Wrap(log zerolog.Logger) *LogAdapter {
	return &LogAdapter{log}
}

type LogAdapter struct {
	log zerolog.Logger
}

func (s *LogAdapter) Debug(v ...interface{}) {
	s.log.Debug().Msg(fmt.Sprint(v...))
}

func (s *LogAdapter) Debugf(format string, v ...interface{}) {
	s.log.Debug().Msg(fmt.Sprintf(format, v...))
}

func (s *LogAdapter) Info(v ...interface{}) {
	s.log.Info().Msg(fmt.Sprint(v...))
}

func (s *LogAdapter) Infof(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}

func (s *LogAdapter) Warn(v ...interface{}) {
	s.log.Warn().Msg(fmt.Sprint(v...))
}

func (s *LogAdapter) Warnf(format string, v ...interface{}) {
	s.log.Warn().Msg(fmt.Sprintf(format, v...))
}

func (s *LogAdapter) Error(v ...interface{}) {
	s.log.Error().Msg(fmt.Sprint(v...))
}

func (s *LogAdapter) Errorf(format string, v ...interface{}) {
	s.log.Error().Msg(fmt.Sprintf(format, v...))
}

func (s *LogAdapter) Fatal(v ...interface{}) {
	s.log.Fatal().Msg(fmt.Sprint(v...))
	os.Exit(1)
}

func (s *LogAdapter) Fatalf(format string, v ...interface{}) {
	s.log.Fatal().Msg(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (s *LogAdapter) Print(v ...interface{}) {
	s.log.Info().Msg(fmt.Sprint(v...))
}

func (s *LogAdapter) Println(v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintln(v...))
}

func (s *LogAdapter) Printf(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}
