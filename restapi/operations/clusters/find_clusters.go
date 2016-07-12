package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FindClustersHandlerFunc turns a function with the right signature into a find clusters handler
type FindClustersHandlerFunc func(FindClustersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindClustersHandlerFunc) Handle(params FindClustersParams) middleware.Responder {
	return fn(params)
}

// FindClustersHandler interface for that can handle valid find clusters params
type FindClustersHandler interface {
	Handle(FindClustersParams) middleware.Responder
}

// NewFindClusters creates a new http.Handler for the find clusters operation
func NewFindClusters(ctx *middleware.Context, handler FindClustersHandler) *FindClusters {
	return &FindClusters{Context: ctx, Handler: handler}
}

/*FindClusters swagger:route GET /clusters clusters findClusters

Returns all clusters that the user is able to access

*/
type FindClusters struct {
	Context *middleware.Context
	Handler FindClustersHandler
}

func (o *FindClusters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFindClustersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*FindClustersOKBodyBody find clusters o k body body

swagger:model FindClustersOKBodyBody
*/
type FindClustersOKBodyBody struct {

	/* clusters

	Required: true
	*/
	Clusters []*ClustersItems0 `json:"clusters"`
}

// Validate validates this find clusters o k body body
func (o *FindClustersOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindClustersOKBodyBody) validateClusters(formats strfmt.Registry) error {

	if err := validate.Required("findClustersOK"+"."+"clusters", "body", o.Clusters); err != nil {
		return err
	}

	for i := 0; i < len(o.Clusters); i++ {

		if swag.IsZero(o.Clusters[i]) { // not required
			continue
		}

		if o.Clusters[i] != nil {

			if err := o.Clusters[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*ClustersItems0 clusters items0

swagger:model ClustersItems0
*/
type ClustersItems0 struct {

	/* URL for the cluster information page in oshinko-rest

	Required: true
	*/
	Href *string `json:"href"`

	/* URL to the spark master

	Required: true
	*/
	MasterURL *string `json:"masterUrl"`

	/* Name of the cluster

	Required: true
	*/
	Name *string `json:"name"`

	/* Current state of the cluster

	Required: true
	*/
	Status *string `json:"status"`

	/* Number of worker nodes in the cluster

	Required: true
	*/
	WorkerCount *int64 `json:"workerCount"`
}

// Validate validates this clusters items0
func (o *ClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHref(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMasterURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateWorkerCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", o.Href); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateMasterURL(formats strfmt.Registry) error {

	if err := validate.Required("masterUrl", "body", o.MasterURL); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateWorkerCount(formats strfmt.Registry) error {

	if err := validate.Required("workerCount", "body", o.WorkerCount); err != nil {
		return err
	}

	return nil
}
