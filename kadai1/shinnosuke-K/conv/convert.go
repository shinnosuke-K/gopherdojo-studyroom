package conv

import (
	"errors"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai1/shinnosuke-K/file"
)

// Extensions
const (
	PNG  string = "png"
	JPG  string = "jpg"
	JPEG string = "jpeg"
	GIF  string = "gif"
)

func Do(dirPath string, before string, after string, delImg bool) error {

	if ok := file.ExistDir(dirPath); !ok {
		return errors.New("not found directory")
	}

	if err := checkOpt(before); err != nil {
		return err
	}

	if err := checkOpt(after); err != nil {
		return err
	}

	files, err := file.GetImgFiles(dirPath, before)
	if err != nil {
		return err
	}

	for n := range files {
		if err := convert(after, files[n]); err != nil {
			return err
		}
	}

	if delImg {
		if err := file.DeleteImg(files); err != nil {
			return err
		}
	}
	return nil
}

// 指定した拡張子が正しいか確認
// Check that the extension you specified is correct.
func checkOpt(ex string) error {
	imgExts := []string{PNG, JPG, JPEG, GIF}
	for n := range imgExts {
		if strings.ToLower(ex) == imgExts[n] {
			return nil
		}
	}
	return fmt.Errorf("image convert error: invalid image extension '%s'", ex)
}

func convert(afterExt string, f file.File) error {
	imgFile, err := file.DecodeToImg(f.Dir, f.Name)
	if err != nil {
		return err
	}

	destFileName := f.Name[:strings.LastIndex(f.Name, f.Extension)] + "." + afterExt
	destFile, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	switch afterExt {
	case PNG:
		if err := png.Encode(destFile, imgFile); err != nil {
			return err
		}
	case JPEG, JPG:
		if err := jpeg.Encode(destFile, imgFile, nil); err != nil {
			return err
		}
	case GIF:
		if err := gif.Encode(destFile, imgFile, nil); err != nil {
			return err
		}
	}
	return nil
}
