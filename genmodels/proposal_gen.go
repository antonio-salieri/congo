//************************************************************************//
// API "congo": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package genmodels

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// Proposal Model
type Proposal struct {
	ID        int `gorm:"primary_key"` // This is the Payload Model PK field
	Abstract  *string
	Detail    *string
	Firstname *string
	Reviews   []Review // has many Reviews
	Title     *string  `gorm:"column:proposal_title"`
	UserID    int      // has many Proposal
	Withdrawn *bool
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
	CreatedAt time.Time  // timestamp
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Proposal) TableName() string {
	return "proposals"

}

// ProposalDB is the implementation of the storage interface for
// Proposal.
type ProposalDB struct {
	Db gorm.DB
	log.Logger
}

// NewProposalDB creates a new storage type.
func NewProposalDB(db gorm.DB, logger log.Logger) *ProposalDB {
	glog := logger.New("db", "Proposal")
	return &ProposalDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *ProposalDB) DB() interface{} {
	return &m.Db
}

// ProposalStorage represents the storage interface.
type ProposalStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Proposal
	Get(ctx goa.Context, id int) (Proposal, error)
	Add(ctx goa.Context, proposal Proposal) (Proposal, error)
	Update(ctx goa.Context, proposal Proposal) error
	Delete(ctx goa.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ProposalDB) TableName() string {
	return "proposals"

}

// CRUD Functions

// Get returns a single Proposal as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ProposalDB) Get(ctx goa.Context, id int) Proposal {
	now := time.Now()
	defer ctx.Info("Proposal:Get", "duration", time.Since(now))
	var native Proposal
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	return native
}

// List returns an array of Proposal
func (m *ProposalDB) List(ctx goa.Context) []Proposal {
	now := time.Now()
	defer ctx.Info("Proposal:List", "duration", time.Since(now))
	var objs []Proposal
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.
func (m *ProposalDB) Add(ctx goa.Context, model Proposal) (Proposal, error) {
	now := time.Now()
	defer ctx.Info("Proposal:Add", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Proposal", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ProposalDB) Update(ctx goa.Context, model Proposal) error {
	now := time.Now()
	defer ctx.Info("Proposal:Update", "duration", time.Since(now))
	obj := m.Get(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ProposalDB) Delete(ctx goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("Proposal:Delete", "duration", time.Since(now))
	var obj Proposal

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Proposal", "error", err.Error())
		return err
	}

	return nil
}
