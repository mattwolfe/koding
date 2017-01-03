package j_provisioner

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

// NewPostRemoteAPIJProvisionerDeleteIDParams creates a new PostRemoteAPIJProvisionerDeleteIDParams object
// with the default values initialized.
func NewPostRemoteAPIJProvisionerDeleteIDParams() *PostRemoteAPIJProvisionerDeleteIDParams {
	var ()
	return &PostRemoteAPIJProvisionerDeleteIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProvisionerDeleteIDParamsWithTimeout creates a new PostRemoteAPIJProvisionerDeleteIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProvisionerDeleteIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerDeleteIDParams {
	var ()
	return &PostRemoteAPIJProvisionerDeleteIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProvisionerDeleteIDParamsWithContext creates a new PostRemoteAPIJProvisionerDeleteIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProvisionerDeleteIDParamsWithContext(ctx context.Context) *PostRemoteAPIJProvisionerDeleteIDParams {
	var ()
	return &PostRemoteAPIJProvisionerDeleteIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProvisionerDeleteIDParams contains all the parameters to send to the API endpoint
for the post remote API j provisioner delete ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJProvisionerDeleteIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j provisioner delete ID params
func (o *PostRemoteAPIJProvisionerDeleteIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerDeleteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j provisioner delete ID params
func (o *PostRemoteAPIJProvisionerDeleteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j provisioner delete ID params
func (o *PostRemoteAPIJProvisionerDeleteIDParams) WithContext(ctx context.Context) *PostRemoteAPIJProvisionerDeleteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j provisioner delete ID params
func (o *PostRemoteAPIJProvisionerDeleteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j provisioner delete ID params
func (o *PostRemoteAPIJProvisionerDeleteIDParams) WithID(id string) *PostRemoteAPIJProvisionerDeleteIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j provisioner delete ID params
func (o *PostRemoteAPIJProvisionerDeleteIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProvisionerDeleteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
