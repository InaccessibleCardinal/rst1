package svc

type ServiceResponse[T any] struct {
	IsOk  bool
	Value T
	Error error
}

type BasicCrudService[T any] interface {
	FindAll() ServiceResponse[[]T]
	FindById(id int) ServiceResponse[T]
	Update(t *T) ServiceResponse[*T]
}
