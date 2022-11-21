package utils

import (
	"io/ioutil"
	"path"
)

type FilesList struct {
	Index  int    `json:"index"`
	Name   string `json:"name"`   // 文件名
	Type   string `json:"type"`   // 文件类型 file or dir
	Size   int64  `json:"size"`   // 文件大小
	Time   int64  `json:"time"`   // 创建时间
	Mode   string `json:"mode"`   // 权限
	Suffix string `json:"suffix"` // 文件后缀
}

// GetFilesList 读取文件列表
func GetFilesList(dir string) ([]FilesList, error) {

	var files []FilesList
	var dirs []FilesList

	list, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for index, f := range list {
		if f.IsDir() {
			dirs = append(dirs, FilesList{
				Index: index + 1,
				Name:  f.Name(),
				Type:  "dir",
				Size:  f.Size(),
				Time:  f.ModTime().Unix(),
				Mode:  f.Mode().String(),
			})
		} else {
			files = append(files, FilesList{
				Index:  index + 1,
				Name:   f.Name(),
				Type:   "file",
				Size:   f.Size(),
				Time:   f.ModTime().Unix(),
				Mode:   f.Mode().String(),
				Suffix: path.Ext(f.Name()),
			})
		}
	}

	return append(dirs, files...), nil
}

func Error(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
}

func JSON(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": 200,
		"msg":  msg,
		"data": data,
	}
}
