// Stefan Nilsson 2013-03-17

// This program runs three independent servers at
//
//	http://localhost:8080, http://localhost:8081, and http://localhost:8082.
//
// Each server accepts HTTP Get requests and answers with an integer
// indicating the current (fake) temperature at KTH in Stockholm.
//
// The servers are unreliable. The response time varies from a few seconds
// up to an hour. When the service is down, which happens quite frequently,
// you get HTTP Error 503 - Service unavailable.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func init() {
	log.SetPrefix("Weather: ")
}

// Server

func main() {
	Station = NewWeatherStation("KTH")
	for _, port := range []string{":8080", ":8081", ":8082"} {
		s := &http.Server{
			Addr:    port,
			Handler: http.HandlerFunc(ServeTemperature),
		}
		go func() {
			log.Printf("starting server at localhost%s", s.Addr)
			err := s.ListenAndServe()
			if err != nil {
				log.Fatalf("ListenAndServe: %v", err)
			}
		}()
	}
	select {} // Block forever.
}

func ServeTemperature(w http.ResponseWriter, r *http.Request) {
	switch x := rand.Float32(); true {
	case x < 0.70:
		time.Sleep(time.Duration(rand.Intn(1200)) * time.Millisecond)
		fmt.Fprintln(w, Station.CurrentTemp())
	case x < 0.85:
		time.Sleep(time.Hour)
	default:
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}
}

// Weather station

var Station *WeatherStation

type WeatherStation struct {
	mutex sync.RWMutex
	name  string
	temp  int
}

func NewWeatherStation(name string) *WeatherStation {
	station := &WeatherStation{name: name}
	station.TakeMeasurement()
	go func() {
		for {
			time.Sleep(time.Minute)
			station.TakeMeasurement()
		}
	}()
	return station
}

func (w *WeatherStation) CurrentTemp() int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	return w.temp
}

func (w *WeatherStation) TakeMeasurement() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.temp = fake.currentTemp() // Fake it for now.
	log.Printf("%d˚C at %s\n", w.temp, w.name)
}

// Fake temperature model for Stockholm

var fake fakeTempModel

type fakeTempModel struct {
	temp           int
	monthlyAverage int
	firstTemp      sync.Once
}

func (f *fakeTempModel) currentTemp() int {
	f.firstTemp.Do(func() {
		now := time.Now()
		Stockholm := []int{-3, -3, 0, 5, 11, 16, 17, 16, 12, 8, 3, -1}
		f.monthlyAverage = Stockholm[now.Month()-1]
		rand.Seed(now.UnixNano())
		f.temp = f.monthlyAverage - 4 + rand.Intn(8)
	})
	f.temp += int(rand.NormFloat64() - 0.02*float64(f.temp-f.monthlyAverage))
	return f.temp
}
