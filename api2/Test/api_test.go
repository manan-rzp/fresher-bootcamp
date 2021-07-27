package Test


import (
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"store-api/Config"
	"store-api/Router"
	"strings"
	"testing"
)

var err error

func TestCreateProduct(t *testing.T) {

	Config.DB,err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		panic(err)
	}
	jsonParam := `{"name":"sample","quantity":1, "price" : 1}`
	req, err := http.NewRequest("POST", "/product", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router := Router.SetupRouter()
	handler := http.Handler(Router)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPlaceOrder(t *testing.T) {

	Config.DB,err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		panic(err)
	}
	jsonParam := `{"ProductID":1, "UserID":1}`
	req, err := http.NewRequest("POST", "/order", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router := Router.SetupRouter()
	handler := http.Handler(Router)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreateUser(t *testing.T) {

	Config.DB,err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		panic(err)
	}
	jsonParam := `{"name":"sample"}`
	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(jsonParam)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router := Router.SetupRouter()
	handler := http.Handler(Router)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
