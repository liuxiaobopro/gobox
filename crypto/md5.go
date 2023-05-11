package crypto

// 参考链接: http://www.topgoer.com/%E5%85%B6%E4%BB%96/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86/md5.html

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {

	//方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	// fmt.Println(md5str1)

	return md5str1

	//方法二
	// w := md5.New()
	// io.WriteString(w, str) //将str写入到w中
	// bw := w.Sum(nil)       //w.Sum(nil)将w的hash转成[]byte格式

	// // md5str2 := fmt.Sprintf("%x", bw)    //将 bw 转成字符串
	// md5str2 := hex.EncodeToString(bw) //将 bw 转成字符串
	// fmt.Println(md5str2)
}
