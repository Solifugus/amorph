package main

import (
    "fmt"
	"config"
)

func main() {
    cfg, err := config.LoadConfig("config/config.yaml")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Loaded config: %+v\n", cfg)
    // ...
}

