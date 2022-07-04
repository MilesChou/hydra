// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAdminDeleteJSONWebKeyParams creates a new AdminDeleteJSONWebKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminDeleteJSONWebKeyParams() *AdminDeleteJSONWebKeyParams {
	return &AdminDeleteJSONWebKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteJSONWebKeyParamsWithTimeout creates a new AdminDeleteJSONWebKeyParams object
// with the ability to set a timeout on a request.
func NewAdminDeleteJSONWebKeyParamsWithTimeout(timeout time.Duration) *AdminDeleteJSONWebKeyParams {
	return &AdminDeleteJSONWebKeyParams{
		timeout: timeout,
	}
}

// NewAdminDeleteJSONWebKeyParamsWithContext creates a new AdminDeleteJSONWebKeyParams object
// with the ability to set a context for a request.
func NewAdminDeleteJSONWebKeyParamsWithContext(ctx context.Context) *AdminDeleteJSONWebKeyParams {
	return &AdminDeleteJSONWebKeyParams{
		Context: ctx,
	}
}

// NewAdminDeleteJSONWebKeyParamsWithHTTPClient creates a new AdminDeleteJSONWebKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminDeleteJSONWebKeyParamsWithHTTPClient(client *http.Client) *AdminDeleteJSONWebKeyParams {
	return &AdminDeleteJSONWebKeyParams{
		HTTPClient: client,
	}
}

/* AdminDeleteJSONWebKeyParams contains all the parameters to send to the API endpoint
   for the admin delete Json web key operation.

   Typically these are written to a http.Request.
*/
type AdminDeleteJSONWebKeyParams struct {

	/* Kid.

	   The JSON Web Key ID (kid)
	*/
	Kid string

	/* Set.

	   The JSON Web Key Set
	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin delete Json web key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminDeleteJSONWebKeyParams) WithDefaults() *AdminDeleteJSONWebKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin delete Json web key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminDeleteJSONWebKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) WithTimeout(timeout time.Duration) *AdminDeleteJSONWebKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) WithContext(ctx context.Context) *AdminDeleteJSONWebKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) WithHTTPClient(client *http.Client) *AdminDeleteJSONWebKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKid adds the kid to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) WithKid(kid string) *AdminDeleteJSONWebKeyParams {
	o.SetKid(kid)
	return o
}

// SetKid adds the kid to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) SetKid(kid string) {
	o.Kid = kid
}

// WithSet adds the set to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) WithSet(set string) *AdminDeleteJSONWebKeyParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the admin delete Json web key params
func (o *AdminDeleteJSONWebKeyParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteJSONWebKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kid
	if err := r.SetPathParam("kid", o.Kid); err != nil {
		return err
	}

	// path param set
	if err := r.SetPathParam("set", o.Set); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}