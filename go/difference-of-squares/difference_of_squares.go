package diffsquares

//SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result * result

}

//SumOfSquares returns the sum of the squares of the first n natural number
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + (i * i)
	}
	return result
}

//Difference return the diference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
