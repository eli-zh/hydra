// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// NewUpdateJSONWebKeySetParams creates a new UpdateJSONWebKeySetParams object
// with the default values initialized.
func NewUpdateJSONWebKeySetParams() *UpdateJSONWebKeySetParams {
	var ()
	return &UpdateJSONWebKeySetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateJSONWebKeySetParamsWithTimeout creates a new UpdateJSONWebKeySetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateJSONWebKeySetParamsWithTimeout(timeout time.Duration) *UpdateJSONWebKeySetParams {
	var ()
	return &UpdateJSONWebKeySetParams{

		timeout: timeout,
	}
}

// NewUpdateJSONWebKeySetParamsWithContext creates a new UpdateJSONWebKeySetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateJSONWebKeySetParamsWithContext(ctx context.Context) *UpdateJSONWebKeySetParams {
	var ()
	return &UpdateJSONWebKeySetParams{

		Context: ctx,
	}
}

// NewUpdateJSONWebKeySetParamsWithHTTPClient creates a new UpdateJSONWebKeySetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateJSONWebKeySetParamsWithHTTPClient(client *http.Client) *UpdateJSONWebKeySetParams {
	var ()
	return &UpdateJSONWebKeySetParams{
		HTTPClient: client,
	}
}

/*UpdateJSONWebKeySetParams contains all the parameters to send to the API endpoint
for the update Json web key set operation typically these are written to a http.Request
*/
type UpdateJSONWebKeySetParams struct {

	/*Body*/
	Body *models.JSONWebKeySet
	/*Set
	  The set

	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) WithTimeout(timeout time.Duration) *UpdateJSONWebKeySetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) WithContext(ctx context.Context) *UpdateJSONWebKeySetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) WithHTTPClient(client *http.Client) *UpdateJSONWebKeySetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) WithBody(body *models.JSONWebKeySet) *UpdateJSONWebKeySetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) SetBody(body *models.JSONWebKeySet) {
	o.Body = body
}

// WithSet adds the set to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) WithSet(set string) *UpdateJSONWebKeySetParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the update Json web key set params
func (o *UpdateJSONWebKeySetParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateJSONWebKeySetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
