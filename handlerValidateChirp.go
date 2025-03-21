package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func filter(str *string) {
	filteredWords := [3]string{"kerfuffle", "sharbert", "fornax"}
	wordSlice := strings.Split(*str, " ")
	for i, word := range wordSlice {
		w := strings.ToLower(word)
		for _, profane := range filteredWords {
			if profane == w {
				wordSlice[i] = "****"
				break
			}
		}
	}
	*str =  strings.Join(wordSlice, " ")
}

func validateChripHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type respVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
		return
	}
	filter(&params.Body)
	respondWithJson(w, http.StatusOK, respVals{
		CleanedBody: params.Body,
	})
}
