package convert

import (
	"image"
	"image/color"
	"math"
)

// Convert function:
// 遍历原始图片的每个像素点，计算其在新图片中的位置
func Convert(t *image.RGBA, s image.Image, x, y float64, r int) {
	for i := s.Bounds().Min.X; i <= s.Bounds().Max.X; i++ {
		for j := s.Bounds().Min.Y; j <= s.Bounds().Max.Y; j++ {
			newI, newJ := CoordinateTrans(float64(i), float64(j), x, y, float64(r))
			// fmt.Println(i, j, newI, newJ)
			t.Set(newI, newJ, s.At(i, j))
		}
	}

	// 如果当前执行的是放大操作，则需要对空白像素点做善后处理
	FillEmptyPixel(t)
}

// CoordinateTrans function:
// original pixel (a, b) --> target pixel (A, B)
// x: factorX, y: factorY, r: rotation angle
// 					 | cos r  sin r |   | x  0 |
// (A, B) = (a, b) * | 				| * |      |
// 					 |-sin r  cos r |   | 0  y |
// (A, B) = (x*cosr*a-x*sinr*b, y*sinr*a+y*cosr*b)
func CoordinateTrans(i, j, x, y, r float64) (int, int) {
	rPi := float64(r) / float64(180) * math.Pi

	newI := int(x*math.Cos(rPi)*float64(i) - x*math.Sin(rPi)*float64(j))
	newJ := int(y*math.Sin(rPi)*float64(i) + y*math.Cos(rPi)*float64(j))

	return newI, newJ
}

// GetMinMaxPointAfterTrans function:
// 获取原始图片经过变换之后的最小点和最大点（坐标和最小和最大的点）
func GetMinMaxPointAfterTrans(s image.Image, x, y, r float64) image.Rectangle {
	minX := s.Bounds().Min.X
	minY := s.Bounds().Min.Y
	maxX := s.Bounds().Max.X
	maxY := s.Bounds().Max.Y

	x1, y1 := CoordinateTrans(float64(minX), float64(minY), x, y, r)
	x2, y2 := CoordinateTrans(float64(maxX), float64(maxY), x, y, r)
	x3, y3 := CoordinateTrans(float64(maxX), float64(maxY), x, y, r)
	x4, y4 := CoordinateTrans(float64(maxX), float64(maxY), x, y, r)

	minX = int(math.Min(math.Min(float64(x1), float64(x2)), math.Min(float64(x3), float64(x4))))
	minY = int(math.Min(math.Min(float64(y1), float64(y2)), math.Min(float64(y3), float64(y4))))
	maxX = int(math.Max(math.Max(float64(x1), float64(x2)), math.Max(float64(x3), float64(x4))))
	maxY = int(math.Max(math.Max(float64(y1), float64(y2)), math.Max(float64(y3), float64(y4))))

	return image.Rectangle{
		Min: image.Point{minX, minY},
		Max: image.Point{maxX, maxY},
	}
}

// FillEmptyPixel function:
// 将图片扩大之后会出现一些点缺失的情况，这事需要将这些空点做补充
// 这里采用简单方法，新图片的空像素点的像素值为其左上方两个像素点的平均值
func FillEmptyPixel(t *image.RGBA) {
	for i := t.Bounds().Min.X; i <= t.Bounds().Max.X; i++ {
		for j := t.Bounds().Min.Y; j <= t.Bounds().Max.Y; j++ {
			pColor := t.At(i, j)
			r, g, b, a := pColor.RGBA()
			if r == 0 && g == 0 && b == 0 && a == 0 {
				nColor := newPixelGRBA(i, j, t)
				t.Set(i, j, nColor)
				// fmt.Println(reflect.TypeOf(pColor))
			}
		}
	}
}

// newPixelGRBA function:
func newPixelGRBA(i, j int, t *image.RGBA) (newColor color.Color) {
	switch {
	case i == t.Bounds().Min.X:
		lR, lG, lB, lA := t.At(i, j-1).RGBA()
		newColor = color.RGBA{uint8(lR), uint8(lG), uint8(lB), uint8(lA)}
	case j == t.Bounds().Min.Y:
		uR, uG, uB, uA := t.At(i-1, j).RGBA()
		newColor = color.RGBA{uint8(uR), uint8(uG), uint8(uB), uint8(uA)}
	default:
		lR, lG, lB, lA := t.At(i-1, j).RGBA()
		uR, uG, uB, uA := t.At(i, j-1).RGBA()
		newColor = color.RGBA{uint8((lR + uR) / 2.0), uint8((lG + uG) / 2.0), uint8((lB + uB) / 2.0), uint8((lA + uA) / 2.0)}
	}

	return
}
