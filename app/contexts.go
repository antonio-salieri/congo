//************************************************************************//
// congo: Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strconv"

	"github.com/raphael/goa"
)

// CreateAccountContext provides the account create action context.
type CreateAccountContext struct {
	*goa.Context
	Payload *CreateAccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller create action.
func NewCreateAccountContext(c *goa.Context) (*CreateAccountContext, error) {
	var err error
	ctx := CreateAccountContext{Context: c}
	p, err := NewCreateAccountPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	Name string
}

// NewCreateAccountPayload instantiates a CreateAccountPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateAccountPayload(raw interface{}) (p *CreateAccountPayload, err error) {
	p, err = UnmarshalCreateAccountPayload(raw, err)
	return
}

// UnmarshalCreateAccountPayload unmarshals and validates a raw interface{} into an instance of CreateAccountPayload
func UnmarshalCreateAccountPayload(source interface{}, inErr error) (target *CreateAccountPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateAccountPayload)
		if v, ok := val["name"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp1) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp1, 2, true, err)
				}
			}
			target.Name = tmp1
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAccountContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteAccountContext provides the account delete action context.
type DeleteAccountContext struct {
	*goa.Context
	AccountID int
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller delete action.
func NewDeleteAccountContext(c *goa.Context) (*DeleteAccountContext, error) {
	var err error
	ctx := DeleteAccountContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListAccountContext provides the account list action context.
type ListAccountContext struct {
	*goa.Context
}

// NewListAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller list action.
func NewListAccountContext(c *goa.Context) (*ListAccountContext, error) {
	var err error
	ctx := ListAccountContext{Context: c}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAccountContext) OK(resp AccountCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.account+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	*goa.Context
	AccountID int
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(c *goa.Context) (*ShowAccountContext, error) {
	var err error
	ctx := ShowAccountContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(resp *Account, view AccountViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.account+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateAccountContext provides the account update action context.
type UpdateAccountContext struct {
	*goa.Context
	AccountID int
	Payload   *UpdateAccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller update action.
func NewUpdateAccountContext(c *goa.Context) (*UpdateAccountContext, error) {
	var err error
	ctx := UpdateAccountContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewUpdateAccountPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	Name string
}

// NewUpdateAccountPayload instantiates a UpdateAccountPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateAccountPayload(raw interface{}) (p *UpdateAccountPayload, err error) {
	p, err = UnmarshalUpdateAccountPayload(raw, err)
	return
}

// UnmarshalUpdateAccountPayload unmarshals and validates a raw interface{} into an instance of UpdateAccountPayload
func UnmarshalUpdateAccountPayload(source interface{}, inErr error) (target *UpdateAccountPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateAccountPayload)
		if v, ok := val["name"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp2) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp2, 2, true, err)
				}
			}
			target.Name = tmp2
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAccountContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateInstanceContext provides the instance create action context.
type CreateInstanceContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
	Payload   *CreateInstancePayload
}

// NewCreateInstanceContext parses the incoming request URL and body, performs validations and creates the
// context used by the instance controller create action.
func NewCreateInstanceContext(c *goa.Context) (*CreateInstanceContext, error) {
	var err error
	ctx := CreateInstanceContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	p, err := NewCreateInstancePayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateInstancePayload is the instance create action payload.
type CreateInstancePayload struct {
	Name string
}

// NewCreateInstancePayload instantiates a CreateInstancePayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateInstancePayload(raw interface{}) (p *CreateInstancePayload, err error) {
	p, err = UnmarshalCreateInstancePayload(raw, err)
	return
}

// UnmarshalCreateInstancePayload unmarshals and validates a raw interface{} into an instance of CreateInstancePayload
func UnmarshalCreateInstancePayload(source interface{}, inErr error) (target *CreateInstancePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateInstancePayload)
		if v, ok := val["name"]; ok {
			var tmp3 string
			if val, ok := v.(string); ok {
				tmp3 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp3) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp3, 2, true, err)
				}
			}
			target.Name = tmp3
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateInstanceContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteInstanceContext provides the instance delete action context.
type DeleteInstanceContext struct {
	*goa.Context
	AccountID  int
	InstanceID int
	SeriesID   int
}

