package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")
}

type Car interface {
	
}

type Car struct {
	numOfWheel uint8
	normalSpeed float64

	engine     Engine
	

}

func (c *Car) PrintInfo() {

}

type Engine struct {

}

type Wheel struct {

}

type Seat struct {

}

