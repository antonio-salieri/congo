//************************************************************************//
// congo: Media Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package user

import (
	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/copier"
)

func (m User) ToDefault() *app.User {
	target := app.User{}
	copier.Copy(&target, &m)
	return &target
}
