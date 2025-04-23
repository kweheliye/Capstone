package main

import (
	"errors"
	"fmt"
	"sync"
)

// Error definitions for truck management operations
var (
	ErrTruckNotFound = errors.New("truck not found")
	ErrTruckExist    = errors.New("truck already exists")
	ErrInvalidCargo  = errors.New("invalid cargo value")
	ErrEmptyID       = errors.New("truck ID cannot be empty")
)

// FleetManager defines the interface for managing a fleet of trucks
type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

// Truck represents a truck with an ID and cargo capacity
type Truck struct {
	ID    string
	Cargo int
}

// truckManager implements the FleetManager interface
type truckManager struct {
	trucks map[string]*Truck
	sync.RWMutex
}

// NewTruckManager creates a new instance of FleetManager
func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

// AddTruck adds a new truck to the fleet with the specified ID and cargo capacity
func (tm *truckManager) AddTruck(id string, cargo int) error {
	tm.Lock()
	defer tm.Unlock()

	// Validate input parameters
	if id == "" {
		return ErrEmptyID
	}
	if cargo < 0 {
		return ErrInvalidCargo
	}

	// Check if truck already exists
	if _, exist := tm.trucks[id]; exist {
		return ErrTruckExist
	}

	// Add the new truck

	tm.trucks[id] = &Truck{
		ID:    id,
		Cargo: cargo,
	}

	return nil
}

// GetTruck retrieves a truck by its ID
func (tm *truckManager) GetTruck(id string) (Truck, error) {

	if id == "" {
		return Truck{}, ErrEmptyID
	}

	tm.RLock()
	defer tm.RUnlock()

	truck, exist := tm.trucks[id]
	if !exist {
		return Truck{}, ErrTruckNotFound
	}

	return *truck, nil
}

// UpdateTruckCargo updates the cargo capacity of a truck
func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {

	if id == "" {
		return ErrEmptyID
	}
	if cargo < 0 {
		return ErrInvalidCargo
	}

	tm.Lock()
	defer tm.Unlock()

	// Check if truck exists
	truck, exist := tm.trucks[id]
	if !exist {
		return ErrTruckNotFound
	}

	truck.Cargo = cargo
	return nil
}

// RemoveTruck removes a truck from the fleet
func (tm *truckManager) RemoveTruck(id string) error {
	tm.Lock()
	defer tm.Unlock()

	if id == "" {
		return ErrEmptyID
	}

	// Check if truck exists
	if _, exist := tm.trucks[id]; !exist {
		return ErrTruckNotFound
	}

	delete(tm.trucks, id)
	return nil
}

// Main function to demonstrate the usage of FleetManager
func main() {
	// Create a new truck manager
	manager := NewTruckManager()

	// Add some trucks
	err := manager.AddTruck("truck1", 1000)
	if err != nil {
		fmt.Printf("Error adding truck1: %v\n", err)
	}

	err = manager.AddTruck("truck2", 2000)
	if err != nil {
		fmt.Printf("Error adding truck2: %v\n", err)
	}

	// Try to add a truck with the same ID
	err = manager.AddTruck("truck1", 1500)
	if err != nil {
		fmt.Printf("Expected error when adding duplicate truck: %v\n", err)
	}

	// Get a truck
	truck, err := manager.GetTruck("truck1")
	if err != nil {
		fmt.Printf("Error getting truck1: %v\n", err)
	} else {
		fmt.Printf("Found truck: ID=%s, Cargo=%d\n", truck.ID, truck.Cargo)
	}

	// Update truck cargo
	err = manager.UpdateTruckCargo("truck1", 1500)
	if err != nil {
		fmt.Printf("Error updating truck1 cargo: %v\n", err)
	}

	// Get the updated truck
	truck, err = manager.GetTruck("truck1")
	if err != nil {
		fmt.Printf("Error getting updated truck1: %v\n", err)
	} else {
		fmt.Printf("Updated truck: ID=%s, Cargo=%d\n", truck.ID, truck.Cargo)
	}

	// Remove a truck
	err = manager.RemoveTruck("truck1")
	if err != nil {
		fmt.Printf("Error removing truck1: %v\n", err)
	}

	// Try to get the removed truck
	_, err = manager.GetTruck("truck1")
	if err != nil {
		fmt.Printf("Expected error when getting removed truck: %v\n", err)
	}
}
