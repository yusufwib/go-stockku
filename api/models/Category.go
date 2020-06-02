package models

import (
	"errors"
	"html"
	// "log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID         			uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	Name  	   			string    	`gorm:"size:255;not null;unique" json:"name"`
	CreatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Category) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Category) ValidateCategory() error {

	if p.Name == "" {
		return errors.New("There are still some that haven't been filled yet")
	}
	return nil
}

func (p *Category) SaveCategory(db *gorm.DB) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Create(&p).Error
	if err != nil {
		return &Category{}, err
	}
	return p, nil
}

func (p *Category) FindAllCategory(db *gorm.DB) (*[]Category, error) {
	var err error
	products := []Category{}
	err = db.Debug().Model(&Category{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &products, nil
}
