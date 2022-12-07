package structures

type Data struct {
	RestaurantID      string
	RestaurantName    string
	CountryCode       string
	City              string
	Address           string
	Locality          string
	Verbose           string
	Longitude         string
	Latitude          string
	Cuisines          string
	AverageCost       string
	Currency          string
	HasTableBooking   string
	HasOnlineDelivery string
	IsDeliveringNow   string
	SwitchToOrderMenu string
	PriceRange        string
	AggregateRating   string
	RatingColor       string
	RatingText        string
	Votes             string
}

type TwoFieldData struct {
	Sr_No          int
	RestaurantName string
	Address        string
}

type FourFieldData struct {
	Sr_No          int
	Cuisine        string
	Currency       string
	RestaurantName string
	Address        string
	City           string
}

type Choices struct {
	Cuisine         string `json:"cuisine"`
	City            string `json:"city"`
	HasTableBooking string `json:"hasTableBooking"`
}
