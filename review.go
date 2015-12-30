package main

import (
	"github.com/gopheracademy/congo/app"
	"github.com/raphael/goa"
)

// ReviewController implements the review resource.
type ReviewController struct {
	goa.Controller
}

// NewReviewController creates a review controller.
func NewReviewController(service goa.Service) app.ReviewController {
	return &ReviewController{Controller: service.NewController("ReviewController")}
}

// Create runs the create action.
func (c *ReviewController) Create(ctx *app.CreateReviewContext) error {
	return nil
}

// Delete runs the delete action.
func (c *ReviewController) Delete(ctx *app.DeleteReviewContext) error {
	return nil
}

// List runs the list action.
func (c *ReviewController) List(ctx *app.ListReviewContext) error {
	return nil
}

// Show runs the show action.
func (c *ReviewController) Show(ctx *app.ShowReviewContext) error {
	res := &app.Review{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *ReviewController) Update(ctx *app.UpdateReviewContext) error {
	return nil
}
