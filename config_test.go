package main

import "fmt"

func ExampleLoadConfig() {
    config := LoadConfig("piper.yml")

    fmt.Printf("Loaded Postgres Host: %v\n", config.Source.Host)
    fmt.Printf("Loaded Postgres Port: %v\n", config.Source.Port)
    fmt.Printf("Loaded Postgres Database: %v\n", config.Source.Database)
    fmt.Printf("Loaded Postgres User: %v\n", config.Source.User)
    fmt.Printf("Loaded Postgres Password: %v\n", config.Source.Password)
    fmt.Printf("Loaded Graphite Host: %v\n", config.Graphite.Host)
    fmt.Printf("Loaded Graphite Port: %v\n", config.Graphite.Port)

    for _, pipeConfig := range config.Pipes {
        fmt.Printf("Loaded Query: %v\n", pipeConfig.Query)
        fmt.Printf("Loaded Dest: %v\n", pipeConfig.Dest)
    }


    // Output:
    // Loaded Postgres Host: localhost
    // Loaded Postgres Port: 5432
    // Loaded Postgres Database: postgres
    // Loaded Postgres User: postgres
    // Loaded Postgres Password: postgres
    // Loaded Graphite Host: localhost
    // Loaded Graphite Port: 2003
    // Loaded Query: SELECT count(*) FROM bar
    // Loaded Dest: piper.bar.count
    // Loaded Query: SELECT id FROM bar order by id desc limit 1
    // Loaded Dest: piper.bar.greatest_id
    //

}
