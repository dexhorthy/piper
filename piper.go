package main

import (

    "log"
    "flag"
    "sync"
	"github.com/marpaia/graphite-golang"
	"strconv"
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func main() {


    config_file := flag.String("f", "piper.yml", "Path to config file")

    flag.Parse()

    config := LoadConfig(*config_file)

    log.Print("Loaded config from ", *config_file)

    Graphite, err := graphite.NewGraphite(config.Graphite.Host, config.Graphite.Port)
    if err != nil {
        log.Fatal(err)
    }

    source := config.Source

    db, err := sql.Open(source.Driver, fmt.Sprintf("user=%v dbname=%v password=%v host=%v port=%v", source.User, source.Database, source.Password, source.Host, source.Port))
    if err != nil {
	    log.Fatal(err)
    }


    var wg sync.WaitGroup

    for _, pipe := range config.Pipes {
        wg.Add(1)
        go func(pipe Pipe) {
            defer wg.Done()
            pipe.pipe(db, Graphite)
        }(pipe)
    }

    wg.Wait()

}

func (pipe *Pipe) pipe(db *sql.DB, graphite *graphite.Graphite) {

    rows, err := db.Query(pipe.Query)

    if err != nil {
        log.Print(err)
    }

    result := 0.0

    if rows.Next() {
        rows.Scan(&result)
    }


    err = graphite.SimpleSend(pipe.Dest, strconv.FormatFloat(result, 'E', -1, 64))

    log.Printf("Sent %6.2f \t->\t%v", result, pipe.Dest)

    if err != nil {
        log.Print(err)
    }

}
