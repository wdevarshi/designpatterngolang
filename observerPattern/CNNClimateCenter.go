package main

import "fmt"

type cnnClimate struct {
	temp     int
	humidity int
	pressure int
}

func NewCNN() ClimateObserver {
	return &cnnClimate{}
}

func (c *cnnClimate) UpdateParams(temp, humidity, pressure int) {
	c.pressure = pressure
	c.humidity = humidity
	c.temp = temp
	c.printOnDevice()
}
func (c *cnnClimate) printOnDevice() {
	fmt.Println("CNN center: ", c.temp, c.humidity, c.pressure)
}
