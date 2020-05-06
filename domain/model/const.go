package model

// DomainModelName is Model name for developer.
type DomainModelName string

// String return as string.
func (p DomainModelName) String() string {
	return string(p)
}

// PropertyName is property name for developer.
type PropertyName string

// String return as string.
func (p PropertyName) String() string {
	return string(p)
}

// Property name for developer.
const (
	IDProperty      PropertyName = "ID"
	AccountProperty PropertyName = "Account"
)

// RepositoryMethod is method of Repository.
type RepositoryMethod string

// methods of Repository.
const (
	RepositoryMethodRead   RepositoryMethod = "READ"
	RepositoryMethodInsert RepositoryMethod = "INSERT"
	RepositoryMethodUpdate RepositoryMethod = "UPDATE"
	RepositoryMethodDelete RepositoryMethod = "DELETE"
	RepositoryMethodList   RepositoryMethod = "LIST"
)
