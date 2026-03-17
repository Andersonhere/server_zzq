package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Wechat   WechatConfig   `mapstructure:"wechat"`
	Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
	Port   int            `mapstructure:"port"`
	Mode   string         `mapstructure:"mode"`
	HTTPS  HTTPSConfig    `mapstructure:"https"`
}

type HTTPSConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
}

type DatabaseConfig struct {
	Driver       string `mapstructure:"driver"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type WechatConfig struct {
	AppID  string `mapstructure:"appid"`
	Secret string `mapstructure:"secret"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
}

var GlobalConfig *Config

// Init 初始化配置
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	// 添加多个配置查找路径
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../configs")

	// 设置默认值
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.filename", "stdout") // 默认输出到标准输出，不写文件

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 解析配置
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		return err
	}

	return nil
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	return GlobalConfig
}
