package untils

import (
	"encoding/base64"
	"io"
	"mime/multipart"
	"os"
)

func Base642Img(base string, path string, fileName string) error {
	dec, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return err
	}
	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func File2Img(file *multipart.FileHeader, path string, name string) error {
	w, err := file.Open()
	if err != nil {
		return err
	}
	defer w.Close()
	f, err := os.Create(path + "/" + name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, w)
	if err != nil {
		return err
	}
	return nil
}
