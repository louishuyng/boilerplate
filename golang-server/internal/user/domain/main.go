package domain

var _ UserDomain = (*Domain)(nil)

type Domain struct {
}

func NewUserDomain() *Domain {
	return &Domain{}
}
