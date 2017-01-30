package operations

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

// NewThreadGetOneParams creates a new ThreadGetOneParams object
// with the default values initialized.
func NewThreadGetOneParams() *ThreadGetOneParams {
	var ()
	return &ThreadGetOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThreadGetOneParamsWithTimeout creates a new ThreadGetOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThreadGetOneParamsWithTimeout(timeout time.Duration) *ThreadGetOneParams {
	var ()
	return &ThreadGetOneParams{

		timeout: timeout,
	}
}

// NewThreadGetOneParamsWithContext creates a new ThreadGetOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewThreadGetOneParamsWithContext(ctx context.Context) *ThreadGetOneParams {
	var ()
	return &ThreadGetOneParams{

		Context: ctx,
	}
}

// NewThreadGetOneParamsWithHTTPClient creates a new ThreadGetOneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThreadGetOneParamsWithHTTPClient(client *http.Client) *ThreadGetOneParams {
	var ()
	return &ThreadGetOneParams{
		HTTPClient: client,
	}
}

/*ThreadGetOneParams contains all the parameters to send to the API endpoint
for the thread get one operation typically these are written to a http.Request
*/
type ThreadGetOneParams struct {

	/*SlugOrID
	  Идентификатор ветки обсуждения.

	*/
	SlugOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the thread get one params
func (o *ThreadGetOneParams) WithTimeout(timeout time.Duration) *ThreadGetOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thread get one params
func (o *ThreadGetOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thread get one params
func (o *ThreadGetOneParams) WithContext(ctx context.Context) *ThreadGetOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thread get one params
func (o *ThreadGetOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thread get one params
func (o *ThreadGetOneParams) WithHTTPClient(client *http.Client) *ThreadGetOneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thread get one params
func (o *ThreadGetOneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlugOrID adds the slugOrID to the thread get one params
func (o *ThreadGetOneParams) WithSlugOrID(slugOrID string) *ThreadGetOneParams {
	o.SetSlugOrID(slugOrID)
	return o
}

// SetSlugOrID adds the slugOrId to the thread get one params
func (o *ThreadGetOneParams) SetSlugOrID(slugOrID string) {
	o.SlugOrID = slugOrID
}

// WriteToRequest writes these params to a swagger request
func (o *ThreadGetOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param slug_or_id
	if err := r.SetPathParam("slug_or_id", o.SlugOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
