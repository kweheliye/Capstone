# Fleet Management System

## Project Overview
The Fleet Management System is a Go-based application designed to manage a fleet of trucks. It provides a simple, thread-safe API for tracking trucks and their cargo capacities. This project demonstrates core Go programming concepts including interfaces, concurrency safety, error handling, and testing.

## Features
- **Add Trucks**: Register new trucks with unique IDs and cargo capacities
- **Retrieve Truck Information**: Look up truck details by ID
- **Update Cargo Capacity**: Modify the cargo capacity of existing trucks
- **Remove Trucks**: Delete trucks from the fleet
- **Thread-Safe Operations**: All operations are protected with read-write mutexes for concurrent access

## Code Structure
The project follows a clean, modular design:

- **Interface-Based Design**: The `FleetManager` interface defines the contract for truck management operations
- **Struct Implementation**: The `truckManager` struct implements the `FleetManager` interface
- **Error Handling**: Custom error types for different failure scenarios
- **Concurrency Safety**: Read-write mutex for safe concurrent access
- **Unit Tests**: Comprehensive test coverage including concurrency testing

## Technical Implementation

### Core Components
1. **FleetManager Interface**: Defines the API contract with four primary operations
2. **Truck Struct**: Represents a truck with an ID and cargo capacity
3. **truckManager Struct**: Implements the FleetManager interface with a thread-safe map of trucks

### Error Handling
The system defines custom errors for different scenarios:
- Empty truck ID
- Invalid cargo value (negative)
- Truck not found
- Duplicate truck ID

### Concurrency
The implementation uses `sync.RWMutex` to ensure thread safety:
- Read locks for retrieval operations
- Write locks for modification operations

## Usage Examples
```go
// Create a new truck manager
manager := NewTruckManager()

// Add a truck
err := manager.AddTruck("truck1", 1000)
if err != nil {
    // Handle error
}

// Get truck information
truck, err := manager.GetTruck("truck1")
if err != nil {
    // Handle error
} else {
    fmt.Printf("Truck ID: %s, Cargo: %d\n", truck.ID, truck.Cargo)
}

// Update truck cargo
err = manager.UpdateTruckCargo("truck1", 1500)
if err != nil {
    // Handle error
}

// Remove a truck
err = manager.RemoveTruck("truck1")
if err != nil {
    // Handle error
}
```

## Testing
The project includes comprehensive unit tests covering:
- Basic functionality for all operations
- Error handling for edge cases
- Concurrent access safety

Run the tests with:
```
go test -v
```

## Future Enhancements
Potential improvements for the system:
- Persistence layer for storing truck data
- Additional truck attributes (location, status, driver info)
- Fleet-wide statistics and reporting
- REST API for remote access
- Authentication and authorization

## Conclusion
This Fleet Management System demonstrates Go's strengths in building concurrent, type-safe applications with clean interfaces. The code showcases best practices in Go programming including proper error handling, interface-based design, and comprehensive testing.