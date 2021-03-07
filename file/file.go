package file

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// PathExist 判断目录是否存在
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func FileCreate(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	log.Printf("h.OldText: %s \n", h.OldText)
	log.Printf("h.NewText: %s \n", h.NewText)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, []byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		if ctx.Err() != nil {
			break
		}
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(ctx, id, group, line)
	}
}

// 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

//获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// 获取文件类型
const (
	documents = "txt doc pdf ppt ppS xl5x xls docx"
	music     = "mp3 wav wma mpa ram ra aac aif m4a"
	video     = "avi mpg mpe mpeg asf wmv mov qt rm mp4 flv m4v webm ogv ogg"
	image     = "bmp dib pcp dif wmf gif jpg tif eps psd cdr iff tga pcd mpt png jpeg"
)

// 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}

}

// GetSize get the img size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt get the img ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the img exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the img has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Open a img according to a specific mode
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the img
func MustOpenFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("img.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("img.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := OpenFile(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

// GetFileType 获取文件类型
func GetFileType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
	}

	filetype := http.DetectContentType(buff)

	//ext := GetExt(p)
	//var list = strings.Split(filetype, "/")
	//filetype = list[0] + "/" + ext
	return filetype, nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// NormalizeEOL will convert Windows (CRLF) and Mac (CR) EOLs to UNIX (LF)
func NormalizeEOL(input []byte) []byte {
	var right, left, pos int
	if right = bytes.IndexByte(input, '\r'); right == -1 {
		return input
	}
	length := len(input)
	tmp := make([]byte, length)

	// We know that left < length because otherwise right would be -1 from IndexByte.
	copy(tmp[pos:pos+right], input[left:left+right])
	pos += right
	tmp[pos] = '\n'
	left += right + 1
	pos++

	for left < length {
		if input[left] == '\n' {
			left++
		}

		right = bytes.IndexByte(input[left:], '\r')
		if right == -1 {
			copy(tmp[pos:], input[left:])
			pos += length - left
			break
		}
		copy(tmp[pos:pos+right], input[left:left+right])
		pos += right
		tmp[pos] = '\n'
		left += right + 1
		pos++
	}
	return tmp[:pos]
}

func IsDirectory(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func MkdirIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
}

func MkFileIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, _ = os.Create(path)
	}
}

func UploadFileTo(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}

func GetParentDirectory(directory string) string {
	runes := []rune(directory)
	l := 0 + strings.LastIndex(directory, "/")
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[0:l])
}

func ParseFileContentType(fileName string) string {
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if strings.HasPrefix(contentType, "text/") {
		contentType = "text/plain"
	}
	return contentType
}

func IsHiddenFile(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}

	return strings.HasPrefix(name, ".")
}

func ByteCountIEC(b int) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

type (
	FileInfoList []*FileInfo
	FileInfo     struct {
		Filename string `form:"filename" json:"filename"` //原始文件名（test.jpg）
		Data     []byte `form:"data" json:"data"`         //文件字节切片
		Size     int64  `form:"size" json:"size"`         //大小（单位：字节）
		Duration int64  `form:"duration" json:"duration"` //时长（单位：秒）
		Path     string `form:"path" json:"path"`         //全路径（本地磁盘或第三方文件系统）
	}

	IFileSize interface {
		Size() int64
	}
)

/*
 * 获取文件绝对全路径
 *  */
func GetAbsolutePath(filePath string) string {
	if !strings.HasPrefix(filePath, string(os.PathSeparator)) {
		filePath = fmt.Sprintf("%s%s%s", GetCurrentPath(), string(os.PathSeparator), filePath)
	}

	return filePath
}

/*
 * 获取全文件路径的相对路径
 *  */
func GetRelativePath(fullpath string) string {
	currentFullPath := GetCurrentPath()
	//path := strings.Replace(fullpath, currentFullPath, "", -1)
	//splitstring := strings.Split(path, string(os.PathSeparator))
	//return strings.Join(splitstring[1:], string(os.PathSeparator))
	relPath, _ := filepath.Rel(currentFullPath, fullpath)
	return relPath
}

/*
 * 创建多级目录
 *  */
func CreateDir(perm os.FileMode, args ...string) (string, error) {
	dirs := strings.Join(args, string(os.PathSeparator))
	err := os.MkdirAll(dirs, perm)
	if err != nil {
		return "", err
	}

	return dirs, nil
}

/*
 * 根据指定的日期在指定的目录下创建多级目录
 *  */
func CreateDateDir(rootPath string, datetime time.Time, perm os.FileMode) (string, error) {
	year, month, day := datetime.Date()
	sYear := fmt.Sprintf("%d", year)
	sMonth := fmt.Sprintf("%02d", month)
	sDay := fmt.Sprintf("%02d", day)

	return CreateDir(perm, rootPath, sYear, sMonth, sDay)
}

/*
 * 根据当前日期在指定根目录下创建多级目录
 *  */
func CreateCurrentDateDir(rootPath string, perm os.FileMode) (string, error) {
	nowDate := time.Now()
	return CreateDateDir(rootPath, nowDate, perm)
}

/*
 * 获取路径里的路径和文件名
 *  */
func GetFilePath(filePath string) (string, string) {
	path := ""
	filename := ""
	paths := strings.Split(filePath, string(os.PathSeparator))
	length := len(paths)
	if length > 0 {
		path = strings.Join(paths[0:length-1], string(os.PathSeparator))
		filename = paths[length-1 : length][0]
	}

	return path, filename
}

