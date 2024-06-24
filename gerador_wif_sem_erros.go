package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"meugo/crypto/base58"
)

const (
	prefix            = "000000000000000000000000000000000000000000000000000000000000"
	maximoCombinacoes = 65536
)

func geradorChaves() []string {
	chavesGeradas := make(map[string]struct{})
	chaves := make([]string, 0, maximoCombinacoes)

	for contador := 0; contador < maximoCombinacoes; {
		suffix := make([]byte, 2)
		_, err := rand.Read(suffix)
		if err != nil {
			log.Fatalf("Falha ao gerar chave: %v", err)
		}

		chaveGerada := prefix + hex.EncodeToString(suffix)

		if _, ok := chavesGeradas[chaveGerada]; !ok {
			chavesGeradas[chaveGerada] = struct{}{}
			contador++
			chaves = append(chaves, chaveGerada)
		}
	}

	return chaves
}

func generateWif(privKeyHex string) string {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	extendedKey := append([]byte{byte(0x80)}, privKeyBytes...)
	extendedKey = append(extendedKey, byte(0x01))

	firstSHA := sha256.Sum256(extendedKey)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	finalKey := append(extendedKey, checksum...)

	wif := base58.Encode(finalKey)
	fmt.Print("\n", wif)
	return wif
}

func main() {
	chaves := geradorChaves()
	for _, chave := range chaves {
		generateWif(chave)
	}
	fmt.Print("o Maximo de Combinações foi atingido.")
}
