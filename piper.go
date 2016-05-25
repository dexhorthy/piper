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

    for _, pipe := range config.Pipes {
        executePipe(pipe, config.Source, config.Graphite)
    }

}

func executePipe(pipe PipeConfig, dataSource DataSourceConfig, graphite GraphiteConfig) {

    log.Print("Executing query: ", pipe.Query)
    value, err := Extract(dataSource, pipe.Query)

    if err != nil {
        log.Fatal(err)
    }

    log.Print("Loaded value: ", value)

    err = Report(value, pipe.Dest, graphite.Host, graphite.Port)

    if err != nil {
        log.Fatal(err)
    }

}
