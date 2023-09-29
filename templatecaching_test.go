package web_golang

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

////go:embed templates/*.gohtml

//var templates embed.FS
//
////simpan diluar global function agar sekali eksekusi
//var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "hello template caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
