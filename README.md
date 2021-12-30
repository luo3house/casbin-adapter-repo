# Casbin Adapter Repo

参考 github.com/casbin/gorm-adapter，因为这个适配器把数据库操作也写进去了（如: Create, Delete）。

这个仓库把数据库操作的那部分抽离出来，方便给 Casbin 的模型拓展其他的字段。

e.g. 简单的 gorm 实现 Repo，但是不限 gorm
~~~ go
type GormRepo struct { 
  db *gorm.DB 
}

func (r *GormRepo) Save(mdl *CasbinRule) error {
  return r.db.Save(mdl).Error
}

func (r *GormRepo) BatchSave(rules *[]CasbinRule) error {
  tx := r.db.Begin()
  for _, rule := range rules {
    tx = tx.Save(*rules)
  }
  if err := tx.Commit().Error; err != nil {
    tx.Rollback()
    return err
  }
  return nil
}

func (r *GormRepo) Delete(rule *CasbinRule) error {
  // e.g. 可以实现自己的删除逻辑
  Id := rule.ID
  // @BizExamine
  return r.db.Raw(`DELETE FROM my_casbin_table WHERE id = ?`, Id).Error
}

func (r *GormRepo) FindAll() ([]CasbinRule, error) {
  rules := []CasbinRule{}
  // e.g. 可以限定某个子系统的 policy 记录
  return r.db.Where("subsystem = ?", "财务小组").Find(&rules)
  return rules, nil
}
~~~

嵌入 Enforcer
~~~ go
Adapter, _ := new(adapter.CasbinAdapter)
Adapter.SetRepo(&GormRepo{ 
  // 自己传数据库
  db: gorm.Open(`<URL>`) 
})
Enforcer, _ := casbin.NewEnforcer(<Matcher>, Adapter)
// 开始验证
Enforcer.Enforce(...)
~~~
