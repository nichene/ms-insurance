package product

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"time"
)

type Product struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name          string    `json:"name" validate:"required"`
	Category      string    `json:"category" validate:"required"`
	BasePrice     float64   `json:"base_price" validate:"required"`
	TariffedPrice float64   `json:"tariffed_price,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CategoryEnum int

const (
	LifeCategory CategoryEnum = iota
	AutoCategory
	TravelCategory
	ResidencialCategory
	HeritageCategory
)

func (c CategoryEnum) String() string {
	return [...]string{"VIDA", "AUTO", "VIAGEM", "RESIDENCIAL", "PATRIMONIAL"}[c]
}

func (p *Product) containsIdentifiableCategory() bool {
	switch category := p.Category; category {
	case LifeCategory.String(), AutoCategory.String(), TravelCategory.String(), ResidencialCategory.String(), HeritageCategory.String():
		return true

	default:
		return false
	}
}

func (p *Product) calculateInsuranceProductTariff() {
	switch category := p.Category; category {
	case LifeCategory.String():
		p.TariffedPrice = p.BasePrice + (p.BasePrice * LifeCategoryIOF) + (p.BasePrice * LifeCategoryPIS) + (p.BasePrice * LifeCategoryCOFINS)
		return
	case AutoCategory.String():
		p.TariffedPrice = p.BasePrice + (p.BasePrice * AutoCategoryIOF) + (p.BasePrice * AutoCategoryPIS) + (p.BasePrice * AutoCategoryCOFINS)
		return
	case TravelCategory.String():
		p.TariffedPrice = p.BasePrice + (p.BasePrice * TravelCategoryIOF) + (p.BasePrice * TravelCategoryPIS) + (p.BasePrice * TravelCategoryCOFINS)
		return
	case ResidencialCategory.String():
		p.TariffedPrice = p.BasePrice + (p.BasePrice * ResidencialCategoryIOF) + (p.BasePrice * ResidencialCategoryPIS) + (p.BasePrice * ResidencialCategoryCOFINS)
		return
	case HeritageCategory.String():
		p.TariffedPrice = p.BasePrice + (p.BasePrice * HeritageCategoryIOF) + (p.BasePrice * HeritageCategoryPIS) + (p.BasePrice * HeritageCategoryCOFINS)
		return
	default:
		log.Default().Print("Invalid category")
		return
	}
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) GormDataType() string {
	return "json"
}

func (p *Product) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Product) Scan(src interface{}) error {
	*p = Product{}

	switch src := src.(type) {
	case nil:
		return nil

	case string:
		return json.Unmarshal([]byte(src), p)

	case []byte:
		return json.Unmarshal(src, p)

	default:
		return errors.New("scan: incompatible type for Product")
	}
}
