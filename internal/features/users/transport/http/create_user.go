package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
	core_logger "github.com/AmongAster/golang-todoapp/internal/core/logger"
	core_http_request "github.com/AmongAster/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/AmongAster/golang-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

// POST /users
// request_id

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "faled to decode and valodate HTTP request")

		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUsers(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("Произошла ошибка...")
	}

	rw.WriteHeader(http.StatusOK)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitializer(dto.FullName, dto.PhoneNumber)
}
