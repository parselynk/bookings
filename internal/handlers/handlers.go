package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/parselynk/bookings/internal/config"
	"github.com/parselynk/bookings/internal/render"
	"github.com/parselynk/bookings/models"
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
	render.RenderTeplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	remoteIP := m.App.Session.GetString(r.Context(), "userIp")

	stringMap["userIp"] = remoteIP

	render.RenderTeplate(w, r, "about.page.tmpl", &models.TemplateData{
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

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTeplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals is the generals-quarters page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTeplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the majors-suite page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTeplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the search-availability page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTeplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the post-availability page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("start")
	w.Write([]byte(fmt.Sprintf("start date is: %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message`
}

// AvailabilityJson handles request for Availability and send JSON response
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTeplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
