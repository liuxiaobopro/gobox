package pdf_writer

import (
	"fmt"
	"unicode"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// PDFWriter 封装PDF生成相关功能
type PDFWriter struct {
	m pdf.Maroto // 用于生成PDF的Maroto实例

	fontName    string // 字体名称
	fontRegular string // 字体常规
	fontBold    string // 字体加粗

	titleSize       float64 // 标题字体大小
	fontSize        float64 // 字体大小
	subTitleSize    float64 // 二级标题字体大小
	thirdTitleSize  float64 // 三级标题字体大小
	fourthTitleSize float64 // 四级标题字体大小
	fifthTitleSize  float64 // 五级标题字体大小
	sixthTitleSize  float64 // 六级标题字体大小

	pageWidth    float64 // 页面宽度
	marginLeft   float64 // 左边距
	marginRight  float64 // 右边距
	marginTop    float64 // 上边距
	textWidth    float64 // 文本宽度
	chineseWidth float64 // 中文字符宽度
	digitWidth   float64 // 数字宽度
	letterWidth  float64 // 英文字符宽度
	spaceWidth   float64 // 空格宽度
}

type PDFWriterOption func(w *PDFWriter)

func WithFontName(fontName string) PDFWriterOption {
	return func(w *PDFWriter) {
		w.fontName = fontName
	}
}

func WithFontRegular(fontRegular string) PDFWriterOption {
	return func(w *PDFWriter) {
		w.fontRegular = fontRegular
	}
}

func WithFontBold(fontBold string) PDFWriterOption {
	return func(w *PDFWriter) {
		w.fontBold = fontBold
	}
}

// NewPDFWriter 创建新的PDFWriter实例
func NewPDFWriter(opts ...PDFWriterOption) *PDFWriter {
	fontSize := 12.0

	marginLeft := 20.0
	marginRight := 10.0
	marginTop := 20.0
	pageWidth := 500.0

	w := &PDFWriter{
		m:               pdf.NewMaroto(consts.Portrait, consts.A4),
		titleSize:       fontSize * 2.0,
		fontSize:        fontSize,
		subTitleSize:    fontSize * 1.8,
		thirdTitleSize:  fontSize * 1.6,
		fourthTitleSize: fontSize * 1.4,
		fifthTitleSize:  fontSize * 1.2,
		sixthTitleSize:  fontSize * 1.0,

		pageWidth:    pageWidth,
		marginLeft:   marginLeft,
		marginRight:  marginRight,
		marginTop:    marginTop,
		textWidth:    pageWidth - marginLeft - marginRight,
		chineseWidth: fontSize * 1.0,
		digitWidth:   fontSize * 0.6,
		letterWidth:  fontSize * 0.6,
		spaceWidth:   fontSize * 0.3,
	}

	for _, opt := range opts {
		opt(w)
	}

	w.init()
	return w
}

// init 初始化PDF设置
func (w *PDFWriter) init() {
	w.m.SetPageMargins(w.marginLeft, w.marginTop, w.marginRight)
	w.m.AddUTF8Font(w.fontName, "", w.fontRegular)
	w.m.AddUTF8Font(w.fontName, "B", w.fontBold)
}

// getCharWidth 计算字符宽度
func (w *PDFWriter) getCharWidth(r rune) float64 {
	switch {
	case unicode.Is(unicode.Han, r):
		return w.chineseWidth
	case unicode.IsDigit(r):
		return w.digitWidth
	case unicode.IsLetter(r):
		return w.letterWidth
	case unicode.IsSpace(r):
		return w.spaceWidth
	default:
		return w.letterWidth
	}
}

// splitTextIntoLines 文本分行
func (w *PDFWriter) splitTextIntoLines(text string) []string {
	var lines []string
	var currentLine string
	var currentWidth float64

	for _, r := range text {
		charWidth := w.getCharWidth(r)

		if currentWidth+charWidth > w.textWidth {
			lines = append(lines, currentLine)
			currentLine = string(r)
			currentWidth = charWidth
		} else {
			currentLine += string(r)
			currentWidth += charWidth
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// AddTitle 添加标题
func (w *PDFWriter) AddTitle(title string) {
	w.m.Row(20, func() {
		w.m.Col(12, func() {
			w.m.Text(title, props.Text{
				Size:   w.titleSize,
				Style:  consts.Bold,
				Align:  consts.Center,
				Family: w.fontName,
			})
		})
	})
}

// 添加二级标题
func (w *PDFWriter) AddSubTitle(subTitle string, level int) {
	var size float64
	switch level {
	case 2:
		size = w.subTitleSize
	case 3:
		size = w.thirdTitleSize
	case 4:
		size = w.fourthTitleSize
	case 5:
		size = w.fifthTitleSize
	case 6:
		size = w.sixthTitleSize
	default:
		size = w.fontSize
	}

	w.m.Row(8, func() {
		w.m.Col(12, func() {
			w.m.Text(subTitle, props.Text{
				Size:   size,
				Style:  consts.Bold,
				Align:  consts.Left,
				Family: w.fontName,
			})
		})
	})
}

// AddText 添加文本内容
func (w *PDFWriter) AddText(text string) {
	lines := w.splitTextIntoLines(text)

	for _, line := range lines {
		w.m.Row(8, func() {
			w.m.Col(12, func() {
				w.m.Text(line, props.Text{
					Size:            w.fontSize,
					Align:           consts.Left,
					Family:          w.fontName,
					VerticalPadding: 1,
				})
			})
		})
	}
}

// AddBlankLine 添加空行
func (w *PDFWriter) AddBlankLine() {
	w.m.Row(8, func() {
		w.m.Col(12, func() {
			w.m.Text("", props.Text{
				Size: w.fontSize,
			})
		})
	})
}

// AddImage 添加图片
func (w *PDFWriter) AddImage(imagePath string, percent float64) {
	w.m.Row(8, func() {
		w.m.Col(12, func() {
			w.m.FileImage(imagePath, props.Rect{
				Percent: percent,
				Center:  true,
			})
		})
	})
}

// Save 保存PDF文件
func (w *PDFWriter) Save(filename string) error {
	if err := w.m.OutputFileAndClose(filename); err != nil {
		return fmt.Errorf("保存PDF失败: %v", err)
	}
	return nil
}
