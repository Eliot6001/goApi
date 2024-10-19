package handlers

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/Eliot6001/goApi/api"
	"github.com/Eliot6001/goApi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetNumberOfArticles(w http.ResponseWriter, r *http.Request) {
	var params = api.ArticleParams{}

	//one not sure if good, thing i did is add Username to the api/api so i can actually 
	// be able to retrieve data without schema error

	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	fmt.Print(params)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.AuthorDetails;
	tokenDetails = (*database).GetAuthorArticles(params.Author);

	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	response := api.ArticleResponse{
		NumberOfArticles: (*tokenDetails).NumberOfArticles,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
