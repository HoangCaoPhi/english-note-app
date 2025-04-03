package initialization

import (
	"fmt"
	"hoangcaophi/english-note-app/src/backend/global"

	"github.com/spf13/viper"
)

func loadConfiguration() {
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fail to read configuration file %w ", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("fail to unmarshal configuration file %w ", err))
	}
}
