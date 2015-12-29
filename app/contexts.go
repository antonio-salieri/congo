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

// RefreshAuthContext provides the auth refresh action context.
type RefreshAuthContext struct {
	*goa.Context
	Payload *RefreshAuthPayload
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(c *goa.Context) (*RefreshAuthContext, error) {
	var err error
	ctx := RefreshAuthContext{Context: c}
	p, err := NewRefreshAuthPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// RefreshAuthPayload is the auth refresh action payload.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application string
	// email
	Email string
	// password
	Password string
}

// NewRefreshAuthPayload instantiates a RefreshAuthPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewRefreshAuthPayload(raw interface{}) (p *RefreshAuthPayload, err error) {
	p, err = UnmarshalRefreshAuthPayload(raw, err)
	return
}

// UnmarshalRefreshAuthPayload unmarshals and validates a raw interface{} into an instance of RefreshAuthPayload
func UnmarshalRefreshAuthPayload(source interface{}, inErr error) (target *RefreshAuthPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(RefreshAuthPayload)
		if v, ok := val["application"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Application`, v, "string", err)
			}
			target.Application = tmp1
		}
		if v, ok := val["email"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = tmp2
		}
		if v, ok := val["password"]; ok {
			var tmp3 string
			if val, ok := v.(string); ok {
				tmp3 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Password`, v, "string", err)
			}
			target.Password = tmp3
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *RefreshAuthContext) Created(resp *Authorize) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json; charset=utf-8")
	return ctx.JSON(201, r)
}

// TokenAuthContext provides the auth token action context.
type TokenAuthContext struct {
	*goa.Context
	Payload *TokenAuthPayload
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(c *goa.Context) (*TokenAuthContext, error) {
	var err error
	ctx := TokenAuthContext{Context: c}
	p, err := NewTokenAuthPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application string
	// email
	Email string
	// password
	Password string
}

// NewTokenAuthPayload instantiates a TokenAuthPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewTokenAuthPayload(raw interface{}) (p *TokenAuthPayload, err error) {
	p, err = UnmarshalTokenAuthPayload(raw, err)
	return
}

// UnmarshalTokenAuthPayload unmarshals and validates a raw interface{} into an instance of TokenAuthPayload
func UnmarshalTokenAuthPayload(source interface{}, inErr error) (target *TokenAuthPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(TokenAuthPayload)
		if v, ok := val["application"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Application`, v, "string", err)
			}
			target.Application = tmp4
		}
		if v, ok := val["email"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = tmp5
		}
		if v, ok := val["password"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Password`, v, "string", err)
			}
			target.Password = tmp6
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *TokenAuthContext) Created(resp *Authorize) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json; charset=utf-8")
	return ctx.JSON(201, r)
}

// CreateProposalContext provides the proposal create action context.
type CreateProposalContext struct {
	*goa.Context
	UserID  int
	Payload *CreateProposalPayload
}

// NewCreateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller create action.
func NewCreateProposalContext(c *goa.Context) (*CreateProposalContext, error) {
	var err error
	ctx := CreateProposalContext{Context: c}
	rawUserID, ok := c.Get("userID")
	if ok {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	p, err := NewCreateProposalPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  string
	Detail    string
	Firstname string
	Title     string
	Withdrawn bool
}

// NewCreateProposalPayload instantiates a CreateProposalPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateProposalPayload(raw interface{}) (p *CreateProposalPayload, err error) {
	p, err = UnmarshalCreateProposalPayload(raw, err)
	return
}

