package constatnts

const (
	AuthorizationHeader = "Authorization"
	UnAuthorized        = "unauthorized"

	JWTSecretKey             = "secret-key"
	JWTKEY                   = "Apex-Hello-09092834j)(*OIJ1k0989)"
	JWTRefreshKeyExpireHours = 24
	SuperAdminPassword       = "user"
	SuperAdminUsername       = "pass"
)

var Roles = map[int]string{
	1: "user",
	2: "admin",
}
