// Package pf does pretty formatting of http errors
package pf

import (
	"fmt"
	"net/http"
)

//go:generate python generate.py && gofmt -s -w generated.go

type PrettyFail interface {
	PrettyFail() PrettyFail
}

type PrettyFailSrc interface {
	Source() *Fail
}

type Fail struct {
	Slug    string   `json:"code"`
	Code    int      `json:"-"`
	Message string   `json:"message,omitempty"`
	Detail  []string `json:"detail,omitempty"`
	safe    bool
}

func (f *Fail) Error() string {
	if len(f.Detail) > 0 {
		return fmt.Sprintf("HTTP %d: %s: %v", f.Code, f.Message, f.Detail)
	}
	return fmt.Sprintf("HTTP %d: %s", f.Code, f.Message)
}

func (f *Fail) PrettyFail() PrettyFail {
	if f.Message == "" || !f.safe {
		f.Message = http.StatusText(f.Code)
	}
	return f
}

func (f *Fail) Safe() *Fail {
	f.safe = true
	return f
}

func (f *Fail) Source() *Fail {
	return f
}

func formatMessage(msg ...interface{}) string {
	if len(msg) == 0 {
		return ""
	} else if len(msg) == 1 {
		return fmt.Sprintf("%v", msg[0])
	} else {
		format, ok := msg[0].(string)
		if ok {
			return fmt.Sprintf(format, msg[1:])
		}
		return fmt.Sprintf("%v", msg)
	}
}

// func BadRequest(msg ...interface{}) *Fail {
// 	message := formatMessage(msg...)
// 	return &Fail{
// 		Message: message,
// 	}
// }
//
// func BadRequestDetail(details []string) *Fail {
// 	return &Fail{
// 		Detail: details,
// 	}
// }
