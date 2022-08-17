package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Users struct {
	bun.BaseModel `bun:"table:users,alias:order_items"`

	ID         string    `bun:"id,pk,type:uuid"`
	IsActive   bool      `bun:"is_active"`
	CreatedBy  string    `bun:"created_by"`
	ModifiedBy string    `bun:"modified_by"`
	UserName   string    `bun:"user_name"`
	Password   string    `bun:"password,type:uuid"`
	CreatedAt  time.Time `bun:"created_at"`
	UpdatedAt  time.Time `bun:"updated_at"`
}
