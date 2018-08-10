package main

//We are encapulating this struct. The new method will return the ClimateDataCenter interface.
//No one would know about the existences
type nationalClimateData struct {
	climateObserver []ClimateObserver
	temp            int
	humidity        int
	pressure        int
}

func (n *nationalClimateData) UpdateParams(temp, humidity, pressure int) {
	n.pressure = pressure
	n.temp = temp
	n.humidity = humidity
	n.notifyAllObserver()
}

func NewNationalClimateData() ClimateDataCenter {
	return &nationalClimateData{}
}

func (n *nationalClimateData) RegisterObserver(observer ClimateObserver) {
	n.climateObserver = append(n.climateObserver, observer)
}

func (n *nationalClimateData) notifyAllObserver() {
	for _, o := range n.climateObserver {
		o.UpdateParams(n.temp, n.humidity, n.pressure)
	}
}
