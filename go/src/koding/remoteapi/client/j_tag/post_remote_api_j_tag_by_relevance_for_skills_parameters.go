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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJTagByRelevanceForSkillsParams creates a new PostRemoteAPIJTagByRelevanceForSkillsParams object
// with the default values initialized.
func NewPostRemoteAPIJTagByRelevanceForSkillsParams() *PostRemoteAPIJTagByRelevanceForSkillsParams {
	var ()
	return &PostRemoteAPIJTagByRelevanceForSkillsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJTagByRelevanceForSkillsParamsWithTimeout creates a new PostRemoteAPIJTagByRelevanceForSkillsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJTagByRelevanceForSkillsParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJTagByRelevanceForSkillsParams {
	var ()
	return &PostRemoteAPIJTagByRelevanceForSkillsParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJTagByRelevanceForSkillsParamsWithContext creates a new PostRemoteAPIJTagByRelevanceForSkillsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJTagByRelevanceForSkillsParamsWithContext(ctx context.Context) *PostRemoteAPIJTagByRelevanceForSkillsParams {
	var ()
	return &PostRemoteAPIJTagByRelevanceForSkillsParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJTagByRelevanceForSkillsParams contains all the parameters to send to the API endpoint
for the post remote API j tag by relevance for skills operation typically these are written to a http.Request
*/
type PostRemoteAPIJTagByRelevanceForSkillsParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j tag by relevance for skills params
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJTagByRelevanceForSkillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j tag by relevance for skills params
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j tag by relevance for skills params
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) WithContext(ctx context.Context) *PostRemoteAPIJTagByRelevanceForSkillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j tag by relevance for skills params
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j tag by relevance for skills params
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJTagByRelevanceForSkillsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j tag by relevance for skills params
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJTagByRelevanceForSkillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
