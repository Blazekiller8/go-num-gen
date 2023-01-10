package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println(rand.Intn(math.MaxInt64))
}
