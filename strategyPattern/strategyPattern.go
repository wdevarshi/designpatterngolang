package strategyPattern

import "strconv"

type Strategy interface {
	Resolve(ccid string) (promotionId string)
}

type DumbStrategy struct {
	dumbCount int
}

func (d *DumbStrategy) Resolve(ccid string) (promotionId string) {
	return "123"
}

type SmartStrategy struct {
	smartnessFactor string
}

func (s *SmartStrategy) Resolve(ccid string) (promotionId string) {
	return s.smartnessFactor + ccid
}

func GetPromotions(ccid string) (promotionIds []string) {
	ds := &DumbStrategy{dumbCount: 10}
	ss := &SmartStrategy{smartnessFactor: "OutOfBound"}

	strategies := make([]Strategy, 2)
	strategies[0] = ds
	strategies[1] = ss

	var pids []string

	for i, s := range strategies {
		pids = append(pids, s.Resolve(strconv.Itoa(i)))
	}

	return pids
}
