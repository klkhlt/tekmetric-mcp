package tekmetric

import (
	"encoding/json"
	"time"
)

// Currency represents a monetary value in cents that outputs as dollars
type Currency int

// MarshalJSON formats Currency as dollars (dividing cents by 100)
func (c Currency) MarshalJSON() ([]byte, error) {
	dollars := float64(c) / 100.0
	return json.Marshal(dollars)
}

// UnmarshalJSON parses Currency from cents
func (c *Currency) UnmarshalJSON(data []byte) error {
	var cents int
	if err := json.Unmarshal(data, &cents); err != nil {
		return err
	}
	*c = Currency(cents)
	return nil
}

// TokenResponse represents the OAuth token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"` // Token lifetime in seconds
	Scope       string `json:"scope"`      // Space-separated shop IDs
}

// Shop represents a Tekmetric shop
type Shop struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Nickname             string  `json:"nickname"`
	Phone                string  `json:"phone"`
	Email                string  `json:"email"`
	Website              string  `json:"website"`
	Address              Address `json:"address"`
	ROCustomLabelEnabled bool    `json:"roCustomLabelEnabled"`
}

// Address represents a physical address
type Address struct {
	ID            int    `json:"id"`
	Address1      string `json:"address1"`
	Address2      string `json:"address2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	StreetAddress string `json:"streetAddress"`
	FullAddress   string `json:"fullAddress"`
}

// Phone represents a phone number
type Phone struct {
	ID      int    `json:"id,omitempty"`
	Number  string `json:"number"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

// CustomerType represents the type of customer
type CustomerType struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// Customer represents a Tekmetric customer
type Customer struct {
	ID                            int           `json:"id"`
	FirstName                     string        `json:"firstName"`
	LastName                      string        `json:"lastName"`
	Email                         string        `json:"email"`
	Phone                         []Phone       `json:"phone"`
	CustomerType                  *CustomerType `json:"customerType,omitempty"`
	ContactFirstName              *string       `json:"contactFirstName"`
	ContactLastName               *string       `json:"contactLastName"`
	Address                       *Address      `json:"address"`
	ShopID                        int           `json:"shopId"`
	EligibleForAccountsReceivable bool          `json:"eligibleForAccountsReceivable"`
	CreditLimit                   float64       `json:"creditLimit"`
	OkForMarketing                bool          `json:"okForMarketing"`
	Notes                         string        `json:"notes,omitempty"`
	CreatedDate                   time.Time     `json:"createdDate"`
	UpdatedDate                   time.Time     `json:"updatedDate"`
	DeletedDate                   *time.Time    `json:"deletedDate"`
}

// Vehicle represents a vehicle
type Vehicle struct {
	ID             int        `json:"id"`
	CustomerID     int        `json:"customerId"`
	ShopID         int        `json:"shopId"`
	Year           int        `json:"year"`
	Make           string     `json:"make"`
	Model          string     `json:"model"`
	SubModel       string     `json:"subModel,omitempty"`
	VIN            string     `json:"vin"`
	LicensePlate   string     `json:"licensePlate,omitempty"`
	Color          string     `json:"color,omitempty"`
	UnitNumber     string     `json:"unitNumber,omitempty"`
	ProductionDate *string    `json:"productionDate,omitempty"`
	Mileage        float64    `json:"mileage"`
	Engine         string     `json:"engine,omitempty"`
	Transmission   string     `json:"transmission,omitempty"`
	DriveType      string     `json:"driveType,omitempty"`
	Notes          string     `json:"notes,omitempty"`
	CreatedDate    time.Time  `json:"createdDate"`
	UpdatedDate    time.Time  `json:"updatedDate"`
	DeletedDate    *time.Time `json:"deletedDate"`
}

// RepairOrderStatus represents the status of a repair order
type RepairOrderStatus struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// RepairOrderLabel represents a label for a repair order
type RepairOrderLabel struct {
	ID     int                `json:"id"`
	Code   string             `json:"code"`
	Name   string             `json:"name"`
	Status *RepairOrderStatus `json:"status,omitempty"`
}

// RepairOrderCustomLabel represents a custom label
type RepairOrderCustomLabel struct {
	Name string `json:"name"`
}

// RepairOrder represents a repair order
type RepairOrder struct {
	ID                     int                     `json:"id"`
	RepairOrderNumber      int                     `json:"repairOrderNumber"`
	ShopID                 int                     `json:"shopId"`
	RepairOrderStatus      RepairOrderStatus       `json:"repairOrderStatus"`
	RepairOrderLabel       *RepairOrderLabel       `json:"repairOrderLabel,omitempty"`
	RepairOrderCustomLabel *RepairOrderCustomLabel `json:"repairOrderCustomLabel,omitempty"`
	Color                  string                  `json:"color,omitempty"`
	AppointmentStartTime   *time.Time              `json:"appointmentStartTime,omitempty"`
	CustomerID             int                     `json:"customerId"`
	TechnicianID           *int                    `json:"technicianId"`
	ServiceWriterID        *int                    `json:"serviceWriterId"`
	VehicleID              int                     `json:"vehicleId"`
	MilesIn                *float64                `json:"milesIn"`
	MilesOut               *float64                `json:"milesOut"`
	Keytag                 *string                 `json:"keytag"`
	CompletedDate          *time.Time              `json:"completedDate"`
	PostedDate             *time.Time              `json:"postedDate"`
	LaborSales             Currency                `json:"laborSales"`
	PartsSales             Currency                `json:"partsSales"`
	SubletSales            Currency                `json:"subletSales"`
	DiscountTotal          Currency                `json:"discountTotal"`
	FeeTotal               Currency                `json:"feeTotal"`
	Taxes                  Currency                `json:"taxes"`
	AmountPaid             Currency                `json:"amountPaid"`
	TotalSales             Currency                `json:"totalSales"`
	Jobs                   []Job                   `json:"jobs,omitempty"`
	Sublets                []Sublet                `json:"sublets,omitempty"`
	Fees                   []Fee                   `json:"fees,omitempty"`
	Discounts              []Discount              `json:"discounts,omitempty"`
	CustomerConcerns       []CustomerConcern       `json:"customerConcerns,omitempty"`
	CreatedDate            time.Time               `json:"createdDate"`
	UpdatedDate            time.Time               `json:"updatedDate"`
	DeletedDate            *time.Time              `json:"deletedDate"`
}

// PartType represents the type of a part
type PartType struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// Part represents a vehicle part
type Part struct {
	ID               int       `json:"id"`
	Quantity         float64   `json:"quantity"`
	Brand            string    `json:"brand,omitempty"`
	Name             string    `json:"name,omitempty"`
	PartNumber       string    `json:"partNumber,omitempty"`
	Description      string    `json:"description,omitempty"`
	Cost             Currency  `json:"cost"`
	Retail           Currency  `json:"retail"`
	Model            *string   `json:"model,omitempty"`
	Width            *string   `json:"width,omitempty"`
	Ratio            *float64  `json:"ratio,omitempty"`
	Diameter         *float64  `json:"diameter,omitempty"`
	ConstructionType *string   `json:"constructionType,omitempty"`
	LoadIndex        *string   `json:"loadIndex,omitempty"`
	SpeedRating      *string   `json:"speedRating,omitempty"`
	PartType         *PartType `json:"partType,omitempty"`
	DOTNumbers       []string  `json:"dotNumbers,omitempty"`
}

// Labor represents labor on a job
type Labor struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Rate     Currency `json:"rate"`
	Hours    float64  `json:"hours"`
	Complete bool     `json:"complete"`
}

