package example

import "fmt"

func main() {
	const hello string = "Hello go"
	const thisYear int = 2022
	var year []int
	var yearConv = map[int]int{2022: 2565}
	year = append(year, thisYear)
	yearConv[2023] = 2566

	fmt.Println(hello)
	fmt.Println(thisYear)
	fmt.Println(year)
	fmt.Println(year[0])
	fmt.Println(yearConv)
	fmt.Println(yearConv[2022])

	fmt.Println()
	var loopData []string = []string{"Hello", "World", "2022"}
	for idx := 0; idx < len(loopData); idx += 1 {
		fmt.Println(loopData[idx])
	}
	for idx, val := range loopData {
		fmt.Println("Index = ", idx, "Value = ", val)
	}
}
