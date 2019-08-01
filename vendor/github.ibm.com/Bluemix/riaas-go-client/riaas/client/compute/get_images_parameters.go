// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetImagesParams creates a new GetImagesParams object
// with the default values initialized.
func NewGetImagesParams() *GetImagesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetImagesParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesParamsWithTimeout creates a new GetImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesParamsWithTimeout(timeout time.Duration) *GetImagesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetImagesParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetImagesParamsWithContext creates a new GetImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesParamsWithContext(ctx context.Context) *GetImagesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetImagesParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetImagesParamsWithHTTPClient creates a new GetImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImagesParamsWithHTTPClient(client *http.Client) *GetImagesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetImagesParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetImagesParams contains all the parameters to send to the API endpoint
for the get images operation typically these are written to a http.Request
*/
type GetImagesParams struct {

	/*FutureVersion
	  allows any date string to be accepted, enabling testing of features still under development

	*/
	FutureVersion *bool
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*ResourceGroupID
	  The resource identifier

	*/
	ResourceGroupID *string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*Visibility
	  Filters the collection to images with the specified visibility

	*/
	Visibility *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get images params
func (o *GetImagesParams) WithTimeout(timeout time.Duration) *GetImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images params
func (o *GetImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images params
func (o *GetImagesParams) WithContext(ctx context.Context) *GetImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images params
func (o *GetImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images params
func (o *GetImagesParams) WithHTTPClient(client *http.Client) *GetImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images params
func (o *GetImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFutureVersion adds the futureVersion to the get images params
func (o *GetImagesParams) WithFutureVersion(futureVersion *bool) *GetImagesParams {
	o.SetFutureVersion(futureVersion)
	return o
}

// SetFutureVersion adds the futureVersion to the get images params
func (o *GetImagesParams) SetFutureVersion(futureVersion *bool) {
	o.FutureVersion = futureVersion
}

// WithGeneration adds the generation to the get images params
func (o *GetImagesParams) WithGeneration(generation int64) *GetImagesParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get images params
func (o *GetImagesParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get images params
func (o *GetImagesParams) WithLimit(limit *int32) *GetImagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get images params
func (o *GetImagesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithResourceGroupID adds the resourceGroupID to the get images params
func (o *GetImagesParams) WithResourceGroupID(resourceGroupID *string) *GetImagesParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get images params
func (o *GetImagesParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithStart adds the start to the get images params
func (o *GetImagesParams) WithStart(start *string) *GetImagesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get images params
func (o *GetImagesParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get images params
func (o *GetImagesParams) WithVersion(version string) *GetImagesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get images params
func (o *GetImagesParams) SetVersion(version string) {
	o.Version = version
}

// WithVisibility adds the visibility to the get images params
func (o *GetImagesParams) WithVisibility(visibility *string) *GetImagesParams {
	o.SetVisibility(visibility)
	return o
}

// SetVisibility adds the visibility to the get images params
func (o *GetImagesParams) SetVisibility(visibility *string) {
	o.Visibility = visibility
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FutureVersion != nil {

		// query param future_version
		var qrFutureVersion bool
		if o.FutureVersion != nil {
			qrFutureVersion = *o.FutureVersion
		}
		qFutureVersion := swag.FormatBool(qrFutureVersion)
		if qFutureVersion != "" {
			if err := r.SetQueryParam("future_version", qFutureVersion); err != nil {
				return err
			}
		}

	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.ResourceGroupID != nil {

		// query param resourceGroupID
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resourceGroupID", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if o.Visibility != nil {

		// query param visibility
		var qrVisibility string
		if o.Visibility != nil {
			qrVisibility = *o.Visibility
		}
		qVisibility := qrVisibility
		if qVisibility != "" {
			if err := r.SetQueryParam("visibility", qVisibility); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
