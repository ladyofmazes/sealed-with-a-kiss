package main

import (
	_ "embed"
	"log"
	"net/http"

	lbook "github.com/ladyofmazes/linkbook/lib"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type figure1 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure1) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-grass")

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"It is a beautiful day",
		"To your left is your local pub",
		"To your right is a bakery",
		"Or you could go home"}
	h.figurepage.Ilinks = []string{"", "", "/sealed-with-a-kiss-drinks", "/sealed-with-a-kiss-bakery", "/sealed-with-a-kiss-room"}
	// Load the stored value
	for i, val := range h.figurepage.Ipage {
		ctx.Dispatch(func(ctx app.Context) {
			var value int
			ctx.SessionStorage().Get(h.figurepage.Ipage[i], &value)

			h.figurepage.IpageScore[val] = value

			var visits int
			ctx.SessionStorage().Get(h.figurepage.Ipage[i]+"Visits", &visits)

			h.figurepage.IpageVisits[val] = visits
		})
	}
}

func (h *figure1) Render() app.UI {
	if h.figurepage == nil {
		h.figurepage = lbook.NewFigurePage()
	}
	if len(h.figurepage.Icaptions) == 0 {
		h.figurepage.Icaptions = []string{""}
	}
	if len(h.figurepage.Ilinks) == 0 {
		h.figurepage.Ilinks = []string{""}
	}
	return h.figurepage.
		Name("sealed-with-a-kiss-grass").
		Figure(
			"/web/20251224_144535.png",
		).Caption(h.figurepage.Icaptions...).Audio("/web/BriskWalk.wav").Links(h.figurepage.Ilinks...)
}

type figure2 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure2) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-drinks")

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"You go to the local pub",
		"After several drinks you don't feel well",
		"Somehow you forgot you are allergic to alcohol",
		"You died (but luckily were resurrected by a cleric because this is a magical world)"}
	h.figurepage.Ilinks = []string{"", "", "", "", "/"}
	// Load the stored value
	for i, val := range h.figurepage.Ipage {
		ctx.Dispatch(func(ctx app.Context) {
			var value int
			ctx.SessionStorage().Get(h.figurepage.Ipage[i], &value)

			h.figurepage.IpageScore[val] = value

			var visits int
			ctx.SessionStorage().Get(h.figurepage.Ipage[i]+"Visits", &visits)

			h.figurepage.IpageVisits[val] = visits
		})
	}
}

func (h *figure2) Render() app.UI {
	if h.figurepage == nil {
		h.figurepage = lbook.NewFigurePage()
	}
	if len(h.figurepage.Icaptions) == 0 {
		h.figurepage.Icaptions = []string{""}
	}
	if len(h.figurepage.Ilinks) == 0 {
		h.figurepage.Ilinks = []string{""}
	}
	return h.figurepage.
		Name("sealed-with-a-kiss-drinks").
		Figure(
			"/web/20251129_173150.png",
		).Caption(h.figurepage.Icaptions...).Audio("/web/Forsaken.wav").Links(h.figurepage.Ilinks...)
}

type figure3 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure3) Render() app.UI {
	if h.figurepage == nil {
		h.figurepage = lbook.NewFigurePage()
	}
	if len(h.figurepage.Icaptions) == 0 {
		h.figurepage.Icaptions = []string{""}
	}
	if len(h.figurepage.Ilinks) == 0 {
		h.figurepage.Ilinks = []string{""}
	}
	return h.figurepage.
		Name("sealed-with-a-kiss-bakery").
		Figure(
			"/web/20251208_121710.png",
		).Audio("/web/ASongForRoss.wav")
}

type figure4 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure4) Render() app.UI {
	if h.figurepage == nil {
		h.figurepage = lbook.NewFigurePage()
	}
	if len(h.figurepage.Icaptions) == 0 {
		h.figurepage.Icaptions = []string{""}
	}
	if len(h.figurepage.Ilinks) == 0 {
		h.figurepage.Ilinks = []string{""}
	}
	return h.figurepage.
		Name("sealed-with-a-kiss-room").
		Figure(
			"/web/20251201_174639.png",
		).Audio("/web/QuietChat.wav")
}

type figure5 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure5) Render() app.UI {
	if h.figurepage == nil {
		h.figurepage = lbook.NewFigurePage()
	}
	if len(h.figurepage.Icaptions) == 0 {
		h.figurepage.Icaptions = []string{""}
	}
	if len(h.figurepage.Ilinks) == 0 {
		h.figurepage.Ilinks = []string{""}
	}
	return h.figurepage.
		Name("sealed-with-a-kiss-door").
		Figure(
			"/web/20251230_171946.png",
		).Audio("/web/Trepidation.wav")
}

type figure6 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure6) Render() app.UI {
	if h.figurepage == nil {
		h.figurepage = lbook.NewFigurePage()
	}
	if len(h.figurepage.Icaptions) == 0 {
		h.figurepage.Icaptions = []string{""}
	}
	if len(h.figurepage.Ilinks) == 0 {
		h.figurepage.Ilinks = []string{""}
	}
	return h.figurepage.
		Name("sealed-with-a-kiss-kiss").
		Figure(
			"/web/20260104_155132.png",
		).Audio("/web/CheerfulSunshine.wav")
}

func main() {
	app.Route("/", func() app.Composer { return &figure1{} })
	app.Route("/sealed-with-a-kiss-drinks", func() app.Composer { return &figure2{} })
	app.Route("/sealed-with-a-kiss-bakery", func() app.Composer { return &figure3{} })
	app.Route("/sealed-with-a-kiss-room", func() app.Composer { return &figure4{} })
	app.Route("/sealed-with-a-kiss-door", func() app.Composer { return &figure5{} })
	app.Route("/sealed-with-a-kiss-kiss", func() app.Composer { return &figure6{} })

	app.RunWhenOnBrowser()

	if app.IsClient {
		select {}
	}

	http.Handle("/", &app.Handler{
		Name:        "That Time I Gave a Dog a Cookie and Then I Did It 14 More Times and Was Transported to Another World",
		Description: "That time I gave a dog a cookie and then I did it 14 more times and was transported to another world",
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
