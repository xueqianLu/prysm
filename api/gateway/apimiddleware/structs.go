package apimiddleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v3/encoding/bytesutil"
)

type HexString string

func (b *HexString) UnmarshalJSON(enc []byte) error {
	if len(enc) == 0 {
		// Empty hex values are represented as "0x".
		*b = "0x"
		return nil
	}
	var s string
	if err := json.Unmarshal(enc, &s); err != nil {
		return err
	}
	if bytesutil.IsHex([]byte(s)) {
		*b = HexString(s)
		return nil
	}
	if s == "" {
		// Empty hex values are represented as "0x".
		*b = "0x"
		return nil
	}
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*b = HexString(hexutil.Encode(bytes))
	return nil
}

// ---------------
// Error handling.
// ---------------

// ErrorJson describes common functionality of all JSON error representations.
type ErrorJson interface {
	StatusCode() int
	SetCode(code int)
	Msg() string
	SetMsg(msg string)
}

// DefaultErrorJson is a JSON representation of a simple error value, containing only a message and an error code.
type DefaultErrorJson struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// InternalServerErrorWithMessage returns a DefaultErrorJson with 500 code and a custom message.
func InternalServerErrorWithMessage(err error, message string) *DefaultErrorJson {
	e := errors.Wrapf(err, message)
	return &DefaultErrorJson{
		Message: e.Error(),
		Code:    http.StatusInternalServerError,
	}
}

// InternalServerError returns a DefaultErrorJson with 500 code.
func InternalServerError(err error) *DefaultErrorJson {
	return &DefaultErrorJson{
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	}
}

func TimeoutError() *DefaultErrorJson {
	return &DefaultErrorJson{
		Message: "Request timeout",
		Code:    http.StatusRequestTimeout,
	}
}

// StatusCode returns the error's underlying error code.
func (e *DefaultErrorJson) StatusCode() int {
	return e.Code
}

// Msg returns the error's underlying message.
func (e *DefaultErrorJson) Msg() string {
	return e.Message
}

// SetCode sets the error's underlying error code.
func (e *DefaultErrorJson) SetCode(code int) {
	e.Code = code
}

// SetMsg sets the error's underlying message.
func (e *DefaultErrorJson) SetMsg(msg string) {
	e.Message = msg
}
