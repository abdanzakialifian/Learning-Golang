package simple

type FooBarService struct {
	*FooService
	*BarSevice
}

func NewFooBarService(fooService *FooService, barService *BarSevice) *FooBarService {
	return &FooBarService{FooService: fooService, BarSevice: barService}
}
