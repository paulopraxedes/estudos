package goalfabeto

import (
	"fmt"
	"sort"
)

/**
 * mapDados = Objeto do tipo Mapa, que representa os alfabetos possíveis
 */
var mapDados = map[string][]string{
	"A": {"Alpha", "Apples", "Ack", "Ace", "Apple", "Able/Affirm", "Able", "Aveiro", "Alan", "Adam", ".-"},
	"B": {"Bravo", "Butter", "Beer", "Beer", "Beer", "Baker", "Baker", "Bragança", "Bobby", "Boy", "-..."},
	"C": {"Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Coimbra", "Charlie", "Charles", "-.-."},
	"D": {"Delta", "Duff", "Don", "Don", "Dog", "Dog", "Dog", "Dafundo", "David", "David", "-.."},
	"E": {"Echo", "Edward", "Edward", "Edward", "Edward", "Easy", "Easy", "Évora", "Edward", "Edward", "."},
	"F": {"Foxtrot", "Freddy", "Freddie", "Freddie", "Freddy", "Fox", "Fox", "Faro", "Frederick", "Frank", "..-."},
	"G": {"Golf", "George", "Gee", "George", "George", "George", "George", "Guarda", "George", "George", "--."},
	"H": {"Hotel", "Harry", "Harry", "Harry", "Harry", "How", "How", "Horta", "Howard", "Henry", "...."},
	"I": {"India", "Ink", "Ink", "Ink", "In", "Item/Interrogatory", "Item", "Itália", "Isaac", "Ida", ".."},
	"J": {"Juliet", "Johnnie", "Johnnie", "Johnnie", "Jug/Johnny", "Jig/Johnny", "Jig", "José", "James", "John", ".---"},
	"K": {"Kilo", "King", "King", "King", "King", "King", "King", "Kilograma", "Kevin", "King", "-.-"},
	"L": {"Lima", "London", "London", "London", "Love", "Love", "Love", "Lisboa", "Larry", "Lincoln", ".-.."},
	"M": {"Mike", "Monkey", "Emma", "Monkey", "Mother", "Mike", "Mike", "Maria", "Michael", "Mary", "--"},
	"N": {"November", "Nuts", "Nuts", "Nuts", "Nuts", "Nab/Negat", "Nan", "Nazaré", "Nicholas", "Nora", "-."},
	"O": {"Oscar", "Orange", "Oranges", "Orange", "Orange", "Oboe", "Oboe", "Ovar", "Oscar", "Ocean", "---"},
	"P": {"Papa", "Pudding", "Pip", "Pip", "Peter", "Peter/Prep", "Peter", "Porto", "Peter", "Paul", ".--."},
	"Q": {"Quebec", "Queenie", "Queen", "Queen", "Queen", "Queen", "Queen", "Queluz", "Quincy", "Queen", "--.-"},
	"R": {"Romeo", "Robert", "Robert", "Robert", "Roger/Robert", "Roger", "Roger", "Rossio", "Robert", "Robert", ".-."},
	"S": {"Sierra", "Sugar", "Esses", "Sugar", "Sugar", "Sugar", "Sugar", "Setúbal", "Stephen", "Sam", "..."},
	"T": {"Tango", "Tommy", "Toc", "Toc", "Tommy", "Tare", "Tare", "Tavira", "Trevor", "Tom", "-"},
	"U": {"Uniform", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Unidade", "Ulysses", "Union", "..-"},
	"V": {"Victor", "Vinegar", "Vic", "Vic", "Vic", "Victor", "Victor", "Viseu", "Vincent", "Victor", "...-"},
	"W": {"Whiskey", "Willie", "William", "William", "William", "William", "William", "Washington", "William", "William", ".--"},
	"X": {"X-ray/Xadrez", "Xerxes", "X-ray", "X-ray", "X-ray", "X-ray", "X-ray", "Xavier", "Xavier", "X-ray", "-..-"},
	"Y": {"Yankee", "Yellow", "Yorker", "Yorker", "Yoke/Yorker", "Yoke", "Yoke", "York", "Yaakov", "Young", "-.--"},
	"Z": {"Zulu", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zulmira", "Zebedee", "Zebra", "--.."},
}

/*MostraTabela - A função pública exibe todo o conteúdo inserido na Tabela de Alfabeto.
 * @Author: Paulo Praxedes
 */
func MostraTabela() {
	/*
	 * Iteramos o map, baseados na ordem das chaves do array. Desta forma, garantimos a exibição dos dados ordenados
	 */
	for _, valor := range OrdenaMapa(mapDados) {
		fmt.Printf("%s - %+v\r\n", valor, mapDados[valor])
	}
}

/*MostraAlfabeto - A função pública exibe todo o conteúdo inserido na Tabela de Alfabeto, baseado no Tipo informado.
 * @Author: Paulo Praxedes
 */
func MostraAlfabeto(tipo string) {
	switch tipo {
	case "--militar", "--radio", "--fone", "--telefone", "--otan", "--nato", "--icao", "--itu", "--imo", "--faa", "--ansi":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][0])
		}
	case "--romano", "--latino":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %s\r\n", valor, valor)
		}
	case "--royal", "--royal-navy":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][1])
		}
	case "--signalese", "--western-front":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][2])
		}
	case "--raf24":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][3])
		}
	case "--raf42":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][4])
		}
	case "--raf43", "--raf":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][5])
		}
	case "--us41", "--us":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][6])
		}
	case "--pt", "--portugal":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][7])
		}
	case "--name", "--names":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][8])
		}
	case "--lapd":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][9])
		}
	case "--morse":
		for _, valor := range OrdenaMapa(mapDados) {
			fmt.Printf("%s - %+v\r\n", valor, mapDados[valor][10])
		}
	}
}

/*MostraAlfabetoFormatado - A função pública exibe todo o conteúdo inserido na Tabela de Alfabeto, baseado no Tipo e
 * palvra informado.
 * @Author: Paulo Praxedes
 */
func MostraAlfabetoFormatado(tipo string, valor string) {
	switch tipo {
	case "--militar", "--radio", "--fone", "--telefone", "--otan", "--nato", "--icao", "--itu", "--imo", "--faa", "--ansi":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][0])
		}
	case "--romano", "--latino":
		for _, caractere := range valor {
			fmt.Printf("%s\r\n", string(caractere))
		}
	case "--royal", "--royal-navy":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][1])
		}
	case "--signalese", "--western-front":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][2])
		}
	case "--raf24":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][3])
		}
	case "--raf42":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][4])
		}
	case "--raf43", "--raf":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][5])
		}
	case "--us41", "--us":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][6])
		}
	case "--pt", "--portugal":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][7])
		}
	case "--name", "--names":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][8])
		}
	case "--lapd":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][9])
		}
	case "--morse":
		for _, caractere := range valor {
			fmt.Printf("%s - %s\r\n", string(caractere), mapDados[string(caractere)][10])
		}
	default:
		fmt.Printf("Valor digitado, não encontrado na tabela. :(")
	}
}

/*OrdenaMapa - A função privada ordena o mapa
 * @Author: Paulo Praxedes
 */
func OrdenaMapa(mapa map[string][]string) (keys []string) {
	for k := range mapa {
		keys = append(keys, k)
	}

	sort.Strings(keys) // Colocamos em ordem alfabética, para uma melhor legibilidade

	return
}
