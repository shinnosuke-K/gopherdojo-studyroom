package file

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Dir       string // absolute directory path which files exists
	Name      string // File Name
	Extension string
}

// GetImgFiles 関数のみ使用
// グローバルで宣言していることで、関数が再帰的に呼ばれてもファイルをスライスへ追加していく
var fileList = make([]File, 0)

func GetImgFiles(path string, beforeEx string) ([]File, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			_, err := GetImgFiles(filepath.Join(path, f.Name()), beforeEx)
			if err != nil {
				return nil, err
			}

		} else if filepath.Ext(f.Name()) != "" && filepath.Ext(f.Name())[1:] == beforeEx {
			fileList = append(fileList, File{
				Dir:       path,
				Name:      f.Name(),
				Extension: filepath.Ext(f.Name()),
			})
		}
	}
	return fileList, nil
}

func ExistDir(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}

func DecodeToImg(dir string, name string) (image.Image, error) {
	file, err := os.Open(filepath.Join(dir, name))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imgFile, _, err := image.Decode(file)
	return imgFile, err
}

// ファイルを削除
// delete file
func DeleteImg(files []File) error {
	for n := range files {
		path := filepath.Join(files[n].Dir, files[n].Name)
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("%w: Couldn't delete %s", err, path)
		}
	}
	return nil
}
