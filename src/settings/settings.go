package settings

import "github.com/igrybkov/leosync/src/lingualeo"

// Config is a representation of credentials for LinguaLeo
type Config struct {
	LinguaLeo lingualeo.ConnectionConfig
}

// GetSettings return configuration instance
//func GetSettings() Config {
//	config := Config{
//		LinguaLeo: lingualeo.ConnectionConfig{
//			Email:    cast.ToString(viper.Get("lingualeo.email")),
//			Password: cast.ToString(viper.Get("lingualeo.password")),
//		},
//	}
//	return config
//}
