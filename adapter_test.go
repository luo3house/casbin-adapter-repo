package adapter

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

const Matcher = `
[request_definition]
r = user, perm

[policy_definition]
p = user, perm

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.user == p.user && r.perm == p.perm
`

var (
	Enforcer *casbin.Enforcer
	Adapter  persist.Adapter
	Repo     IRepo
)

type CustomizeRepo struct {
	FnSave      func(*CasbinRule) error
	FnBatchSave func(*[]CasbinRule) error
	FnDelete    func(*CasbinRule) error
	FnFindAll   func() ([]CasbinRule, error)
}

func (r *CustomizeRepo) Save(rule *CasbinRule) error {
	return r.FnSave(rule)
}
func (r *CustomizeRepo) BatchSave(rules *[]CasbinRule) error {
	return r.FnBatchSave(rules)
}
func (r *CustomizeRepo) Delete(rule *CasbinRule) error {
	return r.FnDelete(rule)
}
func (r *CustomizeRepo) FindAll() ([]CasbinRule, error) {
	return r.FnFindAll()
}

func InitEnforcer() {
	mdl, _ := model.NewModelFromString(Matcher)
	Enforcer, _ = casbin.NewEnforcer(mdl, Adapter)
}

func InitAdapter() {
	Adapter = new(CasbinAdapter).SetRepo(Repo)
}

func TestLoad(t *testing.T) {
	Repo = &CustomizeRepo{
		FnFindAll: func() ([]CasbinRule, error) {
			return []CasbinRule{
				// ID, PType, V0, V1, V2, V3, V4, V5
				{1, "p", "luo1", "learn", "", "", "", ""},
				{2, "p", "luo1", "gaming", "", "", "", ""},
				{3, "p", "luo2", "learn", "", "", "", ""},
				{4, "p", "luo3", "gaming", "", "", "", ""},
			}, nil
		},
	}
	InitAdapter()
	InitEnforcer()
	var ok bool
	// luo3 can only gaming
	ok, _ = Enforcer.Enforce("luo3", "learn")
	if ok {
		t.Error(`luo3 can only gaming`)
	}
	ok, _ = Enforcer.Enforce("luo3", "gaming")
	if !ok {
		t.Error(`luo3 can gaming`)
	}
	// luo2 can only learn
	ok, _ = Enforcer.Enforce("luo2", "gaming")
	if ok {
		t.Error(`luo2 can only learn`)
	}
	ok, _ = Enforcer.Enforce("luo2", "learn")
	if !ok {
		t.Error(`luo2 can learn`)
	}
	// luo1 can not go outside
	ok, _ = Enforcer.Enforce("luo1", "gaming")
	if !ok {
		t.Error(`luo1 can gaming`)
	}
	ok, _ = Enforcer.Enforce("luo1", "learn")
	if !ok {
		t.Error(`luo1 can learn`)
	}
	ok, _ = Enforcer.Enforce("luo1", "buybuybuy")
	if ok {
		t.Error(`luo1 can not buybuybuy`)
	}
}
