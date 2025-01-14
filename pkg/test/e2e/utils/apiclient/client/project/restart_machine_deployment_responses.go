// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// RestartMachineDeploymentReader is a Reader for the RestartMachineDeployment structure.
type RestartMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestartMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestartMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRestartMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartMachineDeploymentOK creates a RestartMachineDeploymentOK with default headers values
func NewRestartMachineDeploymentOK() *RestartMachineDeploymentOK {
	return &RestartMachineDeploymentOK{}
}

/* RestartMachineDeploymentOK describes a response with status code 200, with default header values.

NodeDeployment
*/
type RestartMachineDeploymentOK struct {
	Payload *models.NodeDeployment
}

func (o *RestartMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] restartMachineDeploymentOK  %+v", 200, o.Payload)
}
func (o *RestartMachineDeploymentOK) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *RestartMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartMachineDeploymentUnauthorized creates a RestartMachineDeploymentUnauthorized with default headers values
func NewRestartMachineDeploymentUnauthorized() *RestartMachineDeploymentUnauthorized {
	return &RestartMachineDeploymentUnauthorized{}
}

/* RestartMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type RestartMachineDeploymentUnauthorized struct {
}

func (o *RestartMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] restartMachineDeploymentUnauthorized ", 401)
}

func (o *RestartMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartMachineDeploymentForbidden creates a RestartMachineDeploymentForbidden with default headers values
func NewRestartMachineDeploymentForbidden() *RestartMachineDeploymentForbidden {
	return &RestartMachineDeploymentForbidden{}
}

/* RestartMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type RestartMachineDeploymentForbidden struct {
}

func (o *RestartMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] restartMachineDeploymentForbidden ", 403)
}

func (o *RestartMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartMachineDeploymentDefault creates a RestartMachineDeploymentDefault with default headers values
func NewRestartMachineDeploymentDefault(code int) *RestartMachineDeploymentDefault {
	return &RestartMachineDeploymentDefault{
		_statusCode: code,
	}
}

/* RestartMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type RestartMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the restart machine deployment default response
func (o *RestartMachineDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *RestartMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] restartMachineDeployment default  %+v", o._statusCode, o.Payload)
}
func (o *RestartMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RestartMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
