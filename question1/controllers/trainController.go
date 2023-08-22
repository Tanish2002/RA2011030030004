package controllers

import (
	"sort"
	"time"
	"train-service/services"
)

type Controllers struct {
	Service services.TrainAuthResponse
}

func (c *Controllers) GetTrains() []services.TrainResponse {
	trains := c.Service.GetTrains()

	// Get the current time
	currentTime := time.Now()

	// Filter out trains departing in the next 30 minutes
	var filteredTrains []services.TrainResponse
	for _, train := range *trains {
		departureTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), train.DepartureTime.Hours, train.DepartureTime.Minutes, train.DepartureTime.Seconds, 0, time.UTC)
		if departureTime.Sub(currentTime) > 30*time.Minute {
			filteredTrains = append(filteredTrains, train)
		}
	}

	// Sort filteredTrains based on the given criteria
	sort.Slice(filteredTrains, func(i, j int) bool {
		// Sort by ascending price
		if filteredTrains[i].Price.Sleeper != filteredTrains[j].Price.Sleeper {
			return filteredTrains[i].Price.Sleeper < filteredTrains[j].Price.Sleeper
		}
		// If prices are equal, sort by descending tickets available
		if filteredTrains[i].SeatsAvailable.Sleeper != filteredTrains[j].SeatsAvailable.Sleeper {
			return filteredTrains[i].SeatsAvailable.Sleeper > filteredTrains[j].SeatsAvailable.Sleeper
		}
		// If tickets available are equal, sort by descending departure time
		iDeparture := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), filteredTrains[i].DepartureTime.Hours, filteredTrains[i].DepartureTime.Minutes, filteredTrains[i].DepartureTime.Seconds, 0, time.UTC).Add(time.Duration(filteredTrains[i].DelayedBy) * time.Minute)
		jDeparture := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), filteredTrains[j].DepartureTime.Hours, filteredTrains[j].DepartureTime.Minutes, filteredTrains[j].DepartureTime.Seconds, 0, time.UTC).Add(time.Duration(filteredTrains[j].DelayedBy) * time.Minute)
		return iDeparture.After(jDeparture)
	})

	// Now return the filtered trains
	return filteredTrains
}

func (c *Controllers) GetTrain(id int) *services.TrainResponse {
	train := c.Service.GetTrain(id)
	return train
}