// NewDeleteInstanceContext parses the incoming request URL and body, performs validations and creates the
// context used by the instance controller delete action.
func NewDeleteInstanceContext(c *goa.Context) (*DeleteInstanceContext, error) {
	var err error
	ctx := DeleteInstanceContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawInstanceID, ok := c.Get("instanceID")
	if ok {
		if instanceID, err2 := strconv.Atoi(rawInstanceID); err2 == nil {
			ctx.InstanceID = int(instanceID)
		} else {
			err = goa.InvalidParamTypeError("instanceID", rawInstanceID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteInstanceContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteInstanceContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListInstanceContext provides the instance list action context.
type ListInstanceContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
}

// NewListInstanceContext parses the incoming request URL and body, performs validations and creates the
// context used by the instance controller list action.
func NewListInstanceContext(c *goa.Context) (*ListInstanceContext, error) {
	var err error
	ctx := ListInstanceContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListInstanceContext) OK(resp InstanceCollection, view InstanceCollectionViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.instance+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowInstanceContext provides the instance show action context.
type ShowInstanceContext struct {
	*goa.Context
	AccountID  int
	InstanceID int
	SeriesID   int
}

// NewShowInstanceContext parses the incoming request URL and body, performs validations and creates the
// context used by the instance controller show action.
func NewShowInstanceContext(c *goa.Context) (*ShowInstanceContext, error) {
	var err error
	ctx := ShowInstanceContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawInstanceID, ok := c.Get("instanceID")
	if ok {
		if instanceID, err2 := strconv.Atoi(rawInstanceID); err2 == nil {
			ctx.InstanceID = int(instanceID)
		} else {
			err = goa.InvalidParamTypeError("instanceID", rawInstanceID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowInstanceContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowInstanceContext) OK(resp *Instance, view InstanceViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.instance+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateInstanceContext provides the instance update action context.
type UpdateInstanceContext struct {
	*goa.Context
	AccountID  int
	InstanceID int
	SeriesID   int
	Payload    *UpdateInstancePayload
}

// NewUpdateInstanceContext parses the incoming request URL and body, performs validations and creates the
// context used by the instance controller update action.
func NewUpdateInstanceContext(c *goa.Context) (*UpdateInstanceContext, error) {
	var err error
	ctx := UpdateInstanceContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawInstanceID, ok := c.Get("instanceID")
	if ok {
		if instanceID, err2 := strconv.Atoi(rawInstanceID); err2 == nil {
			ctx.InstanceID = int(instanceID)
		} else {
			err = goa.InvalidParamTypeError("instanceID", rawInstanceID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	p, err := NewUpdateInstancePayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateInstancePayload is the instance update action payload.
type UpdateInstancePayload struct {
	Name string
}

// NewUpdateInstancePayload instantiates a UpdateInstancePayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateInstancePayload(raw interface{}) (p *UpdateInstancePayload, err error) {
	p, err = UnmarshalUpdateInstancePayload(raw, err)
	return
}

// UnmarshalUpdateInstancePayload unmarshals and validates a raw interface{} into an instance of UpdateInstancePayload
func UnmarshalUpdateInstancePayload(source interface{}, inErr error) (target *UpdateInstancePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateInstancePayload)
		if v, ok := val["name"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp4) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp4, 2, true, err)
				}
			}
			target.Name = tmp4
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateInstanceContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateInstanceContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateSeriesContext provides the series create action context.
type CreateSeriesContext struct {
	*goa.Context
	AccountID int
	Payload   *CreateSeriesPayload
}

// NewCreateSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller create action.
func NewCreateSeriesContext(c *goa.Context) (*CreateSeriesContext, error) {
	var err error
	ctx := CreateSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewCreateSeriesPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateSeriesPayload is the series create action payload.
type CreateSeriesPayload struct {
	Name string
}

// NewCreateSeriesPayload instantiates a CreateSeriesPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateSeriesPayload(raw interface{}) (p *CreateSeriesPayload, err error) {
	p, err = UnmarshalCreateSeriesPayload(raw, err)
	return
}

// UnmarshalCreateSeriesPayload unmarshals and validates a raw interface{} into an instance of CreateSeriesPayload
func UnmarshalCreateSeriesPayload(source interface{}, inErr error) (target *CreateSeriesPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateSeriesPayload)
		if v, ok := val["name"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp5) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp5, 2, true, err)
				}
			}
			target.Name = tmp5
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSeriesContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteSeriesContext provides the series delete action context.
type DeleteSeriesContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
}

// NewDeleteSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller delete action.
func NewDeleteSeriesContext(c *goa.Context) (*DeleteSeriesContext, error) {
	var err error
	ctx := DeleteSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteSeriesContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteSeriesContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListSeriesContext provides the series list action context.
type ListSeriesContext struct {
	*goa.Context
	AccountID int
}

// NewListSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller list action.
func NewListSeriesContext(c *goa.Context) (*ListSeriesContext, error) {
	var err error
	ctx := ListSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSeriesContext) OK(resp SeriesCollection, view SeriesCollectionViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.series+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowSeriesContext provides the series show action context.
type ShowSeriesContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
}

// NewShowSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller show action.
func NewShowSeriesContext(c *goa.Context) (*ShowSeriesContext, error) {
	var err error
	ctx := ShowSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSeriesContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSeriesContext) OK(resp *Series, view SeriesViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.series+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateSeriesContext provides the series update action context.
type UpdateSeriesContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
	Payload   *UpdateSeriesPayload
}

// NewUpdateSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller update action.
func NewUpdateSeriesContext(c *goa.Context) (*UpdateSeriesContext, error) {
	var err error
	ctx := UpdateSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	p, err := NewUpdateSeriesPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateSeriesPayload is the series update action payload.
type UpdateSeriesPayload struct {
	Name string
}

// NewUpdateSeriesPayload instantiates a UpdateSeriesPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateSeriesPayload(raw interface{}) (p *UpdateSeriesPayload, err error) {
	p, err = UnmarshalUpdateSeriesPayload(raw, err)
	return
}

// UnmarshalUpdateSeriesPayload unmarshals and validates a raw interface{} into an instance of UpdateSeriesPayload
func UnmarshalUpdateSeriesPayload(source interface{}, inErr error) (target *UpdateSeriesPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateSeriesPayload)
		if v, ok := val["name"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp6) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp6, 2, true, err)
				}
			}
			target.Name = tmp6
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateSeriesContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSeriesContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	*goa.Context
	AccountID int
	Payload   *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewCreateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Email     string
	FirstName string
	ID        int
	LastName  string
	Role      string
}

// NewCreateUserPayload instantiates a CreateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateUserPayload(raw interface{}) (p *CreateUserPayload, err error) {
	p, err = UnmarshalCreateUserPayload(raw, err)
	return
}

