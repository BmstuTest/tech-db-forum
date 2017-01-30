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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostGetOneParams creates a new PostGetOneParams object
// with the default values initialized.
func NewPostGetOneParams() *PostGetOneParams {
	var ()
	return &PostGetOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGetOneParamsWithTimeout creates a new PostGetOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGetOneParamsWithTimeout(timeout time.Duration) *PostGetOneParams {
	var ()
	return &PostGetOneParams{

		timeout: timeout,
	}
}

// NewPostGetOneParamsWithContext creates a new PostGetOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGetOneParamsWithContext(ctx context.Context) *PostGetOneParams {
	var ()
	return &PostGetOneParams{

		Context: ctx,
	}
}

// NewPostGetOneParamsWithHTTPClient creates a new PostGetOneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGetOneParamsWithHTTPClient(client *http.Client) *PostGetOneParams {
	var ()
	return &PostGetOneParams{
		HTTPClient: client,
	}
}

/*PostGetOneParams contains all the parameters to send to the API endpoint
for the post get one operation typically these are written to a http.Request
*/
type PostGetOneParams struct {

	/*ID
	  Идентификатор сообщения.

	*/
	ID int64
	/*Related
	  Включение полной информации о соответвующем объекте сообщения.

	Если тип объекта не указан, то полная информация об этих объектах не
	передаётся.


	*/
	Related []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post get one params
func (o *PostGetOneParams) WithTimeout(timeout time.Duration) *PostGetOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post get one params
func (o *PostGetOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post get one params
func (o *PostGetOneParams) WithContext(ctx context.Context) *PostGetOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post get one params
func (o *PostGetOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post get one params
func (o *PostGetOneParams) WithHTTPClient(client *http.Client) *PostGetOneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post get one params
func (o *PostGetOneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post get one params
func (o *PostGetOneParams) WithID(id int64) *PostGetOneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post get one params
func (o *PostGetOneParams) SetID(id int64) {
	o.ID = id
}

// WithRelated adds the related to the post get one params
func (o *PostGetOneParams) WithRelated(related []string) *PostGetOneParams {
	o.SetRelated(related)
	return o
}

// SetRelated adds the related to the post get one params
func (o *PostGetOneParams) SetRelated(related []string) {
	o.Related = related
}

// WriteToRequest writes these params to a swagger request
func (o *PostGetOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	valuesRelated := o.Related

	joinedRelated := swag.JoinByFormat(valuesRelated, "")
	// query array param related
	if err := r.SetQueryParam("related", joinedRelated...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
