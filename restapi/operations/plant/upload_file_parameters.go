// Code generated by go-swagger; DO NOT EDIT.

package plant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UploadFileMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var UploadFileMaxParseMemory int64 = 32 << 20

// NewUploadFileParams creates a new UploadFileParams object
//
// There are no default values defined in the spec.
func NewUploadFileParams() UploadFileParams {

	return UploadFileParams{}
}

// UploadFileParams contains all the bound params for the upload file operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadFile
type UploadFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Additional data to pass to server
	  In: formData
	*/
	AdditionalMetadata *string
	/*file to upload
	  In: formData
	*/
	File io.ReadCloser
	/*ID of plant to update
	  Required: true
	  In: path
	*/
	PlantID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadFileParams() beforehand.
func (o *UploadFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(UploadFileMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdAdditionalMetadata, fdhkAdditionalMetadata, _ := fds.GetOK("additionalMetadata")
	if err := o.bindAdditionalMetadata(fdAdditionalMetadata, fdhkAdditionalMetadata, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindFile(file, fileHeader); err != nil {
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	rPlantID, rhkPlantID, _ := route.Params.GetOK("plantId")
	if err := o.bindPlantID(rPlantID, rhkPlantID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAdditionalMetadata binds and validates parameter AdditionalMetadata from formData.
func (o *UploadFileParams) bindAdditionalMetadata(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.AdditionalMetadata = &raw

	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadFileParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindPlantID binds and validates parameter PlantID from path.
func (o *UploadFileParams) bindPlantID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("plantId", "path", "int64", raw)
	}
	o.PlantID = value

	return nil
}
