package helpers

import (
	"errors"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"net/http"
	"strings"
)

func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func HandleCommonRepositoryError(err error) (int, error) {
	if strings.EqualFold(err.Error(), constants.ErrDBNoRowsFound.Error()) {
		return http.StatusNotFound, errors.New("no rows found")
	}
	return http.StatusInternalServerError, err
}

func TrueValuePointer() *bool {
	value := true
	return &value
}

func FalseValuePointer() *bool {
	value := false
	return &value
}

var BooleanToInteger = map[bool]int8{false: 0, true: 1}
var IntegerToBoolean = map[int8]*bool{0: FalseValuePointer(), 1: TrueValuePointer()}
