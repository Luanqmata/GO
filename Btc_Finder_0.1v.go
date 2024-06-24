// codigo com erros e quadruplicação de chaves
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

var chaves_desejadas = map[string]bool{
	"1BDyrQ6WoF8VN3g9SAS1iKZcPzFfnDVieY": true,
	"1QCbW9HWnwQWiQqVo5exhAnmfqKRrCRsvW": true,
	"1ErZWg5cFCe4Vw5BzgfzB74VNLaXEiEkhk": true,
	"1Pie8JkxBT6MGPz9Nvi3fsPkr2D8q3GBc1": true,
}

func geradorChaves() string {
	rand.Seed(time.Now().UnixNano())
	for {
		var caracteresUnicos []byte
		for i := 0; i < 4; i++ {
			randomIndex := rand.Intn(len(characters))
			caracteresUnicos = append(caracteresUnicos, characters[randomIndex])
		}
		chaveGerada := prefix + string(caracteresUnicos)
		return chaveGerada
	}
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

func hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}

func encodeAddress(pubKeyHash []byte) string {
	version := byte(0x00)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := doubleSha256(versionedPayload)[:4]
	fullPayload := append(versionedPayload, checksum...)
	return base58.Encode(fullPayload)
}

func doubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func main() {
	for i := 0; i < 65536; i++ {
		privateKeyHex := geradorChaves()
		generateWif(privateKeyHex)
		pubKeyHash := createPublicHash160(privateKeyHex)
		address := encodeAddress(pubKeyHash)

		if chaves_desejadas[address] {
			fmt.Printf("\nEndereço encontrado: %s\n", address)
			fmt.Printf("Chave privada correspondente: %s\n", privateKeyHex)
			// Aqui você pode adicionar a lógica para imprimir a carteira correspondente ao hexa, se necessário
			break
		} else {
			fmt.Printf("%s\n", privateKeyHex)
		}
	}
	fmt.Print("O maximo de combinaçoes posiveis foi atingido")
}
