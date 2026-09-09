package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/AmongAster/golang-todoapp/internal/core/errors"
)

func GetIntPathValu(r *http.Request, key string) (int, error) {
	pathValu := r.PathValue(key)
	if pathValu == "" {
		return 0, fmt.Errorf("no key '%s' in path valus: %w", key, core_errors.ErrinvalidArgument)
	}

	val, err := strconv.Atoi(pathValu)

	if err != nil {
		return 0, fmt.Errorf("path value '%s' by key=%s not a valid int:%v: %w", pathValu, key, err, core_errors.ErrinvalidArgument)
	}

	return val, nil
}
