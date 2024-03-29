package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AliIsmoilov/golang_monolight/pkg/httpErrors"
	"github.com/AliIsmoilov/golang_monolight/pkg/logger"
	"github.com/AliIsmoilov/golang_monolight/pkg/sanitize"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get ctx with timeout and request id from echo context
func GetCtxWithReqID(c echo.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// Get context  with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

// UserCtxKey is a key used for the User object in the context
type UserCtxKey struct{}

// Get user from context
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

// Error response with logging error for echo context
func ErrResponseWithLog(ctx echo.Context, logger logger.Logger, err error) error {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
	return ctx.JSON(httpErrors.ErrorResponse(err))
}

// Error response with logging error for echo context
func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}

// Read request body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

// Read sanitize and validate request
func SanitizeRequest(ctx echo.Context, request interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer ctx.Request().Body.Close()

	sanBody, err := sanitize.SanitizeJSON(body)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err = json.Unmarshal(sanBody, request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}

var allowedImagesContentTypes = map[string]string{
	"image/bmp":                "bmp",
	"image/gif":                "gif",
	"image/png":                "png",
	"image/jpeg":               "jpeg",
	"image/jpg":                "jpg",
	"image/svg+xml":            "svg",
	"image/webp":               "webp",
	"image/tiff":               "tiff",
	"image/vnd.microsoft.icon": "ico",
}

func CheckImageFileContentType(fileContent []byte) (string, error) {
	contentType := http.DetectContentType(fileContent)

	extension, ok := allowedImagesContentTypes[contentType]
	if !ok {
		return "", errors.New("this content type is not allowed")
	}

	return extension, nil
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}

func HandleUnauthorizedWithMessage(w http.ResponseWriter, message string) {
	w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
	w.WriteHeader(http.StatusUnauthorized)

	log.Println(message)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusUnauthorized,
			Message: message,
		}})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "  ")

	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)
}

type response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// errorInfo ...
type errorInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleForbiddenErrWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Println(message+" ", err)
	w.WriteHeader(http.StatusForbidden)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusForbidden,
			Message: message + ": " + err.Error(),
		}})
	return err
}
