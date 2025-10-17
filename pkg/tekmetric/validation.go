package tekmetric

import (
	"fmt"
	"strings"
)

// validateSortParams validates sort field and direction parameters
func validateSortParams(sort, sortDirection string, validSorts []string) error {
	// Validate sort direction
	if sortDirection != "" {
		upper := strings.ToUpper(sortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", sortDirection)
		}
	}

	// Validate sort field
	if sort != "" {
		valid := false
		for _, validSort := range validSorts {
			if sort == validSort {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid sort field '%s'", sort)
		}
	}

	return nil
}

// Validate validates the RepairOrderQueryParams
func (p *RepairOrderQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Validate sort field - based on Tekmetric API documentation
	if p.Sort != "" {
		validSorts := map[string]bool{
			"createdDate":        true,
			"repairOrderNumber":  true,
			"customer.firstName": true,
			"customer.lastName":  true,
		}
		if !validSorts[p.Sort] {
			return fmt.Errorf("invalid sort field '%s': supported fields are createdDate, repairOrderNumber, customer.firstName, customer.lastName", p.Sort)
		}
	}

	// Validate repair order status IDs
	for _, statusID := range p.RepairOrderStatusIds {
		if statusID < 1 || statusID > 7 {
			return fmt.Errorf("invalid repairOrderStatusId '%d': must be 1-7 (1=Estimate, 2=WIP, 3=Complete, 4=Saved, 5=Posted, 6=AR, 7=Deleted)", statusID)
		}
	}

	return nil
}

// Validate validates the CustomerQueryParams
func (p *CustomerQueryParams) Validate() error {
	// Validate customer type ID
	if p.CustomerTypeID != 0 && p.CustomerTypeID != 1 && p.CustomerTypeID != 2 {
		return fmt.Errorf("invalid customerTypeId '%d': must be 1 (Customer) or 2 (Business)", p.CustomerTypeID)
	}

	// Validate sort - can be comma-separated list
	if p.Sort != "" {
		sortFields := strings.Split(p.Sort, ",")
		validSorts := map[string]bool{
			"lastName":  true,
			"firstName": true,
			"email":     true,
		}
		for _, field := range sortFields {
			trimmed := strings.TrimSpace(field)
			if !validSorts[trimmed] {
				return fmt.Errorf("invalid sort field '%s': supported fields are lastName, firstName, email", trimmed)
			}
		}
	}

	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	return nil
}

// Validate validates the VehicleQueryParams
func (p *VehicleQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Note: API documentation doesn't specify allowed sort fields for vehicles
	// So we don't validate the Sort field - let the API reject invalid values

	return nil
}

// Validate validates the AppointmentQueryParams
func (p *AppointmentQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Note: API documentation doesn't specify allowed sort fields for appointments
	// So we don't validate the Sort field - let the API reject invalid values

	return nil
}

// Validate validates the JobQueryParams
func (p *JobQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Validate sort field - based on Tekmetric API documentation
	if p.Sort != "" && p.Sort != "authorizedDate" {
		return fmt.Errorf("invalid sort field '%s': only 'authorizedDate' is supported", p.Sort)
	}

	// Validate repair order status IDs (jobs don't support status 7 - Deleted)
	for _, statusID := range p.RepairOrderStatusIds {
		if statusID < 1 || statusID > 6 {
			return fmt.Errorf("invalid repairOrderStatusId '%d': must be 1-6 (1=Estimate, 2=WIP, 3=Complete, 4=Saved, 5=Posted, 6=AR)", statusID)
		}
	}

	return nil
}

// Validate validates the EmployeeQueryParams
func (p *EmployeeQueryParams) Validate() error {
	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Note: API documentation doesn't specify allowed sort fields for employees
	// So we don't validate the Sort field - let the API reject invalid values

	return nil
}

// Validate validates the InventoryQueryParams
func (p *InventoryQueryParams) Validate() error {
	// Validate required fields
	if p.Shop == 0 {
		return fmt.Errorf("shop is required for inventory queries")
	}
	if p.PartTypeID == 0 {
		return fmt.Errorf("partTypeId is required for inventory queries")
	}

	// Validate part type ID
	if p.PartTypeID != 1 && p.PartTypeID != 2 && p.PartTypeID != 5 {
		return fmt.Errorf("invalid partTypeId '%d': must be 1 (Part), 2 (Tire), or 5 (Battery)", p.PartTypeID)
	}

	// Validate sort direction
	if p.SortDirection != "" {
		upper := strings.ToUpper(p.SortDirection)
		if upper != "ASC" && upper != "DESC" {
			return fmt.Errorf("invalid sort direction '%s': must be ASC or DESC", p.SortDirection)
		}
		p.SortDirection = upper // Normalize
	}

	// Validate sort fields - can be comma-separated
	if p.Sort != "" {
		sortFields := strings.Split(p.Sort, ",")
		validSorts := map[string]bool{
			"id":         true,
			"name":       true,
			"brand":      true,
			"partNumber": true,
		}
		for _, field := range sortFields {
			trimmed := strings.TrimSpace(field)
			if !validSorts[trimmed] {
				return fmt.Errorf("invalid sort field '%s': supported fields are id, name, brand, partNumber", trimmed)
			}
		}
	}

	return nil
}
