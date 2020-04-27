package interactors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	. "common_dashboard_backend/common/logger"
	"common_dashboard_backend/entities"

	"golang.org/x/crypto/bcrypt"
)

var errorMap map[string]map[int]string

func init() {

	//CachedData = make(map[int]*summary.Data)

	errorMap = make(map[string]map[int]string)

	errorMap["tr"] = make(map[int]string)
	errorMap["en"] = make(map[int]string)

	errorMap["tr"][10001] = "Kullanici Email adresi bulunamadi"
	errorMap["tr"][10002] = "Kullanici bulunamadi"
	errorMap["tr"][10003] = "Yanlis sifre girildi"
	errorMap["tr"][10004] = "Token olusturulamadi"
	errorMap["tr"][10005] = "Kullanici adi bulunamadi"
	errorMap["tr"][10006] = "Kullanıcı rolü belirtilmeli"

	errorMap["tr"][10010] = "İşlem başarısız"
	errorMap["tr"][10011] = "Güncelleme başarısız"
	errorMap["tr"][10012] = "Silme işlemi başarısız"
	errorMap["tr"][10013] = "Hiçbir veri etkilenmedi"
	errorMap["tr"][10014] = "Bu isim zaten kullanılıyor"

	//ENG
	errorMap["en"][10001] = "User email address couldn't found"
	errorMap["en"][10002] = "User couldn't found"
	errorMap["en"][10003] = "Wrong password is entered"
	errorMap["en"][10004] = "Couldn't generate found"
	errorMap["en"][10005] = "Username not found"
	errorMap["en"][10006] = "User role can not be blank"

	errorMap["en"][10010] = "Operation failed"
	errorMap["en"][10011] = "Updating failed"
	errorMap["en"][10012] = "Deleting failed"
	errorMap["en"][10013] = "No data affected"
	errorMap["en"][10014] = "This name already used"
}

func Post(url string, jsonData string) string {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return string(contents)
	}
	return ""
}

func GetError(code int, err ...error) *entities.ErrorType {

	if code == 0 {
		return &entities.ErrorType{Code: code, Message: err[0].Error()}
	}

	return &entities.ErrorType{Code: code, Message: errorMap["tr"][code]}
}

func hashAndSalt(pwd []byte) []byte {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		LogError(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return hash
}
