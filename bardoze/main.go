package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type dish struct {
	Code    string  `json:"code"`
	Name    string  `json:"name"`
	Desc    string  `json:"description"`
	Price   float32 `json:"price"`
	Special bool    `json:"special"`
}

type dishesFromJSON struct {
	Dishes []dish
}

type Menu map[string]*dish

type HTMLPages map[string]*template.Template

var (
	dishes    Menu
	templates HTMLPages
)

func init() {

	dishes = make(map[string]*dish)

	bytes, err := os.ReadFile("dishes.json")
	nilOrPanic(err)

	var data dishesFromJSON
	err = json.Unmarshal(bytes, &data)
	nilOrPanic(err)

	for _, d := range data.Dishes {
		dishes[d.Code] = &dish{
			Code:    d.Code,
			Name:    d.Name,
			Desc:    d.Desc,
			Price:   d.Price,
			Special: d.Special,
		}
	}

	views := []string{"welcome", "dishes"}
	templates = make(map[string]*template.Template)

	for _, p := range views {
		tmpl, err := template.ParseFiles("layout.html", p+".html")
		nilOrPanic(err)
		templates[p] = tmpl
	}
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/dishes", dishesHandler)

	err := http.ListenAndServe(":3000", nil)
	nilOrPanic(err)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, len(dishes))
}

func dishesHandler(w http.ResponseWriter, r *http.Request) {
	templates["dishes"].Execute(w, dishes)
}

func (d Menu) GetDishes() []dish {
	dishes := make([]dish, 0)
	for _, dish := range d {
		dishes = append(dishes, *dish)
	}
	return dishes
}

func nilOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
