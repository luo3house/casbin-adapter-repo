package adapter

// https://github.com/casbin/gorm-adapter
type CasbinRule struct {
	ID    uint   `json:"id"`
	Ptype string `json:"ptype"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

func NewModelInstance() *CasbinRule {
	return new(CasbinRule)
}
