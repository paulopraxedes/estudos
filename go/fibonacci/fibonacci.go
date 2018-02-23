package fibonacci

/*
CalculaFibonacci - Calcula o valor da sequência de Fibonacci, baseado na posição informada
*/
func CalculaFibonacci(posicao int) int {
	switch posicao {
	case 1, 2:
		return 1
	default:
		return CalculaFibonacci(posicao-1) + CalculaFibonacci(posicao-2)
	}
}
