package models

var (
	// hot enough for the beach
	beachVacationThreshold float64 = 22

	// cold enough to snow
	skiVacationThreshold float64 = -2
)

type city struct {
	Name  string
	TempC float64
}

type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64

	BeachVacationReady() bool
	SkiVacationReady() bool
}

//func (c city) SkiVacationReady() bool {
//
//}
//
//func (c city) BeachVacationReady() bool {
//
//}
//
//func NewCity(n string, t float64, hasBeach bool, hasMountain bool) CityTemp {
//	return &city{
//		Name:  n,
//		TempC: t,
//	}
//}

//// NewCity creates a new City instance with
//// given name
//func NewCity(n string) *City {
//	return &City{
//		Name: n,
//	}
//}
//
//// Temf returns the city temperature converted
//// to Fahrenheit
//
//func (c City) TempF() float64 {
//	return (c.TempC * 9 / 5) + 32
//}
