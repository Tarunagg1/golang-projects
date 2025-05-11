package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Tarunagg1/student-api/internal/types"
	"github.com/Tarunagg1/student-api/internal/utils/response"
	"github.com/go-playground/validator"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)

			response.WriteJson(w, http.StatusBadRequest, response.GenralError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"Success": "Ok"})
	}
}
