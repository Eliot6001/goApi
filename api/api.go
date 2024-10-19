package api
import (
	"encoding/json"
	"net/http"
)

type ArticleParams struct {
	Author string
}

type ArticleRenspose struct {
	Code int // response 2xx
	NumberOfArticles int64  
} 
type Error struct{
 Code int //response of 4xx / 5xx
 Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(code);
	json.NewEncoder(w).Encode(resp);
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadResponse)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An unexpected error occured.", http.StatusInternalServerError);
	}
)
