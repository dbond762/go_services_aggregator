package theme

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/dbond762/go_services_aggregator/src/theme/models"
)

type Theme struct {
	data models.App
}

func NewTheme(title string) *Theme {
	return &Theme{
		data: models.App{
			Title:   title,
			Menu:    make([]models.MenuItem, 0),
			Content: nil,
		},
	}
}

func (t Theme) Init() {
	fs := http.FileServer(http.Dir("src/theme/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func (t *Theme) AddMenuItem(link, caption string, priority int) {
	menuItem := models.MenuItem{
		Link:     link,
		Caption:  caption,
		Priority: priority,
	}

	t.data.Menu = append(t.data.Menu, menuItem)
}

func (t Theme) Display(w http.ResponseWriter, paths []string, data interface{}) {
	themePaths := []string{
		"src/theme/templates/main.html",
		"src/theme/templates/header.html",
		"src/theme/templates/menu.html",
	}

	paths = append(themePaths, paths...)

	tmpl := template.Must(template.New("main.html").ParseFiles(paths...))

	t.data.Content = data

	sort.Slice(t.data.Menu, func(i, j int) bool {
		return t.data.Menu[i].Priority < t.data.Menu[j].Priority
	})

	if err := tmpl.Execute(w, t.data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}
}
