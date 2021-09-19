package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/parselynk/bookings/models"
	"github.com/parselynk/bookings/pkg/config"
	"github.com/parselynk/bookings/pkg/render"
)

// Repo the repository used by the handler
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "userIp", remoteIP)
	render.RenderTeplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	remoteIP := m.App.Session.GetString(r.Context(), "userIp")

	stringMap["userIp"] = remoteIP

	render.RenderTeplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page handler
func Divide(w http.ResponseWriter, r *http.Request) {

	var x, y float32
	x = 2.0
	y = 4.0

	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f  divided by %f is %f", x, y, f))
}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("Cannot divide by zero.")
		return 0, err
	}

	result := x / y

	return result, nil
}
