package main

import (
    "log"
    "flag"
)

func main() {


    config_file := flag.String("f", "piper.yml", "Path to config file")

    flag.Parse()

    config := LoadConfig(*config_file)

    log.Print("Loaded config from ", *config_file)

    for _, pipeConfig := range config {
        executePipe(pipeConfig)
    }

}

func executePipe(config PipeConfig) {

    log.Print("Executing query: ", config.Query)
    value, err := Extract(config.Source, config.Query)

    if err != nil {
        log.Fatal(err)
    }

    log.Print("Loaded value: ", value)

    err = Report(value, config.Dest, config.Graphite.Host, config.Graphite.Port)

    if err != nil {
        log.Fatal(err)
    }

}
