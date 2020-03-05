// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudCloudinstancesDeleteParams creates a new PcloudCloudinstancesDeleteParams object
// with the default values initialized.
func NewPcloudCloudinstancesDeleteParams() *PcloudCloudinstancesDeleteParams {
	var ()
	return &PcloudCloudinstancesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesDeleteParamsWithTimeout creates a new PcloudCloudinstancesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesDeleteParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesDeleteParams {
	var ()
	return &PcloudCloudinstancesDeleteParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesDeleteParamsWithContext creates a new PcloudCloudinstancesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesDeleteParamsWithContext(ctx context.Context) *PcloudCloudinstancesDeleteParams {
	var ()
	return &PcloudCloudinstancesDeleteParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesDeleteParamsWithHTTPClient creates a new PcloudCloudinstancesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesDeleteParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesDeleteParams {
	var ()
	return &PcloudCloudinstancesDeleteParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesDeleteParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances delete operation typically these are written to a http.Request
*/
type PcloudCloudinstancesDeleteParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) WithContext(ctx context.Context) *PcloudCloudinstancesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances delete params
func (o *PcloudCloudinstancesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