/*
 * 获取路径里的文件名，不带扩展名的文件名，扩展名
 *  */
func GetFilename(filePath string) (string, string, string) {
	paths := strings.Split(filePath, string(os.PathSeparator))
	filename := ""
	filenameWithoutExtname := ""
	extname := ""

	if len(paths) > 0 {
		filename = paths[len(paths)-1]
		filenames := strings.Split(filename, ".")
		filenameWithoutExtname = filenames[0]
		extname = filenames[1]
	}

	return filename, filenameWithoutExtname, extname
}

/*
 * 获取文件内容
 *  */
func GetFileContent(fullFilename string) ([]byte, error) {
	fileStream, err := os.Open(fullFilename)
	if err != nil {
		return nil, err
	}

	defer fileStream.Close()

	fileContent, err := ioutil.ReadAll(fileStream)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

/*
 * 获取Http请求里的文件数据
 * maxSize: 文件大小限制，0表示不限制
 *  */
func GetHttpRequestFile(req *http.Request, args ...int32) (*FileInfo, error) {
	//获取请求文件
	fileStream, fileHeader, err := req.FormFile("img")
	if err != nil {
		return nil, err
	}
	defer fileStream.Close()

	var maxSize int32
	if len(args) > 0 {
		maxSize = args[0]
	}

	//判断大小是否超出限制
	size := fileStream.(IFileSize).Size()
	if maxSize != 0 && size > int64(maxSize) {
		return nil, errors.New("img is too large")
	}

	//读取文件数据到字节切片
	dataBytes, err := ioutil.ReadAll(fileStream)
	if err != nil {
		log.Fatal(err)
	}

	//返回数据结果
	fileInfo := &FileInfo{
		Filename: fileHeader.Filename,
		Data:     dataBytes,
		Size:     size,
	}

	return fileInfo, nil
}

/*
 * 保存Http 上传的文件到磁盘指定目录（返回客户端原文件名，大小，全文件路径，错误）
 *  */
func SaveHttpFile(req *http.Request, filename, basePath string, maxSize int32, args ...string) (*FileInfo, error) {
	fileInfo, err := GetHttpRequestFile(req, maxSize)
	if err != nil {
		return nil, err
	}

	fullFilename, err := SaveFile(fileInfo.Data, filename, basePath, args...)
	if err != nil {
		return nil, err
	}

	fileInfo.Path = fullFilename
	log.Printf("SaveHttpFile Filename: %s, fullFilename: %s", fileInfo.Filename, fileInfo.Path)

	return fileInfo, nil
}

/*
 * 保存文件
 *  */
func SaveFile(data []byte, filename, basePath string, args ...string) (string, error) {
	rootPath, _ := filepath.Abs(basePath)

	if len(args) == 0 {
		//默认保存在上传目录下，且生成当前日期目录
		rootPath, _ = CreateDateDir(rootPath, time.Now(), 0755)
	} else if len(args) == 1 {
		//保存到指定目录
		basePath = args[0]
		rootPath, _ = filepath.Abs(basePath)
		if _, err := CreateDir(0755, rootPath); err != nil {
			return "", err
		}
	}

	fullFilename := rootPath + string(os.PathSeparator) + filename
	log.Printf("SaveFile basePath: %s, filename: %s, fullFilename: %s", basePath, filename, fullFilename)

	//写入文件数据
	outputFile, err := os.OpenFile(fullFilename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	if _, err = io.Copy(outputFile, bytes.NewReader(data)); err != nil {
		return "", err
	}

	return fullFilename, nil
}

/*
 * 移动文件（全路径源文件，目的路径，日期，会根据日期自动创建路径然后连接到目的路径后）
 *  */
func MoveFile(srcFilename, dstPath string, creationDate time.Time) (string, error) {
	var err error

	srcFullFilename, err := filepath.Abs(srcFilename)
	if err != nil {
		return "", err
	}

	srcFile, err := os.OpenFile(srcFullFilename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	rootPath, err := filepath.Abs(dstPath)
	if err != nil {
		return "", err
	}

	if !creationDate.IsZero() {
		//根据日期创建yyyy/mm/dd目录
		rootPath, _ = CreateDateDir(rootPath, creationDate, 0755)
	}

	dstFullFilename := rootPath + string(os.PathSeparator) + filepath.Base(srcFullFilename)
	dstFile, err := os.OpenFile(dstFullFilename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", err
	} else {
		os.Remove(srcFullFilename)
	}

	return dstFullFilename, err
}

/*
 * 删除文件
 *  */
func DeleteFile(filename string, args ...string) error {
	log.Printf("0 DeleteFile filename: %s", filename)

	fullFilename := filename
	if !filepath.IsAbs(filename) {
		if len(args) > 0 {
			filename = filepath.Join(args[0], filename)
			log.Printf("1 DeleteFile filename: %s", filename)

			fullFilename, _ = filepath.Abs(filename)

			log.Printf("2 DeleteFile fullFilename: %s", fullFilename)
		} else {
			fullFilename, _ = filepath.Abs(filename)
		}
	}

	//删除文件
	return os.Remove(fullFilename)
}

/*
 * 判断文件是否存在
 *  */
func FileIsExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
