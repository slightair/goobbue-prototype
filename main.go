package main

import (
    "github.com/slightair/goobbue-prototype/goobbue"
    "os"
    "strconv"
)

func main() {
    port, error := strconv.Atoi(os.Getenv("PORT"))
    if error != nil {
        port = 8080
    }
    goobbue.StartServer(port)
}