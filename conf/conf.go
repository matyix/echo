package conf

import (
	"time"
	"github.com/spf13/viper"
)

type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConf *viper.Viper

func init() {
	defaultConf = readConfig("echo")
}
func Config() Provider {
	return defaultConf
}

func readConfig(app string) *viper.Viper {

	v := viper.New()
	v.SetEnvPrefix(app)
	v.AutomaticEnv()

	v.SetDefault("mode", "debug") // debug, release, test
	v.SetDefault("loglevel", "debug")
	v.SetDefault("json_logs", false)
	v.SetDefault("listen_address", ":5037")
	v.SetDefault("secret", "secret123")
	v.SetDefault("secure", false)
	v.SetDefault("read_timeout", "0m10s")
	v.SetDefault("write_timeout", "0m10s")
	v.SetDefault("max_header_bytes", 1048576) //1MB
	v.SetDefault("cert_file", "ssl/server.crt")
	v.SetDefault("key_file", "ssl/server.key")

	return v
}

func LoadConfigProvider(app string) Provider {
	return readConfig(app)
}
