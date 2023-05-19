package utils

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Hostname() (string, error) {
	return os.Hostname()
}

func ReadBinaryFile(filename string) ([]byte, error) {
	buf1 := bytes.NewBuffer([]byte{})
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cb := make([]byte, 1024)

	for {
		n, err := f.Read(cb)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		buf1.Write(cb[:n])
	}
	return buf1.Bytes(), nil
}

func MakeDirAll(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func WriteDataToFile(filePath string, data []byte) (int, error) {
	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		return 0, err
	}
	return f.Write(data)
}

func ZipCompress(filePath, dest string) error {
	oriFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer oriFile.Close()

	d, _ := os.Create(dest)
	defer d.Close()

	w := zip.NewWriter(d)
	defer w.Close()

	err = zipCompress(oriFile, "", w)
	if err != nil {
		return err
	}

	return nil
}

func zipCompress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = zipCompress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func FileInfo(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}

func WriteStringsToFile(datas []string, filename string) error {
	// 建sql保存到文件
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开文件出错. %s. %s", filename, err.Error())
	}
	defer f.Close()

	for i, data := range datas {
		_, err := f.WriteString(data)
		if err != nil {
			return fmt.Errorf("第%d行写入文件出错. %s. %s", i, filename, err.Error())
		}
	}
	return nil
}

// 获取文件绝对路径
func FileAbs(fileName string) (string, error) {
	return filepath.Abs(fileName)
}

func Filename(path string) string {
	return filepath.Base(path)
}

func FileDir(fileName string) string {
	return filepath.Dir(fileName)
}

// 文件/目录 是否存在
func PathExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 检测和创建目录是否存在
func CheckAndCreateDir(dir string) error {
	exists, err := PathExists(dir)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// 路径不存在则创建
	if err = MakeDirAll(dir); err != nil {
		return fmt.Errorf("目录: %s 创建失败: %s", dir, err.Error())
	}

	return nil
}

func FileLineCount(name string) (int64, error) {
	f, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fd := bufio.NewReader(f)
	var count int64
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return count, nil
}
