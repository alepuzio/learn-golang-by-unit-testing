package main
/**
refactoring on range
remember:
1) _ allows you to dont care about index
2) if you pass [4] int as parameter, you got an error
*/

func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
/*version with array with fixed size
func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
*/

func SumAll(numbersToSum ...[]int) (sums []int) {
    return
}
/*
func SumAll(numbersToSum ...[]int) []int {
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)

    for i, numbers := range numbersToSum {
        sums[i] = Sum(numbers)
    }

    return sums
}
*/
