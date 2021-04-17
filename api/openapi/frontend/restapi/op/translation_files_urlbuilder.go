// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// TranslationFilesURL generates an URL for the translation files operation
type TranslationFilesURL struct {
	LocalizationID *int64
	Platform       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TranslationFilesURL) WithBasePath(bp string) *TranslationFilesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TranslationFilesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *TranslationFilesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/translation-files"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/internal"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var localizationIDQ string
	if o.LocalizationID != nil {
		localizationIDQ = swag.FormatInt64(*o.LocalizationID)
	}
	if localizationIDQ != "" {
		qs.Set("localization_id", localizationIDQ)
	}

	var platformQ string
	if o.Platform != nil {
		platformQ = *o.Platform
	}
	if platformQ != "" {
		qs.Set("platform", platformQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *TranslationFilesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *TranslationFilesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *TranslationFilesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on TranslationFilesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on TranslationFilesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *TranslationFilesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
