package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type spreadSheet struct {
	name  string
	cells [][]Cell
}

func NewSpreadSheet(name string, rows, cols int) *spreadSheet {
	c := make([][]Cell, rows)
	for a := range c {
		c[a] = make([]Cell, cols)
	}
	return &spreadSheet{
		name:  name,
		cells: c,
	}
}

type SpreadSheet interface {
	AddValue(row, col int, value interface{}) error
	DeleteValue(row, col int, value interface{}) error
	Display(row, col int)
}

func (s *spreadSheet) AddValue(row, col int, value interface{}) error {
	s.cells[row][col].value = value
	return nil
}

func (s *spreadSheet) DeleteValue(row, col int, value interface{}) error {
	s.cells[row][col].value = nil
	return nil
}

func (s *spreadSheet) Display(row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			a := s.cells[i][j].value
			switch a.(type) {
			case int:
				fmt.Print(a.(int))
			case float64:
				fmt.Print(a.(float64))
			case string:
				fmt.Print(a.(string))

			}
		}
		fmt.Println()
	}
}

func (s *spreadSheet) Sum(sourceRow, sourceCol, destRow, destCol int) (float64, error) {
	result := 0.0
	for i := sourceRow; i <= destRow; i++ {
		for j := sourceCol; j <= destCol; j++ {
			a := s.cells[i][j].value
			switch a.(type) {
			case int, int32, int64:
				result += float64(a.(int))
			case float64:
				result += float64(a.(int))
			default:
				return 0.0, errors.Errorf("non numerical data type received, val: %+v", a)
			}
		}
	}

	return result, nil
}
