package logger

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

var Log = &log.Logger{
	Out:   os.Stdout,
	Level: log.DebugLevel,
	Formatter: &easy.Formatter{
		TimestampFormat: "02-01-2006 15:04:05",
		LogFormat:       "[%lvl%] %time% - %msg% \n",
	},
}
