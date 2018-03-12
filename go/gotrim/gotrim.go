package gotrim

import (
	"unicode"
)

//LeftTrim - Função recursiva que percorre o slice e remove os caracteres "espaços", da esquerda para a direita
func LeftTrim(text string) string {
	runes := []rune(text)
	if !unicode.IsSpace(runes[0]) || rune([]byte("\n")[0]) == runes[0] {
		return text
	}
	return LeftTrim(text[1:])
}

//RightTrim - Função recursiva que percorre o slice e remove os caracteres "espaços", da direita para a esquerda
func RightTrim(text string) string {
	runes := []rune(text)
	if !unicode.IsSpace(runes[len(runes)-1]) || rune([]byte("\n")[0]) == runes[len(runes)-1] {
		return text
	}
	return RightTrim(text[:len(text)-1])
}

//TopTrim - Função recursiva que verifica se a primeira posição do texto é uma quebra de linha.
//Sendo, ela é removida até que seja retornado um texto válido. Um "espaço" é válido.
//A direção é do topo do texto para baixo
func TopTrim(text string) string {
	b := []byte("\n")
	if text[0] == b[0] {
		return TopTrim(text[1:])
	}
	return text
}

//BottomTrim - Função recursiva que verifica se a última posição do texto é uma quebra de linha.
//Sendo, ela é removida até que seja retornado um texto válido. Um "espaço" é válido.
//A direção é de baixo do texto para o topo
func BottomTrim(text string) string {
	b := []byte("\n")
	if text[len(text)-1] == b[0] {
		return BottomTrim(text[:len(text)-1])
	}
	return text
}
