func setupRanges() {
	rangeStr, existe := ranges[tamanhoChave]
	if !existe {
		log.Fatalf("Só existem 160 chaves. %d este número não é aceito.", tamanhoChave)
	}

	valores := strings.Split(rangeStr, "-")
	minRange = new(big.Int)
	minRange.SetString(valores[0], 16)
	maxRange = new(big.Int)
	maxRange.SetString(valores[1], 16)

	// Calcula o meio do intervalo
	midRange := new(big.Int).Set(maxRange)
	midRange.Add(midRange, minRange)
	midRange.Rsh(midRange, 1) // Divide por 2

	fmt.Println("Escolha um dos dois blocos abaixo:")
	fmt.Printf("1: %s - %s\n", minRange.Text(16), midRange.Text(16))
	fmt.Printf("2: %s - %s\n", midRange.Text(16), maxRange.Text(16))

	var escolha int
	fmt.Print("Digite 1 ou 2: ")
	_, err := fmt.Scanf("%d", &escolha)
	if err != nil || (escolha != 1 && escolha != 2) {
		log.Fatalf("Escolha inválida. Por favor, digite 1 ou 2.")
	}

	// Ajusta minRange e maxRange com base na escolha
	if escolha == 1 {
		maxRange.Set(midRange)
	} else {
		minRange.Set(midRange)
	}
}