// UnmarshalCreateUserPayload unmarshals and validates a raw interface{} into an instance of CreateUserPayload
func UnmarshalCreateUserPayload(source interface{}, inErr error) (target *CreateUserPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateUserPayload)
		if v, ok := val["email"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp7) < 2 {
					err = goa.InvalidLengthError(`payload.Email`, tmp7, 2, true, err)
				}
			}
			target.Email = tmp7
		}
		if v, ok := val["first_name"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			if err == nil {
				if len(tmp8) < 2 {
					err = goa.InvalidLengthError(`payload.FirstName`, tmp8, 2, true, err)
				}
			}
			target.FirstName = tmp8
		} else {
			err = goa.MissingAttributeError(`payload`, "first_name", err)
		}
		if v, ok := val["id"]; ok {
			var tmp9 int
			if f, ok := v.(float64); ok {
				tmp9 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp9
		}
		if v, ok := val["last_name"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.LastName`, v, "string", err)
			}
			if err == nil {
				if len(tmp10) < 2 {
					err = goa.InvalidLengthError(`payload.LastName`, tmp10, 2, true, err)
				}
			}
			target.LastName = tmp10
		}
		if v, ok := val["role"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = tmp11
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	*goa.Context
	AccountID int
	UserID    int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(c *goa.Context) (*DeleteUserContext, error) {
	var err error
	ctx := DeleteUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawUserID, ok := c.Get("userID")
	if ok {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	*goa.Context
	AccountID int
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(resp UserCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.user+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	*goa.Context
	AccountID int
	UserID    int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(c *goa.Context) (*ShowUserContext, error) {
	var err error
	ctx := ShowUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawUserID, ok := c.Get("userID")
	if ok {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(resp *User, view UserViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.user+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	*goa.Context
	AccountID int
	UserID    int
	Payload   *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(c *goa.Context) (*UpdateUserContext, error) {
	var err error
	ctx := UpdateUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawUserID, ok := c.Get("userID")
	if ok {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	p, err := NewUpdateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Email     string
	FirstName string
	ID        int
	LastName  string
	Role      string
}

// NewUpdateUserPayload instantiates a UpdateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateUserPayload(raw interface{}) (p *UpdateUserPayload, err error) {
	p, err = UnmarshalUpdateUserPayload(raw, err)
	return
}

// UnmarshalUpdateUserPayload unmarshals and validates a raw interface{} into an instance of UpdateUserPayload
func UnmarshalUpdateUserPayload(source interface{}, inErr error) (target *UpdateUserPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateUserPayload)
		if v, ok := val["email"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp12) < 2 {
					err = goa.InvalidLengthError(`payload.Email`, tmp12, 2, true, err)
				}
			}
			target.Email = tmp12
		}
		if v, ok := val["first_name"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			if err == nil {
				if len(tmp13) < 2 {
					err = goa.InvalidLengthError(`payload.FirstName`, tmp13, 2, true, err)
				}
			}
			target.FirstName = tmp13
		}
		if v, ok := val["id"]; ok {
			var tmp14 int
			if f, ok := v.(float64); ok {
				tmp14 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp14
		}
		if v, ok := val["last_name"]; ok {
			var tmp15 string
			if val, ok := v.(string); ok {
				tmp15 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.LastName`, v, "string", err)
			}
			if err == nil {
				if len(tmp15) < 2 {
					err = goa.InvalidLengthError(`payload.LastName`, tmp15, 2, true, err)
				}
			}
			target.LastName = tmp15
		}
		if v, ok := val["role"]; ok {
			var tmp16 string
			if val, ok := v.(string); ok {
				tmp16 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = tmp16
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}
