package adapter

import "fmt"

type SimpleRepoImpl struct {
	FnSave      func(*CasbinRule) error
	FnBatchSave func(*[]CasbinRule) error
	FnDelete    func(*CasbinRule) error
	FnFindAll   func() ([]CasbinRule, error)
}

func (r *SimpleRepoImpl) Save(rule *CasbinRule) error {
	if r.FnSave == nil {
		return fmt.Errorf("Save not implemented")
	}
	return r.FnSave(rule)
}
func (r *SimpleRepoImpl) BatchSave(rules *[]CasbinRule) error {
	if r.FnBatchSave == nil {
		return fmt.Errorf("FnBatchSave not implemented")
	}
	return r.FnBatchSave(rules)
}
func (r *SimpleRepoImpl) Delete(rule *CasbinRule) error {
	if r.FnDelete == nil {
		return fmt.Errorf("FnDelete not implemented")
	}
	return r.FnDelete(rule)
}
func (r *SimpleRepoImpl) FindAll() ([]CasbinRule, error) {
	if r.FnFindAll == nil {
		return []CasbinRule{}, fmt.Errorf("FnFindAll not implemented")
	}
	return r.FnFindAll()
}
