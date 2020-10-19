package hello

import (
    "fmt"
    log "github.com/sirupsen/logrus"
)

func Hello() string {
    fmt.Println("hello world")
    log.WithFields(log.Fields{
        "animal": "walrus",
        "number": 1,
        "size":   10,
    }).Info("A walrus appears")
    return ""
}
