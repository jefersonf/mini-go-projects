package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	endpoints = [5]string{"home", "form", "list", "thanks", "sorry"}
	templates = make(map[string]*template.Template)

	responses = make([]*Rsvp, 0, 10)
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

type formData struct {
	*Rsvp
	Errors  []string
	Session *Rsvp
}

func main() {
	loadTemplates()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func loadTemplates() {
	for _, page := range endpoints {
		t, err := template.ParseFiles("web/layout.html", fmt.Sprintf("web/%s.html", page))
		if err != nil {
			panic(err)
		}
		templates[page] = t
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates["home"].Execute(w, nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates["form"].Execute(w, formData{
			Rsvp: &Rsvp{}, Errors: []string{}, Session: &Rsvp{},
		})
		if err != nil {
			fmt.Println(err)
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		respData := &Rsvp{
			Name:       r.Form["name"][0],
			Email:      r.Form["email"][0],
			Phone:      r.Form["phone"][0],
			WillAttend: r.Form["willattend"][0] == "true",
		}

		formErrors := []string{}
		if respData.Name == "" {
			formErrors = append(formErrors, "enter your name")
		}
		if respData.Email == "" {
			formErrors = append(formErrors, "enter your email")
		}
		if respData.Phone == "" {
			formErrors = append(formErrors, "enter your phone number")
		}

		if len(formErrors) > 0 {
			currentForm := respData
			templates["form"].Execute(w, formData{
				Rsvp: &Rsvp{}, Errors: formErrors, Session: currentForm,
			})
			return
		}
		responses = append(responses, respData)

		if respData.WillAttend {
			templates["thanks"].Execute(w, respData.Name)
		} else {
			templates["sorry"].Execute(w, respData.Name)
		}
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, responses)
}
