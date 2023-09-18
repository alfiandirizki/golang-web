package web_golang

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)




func TemplateDataMAP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Tittle": "Template data",
		"Name" : "Vian",
		"Address": map[string]interface{}{
			"Street" :"jln fuad",
		},
	})
}
func TestData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMAP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
type Page struct{
	Tittle string
	Name string
	Address Address
}
type Address struct{
	Street string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Tittle: "Amba",
		Name: "Tukam",
		Address: Address{
			Street: "JLN NGAWI",
		},

		
	})
}
func TestDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}