package files

import (
	"errors"
	"net/http"
	"strings"

	"github.com/schattenbrot/go-simple-upload-server/internal/config"
	"github.com/schattenbrot/go-simple-upload-server/packages/explerror"
)

func hasReadWriteAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			explerror.Forbidden(w, errors.New("no token provided"))
			return
		}
		bearerToken := strings.Split(bearer, " ")[1]

		isAllowed := false
		for _, token := range config.Tokens.ReadWriteTokens {
			if token == bearerToken {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			explerror.Forbidden(w, errors.New("token invalid"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func hasReadAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			explerror.Forbidden(w, errors.New("no token provided"))
			return
		}
		bearerToken := strings.Split(bearer, " ")[1]

		var isAllowed bool = false

		// check if the provided token is a readtoken
		for _, token := range config.Tokens.ReadTokens {
			if token == bearerToken {
				isAllowed = true
				break
			}
		}

		// check if the provided token is a writetoken
		if !isAllowed {
			for _, token := range config.Tokens.ReadWriteTokens {
				if token == bearerToken {
					isAllowed = true
					break
				}
			}
		}

		if !isAllowed {
			explerror.Forbidden(w, errors.New("token invalid"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
