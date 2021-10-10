package global

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string `gorm:"primary_key;default:uuid_generate_v4()"`
	IsActive  bool   `gorm:"default:true"`
	Version   int    `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
