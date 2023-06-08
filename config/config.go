package config

type GlobalConfig struct {
	Name    string  `yaml:"name" default:"envoy-go-filter-hub" validate:"required"`
	Host    string  `yaml:"host" default:"0.0.0.0" validate:"required,ip"`
	Port    string  `yaml:"port" default:"8080" validate:"required"`
	RunMode RunMode `yaml:"runMode" default:"debug" validate:"required,oneof=debug release"`
	Prefix  string  `yaml:"prefix" default:"-"`

	Log   LogConfig   `yaml:"log"`
	JWT   JWTConfig   `yaml:"jwt"`
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type LogConfig struct {
	Path       string `yaml:"path" default:"./log" validate:"required"`
	TimeFormat string `yaml:"timeFormat" default:"2006-01-02 15:04:05" validate:"required"`
}

type JWTConfig struct {
	Issuer       string `yaml:"issuer" default:"envoy-go-filter-hub" validate:"required"`
	AccessSecret string `yaml:"accessSecret" default:"ThisIsMySecret" validate:"required"`
	AccessExpire int64  `yaml:"accessExpire" default:"2592000" validate:"required"`
}

type MysqlConfig struct {
	Address  string `yaml:"address" default:"127.0.0.1:3306" validate:"required"`
	Username string `yaml:"username" default:"root" validate:"required"`
	Password string `yaml:"password" default:"12345678"`
	DB       string `yaml:"DB" default:"filter_hub" validate:"required"`
	//MaxIdle     int           `yaml:"maxIdle" default:"20"`
	//MaxOpen     int           `yaml:"maxOpen" default:"20"`
	//MaxLifetime time.Duration `yaml:"maxLifetime" default:"100"`
}

type RedisConfig struct {
	Address  string `yaml:"address" default:"127.0.0.1:6379" validate:"required"`
	Password string `yaml:"password" default:"-"`
	DB       int    `yaml:"DB" default:"-"`
	//MaxIdle     int           `yaml:"maxIdle" default:"20"`
	//MaxOpen     int           `yaml:"maxOpen" default:"20"`
	//IdleTimeOut time.Duration `yaml:"idleTimeOut" default:"100"`
}

/*
RunMode debug|release
debug:
 1. log in console with colorful text
 2. log request and response detail
 3. set gin in gin.DebugMode
 4. zap log level is zapcore.DebugLevel

release:
 1. log in file with json format
 2. log request and response summary
 3. set gin in gin.ReleaseMode
 4. zap log level is zapcore.InfoLevel
*/
type RunMode string

const (
	ModeDebug   RunMode = "debug"
	ModeRelease RunMode = "release"
)
