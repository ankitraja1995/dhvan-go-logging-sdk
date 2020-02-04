package dhvan_go_logging_sdk


type abstractLogger interface {
	Execute(*FluentdLogger, LogLevel, string, map[string]string) error
	setNext(abstractLogger)
}
