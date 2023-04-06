package mongodoc

import (
	"time"

	"github.com/reearth/reearth/server/pkg/id"
	"github.com/reearth/reearth/server/pkg/user"
	user1 "github.com/reearth/reearth/server/pkg/user"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/util"
)

type PasswordResetDocument struct {
	Token     string
	CreatedAt time.Time
}

type UserDocument struct {
	ID            string
	Name          string
	Email         string
	Auth0Sub      string `bson:"auth0sub,omitempty"`
	Auth0SubList  []string
	Workspace     string `bson:"team"` // DON'T CHANGE NAME
	Lang          string
	Theme         string
	Password      []byte
	PasswordReset *PasswordResetDocument
	Verification  *UserVerificationDoc
}

type UserVerificationDoc struct {
	Code       string
	Expiration time.Time
	Verified   bool
}

type UserConsumer = mongox.SliceFuncConsumer[*UserDocument, *user.User]

func NewUserConsumer() *UserConsumer {
	return NewComsumer[*UserDocument, *user.User]()
}

func NewUser(user *user1.User) (*UserDocument, string) {
	id := user.ID().String()
	auths := user.Auths()
	authsdoc := make([]string, 0, len(auths))
	for _, a := range auths {
		authsdoc = append(authsdoc, a.Sub)
	}
	var v *UserVerificationDoc
	if user.Verification() != nil {
		v = &UserVerificationDoc{
			Code:       user.Verification().Code(),
			Expiration: user.Verification().Expiration(),
			Verified:   user.Verification().IsVerified(),
		}
	}
	pwdReset := user.PasswordReset()

	var pwdResetDoc *PasswordResetDocument
	if pwdReset != nil {
		pwdResetDoc = &PasswordResetDocument{
			Token:     pwdReset.Token,
			CreatedAt: pwdReset.CreatedAt,
		}
	}

	return &UserDocument{
		ID:            id,
		Name:          user.Name(),
		Email:         user.Email(),
		Auth0SubList:  authsdoc,
		Workspace:     user.Workspace().String(),
		Lang:          user.Lang().String(),
		Theme:         string(user.Theme()),
		Verification:  v,
		Password:      user.Password(),
		PasswordReset: pwdResetDoc,
	}, id
}

func (d *UserDocument) Model() (*user1.User, error) {
	uid, err := id.UserIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	tid, err := id.WorkspaceIDFrom(d.Workspace)
	if err != nil {
		return nil, err
	}

	auths := util.Map(d.Auth0SubList, func(s string) user.Auth { return user.AuthFrom(s) })
	if d.Auth0Sub != "" {
		auths = append(auths, user.AuthFrom(d.Auth0Sub))
	}
	var v *user.Verification
	if d.Verification != nil {
		v = user.VerificationFrom(d.Verification.Code, d.Verification.Expiration, d.Verification.Verified)
	}

	u, err := user1.New().
		ID(uid).
		Name(d.Name).
		Email(d.Email).
		Auths(auths).
		Workspace(tid).
		LangFrom(d.Lang).
		Verification(v).
		EncodedPassword(d.Password).
		PasswordReset(d.PasswordReset.Model()).
		Theme(user.Theme(d.Theme)).
		Build()

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *PasswordResetDocument) Model() *user1.PasswordReset {
	if d == nil {
		return nil
	}
	return &user1.PasswordReset{
		Token:     d.Token,
		CreatedAt: d.CreatedAt,
	}
}
