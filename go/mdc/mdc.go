package mdc

/*
CalculaMDC - Método que calcula o maior divisor comum dos inteiros estritamente positivos m e n, usando recursividade.
*/
func CalculaMDC(m int, n int) int {
	if n == 0 {
		return m
	}
	return CalculaMDC(n, m%n)
}
