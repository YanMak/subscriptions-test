package req

import (
	"net/http"

	"project1.v0/pkg/http/resp"

	httpContractCommons "project1.v0/contracts/transport/http"

	validatorХ "project1.v0/pkg/validator_x"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request, vld *validatorХ.ValidatorX) (*T, error) {

	body, err := Decode[T](r.Body)
	if err != nil {
		code := httpContractCommons.HttpStatusForError(httpContractCommons.ErrRequestValidationError)
		errPayload := httpContractCommons.ErrorResponse{
			Message: err.Error(),
			Code:    code,
		}
		resp.Json(*w, errPayload, code)
		return nil, err
	}

	//err = IsValid(body)
	err = vld.Struct(body)
	if err != nil {
		code := httpContractCommons.HttpStatusForError(httpContractCommons.ErrRequestValidationError)
		errPayload := httpContractCommons.ErrorResponse{
			Message: err.Error(),
			Code:    code,
		}
		resp.Json(*w, errPayload, code)
		return nil, err
	}

	return &body, nil

}
