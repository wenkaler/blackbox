package server

import (
	"context"
	"io/ioutil"
	"net/http"
)

func decodeByteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return b, nil
}

func decodeToken(_ context.Context, r *http.Request) (interface{}, error) {
	token := r.Header.Get("token")
	return token, nil
}

func decodeTokenVarious(_ context.Context, r *http.Request) (interface{}, error) {
	token := r.Header.Get("token")
	various := r.Header.Get("various")
	return requestFTV{Token: token, Various: various}, nil
}
