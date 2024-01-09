package config

// 要显示的信息
type SiteInfo struct {
	CreatedAt string `yaml:"created_at" json:"created_at"`
	Title     string `yaml:"title" json:"title"`
	Email     string `yaml:"eamil" json:"email"`
}
