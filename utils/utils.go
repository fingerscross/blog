package utils

// 公共方法
// InList 判断key是否存在与列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Reverse 切片翻转
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
