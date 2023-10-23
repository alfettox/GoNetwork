package main

import (
	"fmt"
	"sync"
	"time"
)

// NetworkSlice represents a network slice with its properties.
type NetworkSlice struct {
	Name        string
	Bandwidth   int
	Latency     int
	ServiceType string
	TimeOfDay   string
}

// NetworkSliceDirectory maintains a list of network slices with specific requirements.
type NetworkSliceDirectory struct {
	Slices []*NetworkSlice
}

// ServiceResult represents the result of a service request.
type ServiceResult struct {
	ServiceType string
	TimeOfDay   string
	Efficiency  float64
}

// SelectNetworkSlice simulates the selection of a network slice based on service and time.
func (nsd *NetworkSliceDirectory) SelectNetworkSlice(serviceType string, timeOfDay string) *NetworkSlice {
	var selectedSlice *NetworkSlice

	for _, slice := range nsd.Slices {
		if slice.ServiceType == serviceType && slice.TimeOfDay == timeOfDay {
			selectedSlice = slice
			break
		}
	}

	if selectedSlice == nil {
		selectedSlice = &NetworkSlice{Name: "Default", Bandwidth: 100, Latency: 20}
	}

	return selectedSlice
}

// CalculateEfficiency calculates the efficiency of network resource usage.
func CalculateEfficiency(selectedSlice *NetworkSlice, serviceType string) float64 {
	maxBandwidth := 1000
	maxLatency := 10
	bandwidthEfficiency := float64(selectedSlice.Bandwidth) / float64(maxBandwidth) * 100
	latencyEfficiency := float64(maxLatency) / float64(selectedSlice.Latency) * 100
	overallEfficiency := (bandwidthEfficiency + latencyEfficiency) / 2
	return overallEfficiency
}

// runSimulation simulates service requests and sends results to a channel.
func runSimulation(timeMultiplier int, resultChan chan<- *ServiceResult, wg *sync.WaitGroup) {
	// Initialize network slice directory with different network slices.
	networkSlices := []*NetworkSlice{
		{Name: "Morning Game", Bandwidth: 500, Latency: 10, ServiceType: "Game", TimeOfDay: "Morning"},
		{Name: "Evening Game", Bandwidth: 1000, Latency: 5, ServiceType: "Game", TimeOfDay: "Evening"},
		{Name: "WebBrowsing", Bandwidth: 300, Latency: 150, ServiceType: "WebBrowsing", TimeOfDay: "Morning"},
		{Name: "Morning Video", Bandwidth: 800, Latency: 15, ServiceType: "Video", TimeOfDay: "Morning"},
		{Name: "Evening Video", Bandwidth: 1200, Latency: 10, ServiceType: "Video", TimeOfDay: "Evening"},
	}

	nsDirectory := &NetworkSliceDirectory{Slices: networkSlices}

	serviceRequests := []struct {
		ServiceType string
		TimeOfDay   string
	}{
		{ServiceType: "Game", TimeOfDay: "Morning"},
		{ServiceType: "Video", TimeOfDay: "Evening"},
		{ServiceType: "Game", TimeOfDay: "Evening"},
		{ServiceType: "Video", TimeOfDay: "Morning"},
		{ServiceType: "WebBrowsing", TimeOfDay: "Morning"},
		{ServiceType: "Video", TimeOfDay: "Morning"},
	}

	for _, request := range serviceRequests {
		selectedSlice := nsDirectory.SelectNetworkSlice(request.ServiceType, request.TimeOfDay)
		efficiency := CalculateEfficiency(selectedSlice, request.ServiceType)

		resultChan <- &ServiceResult{
			ServiceType: request.ServiceType,
			TimeOfDay:   request.TimeOfDay,
			Efficiency:  efficiency,
		}

		// Simulate time passing based on the timeMultiplier.
		time.Sleep(time.Duration(timeMultiplier) * time.Second)
	}

	wg.Done()
}

func main() {
	timeMultiplier := 40

	resultChan := make(chan *ServiceResult)

	var wg sync.WaitGroup

	wg.Add(1)
	go runSimulation(timeMultiplier, resultChan, &wg)

	go func() {
		for result := range resultChan {
			fmt.Printf("Service: %s, Time: %s, Efficiency: %.2f%%\n", result.ServiceType, result.TimeOfDay, result.Efficiency)
		}
	}()

	// Wait for the simulation to finish.
	wg.Wait()

	// Close the result channel.
	close(resultChan)
}
