package middleware

import (
	"log"

	constants "test_service_filmoteka/pkg/constatnts"

	casbin "github.com/casbin/casbin/v2"
)

// JWTRoleAuthorizer is a sturcture for a Role Authorizer type
type JWTRoleAuthorizer struct {
	enforcer interface {
		Enforce(rvals ...interface{}) (bool, error)
	}
	signingKey string
	//	logger     logger.Logger
}

// NewCasbinJWTRoleAuthorizer creates and returns new Role Authorizer
func NewCasbinJWTRoleAuthorizer() (*JWTRoleAuthorizer, error) {

	// enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", "config/models.csv")
	if err != nil {
		log.Println("could not initialize new enforcer:", err.Error())
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		signingKey: constants.JWTSecretKey,
		//		logger:     logger,
	}, nil
}

// IsSuperAdmin checks for if user is superAdmin
func IsSuperAdmin(username, password string) bool {
	pass := constants.SuperAdminPassword
	user := constants.SuperAdminUsername

	return password == pass && user == username
}
