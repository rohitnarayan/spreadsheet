package main

type Function interface {
	Sum(sourceRow, sourceCol, destRow, destCol int) (float64, error)
}
