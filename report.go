package main

import (
    "github.com/marpaia/graphite-golang"
    "log"
    "strconv"
)

func Report(value float64, destination string, host string, port int) error {

    Graphite, err := graphite.NewGraphite(host, port)

    if err != nil {
        return err
    }

    log.Printf("Connected to Graphite at %v:%v", host, port)

    Graphite.SimpleSend(destination, strconv.FormatFloat(value, 'E', -1, 64))

    log.Printf("sent %v to %v", value, destination)

    return nil
}
