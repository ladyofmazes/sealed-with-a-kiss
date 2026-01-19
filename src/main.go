package main

import (
	_ "embed"
	"log"
	"net/http"

	lbook "github.com/ladyofmazes/linkbook/lib"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type figure struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure) Render() app.UI {
	var curPage = lbook.NewFigurePage()
	return curPage.
		Name("cookies").
		Figure(
			"/web/20251208_121710.png",
		).Caption(
		"Click Below to Begin", "Caption", "Caption2").Audio("/web/ASongForRoss.wav").Links("", "", "https://blog-2qu.pages.dev")
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the components with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return &figure{} })

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Add this check - if we're in browser, block forever
	if app.IsClient {
		select {} // Block forever - prevent Go runtime from exiting
	}

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Home",
		Description: "Home Page",
		Resources:   app.LocalDir("."),
		Styles: []string{
			"/app.css",
			"/web/css/prism.css",
			"/web/css/docs.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
