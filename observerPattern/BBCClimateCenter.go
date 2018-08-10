package main

import "fmt"

type bbcClimate struct {
	temp     int
	humidity int
	pressure int
}

func NewBBC() ClimateObserver {
	return &bbcClimate{}
}

func (c *bbcClimate) UpdateParams(temp, humidity, pressure int) {
	c.pressure = pressure
	c.humidity = humidity
	c.temp = temp
	c.printOnDevice()
}
func (c *bbcClimate) printOnDevice() {
	fmt.Println("BBC center: ", c.temp, c.humidity, c.pressure)
}
