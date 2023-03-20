package actlogger

import (
	"io"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ActLogger struct {
	*zerolog.Logger
}

type Config struct {
	Application string
	Filepath    string
	MaxSize     int  // MaxSize the max size in MB of the logfile before it's rolled
	MaxBackups  int  // MaxBackups the max number of rolled files to keep
	MaxAge      int  // MaxAge the max age in days to keep a logfile
	DebugLevel  bool //DebugLevel bool
	Compress    bool //Compress logs archive
}

// Configure sets up the logging framework
func Configure(config Config) *ActLogger {
	var writers []io.Writer

	writers = append(writers, &lumberjack.Logger{
		Filename:   config.Filepath,
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
		Compress:   config.Compress,
	})

	//writers = append(writers, os.Stdout, os.Stderr)

	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "t"

	if config.DebugLevel {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := zerolog.New(io.MultiWriter(writers...)).
		With().Str("app", config.Application).
		Timestamp().
		Logger()

	return &ActLogger{
		Logger: &logger,
	}
}
