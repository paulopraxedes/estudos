package goalfabeto

import (
	"fmt"
)

/**
	* Posição 1  = --romano 		|| --latino
	* Posição 2  = --militar		|| --radio || --fone || --telefone || --otan || --nato || --icao || --itu || --imo || --faa || --ansi
	* Posição 3  = --royal 			|| --royal-navy
	* Posição 4  = --signalese 		|| --western-front
	* Posição 5  = --raf24
	* Posição 6  = --raf42
	* Posição 7  = --raf43 			|| --raf
	* Posição 8  = --us41 			|| --us
	* Posição 9  = --pt 			|| --portugal
	* Posição 10 = --name 			|| --names
	* Posição 11 = --lapd
	* Posição 12 = --morse
**/
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

/*
MostraTabela - A função exibe todo o conteúdo inserido na Tabela de Alfabeto.
@Author: Paulo Praxedes
*/
func MostraTabela() {
	for index, valor := range mapDados {
		fmt.Printf("%s - %+v\r\n", index, valor)
	}
}
