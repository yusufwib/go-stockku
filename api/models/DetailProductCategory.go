package models

import (
	"errors"
	"html"

	// // "log"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type DetailProductCategory struct {
	ID         			uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	ProductID			int			`gorm:"not null" json:"product_id"`
	CategoryID			int			`gorm:"not null" json:"category_id"`
	// Category			Category
	CreatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (dpc *DetailProductCategory) Prepare() {

	pidToStr := strconv.Itoa(dpc.ProductID)
	pid, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(pidToStr)))
	if err != nil {
		return 
	}

	cidToStr := strconv.Itoa(dpc.CategoryID)
	cid, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(cidToStr)))
	if err != nil {
		return 
	}

	dpc.ID = 0
	dpc.ProductID = pid
	dpc.CategoryID = cid
	dpc.CreatedAt = time.Now()
	dpc.UpdatedAt = time.Now()
}

func (dpc *DetailProductCategory) Validate() error {

	if dpc.ProductID != 0  {
		return errors.New("Required Title")
	}
	if dpc.CategoryID != 0 {
		return errors.New("Required Content")
	}
	return nil
}

func (dpc *DetailProductCategory) SaveDetailProductCategory(db *gorm.DB) (*DetailProductCategory, error) {

	var err error
	err = db.Debug().Model(&DetailProductCategory{}).Create(&dpc).Error
	if err != nil {
		return &DetailProductCategory{}, err
	}
	return dpc, nil
}