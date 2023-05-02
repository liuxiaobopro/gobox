package string

// Has 判断字符串是否存在某个字符
func Has(s string, c byte) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == c {
            return true
        }
    }
    return false
}

// IsPrefix 判断字符串是否存在某个前缀
func IsPrefix(s string, prefix string) bool {
    if len(s) < len(prefix) {
        return false
    }
    for i := 0; i < len(prefix); i++ {
        if s[i] != prefix[i] {
            return false
        }
    }
    return true
}

// IsSuffix 判断字符串是否存在某个后缀
func IsSuffix(s string, suffix string) bool {
    if len(s) < len(suffix) {
        return false
    }
    for i := 0; i < len(suffix); i++ {
        if s[len(s)-len(suffix)+i] != suffix[i] {
            return false
        }
    }
    return true
}
