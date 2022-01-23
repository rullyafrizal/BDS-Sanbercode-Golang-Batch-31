package middlewares

// auth middleware
import (
	"net/http"
	"quiz-3/utils"

	"github.com/julienschmidt/httprouter"
)

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			response := map[string]string{"message": "Unauthorized, please fill your username and password"}

			utils.ResponseJSON(rw, response, http.StatusUnauthorized)
			return
		}

		if val, ok := utils.ValidateUser(uname, pwd); !ok {
			response := map[string]string{"message": val}

			utils.ResponseJSON(rw, response, http.StatusUnauthorized)
			return
		}

		next(rw, r, p)
	}
}
