package util

import (
	"fmt"
	"strings"

	"goby/pkg/dict"
)

// ImageExts :
var ImageExts = map[string]string{
	"png":  "png",
	"jpeg": "jpeg",
	"jpg":  "jpg",
	"gif":  "gif",
}

// 文件处理相关

// ByteSize => ByteSize
type ByteSize float64

//
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// FileSizeInString : length of Bytes to size of string
func FileSizeInString(size int64) string {
	return ByteSize(size).String()
}

// GetFileExt :
func GetFileExt(fileName string) string {
	fileNameParts := strings.Split(fileName, dict.Dot)
	return strings.ToLower(fileNameParts[len(fileNameParts)-1])
}

// IsImage :
func IsImage(fileExt string) bool {
	_, ok := ImageExts[fileExt]
	return ok
}

// IsFileImage :
func IsFileImage(fileName string) bool {
	ext := GetFileExt(fileName)
	return IsImage(ext)
}
