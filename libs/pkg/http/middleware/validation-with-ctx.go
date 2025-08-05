package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"project1.v0/pkg/decoder"
	validatorХ "project1.v0/pkg/validator_x"
)

type ctxKey[T any] struct{}

func GetFromContext[T any](ctx context.Context) (*T, bool) {
	val, ok := ctx.Value(ctxKey[T]{}).(*T)
	return val, ok
}

func HttpValidationWithCtx[BodyT any, QueryT any, PathT any](vld *validatorХ.ValidatorX, dcd *decoder.DecoderService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// 1. Body
			if typ := reflect.TypeOf((*BodyT)(nil)).Elem(); typ != reflect.TypeOf(struct{}{}) {
				var bodyT BodyT
				if err := decodeBody(r, &bodyT); err != nil {
					http.Error(w, "invalid body: "+err.Error(), http.StatusBadRequest)
					return
				}
				if err := vld.Struct(bodyT); err != nil {
					http.Error(w, "body validation: "+err.Error(), http.StatusBadRequest)
					return
				}
				ctx = context.WithValue(ctx, ctxKey[BodyT]{}, &bodyT)
			}

			// 2. Query
			if typ := reflect.TypeOf((*QueryT)(nil)).Elem(); typ != reflect.TypeOf(struct{}{}) {
				var queryT QueryT
				if err := dcd.MapQueryToStruct(r.URL.Query(), &queryT); err != nil {
					http.Error(w, "query parse: "+err.Error(), http.StatusBadRequest)
					return
				}
				if err := vld.Struct(queryT); err != nil {
					http.Error(w, "query validation: "+err.Error(), http.StatusBadRequest)
					return
				}
				ctx = context.WithValue(ctx, ctxKey[QueryT]{}, &queryT)
			}

			// 3. Path
			if typ := reflect.TypeOf((*PathT)(nil)).Elem(); typ != reflect.TypeOf(struct{}{}) {
				var pathT PathT
				pathParams := extractPathParams(r)
				if err := dcd.MapPathParamsToStruct(pathParams, &pathT); err != nil {
					http.Error(w, "path parse: "+err.Error(), http.StatusBadRequest)
					return
				}
				if err := vld.Struct(pathT); err != nil {
					http.Error(w, "path validation: "+err.Error(), http.StatusBadRequest)
					return
				}
				ctx = context.WithValue(ctx, ctxKey[PathT]{}, &pathT)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func decodeBody[T any](r *http.Request, out *T) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(out)
}

func extractPathParams(r *http.Request) map[string]string {
	// Реализуй под свой роутер (например, chi или gorilla/mux)
	return map[string]string{}
}
