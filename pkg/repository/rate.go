package repository

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/yaroshevichM/software-school-case/pkg/models"
)

const url = "https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5"

type RateRepository struct {
	db *sqlx.DB
}

func newRateRepository(db *sqlx.DB) *RateRepository {
	return &RateRepository{db}
}

func (r *RateRepository) GetRates() ([]models.Rate, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch data")
	}

	var rawRates []struct {
		Ccy     string `json:"ccy"`
		BaseCcy string `json:"base_ccy"`
		Buy     string `json:"buy"`
		Sale    string `json:"sale"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&rawRates); err != nil {
		return nil, err
	}

	var rates []models.Rate
	for _, rawRate := range rawRates {
		buy, err := strconv.ParseFloat(rawRate.Buy, 64)
		if err != nil {
			return nil, err
		}

		rate := models.Rate{
			Currency:     rawRate.Ccy,
			BaseCurrency: rawRate.BaseCcy,
			Amount:       buy,
		}
		rates = append(rates, rate)
	}

	return rates, nil
}
