package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "sync"
    "time"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
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

func updateEfficiencyData(resultChan chan<- *ServiceResult, timeMultiplier int, ctx context.Context, stopUpdating <-chan struct{}, ctx) {
    ticker := time.NewTicker(time.Duration(timeMultiplier) * time.Second)

    for {
        select {
        case <-ctx.Done():
            // The context was canceled, indicating the application is shutting down.
            return
        case <-ticker.C:
            // Fetch the latest efficiency data from an external source or calculate it.
            newEfficiencyData := []*ServiceResult{
                {ServiceType: "Game", TimeOfDay: "Morning", Efficiency: 85.5},
                {ServiceType: "Video", TimeOfDay: "Evening", Efficiency: 92},
                {ServiceType: "Game", TimeOfDay: "Evening", Efficiency: 88.3},
            }

            // Send the new data to the result channel.
            for _, item := range newEfficiencyData {
                resultChan <- item
            }

            // Log the new data.
            for _, item := range newEfficiencyData {
                log.Printf("New Data: ServiceType=%s, TimeOfDay=%s, Efficiency=%.2f", item.ServiceType, item.TimeOfDay, item.Efficiency)
            }
        case <-stopUpdating:
            // Received a signal to stop updating.
            return
        }
    }
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

// Define your APIHandler to return efficiency data
func APIHandler(w http.ResponseWriter, r *http.Request) {
    efficiencyData := []ServiceResult{
        {ServiceType: "Game", TimeOfDay: "Morning", Efficiency: 85.5},
        {ServiceType: "Video", TimeOfDay: "Evening", Efficiency: 92},
        {ServiceType: "Game", TimeOfDay: "Evening", Efficiency: 88.3},
    }

    // Log the new data
    for _, item := range efficiencyData {
        log.Printf("New Data: ServiceType=%s, TimeOfDay=%s, Efficiency=%.2f", item.ServiceType, item.TimeOfDay, item.Efficiency)
    }

    // Encode data as JSON and send it in the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(efficiencyData)
}


func runSimulation(timeMultiplier int, resultChan chan<- *ServiceResult, wg *sync.WaitGroup) {
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

        time.Sleep(time.Duration(timeMultiplier) * time.Second)
    }

    wg.Done()
}

func main() {
    timeMultiplier := 90
    resultChan := make(chan *ServiceResult)

    // Add a Goroutine to continuously update efficiency data
    go func() {
        updateEfficiencyData(resultChan, timeMultiplier, ctx, stopUpdating)
    }()
    // Create a new Gorilla Mux router
    r := mux.NewRouter()

    // Define API endpoints
    r.HandleFunc("/api/efficiency", func(w http.ResponseWriter, r *http.Request) {
        // Read from the resultChan and send the latest data
        select {
        case result := <-resultChan:
            // Encode data as JSON and send it in the response
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(result)
        default:
            // Handle no data available
            w.WriteHeader(http.StatusNoContent)
        }
    }).Methods("GET")

    // Serve the React frontend or any other static files
    fs := http.FileServer(http.Dir("static"))
    r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
    })
    handler := c.Handler(r)

    // Start the HTTP server with CORS configuration
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal("Error starting the server: ", err)
    }
}