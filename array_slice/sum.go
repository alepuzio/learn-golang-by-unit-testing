package main
/**
refactoring on range
remember:
1) _ allows you to dont care about index
2) if you pass [4] int as parameter, you got an error
*/
func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
