package domain

type CityEntity struct {
	Id          int16  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	CountryCode string `db:"countrycode" json:"country_code"`
	District    string `db:"district" json:"district"`
	Population  int32  `db:"population" json:"population"`
}
