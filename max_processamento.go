package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"meugo/encoding"
	"runtime"
	"sync"
	"time"
)

const (
	prefix        = "0000000000000000000000000000000000000000000000000000000000" // Prefixo da chave
	memBufferSize = 2 * 1024 * 1024 * 1024
)

var chaves_desejadas = map[string]bool{
	"14oFNXucftsHiUMY8uctg6N487riuyXs4h": true,
	"1CfZWK1QTQE3eS9qn61dQjV89KDjZzfNcv": true,
	"1L2GM8eE7mJWLdo3HZS6su1832NX2txaac": true,
	"1rSnXMr63jdCuegJFuidJqWxUPV7AtUf7":  true,
}

var (
	contador          int
	encontrado        bool
	mu                sync.Mutex
	wg                sync.WaitGroup
	ultimaChaveGerada string
	memBuffer         = make([]byte, memBufferSize) // Buffer de 1 GB
)

func gerarChavePrivada() string {
	suffix := make([]byte, 3) // Tamanho do sufixo
	_, err := rand.Read(suffix)
	if err != nil {
		log.Fatalf("Falha ao gerar chave: %v", err)
	}
	chaveGerada := prefix + hex.EncodeToString(suffix)

	// Manipulação eficiente do buffer (exemplo genérico)
	copy(memBuffer[:len(suffix)], suffix)

	return chaveGerada
}

func worker(id int) {
	defer wg.Done()

	for {
		mu.Lock()
		if encontrado {
			mu.Unlock()
			return
		}
		mu.Unlock()

		chave := gerarChavePrivada()
		pubKeyHash := encoding.CreatePublicHash160(chave)
		address := encoding.EncodeAddress(pubKeyHash)

		mu.Lock()
		contador++
		ultimaChaveGerada = chave
		if chaves_desejadas[address] {
			fmt.Printf("\n\n|--------------%s----------------|\n", address)
			fmt.Printf("|----------------------ATENÇÃO-PRIVATE-KEY-----------------------|")
			fmt.Printf("\n|%s|\n", chave)
			encontrado = true
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(10)

	// Inicia goroutines
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Goroutine para exibir o contador e a última chave gerada em tempo real
	go func() {
		for {
			time.Sleep(1 * time.Second)
			mu.Lock()
			if encontrado {
				mu.Unlock()
				break
			}
			fmt.Printf("\rChaves Geradas: %d  ", contador)
			mu.Unlock()
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			mu.Lock()
			if encontrado {
				mu.Unlock()
				break
			}
			fmt.Printf("Ultima Chave Gerada: %s ", ultimaChaveGerada)
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Print("\n\n|--------------------------------------------------by-Luan-BSC---|")
	fmt.Print("\n|-----------------------China-LOOP-MENU------------------------- |")
	fmt.Printf("\n|		Threads usados: %d		                 |", numCPU)
	fmt.Print("\n|		Chaves Analisadas:	", contador)
	fmt.Print("\n|________________________________________________________________|")
}

Bits/Bytes
1	1
2	1
3	1
4	1
5	1
6	1
7	1
8	1
9	2
10	2
11	2
12	2
13	2
14	2
15	2
16	2
17	3
18	3
19	3
20	3
21	3
22	3
23	3
24	3
25	4
26	4
27	4
28	4
29	4
30	4
31	4
32	4
33	5
34	5
35	5
36	5
37	5
38	5
39	5
40	5
41	6
42	6
43	6
44	6
45	6
46	6
47	6
48	6
49	7
50	7
51	7
52	7
53	7
54	7
55	7
56	7
57	8
58	8
59	8
60	8
61	8
62	8
63	8
64	8
65	9
66	9
67	9
68	9
69	9
70	9
71	9
72	9
73	10
74	10
75	10
76	10
77	10
78	10
79	10
80	10
81	11
82	11
83	11
84	11
85	11
86	11
87	11
88	11
89	12
90	12
91	12
92	12
93	12
94	12
95	12
96	12
97	13
98	13
99	13
100	13
101	13
102	13
103	13
104	13
105	14
106	14
107	14
108	14
109	14
110	14
111	14
112	14
113	15
114	15
115	15
116	15
117	15
118	15
119	15
120	15
121	16
122	16
123	16
124	16
125	16
126	16
127	16
128	16
129	17
130	17
131	17
132	17
133	17
134	17
135	17
136	17
137	18
138	18
139	18
140	18
141	18
142	18
143	18
144	18
145	19
146	19
147	19
148	19
149	19
150	19
151	19
152	19
153	20
154	20
155	20
156	20
157	20
158	20
159	20
160	20
