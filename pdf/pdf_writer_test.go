package pdf_writer

import (
	"fmt"
	"os"
	"testing"
)

func TestNewPDFWriter(t *testing.T) {
	// 创建 PDFWriter 实例
	writer := NewPDFWriter(
		WithFontName("NotoSansSC"),
		WithFontRegular("Noto_Sans_SC/static/NotoSansSC-Regular.ttf"),
		WithFontBold("Noto_Sans_SC/static/NotoSansSC-Bold.ttf"),
	)

	// 添加标题
	writer.AddTitle("这是一个中文 PDF 文件!")

	// 长文本示例
	longText := `111这是一个很长的中文段落示例。我们现在测试文本的自动换行功能。这是第二个段落，展示了如何处理多个段落的情况。Maroto提供了简单易用的API，可以轻松创建专业的PDF文档。下面是第三个段落：1. 支持自动换行 2. 可以处理多个段落 3. 保持段落之间的间距 4. 支持中文字符显示。最后一个段落，用来测试更多的文本内容。当文本超出页面宽度时，会自动换行到下一行继续显示，这样可以确保所有内容都能正确显示在PDF中。
	这是一个很长的中文段落示例。我们现在测试文本的自动换行功能。这是第二个段落，展示了如何处理多个段落的情况。Maroto提供了简单易用的API，可以轻松创建专业的PDF文档。下面是第三个段落：1. 支持自动换行 2. 可以处理多个段落 3. 保持段落之间的间距 4. 支持中文字符显示。最后一个段落，用来测试更多的文本内容。当文本超出页面宽度时，会自动换行到下一行继续显示，这样可以确保所有内容都能正确显示在PDF中。
	这是一个很长的中文段落示例。我们现在测试文本的自动换行功能。这是第二个段落，展示了如何处理多个段落的情况。Maroto提供了简单易用的API，可以轻松创建专业的PDF文档。下面是第三个段落：1. 支持自动换行 2. 可以处理多个段落 3. 保持段落之间的间距 4. 支持中文字符显示。最后一个段落，用来测试更多的文本内容。当文本超出页面宽度时，会自动换行到下一行继续显示，这样可以确保所有内容都能正确显示在PDF中。`

	// 添加文本内容
	writer.AddText(longText)

	// 添加空行
	writer.AddBlankLine()

	// 添加二级标题
	writer.AddSubTitle("这是一个二级标题示例", 2)

	writer.AddBlankLine()

	writer.AddText(longText)

	// 保存文件
	if err := writer.Save("output.pdf"); err != nil {
		fmt.Printf("生成PDF失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF文件已成功生成！")
}
