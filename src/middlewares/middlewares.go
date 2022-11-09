package middlewares

import (
	"github.com/marcosgfrites/REST-API_Golang/src/databases"
	"net/http"
)

// CheckDatabase checks if the database is accessible
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if databases.ConnectionCheck() == 0 {
			http.Error(w, "Database connection has been lost", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	}
}
