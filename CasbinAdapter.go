package adapter

import (
	"fmt"

	"github.com/casbin/casbin/v2/model"
)

type CasbinAdapter struct {
	repo      IRepo
	batchSize int
}

func (adapter *CasbinAdapter) LoadPolicy(model model.Model) error {
	lines, err := adapter.repo.FindAll()
	if err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return nil
}

func (adapter *CasbinAdapter) SavePolicy(model model.Model) error {
	var lines []CasbinRule
	flushEvery := adapter.GetBatchSize()
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			lines = append(lines, *NewModelInstanceFromPTypeAndRules(ptype, rule))
			if len(lines) > flushEvery {
				if err := adapter.repo.BatchSave(&lines); err != nil {
					return err
				}
				lines = nil
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			lines = append(lines, *NewModelInstanceFromPTypeAndRules(ptype, rule))
			if len(lines) > flushEvery {
				if err := adapter.repo.BatchSave(&lines); err != nil {
					return err
				}
				lines = nil
			}
		}
	}
	return nil
}

func (adapter *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return adapter.repo.Save(NewModelInstanceFromPTypeAndRules(ptype, rule))
}

func (adapter *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return adapter.repo.Delete(NewModelInstanceFromPTypeAndRules(ptype, rule))
}

func (adapter *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return fmt.Errorf("not implemented")
}
