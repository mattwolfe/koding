package j_account

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

// NewPostRemoteAPIJAccountUnblockUserIDParams creates a new PostRemoteAPIJAccountUnblockUserIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountUnblockUserIDParams() *PostRemoteAPIJAccountUnblockUserIDParams {
	var ()
	return &PostRemoteAPIJAccountUnblockUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountUnblockUserIDParamsWithTimeout creates a new PostRemoteAPIJAccountUnblockUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountUnblockUserIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountUnblockUserIDParams {
	var ()
	return &PostRemoteAPIJAccountUnblockUserIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountUnblockUserIDParamsWithContext creates a new PostRemoteAPIJAccountUnblockUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountUnblockUserIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountUnblockUserIDParams {
	var ()
	return &PostRemoteAPIJAccountUnblockUserIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountUnblockUserIDParams contains all the parameters to send to the API endpoint
for the post remote API j account unblock user ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountUnblockUserIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j account unblock user ID params
func (o *PostRemoteAPIJAccountUnblockUserIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountUnblockUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account unblock user ID params
func (o *PostRemoteAPIJAccountUnblockUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account unblock user ID params
func (o *PostRemoteAPIJAccountUnblockUserIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountUnblockUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account unblock user ID params
func (o *PostRemoteAPIJAccountUnblockUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account unblock user ID params
func (o *PostRemoteAPIJAccountUnblockUserIDParams) WithID(id string) *PostRemoteAPIJAccountUnblockUserIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account unblock user ID params
func (o *PostRemoteAPIJAccountUnblockUserIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountUnblockUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
