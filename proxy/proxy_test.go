package proxy

import "testing"

func TestProxy(t *testing.T) {
	station := &Station{}
	station.stock = 1
	stationProxy := &StationProxy{station: station}

	stationProxy.sell("jack")
	stationProxy.sell("lucy")
}
