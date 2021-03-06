package debugconfig

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

// NewGetDebugConfigParams creates a new GetDebugConfigParams object
// with the default values initialized.
func NewGetDebugConfigParams() *GetDebugConfigParams {
	var ()
	return &GetDebugConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebugConfigParamsWithTimeout creates a new GetDebugConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDebugConfigParamsWithTimeout(timeout time.Duration) *GetDebugConfigParams {
	var ()
	return &GetDebugConfigParams{

		timeout: timeout,
	}
}

// NewGetDebugConfigParamsWithContext creates a new GetDebugConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDebugConfigParamsWithContext(ctx context.Context) *GetDebugConfigParams {
	var ()
	return &GetDebugConfigParams{

		Context: ctx,
	}
}

// NewGetDebugConfigParamsWithHTTPClient creates a new GetDebugConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDebugConfigParamsWithHTTPClient(client *http.Client) *GetDebugConfigParams {
	var ()
	return &GetDebugConfigParams{
		HTTPClient: client,
	}
}

/*GetDebugConfigParams contains all the parameters to send to the API endpoint
for the get debug config operation typically these are written to a http.Request
*/
type GetDebugConfigParams struct {

	/*DebugConfigID
	  ID of config to return

	*/
	DebugConfigID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get debug config params
func (o *GetDebugConfigParams) WithTimeout(timeout time.Duration) *GetDebugConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debug config params
func (o *GetDebugConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debug config params
func (o *GetDebugConfigParams) WithContext(ctx context.Context) *GetDebugConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debug config params
func (o *GetDebugConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debug config params
func (o *GetDebugConfigParams) WithHTTPClient(client *http.Client) *GetDebugConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debug config params
func (o *GetDebugConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDebugConfigID adds the debugConfigID to the get debug config params
func (o *GetDebugConfigParams) WithDebugConfigID(debugConfigID string) *GetDebugConfigParams {
	o.SetDebugConfigID(debugConfigID)
	return o
}

// SetDebugConfigID adds the debugConfigId to the get debug config params
func (o *GetDebugConfigParams) SetDebugConfigID(debugConfigID string) {
	o.DebugConfigID = debugConfigID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebugConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param debugConfigId
	if err := r.SetPathParam("debugConfigId", o.DebugConfigID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
