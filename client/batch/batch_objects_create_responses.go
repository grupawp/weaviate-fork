//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/semi-technologies/weaviate/entities/models"
)

// BatchObjectsCreateReader is a Reader for the BatchObjectsCreate structure.
type BatchObjectsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchObjectsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchObjectsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBatchObjectsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBatchObjectsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewBatchObjectsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchObjectsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBatchObjectsCreateOK creates a BatchObjectsCreateOK with default headers values
func NewBatchObjectsCreateOK() *BatchObjectsCreateOK {
	return &BatchObjectsCreateOK{}
}

/*
BatchObjectsCreateOK handles this case with default header values.

Request succeeded, see response body to get detailed information about each batched item.
*/
type BatchObjectsCreateOK struct {
	Payload []*models.ObjectsGetResponse
}

func (o *BatchObjectsCreateOK) Error() string {
	return fmt.Sprintf("[POST /batch/objects][%d] batchObjectsCreateOK  %+v", 200, o.Payload)
}

func (o *BatchObjectsCreateOK) GetPayload() []*models.ObjectsGetResponse {
	return o.Payload
}

func (o *BatchObjectsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchObjectsCreateUnauthorized creates a BatchObjectsCreateUnauthorized with default headers values
func NewBatchObjectsCreateUnauthorized() *BatchObjectsCreateUnauthorized {
	return &BatchObjectsCreateUnauthorized{}
}

/*
BatchObjectsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type BatchObjectsCreateUnauthorized struct {
}

func (o *BatchObjectsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /batch/objects][%d] batchObjectsCreateUnauthorized ", 401)
}

func (o *BatchObjectsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchObjectsCreateForbidden creates a BatchObjectsCreateForbidden with default headers values
func NewBatchObjectsCreateForbidden() *BatchObjectsCreateForbidden {
	return &BatchObjectsCreateForbidden{}
}

/*
BatchObjectsCreateForbidden handles this case with default header values.

Forbidden
*/
type BatchObjectsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *BatchObjectsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /batch/objects][%d] batchObjectsCreateForbidden  %+v", 403, o.Payload)
}

func (o *BatchObjectsCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchObjectsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchObjectsCreateUnprocessableEntity creates a BatchObjectsCreateUnprocessableEntity with default headers values
func NewBatchObjectsCreateUnprocessableEntity() *BatchObjectsCreateUnprocessableEntity {
	return &BatchObjectsCreateUnprocessableEntity{}
}

/*
BatchObjectsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type BatchObjectsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *BatchObjectsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /batch/objects][%d] batchObjectsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BatchObjectsCreateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchObjectsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchObjectsCreateInternalServerError creates a BatchObjectsCreateInternalServerError with default headers values
func NewBatchObjectsCreateInternalServerError() *BatchObjectsCreateInternalServerError {
	return &BatchObjectsCreateInternalServerError{}
}

/*
BatchObjectsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type BatchObjectsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *BatchObjectsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /batch/objects][%d] batchObjectsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *BatchObjectsCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchObjectsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
BatchObjectsCreateBody batch objects create body
swagger:model BatchObjectsCreateBody
*/
type BatchObjectsCreateBody struct {

	// Define which fields need to be returned. Default value is ALL
	Fields []*string `json:"fields"`

	// objects
	Objects []*models.Object `json:"objects"`
}

// Validate validates this batch objects create body
func (o *BatchObjectsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var batchObjectsCreateBodyFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","class","schema","id","creationTimeUnix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		batchObjectsCreateBodyFieldsItemsEnum = append(batchObjectsCreateBodyFieldsItemsEnum, v)
	}
}

func (o *BatchObjectsCreateBody) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, batchObjectsCreateBodyFieldsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *BatchObjectsCreateBody) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateFieldsItemsEnum("body"+"."+"fields"+"."+strconv.Itoa(i), "body", *o.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *BatchObjectsCreateBody) validateObjects(formats strfmt.Registry) error {

	if swag.IsZero(o.Objects) { // not required
		return nil
	}

	for i := 0; i < len(o.Objects); i++ {
		if swag.IsZero(o.Objects[i]) { // not required
			continue
		}

		if o.Objects[i] != nil {
			if err := o.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *BatchObjectsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BatchObjectsCreateBody) UnmarshalBinary(b []byte) error {
	var res BatchObjectsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
