package http

import (
	"context"
	"net/http"
)

type HttpRequestContextKey string

func SetStringInContext(r *http.Request, key HttpRequestContextKey, val string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, val))
}

func SetStringListInContext(r *http.Request, key HttpRequestContextKey, val []string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, val))
}

func GetStringFromContext(r *http.Request, key HttpRequestContextKey) string {
	contextValue := r.Context().Value(key)
	str, ok := contextValue.(string)
	if ok {
		return str
	}
	return ""
}

func GetStringListFromContext(r *http.Request, key HttpRequestContextKey) []string {
	contextValue := r.Context().Value(key)
	str, ok := contextValue.([]string)
	if ok {
		return str
	}
	return []string{}
}
