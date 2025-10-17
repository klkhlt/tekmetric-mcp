package tekmetric

// RepairOrderQueryParams holds query parameters for repair order searches
type RepairOrderQueryParams struct {
	Shop                 int    `url:"shop,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Size                 int    `url:"size,omitempty"`
	Start                string `url:"start,omitempty"`            // Date format: YYYY-MM-DD
	End                  string `url:"end,omitempty"`              // Date format: YYYY-MM-DD
	PostedDateStart      string `url:"postedDateStart,omitempty"`  // Date format: YYYY-MM-DD
	PostedDateEnd        string `url:"postedDateEnd,omitempty"`    // Date format: YYYY-MM-DD
	UpdatedDateStart     string `url:"updatedDateStart,omitempty"` // Date format: YYYY-MM-DD
	UpdatedDateEnd       string `url:"updatedDateEnd,omitempty"`   // Date format: YYYY-MM-DD
	DeletedDateStart     string `url:"deletedDateStart,omitempty"` // Date format: YYYY-MM-DD
	DeletedDateEnd       string `url:"deletedDateEnd,omitempty"`   // Date format: YYYY-MM-DD
	RepairOrderNumber    int    `url:"repairOrderNumber,omitempty"`
	RepairOrderStatusIds []int  `url:"repairOrderStatusId,omitempty"` // 1-Estimate, 2-WIP, 3-Complete, 4-Saved, 5-Posted, 6-AR, 7-Deleted
	CustomerID           int    `url:"customerId,omitempty"`
	VehicleID            int    `url:"vehicleId,omitempty"`
	Search               string `url:"search,omitempty"`        // Search by RO#, customer name, vehicle info
	Sort                 string `url:"sort,omitempty"`          // createdDate, repairOrderNumber, customer.firstName, customer.lastName
	SortDirection        string `url:"sortDirection,omitempty"` // ASC, DESC
}

// CustomerQueryParams holds query parameters for customer searches
type CustomerQueryParams struct {
	Shop                          int    `url:"shop,omitempty"`
	Page                          int    `url:"page,omitempty"`
	Size                          int    `url:"size,omitempty"`
	Search                        string `url:"search,omitempty"`                        // Search by name, email, phone
	Email                         string `url:"email,omitempty"`                         // Filter by email
	Phone                         string `url:"phone,omitempty"`                         // Filter by phone
	EligibleForAccountsReceivable *bool  `url:"eligibleForAccountsReceivable,omitempty"` // Filter by AR eligibility
	OkForMarketing                *bool  `url:"okForMarketing,omitempty"`                // Filter by marketing permission
	UpdatedDateStart              string `url:"updatedDateStart,omitempty"`              // Filter by updated date
	UpdatedDateEnd                string `url:"updatedDateEnd,omitempty"`                // Filter by updated date
	DeletedDateStart              string `url:"deletedDateStart,omitempty"`              // Filter by deleted date
	DeletedDateEnd                string `url:"deletedDateEnd,omitempty"`                // Filter by deleted date
	CustomerTypeID                int    `url:"customerTypeId,omitempty"`                // 1=Customer, 2=Business
	Sort                          string `url:"sort,omitempty"`                          // lastName, firstName, email (can be comma-separated)
	SortDirection                 string `url:"sortDirection,omitempty"`                 // ASC, DESC
}

// VehicleQueryParams holds query parameters for vehicle searches
type VehicleQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	CustomerID       int    `url:"customerId,omitempty"`       // Filter by customer
	Search           string `url:"search,omitempty"`           // Search by year, make, model
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Filter by updated date
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Filter by updated date
	DeletedDateStart string `url:"deletedDateStart,omitempty"` // Filter by deleted date
	DeletedDateEnd   string `url:"deletedDateEnd,omitempty"`   // Filter by deleted date
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// AppointmentQueryParams holds query parameters for appointment searches
type AppointmentQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	CustomerID       int    `url:"customerId,omitempty"`       // Filter by customer
	VehicleID        int    `url:"vehicleId,omitempty"`        // Filter by vehicle
	Start            string `url:"start,omitempty"`            // Start date filter
	End              string `url:"end,omitempty"`              // End date filter
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Filter by updated date
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Filter by updated date
	IncludeDeleted   *bool  `url:"includeDeleted,omitempty"`   // Include deleted appointments (default: true)
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// JobQueryParams holds query parameters for job searches
type JobQueryParams struct {
	Shop                 int    `url:"shop,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Size                 int    `url:"size,omitempty"`
	VehicleID            int    `url:"vehicleId,omitempty"`           // Filter by vehicle ID
	RepairOrderID        int    `url:"repairOrderId,omitempty"`       // Filter by repair order
	CustomerID           int    `url:"customerId,omitempty"`          // Filter by customer ID
	Authorized           *bool  `url:"authorized,omitempty"`          // Filter by authorized jobs
	AuthorizedDateStart  string `url:"authorizedDateStart,omitempty"` // Filter by authorization date
	AuthorizedDateEnd    string `url:"authorizedDateEnd,omitempty"`   // Filter by authorization date
	UpdatedDateStart     string `url:"updatedDateStart,omitempty"`    // Filter by updated date
	UpdatedDateEnd       string `url:"updatedDateEnd,omitempty"`      // Filter by updated date
	RepairOrderStatusIds []int  `url:"repairOrderStatusId,omitempty"` // 1-6 (no Deleted status for jobs)
	Sort                 string `url:"sort,omitempty"`                // authorizedDate
	SortDirection        string `url:"sortDirection,omitempty"`       // ASC, DESC
}

// EmployeeQueryParams holds query parameters for employee searches
type EmployeeQueryParams struct {
	Shop             int    `url:"shop,omitempty"`
	Page             int    `url:"page,omitempty"`
	Size             int    `url:"size,omitempty"`
	Search           string `url:"search,omitempty"`           // Search by name
	UpdatedDateStart string `url:"updatedDateStart,omitempty"` // Filter by updated date
	UpdatedDateEnd   string `url:"updatedDateEnd,omitempty"`   // Filter by updated date
	Sort             string `url:"sort,omitempty"`             // Sort field (API docs don't specify allowed values)
	SortDirection    string `url:"sortDirection,omitempty"`    // ASC, DESC
}

// InventoryQueryParams holds query parameters for inventory searches
type InventoryQueryParams struct {
	Shop          int      `url:"shop"`       // Required: Shop ID
	PartTypeID    int      `url:"partTypeId"` // Required: 1=Part, 2=Tire, 5=Battery
	Page          int      `url:"page,omitempty"`
	Size          int      `url:"size,omitempty"`
	PartNumbers   []string `url:"partNumbers,omitempty"`   // Exact match on part numbers
	Width         string   `url:"width,omitempty"`         // Tire width (tires only)
	Ratio         float64  `url:"ratio,omitempty"`         // Tire ratio (tires only)
	Diameter      float64  `url:"diameter,omitempty"`      // Tire diameter (tires only)
	TireSize      string   `url:"tireSize,omitempty"`      // Tire size: width+ratio+diameter (tires only)
	Sort          string   `url:"sort,omitempty"`          // id, name, brand, partNumber (comma-separated)
	SortDirection string   `url:"sortDirection,omitempty"` // ASC, DESC
}
