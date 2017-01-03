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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJGroupSomeWithRelationshipParams creates a new PostRemoteAPIJGroupSomeWithRelationshipParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupSomeWithRelationshipParams() *PostRemoteAPIJGroupSomeWithRelationshipParams {
	var ()
	return &PostRemoteAPIJGroupSomeWithRelationshipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupSomeWithRelationshipParamsWithTimeout creates a new PostRemoteAPIJGroupSomeWithRelationshipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupSomeWithRelationshipParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupSomeWithRelationshipParams {
	var ()
	return &PostRemoteAPIJGroupSomeWithRelationshipParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupSomeWithRelationshipParamsWithContext creates a new PostRemoteAPIJGroupSomeWithRelationshipParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupSomeWithRelationshipParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupSomeWithRelationshipParams {
	var ()
	return &PostRemoteAPIJGroupSomeWithRelationshipParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupSomeWithRelationshipParams contains all the parameters to send to the API endpoint
for the post remote API j group some with relationship operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupSomeWithRelationshipParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group some with relationship params
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupSomeWithRelationshipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group some with relationship params
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group some with relationship params
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupSomeWithRelationshipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group some with relationship params
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j group some with relationship params
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJGroupSomeWithRelationshipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j group some with relationship params
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupSomeWithRelationshipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
