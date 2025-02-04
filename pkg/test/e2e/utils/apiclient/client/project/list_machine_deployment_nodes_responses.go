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

// ListMachineDeploymentNodesReader is a Reader for the ListMachineDeploymentNodes structure.
type ListMachineDeploymentNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMachineDeploymentNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMachineDeploymentNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListMachineDeploymentNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListMachineDeploymentNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMachineDeploymentNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMachineDeploymentNodesOK creates a ListMachineDeploymentNodesOK with default headers values
func NewListMachineDeploymentNodesOK() *ListMachineDeploymentNodesOK {
	return &ListMachineDeploymentNodesOK{}
}

/* ListMachineDeploymentNodesOK describes a response with status code 200, with default header values.

Node
*/
type ListMachineDeploymentNodesOK struct {
	Payload []*models.Node
}

func (o *ListMachineDeploymentNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesOK  %+v", 200, o.Payload)
}
func (o *ListMachineDeploymentNodesOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *ListMachineDeploymentNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMachineDeploymentNodesUnauthorized creates a ListMachineDeploymentNodesUnauthorized with default headers values
func NewListMachineDeploymentNodesUnauthorized() *ListMachineDeploymentNodesUnauthorized {
	return &ListMachineDeploymentNodesUnauthorized{}
}

/* ListMachineDeploymentNodesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentNodesUnauthorized struct {
}

func (o *ListMachineDeploymentNodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesUnauthorized ", 401)
}

func (o *ListMachineDeploymentNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentNodesForbidden creates a ListMachineDeploymentNodesForbidden with default headers values
func NewListMachineDeploymentNodesForbidden() *ListMachineDeploymentNodesForbidden {
	return &ListMachineDeploymentNodesForbidden{}
}

/* ListMachineDeploymentNodesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentNodesForbidden struct {
}

func (o *ListMachineDeploymentNodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesForbidden ", 403)
}

func (o *ListMachineDeploymentNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentNodesDefault creates a ListMachineDeploymentNodesDefault with default headers values
func NewListMachineDeploymentNodesDefault(code int) *ListMachineDeploymentNodesDefault {
	return &ListMachineDeploymentNodesDefault{
		_statusCode: code,
	}
}

/* ListMachineDeploymentNodesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListMachineDeploymentNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list machine deployment nodes default response
func (o *ListMachineDeploymentNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListMachineDeploymentNodesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodes default  %+v", o._statusCode, o.Payload)
}
func (o *ListMachineDeploymentNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListMachineDeploymentNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
