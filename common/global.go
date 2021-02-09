package common

//Struct API
// Order struct (Model) ...

type Message  struct {
	Code      int  	 `json:"code"`
	Remark    string `json:"remark"`
	OrderID   string `json:"orderID"`	
	Orders   *Orders `json:"orders,omitempty"`
	Result   *Result  `json:"result,omitempty"`
}

type Orders struct {
	OrderID      string        `json:"orderID"`
	CustomerID   string        `json:"customerID"`
	EmployeeID   string        `json:"employeeID"`
	OrderDate    string        `json:"orderDate"`
	OrdersDet []OrdersDetail   `json:"ordersDetail"`
	
}

type OrdersDetail struct {
	OrderID      string  `json:"orderID"`
	ProductID  	 string  `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	UnitPrice    float64 `json:"UnitPrice"`
	Quantity     int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

// Add Customer and Product Struct
type Customer struct {
	CustomerID   string `json:"customerID"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactName"`
	ContactTitle string `json:"contactTitle"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Region			 string `json:"region"`
	PostalCode   string `json:"postalCode"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Fax					 string `json:"fax"`
}

type Product struct {
	ProductID				int			`json:"productID"`
	ProductName			string	`json:"productName"`
	SupplierID			int			`json:"supplierID"`
	CategoryID			int			`json:"categoryID"`
	QuantityPerUnit	string	`json:"quantityPerUnit"`
	UnitPrice				float32	`json:"unitPrice"`
	UnitsInStock			int		`json:"unitsInStock"`
	UnitsOnOrder			int		`json:"unitsOnOrder"`
	ReorderLevel		int			`json:"reorderLevel"`
	Discontinued		int8		`json:"discontinued"`
	Description			string	`json:"description"`
}

type FastPayRequest struct {
	Request    string `json:"request"`
	Merchant_ID string `json:"merchant_id"`
	Merchant   string `json:"merchant"`
	Signature  string `json:"signature"`
}

type FastPayResponse struct {
	Response     	string `json:"response"`
	Merchant_ID   string `json:"merchant_id"`
	Merchant      string `json:"merchant"`
	PaymentChan []PaymentChannel   `json:"payment_channel"`	
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
}

type PaymentChannel struct {
	PgCode    string `json:"pg_code"`
	PgName   	string `json:"pg_name"`
}

// Trip Struct Request Response DataDetail
type MyTrips struct {
	DepatureDate1 string `json:"depature_date_1"`
	DepatureDate2 string `json:"depature_date_2"`
	Provinsi      int64  `json:"provinsi"`
}

type MytripsResponse struct {
	Message string 		  `json:"message"`
	Status  string 		  `json:"status"`
	TripDetail 	[]TripDetail  `json:"data"`	
}

type TripDetail struct {
	AirlineName      string `json:"AirlineName,omitempty"`
	AirportName      string `json:"AirportName,omitempty"`
	CityName         string `json:"CityName,omitempty"`
	Currency         string `json:"Currency,omitempty"`
	DepartureDate    string `json:"DepartureDate,omitempty"`
	Description      string `json:"Description,omitempty"`
	Destination      string `json:"Destination,omitempty"`
	DetailTransit    string `json:"DetailTransit,omitempty"`
	DoubleType       string `json:"DoubleType,omitempty"`
	Duration         string `json:"Duration,omitempty"`
	Goods            string `json:"Goods,omitempty"`
	HotelName        string `json:"HotelName,omitempty"`
	HotelRating      string `json:"HotelRating,omitempty"`
	Lat              string `json:"Lat,omitempty"`
	LicenseNumber    string `json:"LicenseNumber,omitempty"`
	Logo             string `json:"Logo,omitempty"`
	Long             string `json:"Long,omitempty"`
	Origin           string `json:"Origin,omitempty"`
	OriginCity       string `json:"OriginCity,omitempty"`
	Price            string `json:"Price,omitempty"`
	PromoCode        string `json:"PromoCode,omitempty"`
	PromoDescription string `json:"PromoDescription,omitempty"`
	Provinsi         string `json:"Provinsi,omitempty"`
	QuadType         string `json:"QuadType,omitempty"`
	Rating           string `json:"Rating,omitempty"`
	ReturnDate       string `json:"ReturnDate,omitempty"`
	TermCondition    string `json:"TermCondition,omitempty"`
	Transit          string `json:"Transit,omitempty"`
	TravelID         string `json:"TravelID,omitempty"`
	TravelName       string `json:"TravelName,omitempty"`
	TripID           string `json:"TripID,omitempty"`
	TripleType       string `json:"TripleType,omitempty"`
}

//End Struct API
