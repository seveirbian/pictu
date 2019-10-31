package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg" // 注册jpeg解码器
	"image/png"  // 注册png解码器
	"os"
	"path"
	"path/filepath"
	"pictu/convert"
)

var (
	factorX float64
	factorY float64

	rotation int

	sourcePath string
	targetPath string
	targetName = "pictu"
)

func init() {
	// 初始化参数解析
	flag.Float64Var(&factorX, "x", 1.0, "X-axis scaling factor")
	flag.Float64Var(&factorY, "y", 1.0, "Y-axis scaling factor")
	flag.IntVar(&rotation, "r", 0, "Clockwise rotation angle, only supports 90, 180, 270 degrees")
	flag.StringVar(&sourcePath, "s", "", "Source picture path")
	flag.Usage = usage

	// 构建目标图片路径
	tmpPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("Fail to get target path: %v\n", err)
	}
	targetPath = path.Join(tmpPath, targetName)
}

func main() {
	// 解析命令参数
	flag.Parse()

	// 打开源图片
	sourcePic, err := os.Open(sourcePath)
	if err != nil {
		fmt.Printf("Cannot open source picture: %s %v\n", sourcePath, err)
		os.Exit(1)
	}
	defer sourcePic.Close()

	// 解析源图片
	sImage, sType, err := image.Decode(sourcePic)
	if err != nil {
		fmt.Printf("Fail to decode source picture: %v\n", err)
	}

	// 对源图片进行处理
	// 构建目标图片的image结构
	var tRectangle image.Rectangle
	switch {
	case rotation == 0 || rotation == 90 || rotation == 180 || rotation == 270:
		tRectangle = convert.GetMinMaxPointAfterTrans(sImage, factorX, factorY, float64(rotation))
	default:
		fmt.Printf("Cannot support other rotation angles, please use 0, 90, 180 or 270")
	}

	tImage := image.NewRGBA(tRectangle)

	convert.Convert(tImage, sImage, factorX, factorY, rotation)

	// 保存到目标图片
	switch {
	case sType == "jpeg":
		targetPath = targetPath + ".jpg"
	case sType == "png":
		targetPath = targetPath + ".png"
	}
	targetPic, err := os.Create(targetPath)
	if err != nil {
		fmt.Printf("Fail to create target picture: %s, %v\n", targetPath, err)
	}
	defer targetPic.Close()
	switch {
	case sType == "jpeg":
		jpeg.Encode(targetPic, tImage, nil)
	case sType == "png":
		png.Encode(targetPic, tImage)
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `pictu version: pictu/0.1

Usage: pictu [-x x_factor] [-y y_factor] [-r rotation] -s original_file

Options:
`)
	flag.PrintDefaults()
}
