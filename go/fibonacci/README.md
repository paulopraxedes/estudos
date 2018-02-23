# Fibonacci
---
## Um pouco de história
Leonardo Fibonacci, foi um matemático italiano que "popularizou" a sequência que, após a sua morte, herdou seu sobrenome.
Em termos matemáticos, a sequência é definida recursivamente, pela fórmula:

![Function](http://upload.wikimedia.org/wikipedia/pt/math/0/d/5/0d5cce25d67941bb4661afd52609d93c.png)

---
### Exemplos de Implementações
Há maneiras diversas para implementar, contudo implementei de duas formas, mostrando a forma clássica (por um loop) e da forma matematicamente aderente, tendo como base o fundamento matemático

#### Fibonacci, by loop
```go
    /*Esta forma implementa o conceito de loop, onde iremos armazenar os dados, da sequência, em posições do vetor iterativamente.
    A Função espera como parâmetro um número inteiro e seu retorno é um vetor (array) de inteiros assinado. 
    */ 
    func fibo(posicoes int) (numeros []int) {
	numeros = make([]int, posicoes) //Instancio o objeto, criando um "array", baseado na quantidade informada

	for index := 0; index < posicoes; index++ { //Itero, através do for, que variará do numeral 0 até n - 1 posições
        /*
         Neste ponto faço uma comparação, através da instrução switch. Se o index for 0 ou 1 insiro o valor 1. Qualquer outro valor do index, somarei os valores das duas casas anteriores.
         Imaginemos este exemplo:
         Um vetor de 4 posições, no momento do preenchimento se comportará:
         posição 0 = 1
         posição 1 = 1
         posição 2 = posicao 1 + posicao 0 = 1 + 1 = 2 
         posição 3 = posicao 2 + posicao 1 = 2 + 1 = 3
        */
        switch index { 
		case 0, 1:
			numeros[index] = 1
		default:
			numeros[index] = numeros[index-1] + numeros[index-2]
		}
	}
	return //retorno da função
}
```
#### Fibonacci, by recursion

```go
    /*Esta forma implementa o conceito da recursividade. Em geral, este tipo de algoritmo são considerados "mais enxutos" ou "elegantes".
    A função espera como parâmetro um número inteiro e retornará um outro número inteiro. 
    Como é o funcionamento: Imaginemos que queremos o valor da quinta posição da sequência. Logo, teremos o seguinte: 
    F(5) = F(4) + F(3) = 3 + 2 = 5
        F(4) = F(3) + F(2) = 2 + 1 = 3
            F(3) = F(2) + F(1) = 1 + 1 = 2
                F(2) = 1
                    F(1) = 1
    Quando a função chegar ao F(1), ela retornará todos os valores contidos nesta "pilha"
    */
func fiboRec(posicao int) int {
	switch posicao {
	case 1, 2:
		return 1
	default:
		return fiboRec(posicao-1) + fiboRec(posicao-2)
	}
}
```
---

## Uso das funções

As implementações de cada uma das funções tem usabilidades distintas. A Iterativa, retornará um vetor com os respectivos valores inseridos. A Recursiva, retornará o valor da posição desejada, tão somente. Dito isso, seguem-se os exemplos:

#### Função Iterativa
```go
package main

import "fmt"
/*
    Função Iterativa
*/
func main(){
    for index, valor := range fibo(10) {
		fmt.Printf("%d", valor)
		if index+1 < len(fibo(10)) {
			fmt.Printf(",")
		}
	}
	fmt.Println()
}
func fibo(posicoes int) (numeros []int) {
    numeros = make([]int, posicoes) 
    	for index := 0; index < posicoes; index++ { 
        switch index { 
		case 0, 1:
			numeros[index] = 1
		default:
			numeros[index] = numeros[index-1] + numeros[index-2]
		}
	}
	return
}
```
#### Função Recursiva
```go
package main

import "fmt"
/*
    Função Recursiva
*/
func main(){
    // Se quisermos todos os 10 valores da sequência, por exemplo:
    for index := 0; index < 10; index++ {
		fmt.Printf("%d ", fiboRec(index+1))
    }
    
    // Ou se quisermos apenas um determinado valor da sequência, por exemplo:
    fmt.Printf("%d", fiboRec(10))
}

func fiboRec(posicao int) int {
    switch posicao {
    case 1, 2:
        return 1
    default:
        return fiboRec(posicao-1) + fiboRec(posicao-2)
    }
}
```
---

### Playground
[Playground](https://play.golang.org/p/EBw-PwwE54P)