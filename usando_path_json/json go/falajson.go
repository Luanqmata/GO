package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Wallets struct {
	Wallets []string `json:"wallets"`
}

type Range struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type Ranges struct {
	Ranges []Range `json:"ranges"`
}

func main() {
	walletFile := "enderecos/wallets.json"
	rangeFile := "enderecos/ranges.json"

	walletData, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de carteiras: %v", err)
	}

	rangeData, err := ioutil.ReadFile(rangeFile)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de ranges: %v", err)
	}

	var wallets Wallets
	var ranges Ranges

	err = json.Unmarshal(walletData, &wallets)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal dos dados de carteiras: %v", err)
	}

	err = json.Unmarshal(rangeData, &ranges)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal dos dados de ranges: %v", err)
	}

	// Perguntar ao usuário a carteira desejada
	fmt.Println("Digite o número da carteira desejada (1 para a primeira carteira, 2 para a segunda, etc.):")
	var choice int
	_, err = fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(wallets.Wallets) {
		log.Fatal("Escolha inválida.")
	}

	// Encontrar o range correspondente
	rangeIndex := choice // Adicionar 2 ao índice da carteira (1 base)
	if rangeIndex > len(ranges.Ranges) {
		log.Fatal("Índice de range fora dos limites.")
	}

	selectedRange := ranges.Ranges[rangeIndex-1]
	fmt.Printf("Range correspondente: Min = %s, Max = %s\n", selectedRange.Min, selectedRange.Max)
}
