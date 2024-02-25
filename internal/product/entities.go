package product

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Product struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name          string    `json:"name" validate:"required"`
	Category      string    `json:"category" validate:"required"`
	BasePrice     int64     `json:"base_price" validate:"required" gorm:"-"`
	TariffedPrice int64     `json:"tariffedPrice,omitempty" gorm:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p *Product) calculateInsuranceProductTariff() {
	// TODO
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
