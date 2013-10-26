package goobbue

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func StartServer(port int) {
    http.HandleFunc("/api/map.json", func(w http.ResponseWriter, r *http.Request) {
        m := NewMap()
        mapJSON, _ := json.Marshal(m)
        w.Write(mapJSON)
    })
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "public/" + r.URL.Path[1:])
    })
    
    panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}