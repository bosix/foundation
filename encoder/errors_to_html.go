package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support/str"
)

type ErrorToHtml struct {
	Errors []Error `json:"errors"`
}

func (e ErrorToHtml) IsAble(object interface{}) bool {
	_, ok := object.(error)
	return ok
}

func (e ErrorToHtml) EncodeThrough(object interface{}, _ []inter.Encoder) (string, error) {
	err, ok := object.(error)
	if !ok {
		return "", errors.New("can't convert object to html in error format")
	}

	return str.UpperFirst(err.Error()), nil
}