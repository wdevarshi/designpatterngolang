package main

type ClimateObserver interface {
	UpdateParams(temp, humidity, pressure int)
}

type ClimateDataCenter interface {
	RegisterObserver(observer ClimateObserver)
	UpdateParams(temp, humidity, pressure int)
}

