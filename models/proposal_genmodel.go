//************************************************************************//
// congo: Models
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
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// app.Proposal storage type
// Identifier: Proposal
type Proposal struct {
	Abstract  string `json:"abstract,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	ID        int    `json:"id,omitempty" gorm:"primary_key"`
	Title     string `json:"title,omitempty"`
	Withdrawn bool   `json:"withdrawn,omitempty"`

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// Foreign Keys
	UserID int

	// Children
	Reviews []Review
}

type ProposalStorage interface {
	DB() interface{}
	List(ctx context.Context) []Proposal
	One(ctx context.Context, id int) (Proposal, error)
	Add(ctx context.Context, o Proposal) (Proposal, error)
	Update(ctx context.Context, o Proposal) error
	Delete(ctx context.Context, id int) error

	ListByUser(ctx context.Context, parentid int) []Proposal
	OneByUser(ctx context.Context, parentid, id int) (Proposal, error)
}
type ProposalDB struct {
	db gorm.DB
}

func ProposalFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func (m *ProposalDB) ListByUser(ctx context.Context, parentid int) []Proposal {

	var objs []Proposal
	m.db.Scopes(ProposalFilterByUser(parentid, &m.db)).Find(&objs)
	return objs
}

func (m *ProposalDB) OneByUser(ctx context.Context, parentid, id int) (Proposal, error) {

	var obj Proposal

	err := m.db.Scopes(ProposalFilterByUser(parentid, &m.db)).Find(&obj, id).Error

	return obj, err
}

func NewProposalDB(db gorm.DB) *ProposalDB {

	return &ProposalDB{db: db}

}

func (m *ProposalDB) DB() interface{} {
	return &m.db
}

func (m *ProposalDB) List(ctx context.Context) []Proposal {

	var objs []Proposal
	m.db.Find(&objs)
	return objs
}

func (m *ProposalDB) ListByAbstractEqual(ctx context.Context, abstract string) []Proposal {

	var objs []Proposal
	m.db.Where("abstract = ?", abstract).Find(&objs)
	return objs
}
func (m *ProposalDB) ListByAbstractLike(ctx context.Context, abstract string) []Proposal {

	var objs []Proposal
	m.db.Where("abstract like ?", abstract).Find(&objs)
	return objs
}

func (m *ProposalDB) ListByDetailEqual(ctx context.Context, detail string) []Proposal {

	var objs []Proposal
	m.db.Where("detail = ?", detail).Find(&objs)
	return objs
}
func (m *ProposalDB) ListByDetailLike(ctx context.Context, detail string) []Proposal {

	var objs []Proposal
	m.db.Where("detail like ?", detail).Find(&objs)
	return objs
}

func (m *ProposalDB) ListByFirstnameEqual(ctx context.Context, firstname string) []Proposal {

	var objs []Proposal
	m.db.Where("firstname = ?", firstname).Find(&objs)
	return objs
}
func (m *ProposalDB) ListByFirstnameLike(ctx context.Context, firstname string) []Proposal {

	var objs []Proposal
	m.db.Where("firstname like ?", firstname).Find(&objs)
	return objs
}

func (m *ProposalDB) ListByIdEqual(ctx context.Context, id int) []Proposal {

	var objs []Proposal
	m.db.Where("id = ?", id).Find(&objs)
	return objs
}
func (m *ProposalDB) ListByIdLike(ctx context.Context, id int) []Proposal {

	var objs []Proposal
	m.db.Where("id like ?", id).Find(&objs)
	return objs
}

func (m *ProposalDB) ListByTitleEqual(ctx context.Context, title string) []Proposal {

	var objs []Proposal
	m.db.Where("title = ?", title).Find(&objs)
	return objs
}
func (m *ProposalDB) ListByTitleLike(ctx context.Context, title string) []Proposal {

	var objs []Proposal
	m.db.Where("title like ?", title).Find(&objs)
	return objs
}

func (m *ProposalDB) ListByWithdrawnEqual(ctx context.Context, withdrawn bool) []Proposal {

	var objs []Proposal
	m.db.Where("withdrawn = ?", withdrawn).Find(&objs)
	return objs
}
func (m *ProposalDB) ListByWithdrawnLike(ctx context.Context, withdrawn bool) []Proposal {

	var objs []Proposal
	m.db.Where("withdrawn like ?", withdrawn).Find(&objs)
	return objs
}

func (m *ProposalDB) One(ctx context.Context, id int) (Proposal, error) {

	var obj Proposal

	err := m.db.Find(&obj, id).Error

	return obj, err
}

func (m *ProposalDB) Add(ctx context.Context, model Proposal) (Proposal, error) {
	err := m.db.Create(&model).Error

	return model, err
}

func (m *ProposalDB) Update(ctx context.Context, model Proposal) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.db.Model(&obj).Updates(model).Error

	return err
}

func (m *ProposalDB) Delete(ctx context.Context, id int) error {
	var obj Proposal
	err := m.db.Delete(&obj, id).Error
	if err != nil {
		return err
	}

	return nil
}

func FilterProposalByUser(parent int, list []Proposal) []Proposal {
	filtered := make([]Proposal, 0)
	for _, o := range list {
		if o.UserID == int(parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}
