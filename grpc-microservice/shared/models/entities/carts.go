package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Carts struct {
	bun.BaseModel `bun:"table:carts,alias:c"`

	ProductCode string    `bun:"product_code,pk"`
	ProductName string    `bun:"product_name"`
	Quantity    int32     `bun:"quantity"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
