package tools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// filterFunc is a function that determines if an item matches a query
type filterFunc func(item map[string]interface{}, query string) bool

// genericFilter filters items based on a search query using a custom filter function
func genericFilter(items interface{}, query string, filterFn filterFunc) []map[string]interface{} {
	// Marshal to JSON first
	jsonData, err := json.Marshal(items)
	if err != nil {
		return nil
	}

	// Unmarshal to []map[string]interface{} for filtering
	var itemsList []map[string]interface{}
	if err := json.Unmarshal(jsonData, &itemsList); err != nil {
		return nil
	}

	queryLower := strings.ToLower(query)
	var matches []map[string]interface{}

	for _, item := range itemsList {
		if filterFn(item, queryLower) {
			matches = append(matches, item)
		}
	}

	return matches
}

// customerFilterFunc returns true if the customer matches the query
func customerFilterFunc(item map[string]interface{}, queryLower string) bool {
	// Check name
	firstName, _ := item["firstName"].(string)
	lastName, _ := item["lastName"].(string)
	fullName := strings.ToLower(firstName + " " + lastName)
	if strings.Contains(fullName, queryLower) {
		return true
	}

	// Check email
	if email, ok := item["email"].(string); ok {
		if strings.Contains(strings.ToLower(email), queryLower) {
			return true
		}
	}

	// Check phone numbers
	if phones, ok := item["phone"].([]interface{}); ok {
		for _, p := range phones {
			if phone, ok := p.(map[string]interface{}); ok {
				if number, ok := phone["number"].(string); ok {
					// For phone numbers, check both with and without formatting
					if strings.Contains(number, queryLower) {
						return true
					}
				}
			}
		}
	}

	return false
}

// vehicleFilterFunc returns true if the vehicle matches the query
func vehicleFilterFunc(item map[string]interface{}, queryLower string) bool {
	// Check VIN
	if vin, ok := item["vin"].(string); ok {
		if strings.Contains(strings.ToLower(vin), queryLower) {
			return true
		}
	}

	// Check license plate
	if plate, ok := item["licensePlate"].(string); ok {
		if strings.Contains(strings.ToLower(plate), queryLower) {
			return true
		}
	}

	// Check make/model/year
	year, _ := item["year"].(float64)
	make, _ := item["make"].(string)
	model, _ := item["model"].(string)
	makeModel := strings.ToLower(fmt.Sprintf("%d %s %s", int(year), make, model))
	if strings.Contains(makeModel, queryLower) {
		return true
	}

	return false
}

// filterCustomers filters customers based on search query
func filterCustomers(customers interface{}, query string) []map[string]interface{} {
	return genericFilter(customers, query, customerFilterFunc)
}

// filterVehicles filters vehicles based on search query
func filterVehicles(vehicles interface{}, query string) []map[string]interface{} {
	return genericFilter(vehicles, query, vehicleFilterFunc)
}
