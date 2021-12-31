# Casbin Adapter Repo

参考 github.com/casbin/gorm-adapter，因为这个适配器把数据库操作也写进去了（如: Create, Delete）。

这个仓库把数据库操作的那部分抽离出来，方便给 Casbin 的模型拓展其他的字段。

结构: Repo -> Adapter -> Enforcer

## 嵌入 Enforcer

~~~ go
Repo := adapter.SimpleRepoImpl {
  // 自己实现 CasbinRule 的读写
  FnFindAll: func() ([]CasbinRule, error) {
    return []CasbinRule{
      // ID, PType, V0, V1, V2, V3, V4, V5
      {1, "p", "alice", "data1", "read", "", "", ""},
      {2, "p", "alice", "data1", "read", "", "", ""},
      {3, "p", "alice", "data2", "read", "", "", ""},
      {4, "g", "alice", "data1_admin", "", "", "", ""},
    }, nil
  },
}
Adapter, _ := new(adapter.CasbinAdapter)
Adapter.SetRepo(Repo)
Enforcer, _ := casbin.NewEnforcer(<Matcher>, Adapter)
Enforcer.LoadPolicy()
// 开始验证
Enforcer.Enforce(...)
~~~
