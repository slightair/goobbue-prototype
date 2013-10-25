package main

import "github.com/slightair/goobbue-prototype/goobbue"
import "os"
import "encoding/json"

func main() {
    m := goobbue.NewMap()
    mapJSON, _ := json.Marshal(m)
    os.Stdout.Write(mapJSON)
}