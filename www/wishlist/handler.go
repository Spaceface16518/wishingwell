package wishlist

import (
	"net/http"
	"wishingwell/internal/db"
	"wishingwell/internal/db/db_init"
	"wishingwell/internal/db/wish"
	"wishingwell/internal/templates"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("owner")
	if id == "" {
		w.WriteHeader(400)
		_, err := w.Write([]byte("No ID parameter was given"))
		if err != nil {
			panic(err)
		}
		return
	}

	data, err := view(w, id)
	if err != nil {
		w.WriteHeader(500)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}

	err = templates.Templates.ExecuteTemplate(w, "wishingwell-view.gohtml", data)
	if err != nil {
		w.WriteHeader(500)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}
}

type ViewData struct {
	Wishlist []wish.Wish
	Owner    string
}

func view(w http.ResponseWriter, id string) (ViewData, error) {
	err := db_init.EnsureInit()
	if err != nil {
		return ViewData{}, err
	}

	query := wish.OwnerQuery(id)
	wishes, err := db.DB.Find(query)
	if err != nil {
		return ViewData{}, err
	}

	return ViewData{
		wishes,
		id,
	}, nil

}
