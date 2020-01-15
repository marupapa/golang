package main

import (
	"fmt"
	"strings"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func numJewelsInStones(J string, S string) int {

	sArr := strings.Split(S, "")

	count := 0
	for _, sItem := range sArr {

		//fmt.Printf("item:%v\n", sItem)
		if strings.ContainsAny(J, sItem) {

			//fmt.Printf("item:%v\n", sItem)
			count = count + 1
		}

	}

	//fmt.Printf("count:%v \n", count)
	return count
}

func solution1(n int) int {

	var sum int

	for i := 1; i <= n; i++ {

		//fmt.Println(i)
		if (n % i) == 0 {
			sum = sum + i

			fmt.Println(i)
		}
	}

	return sum
}
func solution2(n int, m int) []int {

	var s1 []int
	var s2 []int

	for i := 1; i <= n; i++ {
		if (n % i) == 0 {
			s1 = append(s1, i)
		}
	}

	for i := 1; i <= m; i++ {
		if (n % i) == 0 {
			s2 = append(s2, i)
		}
	}

	/*최대공약수*/
	result1 := []int{}
	max := 0
	for i, item := range s1 {
		if s2[i] == item {
			result1 = append(result1, item)
		}

		if i != 0 && result1[i-1] < result1[i] {
			max = result1[i]
		}

	}

	gcm := n * m / max

	return []int{max, gcm}
}
func solutionX(arr []int) float64 {

	var sum float64

	for _, item := range arr {

		sum = sum + float64(item)
	}

	result := sum / float64(len(arr))

	return result
}
func main() {

	//numJewelsInStones("aA", "aAAbbbb")

	//fmt.Printf("solution value : %v \n", solution1(12))
	//fmt.Printf("solution value : %v \n", solution2(3, 12))

	num := 10

	var result string

	if num%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}

	fmt.Printf("solution value : %v \n", num%2)

}
func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
