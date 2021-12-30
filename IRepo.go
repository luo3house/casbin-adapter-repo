package adapter

type IRepo interface {
	Save(*CasbinRule) error
	BatchSave(*[]CasbinRule) error
	Delete(*CasbinRule) error
	FindAll() ([]CasbinRule, error)
}
