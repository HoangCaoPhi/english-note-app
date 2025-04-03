package global

import "hoangcaophi/english-note-app/src/backend/options"

type Options struct {
	MongoDb options.MongodbOptions `mapstructure:"mongodb"`
}
