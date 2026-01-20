package main

import (
	_ "embed"
	"log"
	"net/http"
	"strings"

	lbook "github.com/ladyofmazes/linkbook/lib"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type figure1 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure1) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-drinks")
	var figurePages []string = []string{"sealed-with-a-kiss-grass",
		"sealed-with-a-kiss-room",
		"sealed-with-a-kiss-door",
		"sealed-with-a-kiss-kiss",
		"sealed-with-a-kiss-drinks",
		"sealed-with-a-kiss-bakery"}

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"It is a beautiful day",
		"To your left is your local pub",
		"To your right is a bakery",
		"Or you could go home"}
	h.figurepage.Ilinks = []string{"", "", "/sealed-with-a-kiss-drinks", "/sealed-with-a-kiss-bakery", "/sealed-with-a-kiss-room"}

	var kissVisits int
	ctx.SessionStorage().Get("sealed-with-a-kiss-kiss"+"Visits", &kissVisits)
	if kissVisits > 11 {

		for _, val := range figurePages {
			ctx.Dispatch(func(ctx app.Context) {
				ctx.SessionStorage().Set(val, 0)
				ctx.SessionStorage().Set(val+"Visits", 0)
			})
		}
	}
	for i, val := range h.figurepage.Ipage {
		ctx.Dispatch(func(ctx app.Context) {
			var value int
			ctx.SessionStorage().Get(h.figurepage.Ipage[i], &value)

			h.figurepage.IpageScore[val] = value

			var visits int
			ctx.SessionStorage().Get(h.figurepage.Ipage[i]+"Visits", &visits)

			h.figurepage.IpageVisits[val] = visits

			if visits > 0 && strings.TrimSpace(val) == "sealed-with-a-kiss-drinks" {
				var newText []string = []string{"Click Below to Begin", "You feel a little sick like you drank alcohol which of course you are allergic to and a cleric brought you back to life"}
				var newLinks []string = []string{"", ""}
				h.figurepage.Icaptions = append(newText, h.figurepage.Icaptions[1:]...)
				h.figurepage.Ilinks = append(newLinks, h.figurepage.Ilinks[1:]...)
			}
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
		Page("sealed-with-a-kiss-drinks").
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
	var figurePages []string = []string{"sealed-with-a-kiss-grass",
		"sealed-with-a-kiss-room",
		"sealed-with-a-kiss-door",
		"sealed-with-a-kiss-kiss",
		"sealed-with-a-kiss-drinks",
		"sealed-with-a-kiss-bakery"}
	h.figurepage.Page(figurePages...)

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"You go to the local pub",
		"After several drinks you don't feel well",
		"Somehow you forgot you are allergic to alcohol",
		"You died (but luckily were resurrected by a cleric because this is a magical world)"}
	h.figurepage.Ilinks = []string{"", "", "", "", "/"}

	var kissVisits int
	ctx.SessionStorage().Get("sealed-with-a-kiss-drinks"+"Visits", &kissVisits)
	if kissVisits > 11 {
		h.figurepage.Icaptions = []string{"Click Below to Begin", "The cleric at the bar has a policy of only resurrecting the same person 12 times. No more no less. People in this world have a weird fixation on the number twelve.", "Start Over"}
		h.figurepage.Ilinks = []string{"", "", "/"}

		for _, val := range figurePages {
			ctx.Dispatch(func(ctx app.Context) {
				ctx.SessionStorage().Set(val, 0)
				ctx.SessionStorage().Set(val+"Visits", 0)
			})
		}
	}
	for _, val := range figurePages {
		ctx.Dispatch(func(ctx app.Context) {
			var value int
			ctx.SessionStorage().Get(val, &value)

			h.figurepage.IpageScore[strings.TrimSpace(val)] = value

			var visits int
			ctx.SessionStorage().Get(val+"Visits", &visits)

			h.figurepage.IpageVisits[strings.TrimSpace(val)] = visits
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

func (h *figure3) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-bakery")

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"You enter the bakery",
		"Obviously you have to buy a dozen of your favorite cookies",
		"Suddenly a dog appears",
		"\"I am an enchanted dog\", she says, \"And I need a kiss and a dozen cookies to return to my human form\"",
		"\"Is that really true?\"",
		"She doesn't say anything, just sitting there looking cute.",
		"Give her a cookie",
		"Tell her to go away"}
	h.figurepage.Ilinks = []string{"", "", "", "", "", "", "", "/sealed-with-a-kiss-door", "/sealed-with-a-kiss-room"}

	var kissVisits int
	ctx.SessionStorage().Get("sealed-with-a-kiss-room"+"Visits", &kissVisits)
	var kissVisits1 int
	ctx.SessionStorage().Get("sealed-with-a-kiss-door"+"Visits", &kissVisits1)

	var newCaption []string
	var newLinks []string
	if kissVisits > 0 {
		newCaption = []string{"You really don't want to go home alone again but you really want to eat all these cookies", "Give her a cookie", "Tell her to go away"}
		newLinks = []string{"", "/sealed-with-a-kiss-door", "/sealed-with-a-kiss-room"}
	}
	if kissVisits1 > 0 {
		var addCaption []string = []string{"You really feel like this dog is just scamming you for your cookies every day. And she seems weirdly into the kissing as well."}
		var addLink []string = []string{""}
		newCaption = append(addCaption, newCaption...)
		newLinks = append(addLink, newLinks...)
	}
	if kissVisits > 0 || kissVisits1 > 0 {
		h.figurepage.Icaptions = append(h.figurepage.Icaptions[0:6], newCaption...)
		h.figurepage.Ilinks = append(h.figurepage.Ilinks[0:6], newLinks...)
	}

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
		).Caption(h.figurepage.Icaptions...).Audio("/web/ASongForRoss.wav").Links(h.figurepage.Ilinks...)
}

type figure4 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure4) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-room")

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"This is your house. Somehow you ended up back here again. Alone.",
		"Leave"}
	h.figurepage.Ilinks = []string{"", "", "/"}

	var kissVisits int
	ctx.SessionStorage().Get("sealed-with-a-kiss-bakery"+"Visits", &kissVisits)
	if kissVisits > 0 {
		var newCaption []string = []string{"Click Below to Begin", "This is your house. Somehow you ended up back here again. Alone.", "You kind of wish you had given that dog a cookie", "Leave"}
		var newLinks []string = []string{"", "", "/sealed-with-a-kiss-bakery", "/"}

		h.figurepage.Icaptions = newCaption
		h.figurepage.Ilinks = newLinks
	}
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
		).Caption(h.figurepage.Icaptions...).Audio("/web/QuietChat.wav").Links(h.figurepage.Ilinks...)
}

