//go:build go1.22

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.0 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/oapi-codegen/runtime"
)

// Defines values for ErrorCode.
const (
	ParameterInvalid      ErrorCode = "parameter_invalid"
	ParameterMissing      ErrorCode = "parameter_missing"
	ProcessingError       ErrorCode = "processing_error"
	ResourceAlreadyExists ErrorCode = "resource_already_exists"
	ResourceMissing       ErrorCode = "resource_missing"
)

// Defines values for ErrorType.
const (
	ApiError            ErrorType = "api_error"
	InvalidRequestError ErrorType = "invalid_request_error"
)

// Bookmark Bookmark is a resource that represents a saved URL.
type Bookmark struct {
	CreateTime  *time.Time `json:"createTime,omitempty"`
	DisplayName *string    `json:"displayName,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Tags        []string   `json:"tags"`
	Uid         *string    `json:"uid,omitempty"`
	UpdateTime  *time.Time `json:"updateTime,omitempty"`
	Url         string     `json:"url"`
}

// BookmarkCreate Resource create operation model.
type BookmarkCreate struct {
	DisplayName *string  `json:"displayName,omitempty"`
	Tags        []string `json:"tags"`
	Url         string   `json:"url"`
}

// Error Error is the response model when an API call is unsuccessful.
type Error struct {
	Code    ErrorCode              `json:"code"`
	Details map[string]interface{} `json:"details"`
	Message string                 `json:"message"`
	Type    ErrorType              `json:"type"`
}

// ErrorCode defines model for ErrorCode.
type ErrorCode string

// ErrorType defines model for ErrorType.
type ErrorType string

// BookmarkParentKeyCollectionId defines model for BookmarkParentKey.collectionId.
type BookmarkParentKeyCollectionId = string

// BookmarkParentKeyUserId defines model for BookmarkParentKey.userId.
type BookmarkParentKeyUserId = string

// BookmarksCreateJSONRequestBody defines body for BookmarksCreate for application/json ContentType.
type BookmarksCreateJSONRequestBody = BookmarkCreate

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /users/{userId}/collections/{collectionId}/bookmarks)
	BookmarksCreate(w http.ResponseWriter, r *http.Request, userId BookmarkParentKeyUserId, collectionId BookmarkParentKeyCollectionId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// BookmarksCreate operation middleware
func (siw *ServerInterfaceWrapper) BookmarksCreate(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId BookmarkParentKeyUserId

	err = runtime.BindStyledParameterWithOptions("simple", "userId", r.PathValue("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	// ------------- Path parameter "collectionId" -------------
	var collectionId BookmarkParentKeyCollectionId

	err = runtime.BindStyledParameterWithOptions("simple", "collectionId", r.PathValue("collectionId"), &collectionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "collectionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.BookmarksCreate(w, r, userId, collectionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("POST "+options.BaseURL+"/users/{userId}/collections/{collectionId}/bookmarks", wrapper.BookmarksCreate)

	return m
}