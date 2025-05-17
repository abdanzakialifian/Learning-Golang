package simple

type BarRepository struct{}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarSevice struct {
	*BarRepository
}

func NewBarService(repository *BarRepository) *BarSevice {
	return &BarSevice{BarRepository: repository}
}
