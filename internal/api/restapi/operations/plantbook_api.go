// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kaatinga/plantbook/internal/api/restapi/operations/health"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/user"
)

// NewPlantbookAPI creates a new Plantbook instance
func NewPlantbookAPI(spec *loads.Document) *PlantbookAPI {
	return &PlantbookAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),
		TxtProducer:  runtime.TextProducer(),

		HealthGetMetricsHandler: health.GetMetricsHandlerFunc(func(params health.GetMetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation health.GetMetrics has not yet been implemented")
		}),
		HealthAPIVersionHandler: health.APIVersionHandlerFunc(func(params health.APIVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation health.APIVersion has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		}),
		PlantGetPlantByIDHandler: plant.GetPlantByIDHandlerFunc(func(params plant.GetPlantByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.GetPlantByID has not yet been implemented")
		}),
		PlantGetPlantsHandler: plant.GetPlantsHandlerFunc(func(params plant.GetPlantsParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.GetPlants has not yet been implemented")
		}),
		UserGetUserByNameHandler: user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
		}),
		HealthHealthAliveHandler: health.HealthAliveHandlerFunc(func(params health.HealthAliveParams) middleware.Responder {
			return middleware.NotImplemented("operation health.HealthAlive has not yet been implemented")
		}),
		HealthHealthReadyHandler: health.HealthReadyHandlerFunc(func(params health.HealthReadyParams) middleware.Responder {
			return middleware.NotImplemented("operation health.HealthReady has not yet been implemented")
		}),
		UserLoginUserHandler: user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		}),
		UserLogoutUserHandler: user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		}),
		UserUpdateUserHandler: user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		}),
	}
}

/*PlantbookAPI This is a Plantbook project API description. You can find out more about us at [https://github.com/kaatinga/plantbook](https://github.com/kaatinga/plantbook). */
type PlantbookAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TxtProducer registers a producer for the following mime types:
	//   - text/plain
	TxtProducer runtime.Producer

	// HealthGetMetricsHandler sets the operation handler for the get metrics operation
	HealthGetMetricsHandler health.GetMetricsHandler
	// HealthAPIVersionHandler sets the operation handler for the api version operation
	HealthAPIVersionHandler health.APIVersionHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// PlantGetPlantByIDHandler sets the operation handler for the get plant by Id operation
	PlantGetPlantByIDHandler plant.GetPlantByIDHandler
	// PlantGetPlantsHandler sets the operation handler for the get plants operation
	PlantGetPlantsHandler plant.GetPlantsHandler
	// UserGetUserByNameHandler sets the operation handler for the get user by name operation
	UserGetUserByNameHandler user.GetUserByNameHandler
	// HealthHealthAliveHandler sets the operation handler for the health alive operation
	HealthHealthAliveHandler health.HealthAliveHandler
	// HealthHealthReadyHandler sets the operation handler for the health ready operation
	HealthHealthReadyHandler health.HealthReadyHandler
	// UserLoginUserHandler sets the operation handler for the login user operation
	UserLoginUserHandler user.LoginUserHandler
	// UserLogoutUserHandler sets the operation handler for the logout user operation
	UserLogoutUserHandler user.LogoutUserHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *PlantbookAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *PlantbookAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *PlantbookAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *PlantbookAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *PlantbookAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *PlantbookAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *PlantbookAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *PlantbookAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *PlantbookAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the PlantbookAPI
func (o *PlantbookAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.HealthGetMetricsHandler == nil {
		unregistered = append(unregistered, "health.GetMetricsHandler")
	}
	if o.HealthAPIVersionHandler == nil {
		unregistered = append(unregistered, "health.APIVersionHandler")
	}
	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}
	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}
	if o.PlantGetPlantByIDHandler == nil {
		unregistered = append(unregistered, "plant.GetPlantByIDHandler")
	}
	if o.PlantGetPlantsHandler == nil {
		unregistered = append(unregistered, "plant.GetPlantsHandler")
	}
	if o.UserGetUserByNameHandler == nil {
		unregistered = append(unregistered, "user.GetUserByNameHandler")
	}
	if o.HealthHealthAliveHandler == nil {
		unregistered = append(unregistered, "health.HealthAliveHandler")
	}
	if o.HealthHealthReadyHandler == nil {
		unregistered = append(unregistered, "health.HealthReadyHandler")
	}
	if o.UserLoginUserHandler == nil {
		unregistered = append(unregistered, "user.LoginUserHandler")
	}
	if o.UserLogoutUserHandler == nil {
		unregistered = append(unregistered, "user.LogoutUserHandler")
	}
	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *PlantbookAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *PlantbookAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *PlantbookAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *PlantbookAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *PlantbookAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/plain":
			result["text/plain"] = o.TxtProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *PlantbookAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the plantbook API
func (o *PlantbookAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *PlantbookAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metrics"] = health.NewGetMetrics(o.context, o.HealthGetMetricsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/version"] = health.NewAPIVersion(o.context, o.HealthAPIVersionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1/user"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/v1/user/{username}"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/plant/{id}"] = plant.NewGetPlantByID(o.context, o.PlantGetPlantByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/plant"] = plant.NewGetPlants(o.context, o.PlantGetPlantsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/user/{username}"] = user.NewGetUserByName(o.context, o.UserGetUserByNameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health/live"] = health.NewHealthAlive(o.context, o.HealthHealthAliveHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health/ready"] = health.NewHealthReady(o.context, o.HealthHealthReadyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1/user/login"] = user.NewLoginUser(o.context, o.UserLoginUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/user/logout"] = user.NewLogoutUser(o.context, o.UserLogoutUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/api/v1/user/{username}"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *PlantbookAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *PlantbookAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *PlantbookAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *PlantbookAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *PlantbookAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
