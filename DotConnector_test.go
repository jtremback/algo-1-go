package main

import (
	"fmt"
	"testing"
)

var definitelyLine = []Dot{
	{X: 4000, Y: 30000},
	{X: 5000, Y: 30000},
	{X: 6000, Y: 30000},
	{X: 7000, Y: 30000},
	{X: 8000, Y: 30000},
}

var input10 = []Dot{
	{X: 4000, Y: 30000},
	{X: 3500, Y: 28000},
	{X: 3000, Y: 26000},
	{X: 2000, Y: 22000},
	{X: 1000, Y: 18000},
	{X: 13000, Y: 21000},
	{X: 23000, Y: 16000},
	{X: 28000, Y: 13500},
	{X: 28000, Y: 5000},
	{X: 28000, Y: 1000},
}

func TestDotConnector(t *testing.T) {
	fmt.Printf("%+v\n", DotConnector(definitelyLine))
}
