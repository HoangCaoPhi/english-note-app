package global

import "hoangcaophi/english-note-app/src/backend/host/options"

type Options struct {
	MongoDb        options.MongodbOptions        `mapstructure:"mongodb"`
	Authentication options.AuthenticationOptions `mapstructure:"authentication"`
}
