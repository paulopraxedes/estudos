# Go Trim

    Go Trim é a versão em Go do zztrim.
    Funções ZZ é um conjunto dos mais variados aplicativos, 
    escritos em shell script, com as mais variadas aplicações. Dentre eles,
    o zztrim, objeto deste estudo.

## O ZZTrim!
    Abaixo, temos a codificação, original, do miniaplicativo.

```sh
zztrim ()
{
	zzzz -h trim "$1" && return

	local top left right bottom
	local delete_top delete_left delete_right delete_bottom

	# Opções de linha de comando
	while test "${1#-}" != "$1"
	do
		case "$1" in
			-l | --left      ) shift; left=1;;
			-r | --right     ) shift; right=1;;
			-t | --top       ) shift; top=1;;
			-b | --bottom    ) shift; bottom=1;;
			-H | --horizontal) shift; left=1; right=1;;
			-V | --vertical  ) shift; top=1; bottom=1;;
			--*) zztool erro "Opção inválida $1"; return 1;;
			*) break;;
		esac
	done

	# Comportamento padrão, quando nenhuma opção foi informada
	if test -z "$top$bottom$left$right"
	then
		top=1
		bottom=1
		left=1
		right=1
	fi

	# Compõe os comandos sed para apagar os brancos,
	# levando em conta quais são as opções ativas
	test -n "$top"    && delete_top='/[^[:blank:]]/,$!d;'
	test -n "$left"   && delete_left='s/^[[:blank:]]*//;'
	test -n "$right"  && delete_right='s/[[:blank:]]*$//;'
	test -n "$bottom" && delete_bottom='
		:loop
		/^[[:space:]]*$/ {
			$ d
			N
			b loop
		}
	'

	# Dados via STDIN ou argumentos
	zztool multi_stdin "$@" |
		# Aplica os filtros
		sed "$delete_top $delete_left $delete_right" |
		# Este deve vir sozinho, senão afeta os outros (comando N)
		sed "$delete_bottom"

		# Nota: Não há problema se as variáveis estiverem vazias,
		#       sed "" é um comando nulo e não fará alterações.
}
```
## Exemplo de Implementação

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/paulopraxedes/estudos/go/gotrim"
)

func main() {


	if len(os.Args) == 2 {
		if isMenu(os.Args[1]) {
			fmt.Println("Args Error! Only menu options inserted. Check the values! ")
			return
		}
		fmt.Println(gotrim.RightTrim(gotrim.LeftTrim(os.Args[1])))

	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "-t", "--top":
			fmt.Println(gotrim.TopTrim(os.Args[2]))
		case "-b", "--bottom":
			fmt.Println(gotrim.BottomTrim(os.Args[2]))
		case "-l", "--left":
			fmt.Println(gotrim.LeftTrim(os.Args[2]))
		case "-r", "--right":
			fmt.Println(gotrim.RightTrim(os.Args[2]))
		case "-V", "--vertical":
			fmt.Println(gotrim.BottomTrim(gotrim.RightTrim(os.Args[2])))
		case "-H", "--horizontal":
			fmt.Println(gotrim.RightTrim(gotrim.LeftTrim(os.Args[2])))
		}
	}

}
func isMenu(argumento string) bool {
	if strings.Contains(argumento, "-t") || strings.Contains(argumento, "--top") ||
		strings.Contains(argumento, "-b") || strings.Contains(argumento, "--bottom") ||
		strings.Contains(argumento, "-l") || strings.Contains(argumento, "--left") ||
		strings.Contains(argumento, "-r") || strings.Contains(argumento, "--right") ||
		strings.Contains(argumento, "-V") || strings.Contains(argumento, "--vertical") ||
		strings.Contains(argumento, "-H") || strings.Contains(argumento, "--horizontal") {
		return true
	}

	return false
}
```

## Testes
    Go já possui uma ferramenta para testes sem a necessidade de usar libs de terceiros.
    Para executar os testes, basta utilizar o comando: go test
    A versão verbosa é utilizando o flag -v: go test -v

    Para verificarmos a cobertura destes,temos:

    - No pacote gotrim, temos a função LeftTrim, que remove os espaços do início do texto, conforme descrita abaixo.

```go

//LeftTrim - Funcao
func LeftTrim(text string) string {
	runes := []rune(text)
	if !unicode.IsSpace(runes[0]) || rune([]byte("\n")[0]) == runes[0] {
		return text
	}
	return LeftTrim(text[1:])
}

```

    - Para testarmos o funcionamento da função, podemos pensar em quantos caracteres (len) deveriam ter após a alteração ou qual seria o texto alocado na variável e compará-lo com o resultado esperado.
    Baseado nestes dois casos, poderiamos codificar desta forma:
    
```go
package main

import (
	"strings"
	"testing"

	"github.com/paulopraxedes/estudos/go/gotrim"
)

const text = " \n\nFoo Bar  Foo  \n\n  "

func TestLeftTrim(t *testing.T) {
	textOk := "\n\nFoo Bar  Foo  \n\n  "
    //Tendo como base o valor da constante text, a alteração esperada
    //será apenas a remoção de um caractere.
    // O primeiro "if" cobre o tamanho esperado com a alteração.
	if len(text)-1 != len(gotrim.LeftTrim(text)) {
		t.Errorf("Tamanho não esperado. [Goal:%d - Atual:%d]", len(textOk), len(gotrim.LeftTrim(text)))
	} else {
	    //O segundo "if" cobre o valor resultante do texto gerado pela função.
		if strings.Compare(textOk, gotrim.LeftTrim(text)) != 0 {
			t.Errorf("Valor Retornado não esperado. Valor esperado : %+v", textOk)
		}
	}
}
```
---
O Projeto ZZ está disponível no [GitHub](https://github.com/funcoeszz/funcoeszz).
Para contribuições, leiam o [README.md](https://github.com/funcoeszz/funcoeszz/blob/master/README.md)
