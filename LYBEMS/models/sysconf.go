package sysconf

// Input 输入
type Input struct {
	//view:用于注释
	//json:json形式
	//from:解释该字段来自哪里，比如那个表
	//binding: required:必须字段 email:email形式
	//grom:数据库中列名
	ID int `view:"ID号" json:"id" from:"id" binding:"required" gorm:"column:id"`
}

//Output 类 用于解析配置中的 数据库信息
type Output struct {
	Database DatabaseConfig `json:"database"`
	Self     SelfConfig     `json:"self"`
}

// DatabaseConfig 数据库详细配置
type DatabaseConfig struct {
	Types  string `json:"types"`
	Local  string `json:"local"`
	Online string `json:"online"`
}

//SelfConfig 配置端口等信心
type SelfConfig struct {
	Port string `json:"port"`
	Flag int    `json:"flag"`
	Tag  int    `json:"tag"`
}
