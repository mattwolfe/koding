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

// NewPostRemoteAPIJGroupEachParams creates a new PostRemoteAPIJGroupEachParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupEachParams() *PostRemoteAPIJGroupEachParams {
	var ()
	return &PostRemoteAPIJGroupEachParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupEachParamsWithTimeout creates a new PostRemoteAPIJGroupEachParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupEachParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupEachParams {
	var ()
	return &PostRemoteAPIJGroupEachParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupEachParamsWithContext creates a new PostRemoteAPIJGroupEachParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupEachParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupEachParams {
	var ()
	return &PostRemoteAPIJGroupEachParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupEachParams contains all the parameters to send to the API endpoint
for the post remote API j group each operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupEachParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group each params
func (o *PostRemoteAPIJGroupEachParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupEachParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group each params
func (o *PostRemoteAPIJGroupEachParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group each params
func (o *PostRemoteAPIJGroupEachParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupEachParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group each params
func (o *PostRemoteAPIJGroupEachParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j group each params
func (o *PostRemoteAPIJGroupEachParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJGroupEachParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j group each params
func (o *PostRemoteAPIJGroupEachParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupEachParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
