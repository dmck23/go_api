package domain

type CityDb interface {
	GetCityById(id string) (*CityEntity, error)
	GetCities(gt string, lt string, countryCode string, district string) (*[]CityEntity, error)
}
