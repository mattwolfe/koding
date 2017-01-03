package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJTagCreateSynonymIDParams creates a new PostRemoteAPIJTagCreateSynonymIDParams object
// with the default values initialized.
func NewPostRemoteAPIJTagCreateSynonymIDParams() *PostRemoteAPIJTagCreateSynonymIDParams {
	var ()
	return &PostRemoteAPIJTagCreateSynonymIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJTagCreateSynonymIDParamsWithTimeout creates a new PostRemoteAPIJTagCreateSynonymIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJTagCreateSynonymIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJTagCreateSynonymIDParams {
	var ()
	return &PostRemoteAPIJTagCreateSynonymIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJTagCreateSynonymIDParamsWithContext creates a new PostRemoteAPIJTagCreateSynonymIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJTagCreateSynonymIDParamsWithContext(ctx context.Context) *PostRemoteAPIJTagCreateSynonymIDParams {
	var ()
	return &PostRemoteAPIJTagCreateSynonymIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJTagCreateSynonymIDParams contains all the parameters to send to the API endpoint
for the post remote API j tag create synonym ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJTagCreateSynonymIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j tag create synonym ID params
func (o *PostRemoteAPIJTagCreateSynonymIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJTagCreateSynonymIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j tag create synonym ID params
func (o *PostRemoteAPIJTagCreateSynonymIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j tag create synonym ID params
func (o *PostRemoteAPIJTagCreateSynonymIDParams) WithContext(ctx context.Context) *PostRemoteAPIJTagCreateSynonymIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j tag create synonym ID params
func (o *PostRemoteAPIJTagCreateSynonymIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j tag create synonym ID params
func (o *PostRemoteAPIJTagCreateSynonymIDParams) WithID(id string) *PostRemoteAPIJTagCreateSynonymIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j tag create synonym ID params
func (o *PostRemoteAPIJTagCreateSynonymIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJTagCreateSynonymIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
