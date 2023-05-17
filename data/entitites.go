package data

type CelestialBody struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Mass               float32  `json:"mass"`
	Density            float32  `json:"density"`
	Diameter           float32  `json:"diameter"`
	Gravity            float32  `json:"gravity"`
	DayInEarthHours    float32    `json:"day_in_earth_hours"`
	YearInEarthDays    float32    `json:"year_in_earth_days"`
	Moons              int    `json:"moons"`
	AverageTemperature int    `json:"average_temperature"`
}
