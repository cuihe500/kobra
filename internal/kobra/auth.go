package kobra

type Authentication struct {
	Role   string
	Path   string
	Method []string
}

var Authentications = make([]Authentication, 0)

func (auth *Authentication) SetRole(role string) *Authentication {
	auth.Role = role
	return auth
}
func (auth *Authentication) SetPath(path string) *Authentication {
	auth.Path = path
	return auth
}
func (auth *Authentication) SetMethod(method ...string) *Authentication {
	auth.Method = method
	return auth
}
func (auth *Authentication) GetCasbinAuthenticationConfig() []string {
	if len(auth.Method) == 1 {
		return []string{auth.Role, auth.Path, auth.Method[0]}
	}
	method := ""
	for n, v := range auth.Method {
		if n < len(auth.Method)-1 {
			method = method + "(" + v + ")" + "|"
			continue
		}
		method = method + "(" + v + ")"
	}
	return []string{auth.Role, auth.Path, method}
}
