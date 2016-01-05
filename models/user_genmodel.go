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

	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// app.UserModel storage type
// Identifier:
type User struct {
	ID        int    `json:"ID,omitempty" gorm:"primary_key"`
	Bio       string `json:"bio,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Role      string `json:"role,omitempty"`
	State     string `json:"state,omitempty"`

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// Children
	Proposals []Proposal
	Reviews   []Review

	// Authboss

	// Auth
	Password string

	// OAuth2
	Oauth2Uid      string
	Oauth2Provider string
	Oauth2Token    string
	Oauth2Refresh  string
	Oauth2Expiry   time.Time

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Lock
	AttemptNumber int64
	AttemptTime   time.Time
	Locked        time.Time

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time
}

func UserFromCreatePayload(ctx *app.CreateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)

	return m
}

func UserFromUpdatePayload(ctx *app.UpdateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)
	return m
}

func (m User) ToApp() *app.User {
	target := app.User{}
	copier.Copy(&target, &m)
	return &target
}

func (m User) GetRole() string {
	return m.Role
}

type UserStorage interface {
	List(ctx context.Context) []User
	One(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, o User) (User, error)
	Update(ctx context.Context, o User) error
	Delete(ctx context.Context, id int) error
}

type UserDB struct {
	DB gorm.DB
}

func NewUserDB(db gorm.DB) *UserDB {

	return &UserDB{DB: db}

}

func (m *UserDB) List(ctx context.Context) []User {

	var objs []User
	m.DB.Find(&objs)
	return objs
}

func (m *UserDB) ListByIDEqual(ctx context.Context, id int) []User {

	var objs []User
	m.DB.Where("id = ?", id).Find(&objs)
	return objs
}
func (m *UserDB) ListByIDLike(ctx context.Context, id int) []User {

	var objs []User
	m.DB.Where("id like ?", id).Find(&objs)
	return objs
}

func (m *UserDB) ListByBioEqual(ctx context.Context, bio string) []User {

	var objs []User
	m.DB.Where("bio = ?", bio).Find(&objs)
	return objs
}
func (m *UserDB) ListByBioLike(ctx context.Context, bio string) []User {

	var objs []User
	m.DB.Where("bio like ?", bio).Find(&objs)
	return objs
}

func (m *UserDB) ListByCityEqual(ctx context.Context, city string) []User {

	var objs []User
	m.DB.Where("city = ?", city).Find(&objs)
	return objs
}
func (m *UserDB) ListByCityLike(ctx context.Context, city string) []User {

	var objs []User
	m.DB.Where("city like ?", city).Find(&objs)
	return objs
}

func (m *UserDB) ListByCountryEqual(ctx context.Context, country string) []User {

	var objs []User
	m.DB.Where("country = ?", country).Find(&objs)
	return objs
}
func (m *UserDB) ListByCountryLike(ctx context.Context, country string) []User {

	var objs []User
	m.DB.Where("country like ?", country).Find(&objs)
	return objs
}

func (m *UserDB) ListByEmailEqual(ctx context.Context, email string) []User {

	var objs []User
	m.DB.Where("email = ?", email).Find(&objs)
	return objs
}
func (m *UserDB) ListByEmailLike(ctx context.Context, email string) []User {

	var objs []User
	m.DB.Where("email like ?", email).Find(&objs)
	return objs
}

func (m *UserDB) ListByFirstnameEqual(ctx context.Context, firstname string) []User {

	var objs []User
	m.DB.Where("firstname = ?", firstname).Find(&objs)
	return objs
}
func (m *UserDB) ListByFirstnameLike(ctx context.Context, firstname string) []User {

	var objs []User
	m.DB.Where("firstname like ?", firstname).Find(&objs)
	return objs
}

func (m *UserDB) ListByLastnameEqual(ctx context.Context, lastname string) []User {

	var objs []User
	m.DB.Where("lastname = ?", lastname).Find(&objs)
	return objs
}
func (m *UserDB) ListByLastnameLike(ctx context.Context, lastname string) []User {

	var objs []User
	m.DB.Where("lastname like ?", lastname).Find(&objs)
	return objs
}

func (m *UserDB) ListByRoleEqual(ctx context.Context, role string) []User {

	var objs []User
	m.DB.Where("role = ?", role).Find(&objs)
	return objs
}
func (m *UserDB) ListByRoleLike(ctx context.Context, role string) []User {

	var objs []User
	m.DB.Where("role like ?", role).Find(&objs)
	return objs
}

func (m *UserDB) ListByStateEqual(ctx context.Context, state string) []User {

	var objs []User
	m.DB.Where("state = ?", state).Find(&objs)
	return objs
}
func (m *UserDB) ListByStateLike(ctx context.Context, state string) []User {

	var objs []User
	m.DB.Where("state like ?", state).Find(&objs)
	return objs
}

func (m *UserDB) One(ctx context.Context, id int) (User, error) {

	var obj User

	err := m.DB.Find(&obj, id).Error

	return obj, err
}

func (m *UserDB) Add(ctx context.Context, model User) (User, error) {
	err := m.DB.Create(&model).Error

	return model, err
}

func (m *UserDB) Update(ctx context.Context, model User) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(model).Error

	return err
}

func (m *UserDB) Delete(ctx context.Context, id int) error {
	var obj User
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		return err
	}

	return nil
}
