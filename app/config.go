package app

import (
    "log"

    "github.com/joeshaw/envdecode"
)

func NewConfig() *ConfigApplication {
    var c ConfigApplication
    if err := envdecode.StrictDecode(&c); err != nil {
        log.Fatalf("Failed to decode: %s", err)
    }

    return &c
}

type ConfigApplication struct {
    API      API
    Database Database
    Server   Server
    Log      Log
}

type API struct {
    Key    string `env:"API_KEY,required"`
    Origin string `env:"ORIGIN,required"`
}

type Server struct {
    Host string `env:"SERVER_HOST,required"`
    Port string `env:"SERVER_PORT,required"`
}

type Database struct {
    User     string `env:"DATABASE_USER,required"`
    Password string `env:"DATABASE_PASSWORD,required"`
    Host     string `env:"DATABASE_HOST,required"`
    Port     int    `env:"DATABASE_PORT,required"`
    Name     string `env:"DATABASE_NAME,required"`
}

type Log struct {
    Level string `env:"LOG_LEVEL,required"`
}
