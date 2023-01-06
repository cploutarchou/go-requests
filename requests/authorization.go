package requests

type Authorization interface {
	//	Basic sets the authorization to basic
	Bearer(token string) Authorization

	// Basic sets the authorization to basic
	Basic(username, password string) Authorization

	// String returns the string representation of the authorization
	String() string

	// Type returns the type of the authorization
	Type() AuthorizationType

	// Value returns the value of the authorization
	Value() string

	// IsBasic returns true if the authorization is basic
	IsBasic() bool

	// IsBearer returns true if the authorization is bearer
	IsBearer() bool

	// IsEmpty returns true if the authorization is empty
	IsEmpty() bool

	// IsSet returns true if the authorization is set
	IsSet() bool
}

// AuthorizationType is the type of the authorization header value
type AuthorizationType string

const (
	//AuthorizationTypeBasic is the basic authorization type
	AuthorizationTypeBasic AuthorizationType = "Basic"
	//AuthorizationTypeBearer is the bearer authorization type
	AuthorizationTypeBearer AuthorizationType = "Bearer"
)

// authorizationImpl is the implementation of the authorization interface
type authorizationImpl struct {
	authorizationType AuthorizationType
	value             string
}

// Bearer sets the authorization to bearer
func (a *authorizationImpl) Bearer(token string) Authorization {
	a.authorizationType = AuthorizationTypeBearer
	a.value = token
	return a
}

// Basic sets the authorization to basic
func (a *authorizationImpl) Basic(username, password string) Authorization {
	a.authorizationType = AuthorizationTypeBasic
	a.value = username + ":" + password
	return a
}

// String returns the string representation of the authorization
func (a *authorizationImpl) String() string {
	return string(a.authorizationType) + " " + a.value
}

// Type returns the type of the authorization
func (a *authorizationImpl) Type() AuthorizationType {
	return a.authorizationType
}

// Value returns the value of the authorization
func (a *authorizationImpl) Value() string {
	return a.value
}

// IsBasic returns true if the authorization is basic
func (a *authorizationImpl) IsBasic() bool {
	return a.authorizationType == AuthorizationTypeBasic
}

// IsBearer returns true if the authorization is bearer
func (a *authorizationImpl) IsBearer() bool {
	return a.authorizationType == AuthorizationTypeBearer
}

// IsEmpty returns true if the authorization is empty
func (a *authorizationImpl) IsEmpty() bool {
	return a.authorizationType == "" && a.value == ""
}

// IsSet returns true if the authorization is set
func (a *authorizationImpl) IsSet() bool {
	return !a.IsEmpty()
}

// NewAuthorization creates a new authorization
func NewAuthorization() Authorization {
	return &authorizationImpl{}
}
