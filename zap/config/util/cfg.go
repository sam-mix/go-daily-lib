package util

import (
	"encoding/json"

	"go.uber.org/zap"
)

func GetLogger() *zap.Logger {
	rawJSON := []byte(`{
		"level":"debug",
		"encoding":"json",
		"outputPaths": ["../server.log"],
		"errorOutputPaths": ["../err.log"],
		"initialFields":{"name":"dj"},
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
