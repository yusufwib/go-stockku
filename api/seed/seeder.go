package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yusufwib/go-stockku/api/models"
)

var users = []models.User{
	models.User{
		Name: "Steveeeeee",
		Username: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
		StoreName: "Kopi",
		Phone: "081334955666",
	},
	models.User{
		Name: "Srema",
		Username: "QSteven victor",
		Email:    "QQsteven@gmail.com",
		Password: "ppassword",
		StoreName: "pKopi",
		Phone: "081134955666",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

var cat = []models.Category{
	models.Category{
		Name: "Sport",
	},
	models.Category{
		Name: "Baju",
	},
	models.Category{
		Name: "Makanan",
	},
}


var detCat = []models.DetailProductCategory{
	models.DetailProductCategory{
		ProductID: 1,
		CategoryID: 2,
	},
	models.DetailProductCategory{
		ProductID: 1,
		CategoryID: 1,
	},
	models.DetailProductCategory{
		ProductID: 2,
		CategoryID: 2,
	},
}

var products = []models.Product{
	models.Product{
		Name: "Kopi",
		Weight: 50,
		PurchasePrice: 10000,
		SellingPrice: 12000,
		Desc: "good",
		Stock: 20,
		MaxSupply: 100,
	},
	models.Product{
		Name: "Kopi Hitam Kupu-Kupu",
		Weight: 50,
		PurchasePrice: 10000,
		SellingPrice: 12000,
		Desc: "cukupan",
		Stock: 20,
		MaxSupply: 100,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.Product{}, &models.DetailProductCategory{}, &models.Category{},  &models.DetailProductImage{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Product{},&models.DetailProductCategory{}, &models.Category{}, &models.DetailProductImage{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

	for i, _ := range products {
		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
	}
	for i, _ := range detCat {
		err = db.Debug().Model(&models.DetailProductCategory{}).Create(&detCat[i]).Error
		if err != nil {
			log.Fatalf("cannot seed detcat table: %v", err)
		}
	}
	for i, _ := range cat {
		err = db.Debug().Model(&models.Category{}).Create(&cat[i]).Error
		if err != nil {
			log.Fatalf("cannot seed cat table: %v", err)
		}
	}
}