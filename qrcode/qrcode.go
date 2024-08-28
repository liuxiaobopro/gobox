package qrcode

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"

	"github.com/liuxiaobopro/gobox/file"
	"github.com/skip2/go-qrcode"
)

type QrCode struct {
	Cotent   string // 二维码内容
	FilePath string // 二维码保存路径

	Margin int // 外边距大小
}

type QrCodeOption func(q *QrCode)

// WithMargin 设置外边距大小
func WithMargin(margin int) QrCodeOption {
	return func(q *QrCode) {
		q.Margin = margin
	}
}

func New(content, filePath string, options ...QrCodeOption) *QrCode {
	q := &QrCode{
		Cotent:   content,
		FilePath: filePath,
	}

	for _, option := range options {
		option(q)
	}

	return q
}

// Create 生成二维码
func (q *QrCode) Create() error {
	if _, err := file.Create(q.FilePath, false); err != nil {
		return err
	}

	if err := qrcode.WriteFile(q.Cotent, qrcode.Medium, 512, q.FilePath); err != nil {
		return err
	}

	var r io.Reader

	r, err := os.Open(q.FilePath)
	if err != nil {
		return err
	}

	// 加载生成的二维码图像
	originalQR, err := png.Decode(r)
	if err != nil {
		return err
	}

	// 设置外边距大小
	margin := q.Margin

	// 创建一个新的图像，包括外边距
	newImg := image.NewRGBA(image.Rect(0, 0, originalQR.Bounds().Dx()+2*margin, originalQR.Bounds().Dy()+2*margin))

	// 设置新图像的背景为白色（或者其他颜色）
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 将原始二维码绘制到新图像的中间位置
	draw.Draw(newImg, originalQR.Bounds().Add(image.Pt(margin, margin)), originalQR, image.Point{}, draw.Over)

	// 保存带有外边距的二维码图像
	qrWithMarginFile, err := os.Create(q.FilePath)
	if err != nil {
		return err
	}
	defer qrWithMarginFile.Close()

	err = png.Encode(qrWithMarginFile, newImg)
	if err != nil {
		return err
	}

	return nil
}
