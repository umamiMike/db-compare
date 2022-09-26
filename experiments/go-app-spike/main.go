package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo

	name string
	bar  string
}

func (h *hello) helloComponent() app.HTMLP {
	evald := foo{Foo: "Foo foo string"}
	return app.P().Body(
		app.Input().
			Type("text").
			Value(h.name).
			Placeholder("What is your name?").
			AutoFocus(true).
			OnChange(h.ValueTo(&h.name)),
		app.Input().
			Type("text").
			Value(h.bar).
			Placeholder("What is your foo?").
			AutoFocus(true).
			OnChange(h.ValueTo(&h.bar)),
		app.Text(evald.evaluate().evaluated),
	)
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		h.helloComponent(),
		app.H1().Body(app.Text(h.bar)),
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "",
				app.Text(h.name),
			).Else(
				app.Text("World!"),
			),
		),
	)
}

func main() {
	// Components routing:
	app.Route("/", &hello{})
	app.Route("/hello", &hello{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
