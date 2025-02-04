// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchAddonV2Reader is a Reader for the PatchAddonV2 structure.
type PatchAddonV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAddonV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAddonV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAddonV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAddonV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchAddonV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAddonV2OK creates a PatchAddonV2OK with default headers values
func NewPatchAddonV2OK() *PatchAddonV2OK {
	return &PatchAddonV2OK{}
}

/* PatchAddonV2OK describes a response with status code 200, with default header values.

Addon
*/
type PatchAddonV2OK struct {
	Payload *models.Addon
}

func (o *PatchAddonV2OK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonV2OK  %+v", 200, o.Payload)
}
func (o *PatchAddonV2OK) GetPayload() *models.Addon {
	return o.Payload
}

func (o *PatchAddonV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAddonV2Unauthorized creates a PatchAddonV2Unauthorized with default headers values
func NewPatchAddonV2Unauthorized() *PatchAddonV2Unauthorized {
	return &PatchAddonV2Unauthorized{}
}

/* PatchAddonV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchAddonV2Unauthorized struct {
}

func (o *PatchAddonV2Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonV2Unauthorized ", 401)
}

func (o *PatchAddonV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAddonV2Forbidden creates a PatchAddonV2Forbidden with default headers values
func NewPatchAddonV2Forbidden() *PatchAddonV2Forbidden {
	return &PatchAddonV2Forbidden{}
}

/* PatchAddonV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchAddonV2Forbidden struct {
}

func (o *PatchAddonV2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonV2Forbidden ", 403)
}

func (o *PatchAddonV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAddonV2Default creates a PatchAddonV2Default with default headers values
func NewPatchAddonV2Default(code int) *PatchAddonV2Default {
	return &PatchAddonV2Default{
		_statusCode: code,
	}
}

/* PatchAddonV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type PatchAddonV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch addon v2 default response
func (o *PatchAddonV2Default) Code() int {
	return o._statusCode
}

func (o *PatchAddonV2Default) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] patchAddonV2 default  %+v", o._statusCode, o.Payload)
}
func (o *PatchAddonV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchAddonV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
