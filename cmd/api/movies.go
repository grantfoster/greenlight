package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"greenlight.grantfoster/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("this handler ran")

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	app.logger.Error(fmt.Sprintf("input is: %+v", input))
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Year:      2000,
		Title:     "Pitch Black",
		Runtime:   108,
		Genres:    []string{"action", "horror", "sci-fi"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}
