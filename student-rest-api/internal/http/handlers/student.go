package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"githug.com/tarunagg1/student-api/internal/types"
	"githug.com/tarunagg1/student-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
		}

		// validator
		// validate := validator.New(validator.WithRequiredStructEnabled())

		// request validation
		if err := validator.New().Struct(student); err != nil {
			// respoe
			validateErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"sucess": "OK"})
	}
}
