//conversor ainda está com erros
package main

import (
	"encoding/hex"
	"fmt"
	"meugo/crypto/base58"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	privateKeyHex := "000000000000000000000000000000000000000000000000000000000001764f"

	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		fmt.Println("Erro ao decodificar chave privada:", err)
		return
	}

	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)

	publicKey := privateKey.PubKey()

	publicKeyBytes := publicKey.SerializeCompressed()
	
	address := base58.Encode(publicKeyBytes)

	fmt.Println("Endereço Bitcoin correspondente à chave privada:", address)
}
