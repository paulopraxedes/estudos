package goalfa

import (
	"fmt"
	"sort"
)

/**
 * mapDados = Objeto do tipo Mapa, que representa os alfabetos possíveis
 */
var mapDados = map[int][]string{
	0:  {"A", "Alpha", "Apples", "Ack", "Ace", "Apple", "Able/Affirm", "Able", "Aveiro", "Alan", "Adam", ".-"},
	1:  {"B", "Bravo", "Butter", "Beer", "Beer", "Beer", "Baker", "Baker", "Bragança", "Bobby", "Boy", "-..."},
	2:  {"C", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Coimbra", "Charlie", "Charles", "-.-."},
	3:  {"D", "Delta", "Duff", "Don", "Don", "Dog", "Dog", "Dog", "Dafundo", "David", "David", "-.."},
	4:  {"E", "Echo", "Edward", "Edward", "Edward", "Edward", "Easy", "Easy", "Évora", "Edward", "Edward", "."},
	5:  {"F", "Foxtrot", "Freddy", "Freddie", "Freddie", "Freddy", "Fox", "Fox", "Faro", "Frederick", "Frank", "..-."},
	6:  {"G", "Golf", "George", "Gee", "George", "George", "George", "George", "Guarda", "George", "George", "--."},
	7:  {"H", "Hotel", "Harry", "Harry", "Harry", "Harry", "How", "How", "Horta", "Howard", "Henry", "...."},
	8:  {"I", "India", "Ink", "Ink", "Ink", "In", "Item/Interrogatory", "Item", "Itália", "Isaac", "Ida", ".."},
	9:  {"J", "Juliet", "Johnnie", "Johnnie", "Johnnie", "Jug/Johnny", "Jig/Johnny", "Jig", "José", "James", "John", ".---"},
	10: {"K", "Kilo", "King", "King", "King", "King", "King", "King", "Kilograma", "Kevin", "King", "-.-"},
	11: {"L", "Lima", "London", "London", "London", "Love", "Love", "Love", "Lisboa", "Larry", "Lincoln", ".-.."},
	12: {"M", "Mike", "Monkey", "Emma", "Monkey", "Mother", "Mike", "Mike", "Maria", "Michael", "Mary", "--"},
	13: {"N", "November", "Nuts", "Nuts", "Nuts", "Nuts", "Nab/Negat", "Nan", "Nazaré", "Nicholas", "Nora", "-."},
	14: {"O", "Oscar", "Orange", "Oranges", "Orange", "Orange", "Oboe", "Oboe", "Ovar", "Oscar", "Ocean", "---"},
	15: {"P", "Papa", "Pudding", "Pip", "Pip", "Peter", "Peter/Prep", "Peter", "Porto", "Peter", "Paul", ".--."},
	16: {"Q", "Quebec", "Queenie", "Queen", "Queen", "Queen", "Queen", "Queen", "Queluz", "Quincy", "Queen", "--.-"},
	17: {"R", "Romeo", "Robert", "Robert", "Robert", "Roger/Robert", "Roger", "Roger", "Rossio", "Robert", "Robert", ".-."},
	18: {"S", "Sierra", "Sugar", "Esses", "Sugar", "Sugar", "Sugar", "Sugar", "Setúbal", "Stephen", "Sam", "..."},
	19: {"T", "Tango", "Tommy", "Toc", "Toc", "Tommy", "Tare", "Tare", "Tavira", "Trevor", "Tom", "-"},
	20: {"U", "Uniform", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Unidade", "Ulysses", "Union", "..-"},
	21: {"V", "Victor", "Vinegar", "Vic", "Vic", "Vic", "Victor", "Victor", "Viseu", "Vincent", "Victor", "...-"},
	22: {"W", "Whiskey", "Willie", "William", "William", "William", "William", "William", "Washington", "William", "William", ".--"},
	23: {"X", "X-ray/Xadrez", "Xerxes", "X-ray", "X-ray", "X-ray", "X-ray", "X-ray", "Xavier", "Xavier", "X-ray", "-..-"},
	24: {"Y", "Yankee", "Yellow", "Yorker", "Yorker", "Yoke/Yorker", "Yoke", "Yoke", "York", "Yaakov", "Young", "-.--"},
	25: {"Z", "Zulu", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zulmira", "Zebedee", "Zebra", "--.."},
}

/**
ordenaTabela - função que ordena a lista, de modo crescente, para uma melhor leitura
*/
func ordenaTabela(tabela map[int][]string) (keys []int) {
	for valor := range tabela {
		keys = append(keys, valor)
	}
	sort.Ints(keys)
	return
}

/*MostraTabela - função que imprime a tabela gerada.
Tendo uma única coluna, imprime o valor do índice
Tendo mais que uma coluna, imprime o slice inteiro
*/
func MostraTabela(tabela map[int][]string) {
	if tabela == nil {
		tabela = mapDados
	}
	for valor := range ordenaTabela(tabela) {
		if len(tabela[valor]) > 1 {
			fmt.Printf("%+s\r\n", tabela[valor])
		} else {
			fmt.Printf("%+s\r\n", tabela[valor][0])
		}
	}
}
