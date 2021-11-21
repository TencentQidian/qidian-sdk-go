package types

type Industry struct {
	Id   int64  `json:"id"`   // 所属行业分类id
	Name string `json:"name"` // 所属行业分类名称
}

type Addr struct {
	Id   int64  `json:"id"`   // 地址id
	Name string `json:"name"` // 地址名称
}

type Category struct {
	Id      int64  `json:"id"`      // 所属行业分类ID
	Name    string `json:"name"`    // 所属行业分类名称
	Display int    `json:"display"` // 所属行业分类开关，0:关闭，1:启用
}

// GetCorpInfoRsp ref: https://api.qidian.qq.com/wiki/doc/open/el9var82n7dpz98vg0ip
type GetCorpInfoRsp struct {
	Status             int       `json:"status"`
	CompanyShortName   int       `json:"companyShortName"`   // 企业简称
	CompanyIntroduce   int       `json:"companyIntroduce"`   // 企业简介
	CompanyIndustry    *Industry `json:"companyIndustry"`    // 所属行业分类信息
	CompanyIndustrySub *Industry `json:"companyIndustrySub"` // 所属行业子分类信息
	Country            *Addr     `json:"country"`            // 国家
	Province           *Addr     `json:"province"`           // 省份
	City               *Addr     `json:"city"`               // 城市
	Section            *Addr     `json:"section"`            // 区域
	Address            string    `json:"address"`            // 区域
	CompanyCall        string    `json:"companyCall"`        // 公司电话
	CompanyEmail       string    `json:"companyEmail"`       // 公司邮箱
	CompanyPage        string    `json:"companyPage"`        // 公司网址
	CompanyIcon        string    `json:"companyIcon"`        // 公司头像
	CompanySignature   string    `json:"companySignature"`   // 公司签名
	CompanyDueDate     string    `json:"companyDueDate"`     // 企业号码过期时间
	Category           *Category `json:"category"`           // 所属行业分类信息（新）
	SubCategory        *Category `json:"subCategory"`        // 所属行业子分类信息（新）
}