// Fee represents a fee
type Fee struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Total Currency `json:"total"`
}

// Discount represents a discount
type Discount struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Total Currency `json:"total"`
}

// Job represents a job on a repair order
type Job struct {
	ID              int        `json:"id"`
	RepairOrderID   int        `json:"repairOrderId"`
	VehicleID       int        `json:"vehicleId"`
	CustomerID      int        `json:"customerId"`
	Name            string     `json:"name"`
	Authorized      bool       `json:"authorized"`
	AuthorizedDate  *string    `json:"authorizedDate,omitempty"`
	Selected        bool       `json:"selected"`
	TechnicianID    *int       `json:"technicianId"`
	Note            string     `json:"note,omitempty"`
	JobCategoryName string     `json:"jobCategoryName,omitempty"`
	PartsTotal      Currency   `json:"partsTotal"`
	LaborTotal      Currency   `json:"laborTotal"`
	DiscountTotal   Currency   `json:"discountTotal"`
	FeeTotal        Currency   `json:"feeTotal"`
	Subtotal        Currency   `json:"subtotal"`
	Archived        bool       `json:"archived"`
	CreatedDate     time.Time  `json:"createdDate"`
	UpdatedDate     time.Time  `json:"updatedDate"`
	CompletedDate   *time.Time `json:"completedDate,omitempty"`
	Labor           []Labor    `json:"labor,omitempty"`
	Parts           []Part     `json:"parts,omitempty"`
	Fees            []Fee      `json:"fees,omitempty"`
	Discounts       []Discount `json:"discounts,omitempty"`
	LaborHours      float64    `json:"laborHours"`
	LoggedHours     float64    `json:"loggedHours"`
	Sort            int        `json:"sort,omitempty"`
}

