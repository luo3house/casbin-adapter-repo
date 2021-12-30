package adapter

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// https://github.com/casbin/gorm-adapter
func loadPolicyLine(line CasbinRule, model model.Model) {
	var p = []string{line.Ptype,
		line.V0, line.V1, line.V2,
		line.V3, line.V4, line.V5}

	index := len(p) - 1
	for p[index] == "" {
		index--
	}
	index += 1
	p = p[:index]

	persist.LoadPolicyArray(p, model)
}
