package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"insights/eath/v2/views/pages"
	views "insights/eath/v2/views/shared"
	"net/http"
)

func main() {
	component := views.Main()

	http.Handle("/", templ.Handler(component))

	http.Handle("/login", http.HandlerFunc(login))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	component := pages.NewPage(email)
	component.Render(context.Background(), w)
}
