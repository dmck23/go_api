package repositories

import (
	"fmt"
	"go_api/world/internal/cities/domain"
	"log"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type CityRepository struct {
	db *sqlx.DB
}

func NewCityRepository() domain.CityDb {

	user := viper.GetString("database.postgres.user")
	db_name := viper.GetString("database.postgres.db_name")
	password := viper.GetString("database.postgres.password")
	host := viper.GetString("database.postgres.host")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", user, db_name, password, host))
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return &CityRepository{db: db}
}

func (cr *CityRepository) GetCityById(id string) (*domain.CityEntity, error) {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "countrycode", "district", "population").From("city")
	sb.Where(sb.Equal("id", id))

	sql, args := sb.Build()

	city := domain.CityEntity{}

	err := cr.db.Get(&city, sql, args...)

	return &city, err
}

func (cr *CityRepository) GetCities(gt string, lt string, countryCode string, district string) (*[]domain.CityEntity, error) {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "name", "countrycode", "district", "population").From("city")

	criteria := []string{}

	if gt != "" {
		criteria = append(criteria, sb.GreaterEqualThan("population", gt))
	}

	if lt != "" {
		criteria = append(criteria, sb.LessEqualThan("population", lt))
	}

	if countryCode != "" {
		criteria = append(criteria, sb.Equal("countrycode", countryCode))
	}

	if district != "" {
		criteria = append(criteria, sb.Equal("district", district))
	}

	if len(criteria) > 0 {
		sb.Where(criteria...)
	}

	sql, args := sb.Build()

	city := []domain.CityEntity{}

	err := cr.db.Select(&city, sql, args...)

	return &city, err
}
