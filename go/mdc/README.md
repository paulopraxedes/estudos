# MDC
---
## Um pouco de história
Significa Máximo Divisor Comum. O MDC comum entre dois ou mais números naturais positivos (e maiores de zero) é o maior de seus divisores. Dois números naturais sempre tem divisores em comum. Um entendimento aplicado ao MDC é dado pelo exemplo abaixo:

![MDC](https://matematicabasica.net/wp-content/uploads/2015/09/mdc-m%C3%A1ximo-divisor-comum.png)

---
### Exemplos de Implementações

#### Por Recursão

```go
func calculaMDC(m int, n int) int {
	if n == 0 {
		return m
	}
	return calculaMDC(n, m%n)
}
```

#### Por Iteratividade

```go
func calculaMDC(m int, n int) int {
	var r int
	for{
		r = m%n
		m = n
		n = r
		
		if r == 0 {
			break
		}
	}
	return m
}
```
---

## Uso das funções

#### Função Iterativa
[Playground](https://play.golang.org/p/jbjzvtYGcWH)
```go
package main

import "fmt"
/*
    Função Iterativa
*/
func main() {
	fmt.Println(calculaMDC(16,64))
}

func calculaMDC(m int, n int) int {
	var r int
	for{
		r = m%n
		m = n
		n = r
		
		if r == 0 {
			break
		}
	}
	return m
}
```
#### Função Recursiva
[Playground](https://play.golang.org/p/xgWO6lvuM_V)
```go
package main

import "fmt"
/*
    Função Recursiva
*/
func main() {
	fmt.Println(calculaMDC(16,64))
}

func calculaMDC(m int, n int) int {
	if n == 0 {
		return m
	}
	return calculaMDC(n, m%n)
}
```
---