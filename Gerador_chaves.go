package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    prefix    = "000000000000000000000000000000000000000000000000000000000000"
    characters = "0123456789abcdef"
)

func geradorChaves() string {
    rand.Seed(time.Now().UnixNano())

    var caracteresUnicos []byte
    for i := 0; i < 4; i++ {
        randomIndex := rand.Intn(len(characters))
        caracteresUnicos = append(caracteresUnicos, characters[randomIndex])
    }

    chaveGerada := prefix + string(caracteresUnicos)
    return chaveGerada
}

func main() {
    chave := geradorChaves()
    fmt.Println(chave)
}
