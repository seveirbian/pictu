# pictu
图片转换工具，能够缩小放大和旋转图片，使用go语言实现

## 功能
1. 缩小图片
2. 放大图片
3. 旋转图片

## 下载安装
```
git clone https://github.com/seveirbian/pictu.git
cd pictu
go build
```

## 使用方法
1. 缩小图片
```
./pictu -x 0.5 -y 0.5 -s ./images/source.jpg
```
![原始图片](/images/source.jpg)
![缩小的图片](/images/shrunk.jpg)

2. 放大图片
```
./pictu -x 1.2 -y 1.2 -s ./images/source.jpg
```
![原始图片](/images/source.jpg)
![放大的图片](/images/enlarged.jpg)

3. 旋转图片
```
./pictu -r 90 -s ./images/source.jpg
```
![原始图片](/images/source.jpg)
![旋转的图片](/images/rotated.jpg)

4. 缩小同时旋转图片
```
./pictu -x 0.5 -y 0.5 -r 90 -s ./images/source.jpg
```
![原始图片](/images/source.jpg)
![缩小同时旋转的图片](/images/shrunk_rotated.jpg)

## License
本项目采用Apache License 2.0进行许可

This repository is licensed under Apache License 2.0.