package options

type AuthenticationOptions struct {
	Jwt JwtOptions `mapstructure:"jwt"`
}

type JwtOptions struct {
	SecretKey string `mapstructure:"secret"`
}
