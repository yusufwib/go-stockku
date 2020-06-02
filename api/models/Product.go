package models

import (
	"errors"
	"fmt"
	"html"

	// "log"
	"strconv"
	"strings"
	"time"
	"reflect"

	"github.com/jinzhu/gorm"
)
type Product struct {
	ID         			uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	Name  	   			string    	`gorm:"size:255;not null;unique" json:"name"`
	Desc  	   			string    	`gorm:"size:255;not null" json:"desc"`
	Weight 	        	int 	   	`gorm:"size:100;not null" json:"weight"`
	PurchasePrice   	int   	 	`gorm:"size:100;not null" json:"purchase_price"`
	SellingPrice   		int   	 	`gorm:"size:100;not null" json:"selling_price"`
	Stock		   		int    		`gorm:"size:100;not null" json:"stock"`
	MaxSupply		    int    		`gorm:"size:100;not null" json:"max_supply"`
	Category			[]Category  `json:"category"`
	ProductImage		[]DetailProductImage  `json:"product_image"`
	CreatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) Prepare() {
	
	weightToStr := strconv.Itoa(p.Weight)
	weight, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(weightToStr)))
	if err != nil {
		return 
	}

	ppToStr := strconv.Itoa(p.PurchasePrice)
	purchasePrice, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(ppToStr)))
	if err != nil {
		return 
	}

	spToStr := strconv.Itoa(p.SellingPrice)
	sellingPrice, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(spToStr)))
	if err != nil {
		return 
	}

	stockToStr := strconv.Itoa(p.SellingPrice)
	stock, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(stockToStr)))
	if err != nil {
		return 
	}

	msToStr := strconv.Itoa(p.MaxSupply)
	maxSupply, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(msToStr)))
	if err != nil {
		return 
	}

	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Desc = html.EscapeString(strings.TrimSpace(p.Desc))
	p.Weight = weight
	p.PurchasePrice = purchasePrice
	p.SellingPrice = sellingPrice
	p.Stock = stock
	p.MaxSupply = maxSupply
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) ValidateProduct() error {

	if p.Name == "" {
		return errors.New("There are still some that haven't been filled yet")
	}
	if p.Desc == "" {
		return errors.New("There are still some that haven't been filled yet")
	}
	if p.Weight < 0 {
		return errors.New("There are still some that haven't been filled yet")
	}
	if p.PurchasePrice < 0 {
		return errors.New("There are still some that haven't been filled yet")
	}
	if p.SellingPrice < 0 {
		return errors.New("There are still some that haven't been filled yet")
	}
	if p.Stock < 0 {
		return errors.New("There are still some that haven't been filled yet")
	}
	if p.MaxSupply < 0 {
		return errors.New("There are still some that haven't been filled yet")
	}
	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	// if errr != nil {
	// 	return &[]DetailProductCategory{}, err
	// }
	t := reflect.TypeOf(Product{})
	fmt.Println(t)

	// multiple := reflect.MakeSlice(reflect.SliceOf(t), 0, 10)

	if len(products) > 0 {
		for i, _ := range products {
			// fmt.Println(products[i].ID)
			err := db.Debug().Table("categories").Joins("inner join detail_product_categories on categories.id = detail_product_categories.category_id").Where("product_id = ?", products[i].ID).Find(&products[i].Category).Error
			if err != nil {
				return &[]Product{}, err
			}
		}
	}
	
	if len(products) > 0 {
		for i, _ := range products {

			errImg := db.Debug().Table("detail_product_images").Joins("inner join products on detail_product_images.product_id = products.id").Where("product_id = ?", products[i].ID).Find(&products[i].ProductImage).Error
			if errImg != nil {
				return &[]Product{}, err
			}
		}
	}
	return &products, nil
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.ID).Take(&p.Name).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}


func (p *Product) UpdateAProduct(db *gorm.DB) (*Product, error) {

	var err error

	err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{Name: p.Name, Weight: p.Weight,PurchasePrice: p.PurchasePrice,SellingPrice: p.SellingPrice,Stock: p.Stock,MaxSupply: p.MaxSupply ,UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteAProduct(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}