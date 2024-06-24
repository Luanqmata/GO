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
// ----------------------------------------- GERADOR CHAVES UNICAS ------------------------------------------------

package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func geradorChaves() {
	const prefix = "000000000000000000000000000000000000000000000000000000000000"
	chavesGeradas := make(map[string]bool)

	for {
		// Buffer de 2 bytes para gerar 4 caracteres hexadecimais (16 bits = 4 hex digits)
		suffix := make([]byte, 2)
		_, err := rand.Read(suffix)
		if err != nil {
			log.Fatalf("Falha ao gerar chave: %v", err)
		}

		// Convertendo os 2 bytes para uma string hexadecimal
		suffixStr := hex.EncodeToString(suffix)
		chaveGerada := prefix + suffixStr

		// Verificar se a chave já existe no mapa
		if _, existe := chavesGeradas[chaveGerada]; existe {
			fmt.Println("Chave duplicada encontrada:", chaveGerada)
			continue
		}

		// Adicionar chave ao mapa de chaves geradas
		chavesGeradas[chaveGerada] = true
		fmt.Println("Chave gerada:", chaveGerada)
	}
}

func main() {
	geradorChaves()
}
