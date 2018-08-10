package main

func main() {
	climateDataCenter := NewNationalClimateData()
	cnn := NewCNN()
	bbc := NewBBC()
	climateDataCenter.RegisterObserver(cnn)
	climateDataCenter.RegisterObserver(bbc)

	climateDataCenter.UpdateParams(10, 10, 10)
	climateDataCenter.UpdateParams(20, 10, 10)
}