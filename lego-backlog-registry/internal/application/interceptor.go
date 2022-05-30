package application

import (
	"context"
	"errors"
	"github.com/mercadolibre/lego-backlog-registry/internal/consts"
	"github.com/mercadolibre/lego-backlog-registry/internal/entities"
	"net/http"
)

type RemoteServiceFinder interface {
	FindService(serviceName string) (entities.RemoteService, error)
}

func RemoteServiceInterceptor(finder RemoteServiceFinder, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		processName := r.Header.Get("x-process-name")
		if processName == "" {
			internalHandleError(w, errors.New("header x-process-name not informed"), http.StatusBadRequest)
			return
		}

		remoteService, err := finder.FindService(processName)
		if err != nil {
			internalHandleError(w, err, http.StatusInternalServerError)
			return
		}

		newCtx := context.WithValue(r.Context(), consts.RemoteServiceKey, remoteService)
		next.ServeHTTP(w, r.Clone(newCtx))
	}
}
