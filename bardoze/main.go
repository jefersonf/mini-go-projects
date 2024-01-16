package main

import (
	"html/template"
	"net/http"
)

type dish struct {
	Name, Desc, Code string
	Price            float32
	Special          bool
}

type Menu map[string]*dish

type HTMLPages map[string]*template.Template

var (
	dishes    Menu
	templates HTMLPages
)

func init() {
	// load dishes sample
	// TODO: move this data to dishes.json
	dishes = map[string]*dish{
		"SUC01": {
			Name:  "Suco de Laranja",
			Desc:  "Suco de laranja natural, 300ml",
			Price: 4.,
		},
		"SUC02": {
			Name:  "Suco de Uva",
			Desc:  "Suco de uva natural, 300ml",
			Price: 5.,
		},
		"SAL01": {
			Name:  "Pastel de Queijo",
			Desc:  "Pastel de Queijo Mussarela, 300g",
			Price: 12.,
		},
		"SAL02": {
			Name:    "Pastel de Carne de Sol",
			Desc:    "Pastel de Carne de Sol, 300g",
			Price:   18.,
			Special: true,
		},
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
