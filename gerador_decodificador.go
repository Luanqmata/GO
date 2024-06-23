//16 bits
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"

	"meugo/crypto/base58"

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
)

const (
	prefix     = "000000000000000000000000000000000000000000000000000000000000"
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

	return wif
}

// Função para criar o hash público a partir de uma chave privada
func createPublicHash160(privKeyHex string) []byte {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)

	compressedPubKey := privKey.PubKey().SerializeCompressed()

	pubKeyHash := hash160(compressedPubKey)
	return pubKeyHash
}

// Função para calcular o hash RIPEMD160(SHA256(b))
func hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}

// Função para codificar o hash da chave pública em um endereço Bitcoin
func encodeAddress(pubKeyHash []byte) string {
	version := byte(0x00)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := doubleSha256(versionedPayload)[:4]
	fullPayload := append(versionedPayload, checksum...)
	return base58.Encode(fullPayload)
}

// Função para calcular SHA256(SHA256(b))
func doubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func main() {
	for i := 0; i < 10; i++ {
		// Gerar a chave privada
		privateKeyHex := geradorChaves()
		fmt.Println("\n\nChave privada gerada:", privateKeyHex)

		// Gerar o WIF a partir da chave privada
		wif := generateWif(privateKeyHex)
		fmt.Println("WIF correspondente à chave privada:", wif)

		// Criar o hash público a partir da chave privada
		pubKeyHash := createPublicHash160(privateKeyHex)
		fmt.Printf("Public Key Hash: %x\n", pubKeyHash)

		// Codificar o hash da chave pública em um endereço Bitcoin
		address := encodeAddress(pubKeyHash)
		fmt.Println("Endereço Bitcoin correspondente à chave privada:", address)
	}
}