// Vendor represents a vendor/supplier
type Vendor struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Nickname string  `json:"nickname,omitempty"`
	Website  *string `json:"website"`
	Phone    *string `json:"phone"`
}

// SubletItem represents an item in a sublet
type SubletItem struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Cost     Currency `json:"cost"`
	Price    Currency `json:"price"`
	Complete bool     `json:"complete"`
}

// Sublet represents subcontracted work
type Sublet struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Vendor         *Vendor      `json:"vendor,omitempty"`
	Authorized     *bool        `json:"authorized"`
	AuthorizedDate *string      `json:"authorizedDate"`
	Selected       bool         `json:"selected"`
	Note           *string      `json:"note"`
	Items          []SubletItem `json:"items,omitempty"`
	Price          Currency     `json:"price"`
	Cost           Currency     `json:"cost"`
}

// CustomerConcern represents a customer's concern
type CustomerConcern struct {
	ID          int     `json:"id"`
	Concern     string  `json:"concern"`
	TechComment *string `json:"techComment"`
}

// Appointment represents an appointment
type Appointment struct {
	ID               int        `json:"id"`
	ShopID           int        `json:"shopId"`
	CustomerID       int        `json:"customerId"`
	VehicleID        int        `json:"vehicleId"`
	ServiceWriterID  *int       `json:"serviceWriterId"`
	TechnicianID     *int       `json:"technicianId"`
	StartTime        time.Time  `json:"startTime"`
	EndTime          time.Time  `json:"endTime"`
	Status           string     `json:"status"`
	CustomerConcerns string     `json:"customerConcerns,omitempty"`
	Notes            string     `json:"notes,omitempty"`
	CreatedDate      time.Time  `json:"createdDate"`
	UpdatedDate      time.Time  `json:"updatedDate"`
	DeletedDate      *time.Time `json:"deletedDate"`
}

// EnrichedAppointment represents an appointment with customer and vehicle details
type EnrichedAppointment struct {
	Appointment
	Customer *Customer `json:"customer,omitempty"`
	Vehicle  *Vehicle  `json:"vehicle,omitempty"`
}

// Employee represents an employee
type Employee struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone,omitempty"`
	Role        string     `json:"role"`
	Active      bool       `json:"active"`
	ShopID      int        `json:"shopId"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate"`
}

// InventoryPart represents an inventory part
type InventoryPart struct {
	ID          int        `json:"id"`
	ShopID      int        `json:"shopId"`
	PartNumber  string     `json:"partNumber"`
	Description string     `json:"description"`
	Brand       string     `json:"brand,omitempty"`
	Cost        Currency   `json:"cost"`
	Retail      Currency   `json:"retail"`
	Quantity    float64    `json:"quantity"`
	Location    string     `json:"location,omitempty"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse[T any] struct {
	Content          []T  `json:"content"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	Last             bool `json:"last"`
	First            bool `json:"first"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	NumberOfElements int  `json:"numberOfElements"`
	Empty            bool `json:"empty"`
}

// APIResponse represents a standard API response with data
type APIResponse[T any] struct {
	Type    string                 `json:"type"`
	Message string                 `json:"message"`
	Data    T                      `json:"data"`
	Details map[string]interface{} `json:"details"`
}

// CannedJob represents a predefined job template
type CannedJob struct {
	ID           int       `json:"id"`
	ShopID       int       `json:"shopId"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	CategoryName string    `json:"categoryName,omitempty"`
	LaborRate    int       `json:"laborRate"`
	LaborHours   float64   `json:"laborHours"`
	CreatedDate  time.Time `json:"createdDate"`
	UpdatedDate  time.Time `json:"updatedDate"`
}
