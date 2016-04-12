//************************************************************************//
// API "congo": Model Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListProposal returns an array of view: default.
func (m *ProposalDB) ListProposal(ctx context.Context, userID int) []*app.Proposal {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "listproposal"}, time.Now())

	var native []*Proposal
	var objs []*app.Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.Error(ctx, "error listing Proposal", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ProposalToProposal())
	}

	return objs
}

// ProposalToProposal returns the Proposal representation of Proposal.
func (m *Proposal) ProposalToProposal() *app.Proposal {
	proposal := &app.Proposal{}
	proposal.Abstract = &m.Abstract
	proposal.Detail = &m.Detail
	proposal.ID = &m.ID
	proposal.Title = &m.Title

	return proposal
}

// OneProposal returns an array of view: default.
func (m *ProposalDB) OneProposal(ctx context.Context, id int, userID int) (*app.Proposal, error) {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "oneproposal"}, time.Now())

	var native Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.Error(ctx, "error getting Proposal", "error", err.Error())
		return nil, err
	}

	view := *native.ProposalToProposal()
	return &view, err
}

// MediaType Retrieval Functions

// ListProposalLink returns an array of view: link.
func (m *ProposalDB) ListProposalLink(ctx context.Context, userID int) []*app.ProposalLink {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "listproposallink"}, time.Now())

	var native []*Proposal
	var objs []*app.ProposalLink
	err := m.Db.Scopes(ProposalFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.Error(ctx, "error listing Proposal", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ProposalToProposalLink())
	}

	return objs
}

// ProposalToProposalLink returns the Proposal representation of Proposal.
func (m *Proposal) ProposalToProposalLink() *app.ProposalLink {
	proposal := &app.ProposalLink{}
	proposal.ID = &m.ID
	proposal.Title = &m.Title

	return proposal
}

// OneProposalLink returns an array of view: link.
func (m *ProposalDB) OneProposalLink(ctx context.Context, id int, userID int) (*app.ProposalLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "oneproposallink"}, time.Now())

	var native Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.Error(ctx, "error getting Proposal", "error", err.Error())
		return nil, err
	}

	view := *native.ProposalToProposalLink()
	return &view, err
}