type figure5 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure5) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	var figurePages []string = []string{"sealed-with-a-kiss-grass",
		"sealed-with-a-kiss-room",
		"sealed-with-a-kiss-door",
		"sealed-with-a-kiss-kiss",
		"sealed-with-a-kiss-drinks",
		"sealed-with-a-kiss-bakery"}
	h.figurepage.Page(figurePages...)

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"You gave the dog a cookie and before you knew it she was inviting you back to her apartment",
		"The building is very impressive. She must have a lot of money",
		"\"The code is 123456\", She tells you at the door.",
		"\"That's the same code as my luggage.\" You say, successfully entering the code and opening the door.",
		"Her apartment is as nice as the building. The floor is a beautiful blond hard wood and it is lushly furnished.",
		"\"I like what you've done with the place.\" You say.",
		"\"Just kiss me!\" She says. ",
		"Kiss the dog",
		"Refuse to kiss the dog"}
	h.figurepage.Ilinks = []string{"", "", "", "", "", "", "", "", "/", "/sealed-with-a-kiss-room"}

	var kissVisits int
	ctx.SessionStorage().Get("sealed-with-a-kiss-door"+"Visits", &kissVisits)
	if kissVisits > 11 {
		h.figurepage.Ilinks[8] = "/sealed-with-a-kiss-kiss"
	}
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
		).Caption(h.figurepage.Icaptions...).Audio("/web/Trepidation.wav").Links(h.figurepage.Ilinks...)
}

type figure6 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure6) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-door")

	h.figurepage.Icaptions = []string{"Click Below to Begin",
		"You kiss the dog, just like you have so many times before",
		"Despite the wetness and the smell you have grown used to the experience",
		"However this time something else happens",
		"Suddenly the dog shifts and grows and becomes a beautiful woman",
		"\"I told you I was a human. And I am also a princess and this apartment is an enchanted castle.\"",
		"As she is speaking the apartment is also shifting, turning from traditional hardwood to cold stone",
		"Embrace and marry the princess",
		"Refuse to marry the princess"}
	h.figurepage.Ilinks = []string{"", "", "", "", "", "", "", "/sealed-with-a-kiss-good-end", "/"}
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
		).Caption(h.figurepage.Icaptions...).Audio("/web/CheerfulSunshine.wav").Links(h.figurepage.Ilinks...)
}

type figure7 struct {
	app.Compo
	figurepage *lbook.FigurePage
}

func (h *figure7) OnMount(ctx app.Context) {
	h.figurepage = lbook.NewFigurePage()
	h.figurepage.Page("sealed-with-a-kiss-kiss")

	h.figurepage.Icaptions = []string{"Click Here to Begin",
		"You feel yourself changing. You grow a frighteningly long beard. Your shoulders grow broader and you can tell if you looked you would have a six pack.",
		"\"What is happening?\" You ask.",
		"\"A princess needs a knight.\" She says, standing on her tiptoes to kiss your now significantly taller frame. \"And of course there is the matter of the dragon.\"",
		"\"The dragon?\" you ask.",
		"In the distance you hear a horrifying roar. Also you suspect there are no cookies in this place.",
		"Restart?"}
	h.figurepage.Ilinks = []string{"", "", "", "", "", "", "/"}
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

func (h *figure7) Render() app.UI {
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
		Name("sealed-with-a-kiss-good-end").
		Caption(h.figurepage.Icaptions...).Links(h.figurepage.Ilinks...)
}

func main() {
	app.Route("/", func() app.Composer { return &figure1{} })
	app.Route("/sealed-with-a-kiss-drinks", func() app.Composer { return &figure2{} })
	app.Route("/sealed-with-a-kiss-bakery", func() app.Composer { return &figure3{} })
	app.Route("/sealed-with-a-kiss-room", func() app.Composer { return &figure4{} })
	app.Route("/sealed-with-a-kiss-door", func() app.Composer { return &figure5{} })
	app.Route("/sealed-with-a-kiss-kiss", func() app.Composer { return &figure6{} })
	app.Route("/sealed-with-a-kiss-good-end", func() app.Composer { return &figure7{} })

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
