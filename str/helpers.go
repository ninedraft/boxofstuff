package str

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func Match(re string) func(str string) bool {
	var reg = regexp.MustCompile(re)
	return func(str string) bool {
		return reg.MatchString(str)
	}
}

func Len(str string) int {
	return utf8.RuneCountInString(str)
}

func StrictLess(a, b string) bool {
	return a < b
}

func InSet(set []string) func(str string) bool {
	return func(str string) bool {
		for _, s := range set {
			if s == str {
				return true
			}
		}
		return false
	}
}

func NotInSet(set []string) func(str string) bool {
	return func(str string) bool {
		for _, s := range set {
			if s == str {
				return false
			}
		}
		return true
	}
}

func Contains(substr string) func(string) bool {
	return func(str string) bool {
		return strings.Contains(str, substr)
	}
}

func NotContains(substr string) func(string) bool {
	return func(str string) bool {
		return !strings.Contains(str, substr)
	}
}

func TrimPrefix(prefix string) func(str string) string {
	return curry2Inverse(prefix, strings.TrimPrefix)
}

func TrimSuffix(suffix string) func(str string) string {
	return curry2Inverse(suffix, strings.TrimSuffix)
}

func Shorter(l uint) func(str string) bool {
	return func(str string) bool {
		return l > uint(utf8.RuneCountInString(str))
	}
}

func Longer(l uint) func(str string) bool {
	return func(str string) bool {
		return l < uint(utf8.RuneCountInString(str))
	}
}

func EqLen(l uint) func(str string) bool {
	return func(str string) bool {
		return l == uint(utf8.RuneCountInString(str))
	}
}

func Chop(l uint) func(str string) []string {
	switch l {
	case 0:
		panic("str.Chop: l must be non-zero value")
	case 1:
		return func(str string) []string {
			var chunks = make([]string, 0, len(str))
			var buf = make([]byte, utf8.UTFMax)
			for _, r := range str {
				var n = utf8.EncodeRune(buf, r)
				chunks = append(chunks, string(buf[:n]))
			}
			return chunks
		}
	default:
		return func(str string) []string {
			var strLen = uint(len(str))
			if strLen <= l {
				return []string{str}
			}
			var chunksN = strLen / l
			var chunks = make([]string, 0, chunksN)
			for i := uint(0); i < chunksN; i++ {
				var right = 3 * (i + 1)
				if right > strLen {
					right = strLen
				}
				chunks = append(chunks, str[3*i:right])
			}
			return chunks
		}
	}
}

func curry2(a string, f func(a, b string) string) func(str string) string {
	return func(b string) string {
		return f(a, b)
	}
}

func curry2Inverse(b string, f func(a, b string) string) func(string) string {
	return func(a string) string {
		return f(a, b)
	}
}
