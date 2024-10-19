package middleware

import (
	"errors"
	"net/http"

	"github.com/eliot6001/goapi/api"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.");

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){
		var username string = r.URL.Query().Get("Author")
		var token = r.Header.Get("Authorization");
		var err error;
		if username == "" || token == "" {
		 log.Error(UnAuthorizedError);
		 api.RequestErrorHandler(w, UnAuthorizedError)
		 return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase();
		if err != nil {
			api.InternalErrorHandler(w);
			return
		}
		var loginDetails *tools.LoginDetails;
		if(loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError);
			return;
		}
		next.ServeHTTP(w, r);
	}
} 
