package mlog

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = map[string]zerolog.Level{
	CONSTANTS.LEVELS.PANIC: zerolog.PanicLevel,
	CONSTANTS.LEVELS.FATAL: zerolog.FatalLevel,
	CONSTANTS.LEVELS.ERROR: zerolog.ErrorLevel,
	CONSTANTS.LEVELS.WARN:  zerolog.WarnLevel,
	CONSTANTS.LEVELS.INFO:  zerolog.InfoLevel,
	CONSTANTS.LEVELS.DEBUG: zerolog.DebugLevel,
	CONSTANTS.LEVELS.TRACE: zerolog.TraceLevel,
}

var Logger zerolog.Logger = log.Logger

type ErrorStackMarshaler = func(err error) interface{}

func Init(cfg Config, errorStackMarsheler ErrorStackMarshaler) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if errorStackMarsheler != nil {
		zerolog.ErrorStackMarshaler = errorStackMarsheler
	}
	zerolog.SetGlobalLevel(levelMap[cfg.Level])
	var logger zerolog.Logger
	var writers []io.Writer

	if cfg.EnablePrettyPrint {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
	} else {
		writers = append(writers, os.Stdout)
	}
	if cfg.FileConfig.Enabled {
		writers = append(writers, newRollingFile(&cfg.FileConfig))
	}
	logger = zerolog.New(zerolog.MultiLevelWriter(writers...)).
		With().
		Timestamp().
		Logger()

	if cfg.EnableCaller {
		logger = logger.With().Caller().Logger()
	}

	if cfg.SampleRate > 1 {
		logger = logger.Sample(&zerolog.BasicSampler{N: uint32(cfg.SampleRate)})
	}

	e := logger.Info().
		Bool("prettyPrint", cfg.EnablePrettyPrint).
		Bool("fileOutputEnabled", cfg.FileConfig.Enabled)
	if cfg.FileConfig.Enabled {
		e = e.Str("fileOutput", path.Join(cfg.FileConfig.Directory, cfg.FileConfig.Filename)).
			Int("fileMaxBackups", cfg.FileConfig.MaxBackups).
			Int("fileMaxSize", cfg.FileConfig.MaxSizeInMb).
			Int("fileMaxAge", cfg.FileConfig.MaxAgeInDays)
	}
	e.Bool("callerEnabled", cfg.EnableCaller).
		Uint("sampleRate", cfg.SampleRate).
		Msg("logger configured")

	log.Logger = logger
	Logger = log.Logger
	return logger
}

func LogErrors(event *zerolog.Event, errs []error) {
	added := false
	for _, e := range errs {
		if e == nil {
			continue
		}
		added = true
		event = event.Err(e)
	}
	if added {
		event.Msg("")
	}
}

func newRollingFile(cfg *FileConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join(cfg.Directory, cfg.Filename),
		MaxBackups: cfg.MaxBackups,
		MaxSize:    cfg.MaxSizeInMb,
		MaxAge:     cfg.MaxAgeInDays,
	}
}
