// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// PutDeviceNameAddrURL generates an URL for the put device name addr operation
type PutDeviceNameAddrURL struct {
	Name string

	Sapool   *bool
	Snapshot *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutDeviceNameAddrURL) WithBasePath(bp string) *PutDeviceNameAddrURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutDeviceNameAddrURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PutDeviceNameAddrURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/device/{name}/addr"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on PutDeviceNameAddrURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var sapoolQ string
	if o.Sapool != nil {
		sapoolQ = swag.FormatBool(*o.Sapool)
	}
	if sapoolQ != "" {
		qs.Set("sapool", sapoolQ)
	}

	var snapshotQ string
	if o.Snapshot != nil {
		snapshotQ = swag.FormatBool(*o.Snapshot)
	}
	if snapshotQ != "" {
		qs.Set("snapshot", snapshotQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PutDeviceNameAddrURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PutDeviceNameAddrURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PutDeviceNameAddrURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PutDeviceNameAddrURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PutDeviceNameAddrURL")
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
func (o *PutDeviceNameAddrURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}