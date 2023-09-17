package main

import "fmt"

func main() {
	s := NewSpreadSheet("New Workbook", 2, 2)
	_ = s.AddValue(0, 0, 2)
	_ = s.AddValue(0, 1, 3)
	_ = s.AddValue(1, 0, 1)
	_ = s.AddValue(1, 1, "string")

	s.Display(2, 2)

	fmt.Println(s.Sum(0, 0, 1, 1))
}
