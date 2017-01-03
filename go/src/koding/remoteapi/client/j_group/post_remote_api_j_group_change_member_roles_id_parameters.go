package j_group

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

// NewPostRemoteAPIJGroupChangeMemberRolesIDParams creates a new PostRemoteAPIJGroupChangeMemberRolesIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupChangeMemberRolesIDParams() *PostRemoteAPIJGroupChangeMemberRolesIDParams {
	var ()
	return &PostRemoteAPIJGroupChangeMemberRolesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupChangeMemberRolesIDParamsWithTimeout creates a new PostRemoteAPIJGroupChangeMemberRolesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupChangeMemberRolesIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupChangeMemberRolesIDParams {
	var ()
	return &PostRemoteAPIJGroupChangeMemberRolesIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupChangeMemberRolesIDParamsWithContext creates a new PostRemoteAPIJGroupChangeMemberRolesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupChangeMemberRolesIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupChangeMemberRolesIDParams {
	var ()
	return &PostRemoteAPIJGroupChangeMemberRolesIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupChangeMemberRolesIDParams contains all the parameters to send to the API endpoint
for the post remote API j group change member roles ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupChangeMemberRolesIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group change member roles ID params
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupChangeMemberRolesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group change member roles ID params
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group change member roles ID params
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupChangeMemberRolesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group change member roles ID params
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j group change member roles ID params
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) WithID(id string) *PostRemoteAPIJGroupChangeMemberRolesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group change member roles ID params
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupChangeMemberRolesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
