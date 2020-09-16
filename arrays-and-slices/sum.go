package main

//Sum adds an array of numbers together
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {

}
