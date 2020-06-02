package models

import (
	// "errors"
	"html"

	// // "log"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type DetailProductImage struct {
	ID         			uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	ProductID			int			`gorm:"not null" json:"product_id"`
	Image				string		`gorm:"size:255;not null" json:"image"`
	CreatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *DetailProductImage) Prepare() {
	pidToStr := strconv.Itoa(p.ProductID)
	pid, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(pidToStr)))
	if err != nil {
		return 
	}
	p.ID = 0
	p.ProductID = pid
	p.Image = html.EscapeString(strings.TrimSpace(p.Image))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *DetailProductImage) Save(db *gorm.DB) (*DetailProductImage, error) {
	var err error
	err = db.Debug().Model(&DetailProductImage{}).Create(&p).Error
	if err != nil {
		return &DetailProductImage{}, err
	}
	return p, nil
}