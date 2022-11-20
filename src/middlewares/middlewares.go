package middlewares

import (
	"github.com/marcosgfrites/REST-API_Golang/src/databases"
	"github.com/marcosgfrites/REST-API_Golang/src/utils/common"
	"net/http"
)

// CheckDatabase checks if the database is accessible
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if databases.ConnectionCheck() == 0 {
			http.Error(w, common.DatabaseConnectionLost, http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	}
}
