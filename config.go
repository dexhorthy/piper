package main

import (
        "log"
        "io/ioutil"
        "gopkg.in/yaml.v2"
)

type DataSourceConfig struct {
    Driver string
    Host string
    Port int
    User string
    Password string
    Database string
}


type GraphiteConfig struct {
    Host string
    Port int
}

type Pipe struct {
    Query string
    Dest string
}

type PiperConfig struct {
    Source DataSourceConfig
    Graphite GraphiteConfig
    Pipes []Pipe
}

func LoadConfig(path string) PiperConfig {
    data, err := ioutil.ReadFile(path)

    if err != nil {
        log.Fatalf("error: %v", err)
    }

    config := PiperConfig{}

    err = yaml.Unmarshal([]byte(data), &config)

    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return config
}