// UnmarshalCreateProposalPayload unmarshals and validates a raw interface{} into an instance of CreateProposalPayload
func UnmarshalCreateProposalPayload(source interface{}, inErr error) (target *CreateProposalPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateProposalPayload)
		if v, ok := val["abstract"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp7) < 50 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp7, 50, true, err)
				}
				if len(tmp7) > 500 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp7, 500, false, err)
				}
			}
			target.Abstract = tmp7
		}
		if v, ok := val["detail"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp8) < 100 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp8, 100, true, err)
				}
				if len(tmp8) > 2000 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp8, 2000, false, err)
				}
			}
			target.Detail = tmp8
		}
		if v, ok := val["firstname"]; ok {
			var tmp9 string
			if val, ok := v.(string); ok {
				tmp9 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp9) < 2 {
					err = goa.InvalidLengthError(`payload.Firstname`, tmp9, 2, true, err)
				}
			}
			target.Firstname = tmp9
		}
		if v, ok := val["title"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp10) < 10 {
					err = goa.InvalidLengthError(`payload.Title`, tmp10, 10, true, err)
				}
				if len(tmp10) > 200 {
					err = goa.InvalidLengthError(`payload.Title`, tmp10, 200, false, err)
				}
			}
			target.Title = tmp10
		} else {
			err = goa.MissingAttributeError(`payload`, "title", err)
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp11 bool
			if val, ok := v.(bool); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = tmp11
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateProposalContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteProposalContext provides the proposal delete action context.
type DeleteProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
}

// NewDeleteProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller delete action.
func NewDeleteProposalContext(c *goa.Context) (*DeleteProposalContext, error) {
	var err error
	ctx := DeleteProposalContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
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
func (ctx *DeleteProposalContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListProposalContext provides the proposal list action context.
type ListProposalContext struct {
	*goa.Context
	UserID int
}

// NewListProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller list action.
func NewListProposalContext(c *goa.Context) (*ListProposalContext, error) {
	var err error
	ctx := ListProposalContext{Context: c}
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

// OK sends a HTTP response with status code 200.
func (ctx *ListProposalContext) OK(resp ProposalCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowProposalContext provides the proposal show action context.
type ShowProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
}

// NewShowProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller show action.
func NewShowProposalContext(c *goa.Context) (*ShowProposalContext, error) {
	var err error
	ctx := ShowProposalContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
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
func (ctx *ShowProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OK(resp *Proposal, view ProposalViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateProposalContext provides the proposal update action context.
type UpdateProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Payload    *UpdateProposalPayload
}

// NewUpdateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller update action.
func NewUpdateProposalContext(c *goa.Context) (*UpdateProposalContext, error) {
	var err error
	ctx := UpdateProposalContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
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
	p, err := NewUpdateProposalPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateProposalPayload is the proposal update action payload.
type UpdateProposalPayload struct {
	Abstract  string
	Detail    string
	Firstname string
	Title     string
	Withdrawn bool
}

// NewUpdateProposalPayload instantiates a UpdateProposalPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateProposalPayload(raw interface{}) (p *UpdateProposalPayload, err error) {
	p, err = UnmarshalUpdateProposalPayload(raw, err)
	return
}

// UnmarshalUpdateProposalPayload unmarshals and validates a raw interface{} into an instance of UpdateProposalPayload
func UnmarshalUpdateProposalPayload(source interface{}, inErr error) (target *UpdateProposalPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateProposalPayload)
		if v, ok := val["abstract"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp12) < 50 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp12, 50, true, err)
				}
				if len(tmp12) > 500 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp12, 500, false, err)
				}
			}
			target.Abstract = tmp12
		}
		if v, ok := val["detail"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp13) < 100 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp13, 100, true, err)
				}
				if len(tmp13) > 2000 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp13, 2000, false, err)
				}
			}
			target.Detail = tmp13
		}
		if v, ok := val["firstname"]; ok {
			var tmp14 string
			if val, ok := v.(string); ok {
				tmp14 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp14) < 2 {
					err = goa.InvalidLengthError(`payload.Firstname`, tmp14, 2, true, err)
				}
			}
			target.Firstname = tmp14
		}
		if v, ok := val["title"]; ok {
			var tmp15 string
			if val, ok := v.(string); ok {
				tmp15 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp15) < 10 {
					err = goa.InvalidLengthError(`payload.Title`, tmp15, 10, true, err)
				}
				if len(tmp15) > 200 {
					err = goa.InvalidLengthError(`payload.Title`, tmp15, 200, false, err)
				}
			}
			target.Title = tmp15
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp16 bool
			if val, ok := v.(bool); ok {
				tmp16 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = tmp16
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateProposalContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateReviewContext provides the review create action context.
type CreateReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Payload    *CreateReviewPayload
}

// NewCreateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller create action.
func NewCreateReviewContext(c *goa.Context) (*CreateReviewContext, error) {
	var err error
	ctx := CreateReviewContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
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
	p, err := NewCreateReviewPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	Comment string
	Rating  int
}

// NewCreateReviewPayload instantiates a CreateReviewPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateReviewPayload(raw interface{}) (p *CreateReviewPayload, err error) {
	p, err = UnmarshalCreateReviewPayload(raw, err)
	return
}

// UnmarshalCreateReviewPayload unmarshals and validates a raw interface{} into an instance of CreateReviewPayload
func UnmarshalCreateReviewPayload(source interface{}, inErr error) (target *CreateReviewPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateReviewPayload)
		if v, ok := val["comment"]; ok {
			var tmp17 string
			if val, ok := v.(string); ok {
				tmp17 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp17) < 10 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp17, 10, true, err)
				}
				if len(tmp17) > 200 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp17, 200, false, err)
				}
			}
			target.Comment = tmp17
		}
		if v, ok := val["rating"]; ok {
			var tmp18 int
			if f, ok := v.(float64); ok {
				tmp18 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp18 < 1 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp18, 1, true, err)
				}
				if tmp18 > 5 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp18, 5, false, err)
				}
			}
			target.Rating = tmp18
		} else {
			err = goa.MissingAttributeError(`payload`, "rating", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateReviewContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteReviewContext provides the review delete action context.
type DeleteReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewDeleteReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller delete action.
func NewDeleteReviewContext(c *goa.Context) (*DeleteReviewContext, error) {
	var err error
	ctx := DeleteReviewContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID, ok := c.Get("reviewID")
	if ok {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
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
func (ctx *DeleteReviewContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListReviewContext provides the review list action context.
type ListReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
}

// NewListReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller list action.
func NewListReviewContext(c *goa.Context) (*ListReviewContext, error) {
	var err error
	ctx := ListReviewContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
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

// OK sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OK(resp ReviewCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.review+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowReviewContext provides the review show action context.
type ShowReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewShowReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller show action.
func NewShowReviewContext(c *goa.Context) (*ShowReviewContext, error) {
	var err error
	ctx := ShowReviewContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID, ok := c.Get("reviewID")
	if ok {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
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
func (ctx *ShowReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OK(resp *Review, view ReviewViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.review+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateReviewContext provides the review update action context.
type UpdateReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Payload    *UpdateReviewPayload
}

// NewUpdateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller update action.
func NewUpdateReviewContext(c *goa.Context) (*UpdateReviewContext, error) {
	var err error
	ctx := UpdateReviewContext{Context: c}
	rawProposalID, ok := c.Get("proposalID")
	if ok {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID, ok := c.Get("reviewID")
	if ok {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
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
	p, err := NewUpdateReviewPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateReviewPayload is the review update action payload.
type UpdateReviewPayload struct {
	Comment string
	Rating  int
}

// NewUpdateReviewPayload instantiates a UpdateReviewPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateReviewPayload(raw interface{}) (p *UpdateReviewPayload, err error) {
	p, err = UnmarshalUpdateReviewPayload(raw, err)
	return
}

// UnmarshalUpdateReviewPayload unmarshals and validates a raw interface{} into an instance of UpdateReviewPayload
func UnmarshalUpdateReviewPayload(source interface{}, inErr error) (target *UpdateReviewPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateReviewPayload)
		if v, ok := val["comment"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp19) < 10 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp19, 10, true, err)
				}
				if len(tmp19) > 200 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp19, 200, false, err)
				}
			}
			target.Comment = tmp19
		}
		if v, ok := val["rating"]; ok {
			var tmp20 int
			if f, ok := v.(float64); ok {
				tmp20 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp20 < 1 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp20, 1, true, err)
				}
				if tmp20 > 5 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp20, 5, false, err)
				}
			}
			target.Rating = tmp20
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateReviewContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	*goa.Context
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	p, err := NewCreateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Bio       string
	City      string
	Country   string
	Email     string
	Firstname string
	Lastname  string
	Password  string
	Role      string
	State     string
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
		if v, ok := val["bio"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp21) > 500 {
					err = goa.InvalidLengthError(`payload.Bio`, tmp21, 500, false, err)
				}
			}
			target.Bio = tmp21
		}
		if v, ok := val["city"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = tmp22
		}
		if v, ok := val["country"]; ok {
			var tmp23 string
			if val, ok := v.(string); ok {
				tmp23 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = tmp23
		}
		if v, ok := val["email"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp24) < 2 {
					err = goa.InvalidLengthError(`payload.Email`, tmp24, 2, true, err)
				}
			}
			target.Email = tmp24
		}
		if v, ok := val["firstname"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = tmp25
		} else {
			err = goa.MissingAttributeError(`payload`, "firstname", err)
		}
		if v, ok := val["lastname"]; ok {
			var tmp26 string
			if val, ok := v.(string); ok {
				tmp26 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Lastname`, v, "string", err)
			}
			target.Lastname = tmp26
		}
		if v, ok := val["password"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Password`, v, "string", err)
			}
			if err == nil {
				if len(tmp27) < 8 {
					err = goa.InvalidLengthError(`payload.Password`, tmp27, 8, true, err)
				}
			}
			target.Password = tmp27
		}
		if v, ok := val["role"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = tmp28
		}
		if v, ok := val["state"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = tmp29
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
	UserID int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(c *goa.Context) (*DeleteUserContext, error) {
	var err error
	ctx := DeleteUserContext{Context: c}
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
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
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
	UserID int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(c *goa.Context) (*ShowUserContext, error) {
	var err error
	ctx := ShowUserContext{Context: c}
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
	UserID  int
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(c *goa.Context) (*UpdateUserContext, error) {
	var err error
	ctx := UpdateUserContext{Context: c}
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
	Bio       string
	City      string
	Country   string
	Email     string
	Firstname string
	Lastname  string
	Password  string
	Role      string
	State     string
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
		if v, ok := val["bio"]; ok {
			var tmp30 string
			if val, ok := v.(string); ok {
				tmp30 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp30) > 500 {
					err = goa.InvalidLengthError(`payload.Bio`, tmp30, 500, false, err)
				}
			}
			target.Bio = tmp30
		}
		if v, ok := val["city"]; ok {
			var tmp31 string
			if val, ok := v.(string); ok {
				tmp31 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = tmp31
		}
		if v, ok := val["country"]; ok {
			var tmp32 string
			if val, ok := v.(string); ok {
				tmp32 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = tmp32
		}
		if v, ok := val["email"]; ok {
			var tmp33 string
			if val, ok := v.(string); ok {
				tmp33 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp33) < 2 {
					err = goa.InvalidLengthError(`payload.Email`, tmp33, 2, true, err)
				}
			}
			target.Email = tmp33
		}
		if v, ok := val["firstname"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = tmp34
		}
		if v, ok := val["lastname"]; ok {
			var tmp35 string
			if val, ok := v.(string); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Lastname`, v, "string", err)
			}
			target.Lastname = tmp35
		}
		if v, ok := val["password"]; ok {
			var tmp36 string
			if val, ok := v.(string); ok {
				tmp36 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Password`, v, "string", err)
			}
			if err == nil {
				if len(tmp36) < 8 {
					err = goa.InvalidLengthError(`payload.Password`, tmp36, 8, true, err)
				}
			}
			target.Password = tmp36
		}
		if v, ok := val["role"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = tmp37
		}
		if v, ok := val["state"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = tmp38
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
