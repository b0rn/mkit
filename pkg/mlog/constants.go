package mlog

type loggerLevels struct {
	PANIC string
	FATAL string
	ERROR string
	WARN  string
	INFO  string
	DEBUG string
	TRACE string
}

type loggerConstants struct {
	ZEROLOG string
	LEVELS  loggerLevels
}

var CONSTANTS = loggerConstants{
	ZEROLOG: "zerolog",
	LEVELS: loggerLevels{
		PANIC: "panic",
		FATAL: "fatal",
		ERROR: "error",
		WARN:  "warn",
		INFO:  "info",
		DEBUG: "debug",
		TRACE: "trace",
	},
}
