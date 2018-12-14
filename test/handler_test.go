package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"simpleapi/forms"
	"simpleapi/handlers"
	"simpleapi/lib"
	"simpleapi/models"
	"testing"
	"time"
)

const connSQL = "demo:demo@/demo?charset=utf8&parseTime=True&loc=Local"

var rsp struct {
	Key     string
	Message string
}

var (
	cmpS, cmpC models.Company
	cntrct     models.Company
)

func init() {
	var err error
	if lib.Env.DB, err = sql.Open("mysql", connSQL); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if err := lib.InitDB(lib.Env.DB); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func TestCompanyCreate(t *testing.T) {

	companyS := forms.Company{
		Name: "seller",
	}

	companyC := forms.Company{
		Name: "client",
	}

	g := gin.Default()
	var body bytes.Buffer

	//wrong request
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	testRecorder := httptest.NewRecorder()
	g.POST("/", handlers.HandleCompany)
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusBadRequest {
		t.Error("wrong response")
		t.Fail()
	}

	//positive case
	bod, err := json.Marshal(companyS)
	if err != nil {
		t.Fatalf("marshalling error %v", err)
	}
	body.Write(bod)

	req, err = http.NewRequest("POST", "/create", &body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	testRecorder = httptest.NewRecorder()
	g.POST("/:action", handlers.HandleCompany)
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Error("not OK response")
		t.Fail()
	}

	if json.Unmarshal(testRecorder.Body.Bytes(), &cmpS) != nil {
		t.Error("unmarshall body ", testRecorder.Body.String())
		t.Fail()
	}

	bod, err = json.Marshal(companyC)
	if err != nil {
		t.Fatalf("marshalling error %v", err)
	}
	body.Write(bod)

	req, err = http.NewRequest("POST", "/create", &body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	testRecorder = httptest.NewRecorder()
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Error("not OK response")
		t.Fail()
	}

	if json.Unmarshal(testRecorder.Body.Bytes(), &cmpC) != nil {
		t.Error("unmarshall body ", testRecorder.Body.String())
		t.Fail()
	}

	t.Log("company create processed")
}

func TestContract(t *testing.T) {

	contract := forms.Contract{
		Seller:      cmpS.ID,
		Client:      cmpC.ID,
		DateSigned:  time.Now(),
		ValidTill:   time.Now().Add(time.Hour * 5),
		CreditsInit: 15,
	}

	g := gin.Default()
	var body bytes.Buffer

	//wrong request
	req, err := http.NewRequest("POST", "/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	testRecorder := httptest.NewRecorder()
	g.POST("/:action", handlers.HandleContract)
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusBadRequest {
		t.Error("wrong response")
		t.Fail()
	}

	//positive case
	bod, err := json.Marshal(contract)
	if err != nil {
		t.Fatalf("marshalling error %v", err)
	}
	body.Write(bod)

	req, err = http.NewRequest("POST", "/create", &body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	testRecorder = httptest.NewRecorder()
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Error("not OK response")
		t.Fail()
	}

	if json.Unmarshal(testRecorder.Body.Bytes(), &cntrct) != nil {
		t.Error("unmarshall body ", testRecorder.Body.String())
		t.Fail()
	}

	t.Log("contract create processed")

}

func TestPurchase(t *testing.T) {

	purchase := forms.Purchase{
		Contract: cntrct.ID,
		Credits:  5,
	}

	g := gin.Default()
	var body bytes.Buffer

	//wrong request
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	testRecorder := httptest.NewRecorder()
	g.POST("/", handlers.MakePurchase)
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusBadRequest {
		t.Error("wrong response")
		t.Fail()
	}

	//positive case
	bod, err := json.Marshal(purchase)
	if err != nil {
		t.Fatalf("marshalling error %v", err)
	}
	body.Write(bod)

	req, err = http.NewRequest("POST", "/", &body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	testRecorder = httptest.NewRecorder()
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Error("not OK response")
		t.Fail()
	}

	if json.Unmarshal(testRecorder.Body.Bytes(), &rsp) != nil {
		t.Error("unmarshall body ", testRecorder.Body.String())
		t.Fail()
	}

	t.Log("purchase processed")

}

func TestDelete(t *testing.T) {

	contract := forms.Contract{
		ID: cntrct.ID,
	}

	company := forms.Company{
		ID: cmpC.ID,
	}

	g := gin.Default()
	var body bytes.Buffer

	//positive case
	bod, err := json.Marshal(contract)
	if err != nil {
		t.Fatalf("marshalling error %v", err)
	}
	body.Write(bod)

	req, err := http.NewRequest("POST", "/contract/delete", &body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	testRecorder := httptest.NewRecorder()
	g.POST("/contract/:action", handlers.HandleContract)
	g.POST("/company/:action", handlers.HandleCompany)
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Error("not OK response")
		t.Fail()
	}

	if json.Unmarshal(testRecorder.Body.Bytes(), &rsp) != nil {
		t.Error("unmarshall body ", testRecorder.Body.String())
		t.Fail()
	}

	bod, err = json.Marshal(company)
	if err != nil {
		t.Fatalf("marshalling error %v", err)
	}
	body.Write(bod)

	req, err = http.NewRequest("POST", "/company/delete", &body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	testRecorder = httptest.NewRecorder()
	g.ServeHTTP(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Error("not OK response")
		t.Fail()
	}

	if json.Unmarshal(testRecorder.Body.Bytes(), &rsp) != nil {
		t.Error("unmarshall body ", testRecorder.Body.String())
		t.Fail()
	}

	t.Log("delete processed")

}

func TestClose(t *testing.T) {
	_, _ = lib.Env.DB.Exec("DELETE FROM company")
	lib.Close(lib.Env.DB)
}
