package constatnts

const (
	AuthorizationHeader = "Authorization"

	JWTSecretKey             = "secret-key"
	JWTRefreshKeyExpireHours = 24
)

var Roles = map[int]string{
	1: "user",
	2: "admin",
}
