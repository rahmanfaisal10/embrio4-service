package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	env        string
	configFile string
)

type Config struct {
	DBUsername   string `envconfig:"DB_USERNAME"`
	DBPassword   string `envconfig:"DB_PASSWORD"`
	DBConnection string `envconfig:"DB_CONNECTION"`
	DBHost       string `envconfig:"DB_HOST"`
	DBPort       string `envconfig:"DB_PORT"`
	DBName       string `envconfig:"DB_NAME"`
	SMTPHost     string `envconfig:"SMTP_HOST"`
	SMTPPort     int    `envconfig:"SMTP_PORT"`
	Email        string `envconfig:"EMAIL"`
	MailPassword string `envconfig:"MAILPASSWORD"`

	Port string `envconfig:"PORT"`

	SecretKey string `envconfig:"SECRET_KEY"`

	Env string
}

func init() {
	flag.StringVar(&env, "env", "local", "server environment")
	flag.Parse()

	configFile = fmt.Sprintf("config/%s.env", env)
}

func Get() Config {
	godotenv.Load(configFile)
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Env = env
	return cfg
}
