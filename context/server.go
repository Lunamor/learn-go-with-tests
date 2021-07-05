package context

import (
    "net/http"
    "fmt"
)

func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, store.Fetch())
    }
}

type Store interface {
    Fetch() string
}

