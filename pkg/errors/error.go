package errors

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/language"
)

type Error struct {
	Message  string
	HttpCode int
	Code     int
	RefID    uuid.UUID

	originalError error
	translates    map[language.Tag]string
	details       map[string]string
}

type jsonError struct {
	RefID   uuid.UUID
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) AddOriginalError(err error) *Error {
	e.originalError = err
	return e
}

func (e *Error) Original() error {
	if e.originalError != nil {
		return e.originalError
	}
	return nil
}

func (e *Error) AddDetail(key string, detail interface{}) *Error {
	if e.details == nil {
		e.details = make(map[string]string)
	}

	if v, ok := e.details[key]; !ok {
		e.details[key] = fmt.Sprintf("%v", detail)
	} else {
		e.details[key] = fmt.Sprintf("%v - %v", v, detail)
	}

	return e
}

func (e *Error) Details() map[string]string {
	if e.details == nil {
		e.details = make(map[string]string)
	}
	return e.details
}

func (e *Error) Json(ctx context.Context) string {
	lang := ctx.Value("language")
	jsonErr := new(jsonError)
	if lang != nil {
		langTag, err := language.Parse(lang.(string))
		if err == nil {
			jsonErr.Message = e.translates[langTag]
		}
	}
	if jsonErr.Message == "" {
		jsonErr.Message = e.Message
	}
	jsonErr.Code = e.Code

	//todo: save error based on ref id
	jsonErr.RefID = e.RefID

	str, err := json.Marshal(jsonErr)
	if err != nil {
		return ""
	}
	return string(str)
}
