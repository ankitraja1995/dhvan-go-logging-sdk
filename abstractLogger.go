package dhvan_go_logging_sdk

import (
	"dhvan-go-logging-sdk/enums"
)

type abstractLogger interface {
	Execute(*FluentdLogger, enums.LogLevel, string, map[string]string) error
	setNext(abstractLogger)
}
