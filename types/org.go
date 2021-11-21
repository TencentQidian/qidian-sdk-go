package types

type Department struct {
	ID       string `json:"dep_id"`    // 部门id
	Name     string `json:"dep_name"`  // 部门名称
	ParentID string `json:"parent_id"` // 父部门id（如为部门为根部门，则父部门id=0）
	Order    int    `json:"order"`     // 在父部门的排序默认为0
	LeaderID string `json:"leader_id"` // 主管openid，无主管为空
}

type GetDepartmentsRsp struct {
	Total string        `json:"departments_total"` // 部门总数
	Items []*Department `json:"departments_list"`  // 部门列表
}
