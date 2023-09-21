package tools

import (
	"os"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // 文件存在
	}
	if os.IsNotExist(err) {
		return false // 文件不存在
	}
	// 其他错误，如权限问题等
	return false
}

func Write(context []byte, path string) error {
	err := os.WriteFile(path, context, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read(path string) ([]byte, error) {
	// 使用 os.ReadFile 读取文件内容
	content, err := os.ReadFile(path)
	if err != nil {
		// 如果出现错误，返回 nil 切片和错误信息
		return nil, err
	}

	// 成功时返回文件内容的字节切片
	return content, nil
}
func CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
