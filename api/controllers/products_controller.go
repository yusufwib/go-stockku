package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"
	"io"
	"os"
	"math/rand"
	"time"
	"path/filepath"
	// "bytes"
	// "encoding/gob"
	"strconv"

	"github.com/yusufwib/go-stockku/api/auth"
	"github.com/yusufwib/go-stockku/api/models"
	"github.com/yusufwib/go-stockku/api/responses"
	"github.com/yusufwib/go-stockku/api/utils/formaterror"
)

type ImgXixixi struct {
	ID int 
}

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Product{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.ValidateProduct()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	
	//need develop
	if uid == 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(""))
		return
	}

	postCreated, err := post.SaveProduct(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {

	post := models.Product{}

	posts, err := post.FindAllProducts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) uploadImg(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames
	productID := r.FormValue("product_id")

	id, _ := strconv.Atoi(productID)

	fmt.Println(productID)

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//validate belum
	
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		filenameAlias := "IMG-PROD-" + RandStringRunes(11) + ".png"
		
		fileLocation := filepath.Join(dir, "files", filenameAlias)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()
		
		if _, err := io.Copy(targetFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		detImg := models.DetailProductImage{}
		
		
		b := models.DetailProductImage{
				ProductID : id,
				Image : "/files/" + filenameAlias,
			}
		body, _ := json.Marshal(b)

		err = json.Unmarshal(body, &detImg)
		fmt.Print(err)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		detImg.Prepare()

		postCreated, err := detImg.Save(server.DB)
		fmt.Print(postCreated)
		if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path))
	responses.JSON(w, http.StatusCreated, postCreated)
	
}
	
}




func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}