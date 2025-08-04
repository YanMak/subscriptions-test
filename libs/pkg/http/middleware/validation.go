package middleware

import (
	"context"
	"net/http"

	"project1.v0/pkg/http/req"
	validatorХ "project1.v0/pkg/validator_x"
)

type requestDtoContextKey struct{}

func GetRequestDto[T interface{}](ctx context.Context) (*T, bool) {
	dto, ok := ctx.Value(requestDtoContextKey{}).(*T)
	return dto, ok
}

func HttpValidation[T interface{}](vld *validatorХ.ValidatorX) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				reqPayload, err := req.HandleBody[T](&w, r, vld)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				ctx := context.WithValue(r.Context(), requestDtoContextKey{}, reqPayload)

				// Call next
				next.ServeHTTP(w, r.WithContext(ctx))

			})
	}
}

// func ValidationMiddleware[T interface{}](next http.Handler) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {

// 			reqPayload, err := req.HandleBody[T](&w, r)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusBadRequest)
// 				return
// 			}

// 			ctx := context.WithValue(r.Context(), requestDtoContextKey{}, reqPayload)

// 			// Call next
// 			next.ServeHTTP(w, r.WithContext(ctx))

// 		})
// }
