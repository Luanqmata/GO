//sucesso
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"meugo/crypto/base58" // Assumindo que a função base58 já está implementada

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
)

// Chave privada hexadecimal
const privateKeyHex = "000000000000000000000000000000000000000000000000000000000003080d"

// Função para gerar WIF a partir de uma chave privada hexadecimal
func generateWif(privKeyHex string) string {
	// Decodifica a chave privada hexadecimal
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Adiciona prefixo e sufixo
	extendedKey := append([]byte{byte(0x80)}, privKeyBytes...)
	extendedKey = append(extendedKey, byte(0x01))

	// Calcula o checksum
	firstSHA := sha256.Sum256(extendedKey)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	// Adiciona o checksum
	finalKey := append(extendedKey, checksum...)

	// Codifica em base58
	wif := base58.Encode(finalKey)

	return wif
}

// Função para criar o hash público a partir de uma chave privada
func createPublicHash160(privKeyHex string) []byte {
	// Decodifica a chave privada hexadecimal
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Cria uma nova chave privada usando o pacote secp256k1
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)

	// Obtém a chave pública correspondente em formato comprimido
	compressedPubKey := privKey.PubKey().SerializeCompressed()

	// Gera um endereço Bitcoin a partir da chave pública
	pubKeyHash := hash160(compressedPubKey)
	return pubKeyHash
}

// hash160 calcula o hash RIPEMD160(SHA256(b))
func hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}

// encodeAddress codifica o hash da chave pública em um endereço Bitcoin
func encodeAddress(pubKeyHash []byte) string {
	version := byte(0x00) // 0x00 é o prefixo para endereços Bitcoin MainNet
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := doubleSha256(versionedPayload)[:4]
	fullPayload := append(versionedPayload, checksum...)
	return base58.Encode(fullPayload)
}

// doubleSha256 calcula SHA256(SHA256(b))
func doubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func main() {
	// Gera o WIF a partir da chave privada
	wif := generateWif(privateKeyHex)
	fmt.Println("\nWIF correspondente à chave privada:", wif, "\n")

	// Cria o hash público a partir da chave privada
	pubKeyHash := createPublicHash160(privateKeyHex)
	fmt.Printf("Public Key Hash: %x\n", pubKeyHash)

	// Codifica o hash da chave pública em um endereço Bitcoin
	address := encodeAddress(pubKeyHash)
	fmt.Printf("Endereço Bitcoin: %s\n", address)
}
