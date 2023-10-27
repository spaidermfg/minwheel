package test

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
	"unsafe"
)

// strings

func TestClone(t *testing.T) {
	str := "Beautiful"
	clone := strings.Clone(str)
	fmt.Printf("str: %v, %p\nclone: %v, %p\n", str, &str, clone, &clone)
	fmt.Printf("sizeof: %v,%v\n", unsafe.Sizeof(str), unsafe.Sizeof(clone))
	fmt.Printf("byte: %v,%v\n", []byte(str), []byte(clone))
}

// 比较的是首字母的ASCII码
// a == b, return 0
// a < b, return -1
// a > b, return 1
func TestCompare(t *testing.T) {
	a := "Beautiful"
	b := "sky"
	compare := strings.Compare(a, b)
	fmt.Println("compare:", compare)
	fmt.Printf("byte: %v,%v\n", []byte(a), []byte(b))
}

// strings.Index的调用
// 子字符串必须连续
func TestContains(t *testing.T) {
	a := "Beautiful"
	b := "ti"
	contains := strings.Contains(a, b)
	fmt.Println("a contains b?", contains)
	fmt.Printf("byte: %v,%v\n", []byte(a), []byte(b))
}

// 只要有一个字符匹配，就返回true
func TestContainsAny(t *testing.T) {
	a := "Beautiful"
	b := "hjkl"
	containsAny := strings.ContainsAny(a, b)
	fmt.Println("a containsAny b?", containsAny)
	fmt.Printf("byte: %v,%v\n", []byte(a), []byte(b))
}

func TestContainsFunc(t *testing.T) {
	a := "Beautiful"
	containsFunc := strings.ContainsFunc(a, func(r rune) bool {
		if a == "Beautiful" {
			return true
		}
		return false
	})
	fmt.Println("a containsFunc b?", containsFunc)
}

// rune 表示Unicode类型
func TestContainsRune(t *testing.T) {
	a := "Beautiful"
	containsRune := strings.ContainsRune(a, 66)
	fmt.Println("a containsFunc b?", containsRune)
	fmt.Printf("byte: %v\n", []byte(a))
}

// 计算字符串中子子字符串的数量，如果不存在，则返回字符串总长度+1
func TestCount(t *testing.T) {
	a := "Beautiful"
	b := "u"
	count := strings.Count(a, b)
	fmt.Println("count:", count)
}

// if not found， return all
func TestCut(t *testing.T) {
	cut := func(a, b string) {
		before, after, found := strings.Cut(a, b)
		fmt.Printf("all: %v,before: %v, after: %v, founc: %v\n", a, before, after, found)
	}

	a := "Beautiful"
	cut(a, "u")
	cut(a, "ti")
	cut(a, "gb")

	// 	output:
	//	all: Beautiful,before: Bea, after: tiful, founc: true
	//	all: Beautiful,before: Beau, after: ful, founc: true
	//	all: Beautiful,before: Beautiful, after: , founc: false
}

// 根据前缀子字符串裁切字符串，返回后缀字符串以及布尔值
func TestCutPrefix(t *testing.T) {
	a := "Beautiful"
	after, found := strings.CutPrefix(a, "Be")
	fmt.Printf("after: %v, found: %v\n", after, found)
}

// 根据后缀子字符串裁切字符串，返回前缀字符串以及布尔值
func TestCutSuffix(t *testing.T) {
	a := "Beautiful"
	before, found := strings.CutSuffix(a, "ful")
	fmt.Printf("before: %v, found: %v\n", before, found)
}

// 比较字符串是否一致，不区分大小写
func TestEqualFold(t *testing.T) {
	a := "Beautiful"
	b := "bEAUTIFUL"
	fold := strings.EqualFold(a, b)
	fmt.Printf("%v=%v?: %v\n", a, b, fold)
}

// 将以空格分隔的字符串转化为字符串数组
func TestFields(t *testing.T) {
	a := "Bea uti ful"
	fields := strings.Fields(a)
	fmt.Printf("fields: %v, len: %v\n", fields, len(fields))
}

func TestFieldsFunc(t *testing.T) {
	a := "bea;uti;ful..."
	fields := strings.FieldsFunc(a, func(r rune) bool {
		return !unicode.IsLower(r) && !unicode.IsNumber(r)
	})
	fmt.Printf("fields: %v, len: %v\n", fields, len(fields))
}

// 检索字符串前缀是否包含某子字符串，返回布尔值
func TestHasPrefix(t *testing.T) {
	a := "Beautiful"
	prefix := strings.HasPrefix(a, "Be")
	fmt.Printf("prefix: %v\n", prefix)

	//	output:
	//	prefix: true
}

// 检索字符串后缀是否包含某子字符串，返回布尔值
func TestHasSuffix(t *testing.T) {
	a := "Beautiful"
	suffix := strings.HasSuffix(a, "ful")
	fmt.Printf("suffix: %v\n", suffix)

	//	output:
	//	suffix: true
}

// 返回子字符串的下标，如果不存在则返回-1
func TestIndex(t *testing.T) {
	a := "Beautiful"
	b := "u"
	index := strings.Index(a, b)
	fmt.Printf("all: %v, son: %v, index: %v\n", a, b, index)

	b = "m"
	index = strings.Index(a, b)
	fmt.Printf("all: %v, son: %v, index: %v\n", a, b, index)

	//	output:
	//	all: Beautiful, son: u, index: 3
	//	all: Beautiful, son: m, index: -1

}

// 返回子字符串中第一个存在于父字符串中的下标
func TestIndexAny(t *testing.T) {
	a := "Beautiful"
	b := "mat"
	index := strings.IndexAny(a, b)
	fmt.Printf("all: %v, son: %v, index: %v\n", a, b, index)

	//	output:
	//	all: Beautiful, son: mat, index: 2
}

// 返回字节在字符串中的位置下标, 不存在返回-1
func TestIndexByte(t *testing.T) {
	a := "Beautiful"
	index := strings.IndexByte(a, 't')
	fmt.Printf("all: %v, son: %v, index: %v\n", a, 't', index)

	//	output:
	//	all: Beautiful, son: 116, index: 4
}

// 根据自定义检索规则返回子字符串在字符串中的下标位置
func TestIndexFunc(t *testing.T) {
	a := "Beau7长城tiful"
	index := strings.IndexFunc(a, func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	})
	fmt.Printf("all: %v, son: %v, index: %v\n", a, '长', index)

	//	output:
	//	all: Beau7长城tiful, son: 38271, index: 5
}

// 返回unicode字符在字符串中的下标位置
func TestIndexRune(t *testing.T) {
	a := "Beautiful"
	index := strings.IndexRune(a, 'i')
	fmt.Printf("all: %v, son: %v, index: %v\n", a, 'i', index)

	//	output:
	//	all: Beautiful, son: 105, index: 5
}

// 使用指定分隔符组合字符串数组
func TestJoin(t *testing.T) {
	a := []string{"I", "love", "you."}
	b := " "
	join := strings.Join(a, b)
	fmt.Printf("join: %v\n", join)

	//	output:
	//	join: I love you.
}

// 返回子字符串在父字符串中最后匹配的下标位置
func TestLastIndex(t *testing.T) {
	a := "Beautiful"
	b := "u"
	lastIndex := strings.LastIndex(a, b)
	firstIndex := strings.Index(a, b)
	fmt.Printf("all: %v, son: %v, first index: %v, last index: %v\n", a, b, firstIndex, lastIndex)

	//	output:
	//	all: Beautiful, son: u, first index: 3, last index: 7
}

// 返回子字符串出现在父字符串中最靠后位置的下标，不存在则返回-1
func TestLastIndexAny(t *testing.T) {
	a := "Beautiful"
	b := "redubsglsa"
	indexAny := strings.LastIndexAny(a, b)
	fmt.Printf("all: %v, son: %v, last index: %v\n", a, b, indexAny)

	//	output:
	//	all: Beautiful, son: redubsglsa, first index: -1, last index: 8
}

func TestLastIndexByte(t *testing.T) {
	//strings.LastIndexByte()
}

func TestLastIndexFunc(t *testing.T) {
	//strings.LastIndexFunc()
}

func TestMap(t *testing.T) {
	//a := "Beautiful"
	//s := strings.Map(func(r rune) rune {
	//
	//}, a)
	//fmt.Printf("s: %v\n", s)
}

func TestNewReader(t *testing.T) {
	//strings.NewReader()
}

func TestNewReplacer(t *testing.T) {
	//strings.NewReplacer()
}

// 返回字符串s重复n次的字符串
func TestRepeat(t *testing.T) {
	a := "Beautiful"
	repeat := strings.Repeat(a, 3)
	fmt.Printf("repeat: %v\n", repeat)

	//	output:
	//	repeat: BeautifulBeautifulBeautiful
}

// 根据count数替换检索到的old字符串，count=-1替换所有检索到的字符串
func TestReplace(t *testing.T) {
	a := strings.Repeat("Beautiful", 4)
	old := "ea"
	newer := "ae"
	replace := strings.Replace(a, old, newer, 2)
	fmt.Printf("replace2: %v\n", replace)

	replace = strings.Replace(a, old, newer, -1)
	fmt.Printf("replace-1: %v\n", replace)

	//	output:
	//	replace2: BaeutifulBaeutifulBeautifulBeautiful
	//	replace-1: BaeutifulBaeutifulBaeutifulBaeutiful
}

// 替换所有匹配到的字符串
func TestReplaceAll(t *testing.T) {
	a := strings.Repeat("Beautiful", 4)
	old := "ea"
	newer := "ae"
	replace := strings.ReplaceAll(a, old, newer)
	fmt.Printf("replace: %v\n", replace)

	//	output:
	//	replace: BaeutifulBaeutifulBaeutifulBaeutiful
}

// 根据标志符分隔字符串，返回分隔后不包含标志符的字符串数组
func TestSplit(t *testing.T) {
	a := "Beautiful space"
	sep := " "
	split := strings.Split(a, sep)
	fmt.Printf("all: %v, sep: %v, split: %v\n", a, sep, split)

	//	output:
	//	all: Beautiful space, sep:  , split: [Beautiful space]
}

// 根据标志符分隔字符串，返回分隔后包含标志符的字符串数组
func TestSplitAfter(t *testing.T) {
	a := "Beautiful"
	sep := "t"
	split := strings.Split(a, sep)
	after := strings.SplitAfter(a, sep)

	fmt.Printf("all: %v, sep: %v, split: %v, split after: %v\n", a, sep, split, after)

	//	output:
	//	all: Beautiful, sep: t, split: [Beau iful], split after: [Beaut iful]
}

func TestSplitAfterN(t *testing.T) {
	//strings.SplitAfterN()
}

// 将字符串根据指定的字符标志分割成n份
func TestSplitN(t *testing.T) {
	a := "this is a new shoes"
	sep := " "
	n := 3
	sn := strings.SplitN(a, sep, n)
	fmt.Printf("all: %v, sep: %v, n: %v, split: %q\n", a, sep, n, sn)

	//	output:
	//	all: this is a new shoes, sep:  , n: 3, split: ["this" "is" "a new shoes"]
}

// 将字符串转为小写字符串
func TestToLower(t *testing.T) {
	a := "Beautiful"
	lower := strings.ToLower(a)
	fmt.Printf("origin: %v, lower: %v\n", a, lower)

	//	output:
	//	origin: Beautiful, lower: beautiful
}

// 返回指定unicode类型的小写格式
func TestToLowerSpecial(t *testing.T) {
	a := "асГДГССШШсцсшешшс"
	special := strings.ToLowerSpecial(unicode.TurkishCase, a)
	fmt.Printf("origin: %v, lower: %v\n", a, special)

	//	output:
	//	origin: асГДГССШШсцсшешшс, lower: асгдгссшшсцсшешшс
}

// 返回字符串的大写格式
func TestToTitle(t *testing.T) {
	a := "Beautiful"
	title := strings.ToTitle(a)
	fmt.Printf("title: %v\n", title)

	//	output:
	//	title: BEAUTIFUL
}

// 返回指定unicode类型的大写格式
func TestTitleSpecial(t *testing.T) {
	a := "асдфагаеср"
	special := strings.ToTitleSpecial(unicode.TurkishCase, a)
	fmt.Printf("special: %v", special)

	//	output:
	//	title: АСДФАГАЕСР
}

// 返回字符串的大写格式
func TestToUpper(t *testing.T) {
	a := "Beautiful"
	upper := strings.ToUpper(a)
	fmt.Printf("origin: %v, upper: %v\n", a, upper)

	//	output:
	//	origin: Beautiful, upper: BEAUTIFUL
}

// 返回指定unicode类型的大写格式
func TestToUpperSpecial(*testing.T) {
	a := "асдфагаеср"
	upper := strings.ToUpperSpecial(unicode.TurkishCase, a)
	fmt.Printf("origin: %v, upper: %v\n", a, upper)

	//	output:
	//	origin: асдфагаеср, upper: АСДФАГАЕСР
}

// 返回字符串删除前缀和后缀中的指定字符后的字符串
func TestTrim(t *testing.T) {
	a := "+-+Beautiful-+-"
	trim := strings.Trim(a, "+-")
	fmt.Printf("origin: %v, trim: %v\n", a, trim)

	//	output:
	//	origin: +-+Beautiful-+-, trim: Beautiful
}

// 自定义cut条件，cut掉字符串两侧字符，返回cut后的字符
func TestTrimFunc(t *testing.T) {
	a := "+-+Beautiful-+-"
	trim := strings.TrimFunc(a, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})
	fmt.Printf("origin: %v, trim: %v\n", a, trim)

	//	output:
	//	origin: +-+Beautiful-+-, trim: Beautiful
}

// cut掉字符串左边所有匹配到的指定字符，保留右边的字符串
func TestTrimLeft(t *testing.T) {
	a := "+-+Beautiful-+-"
	b := "+-"
	left := strings.TrimLeft(a, b)
	fmt.Printf("origin: %v, trim: %v, left: %v\n", a, b, left)

	//	output:
	// 	origin: +-+Beautiful-+-, trim: +-, left: Beautiful-+-
}

// 自定义cut条件，cut掉字符串左侧字符，返回右侧保留字符
func TestTrimLeftFunc(t *testing.T) {
	a := "+-+Beautiful-+-"
	left := strings.TrimLeftFunc(a, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})

	fmt.Printf("origin: %v, left: %v\n", a, left)

	//	output:
	//	origin: +-+Beautiful-+-, left: Beautiful-+-
}

// cut掉字符串指定的前缀字符串，返回cut后的字符串
func TestPrefix(t *testing.T) {
	a := "+-Beautiful-+-"
	b := "+-"
	prefix := strings.TrimPrefix(a, b)
	fmt.Printf("origin: %v, trim: %v, prefix: %v\n", a, b, prefix)

	//	output:
	//	origin: +-Beautiful-+-, trim: +-, prefix: Beautiful-+-
}

// cut掉字符串右边所有匹配到的指定字符，保留左边的字符串
func TestRight(t *testing.T) {
	a := "+-+Beautiful-+-"
	b := "+-"
	right := strings.TrimRight(a, b)

	fmt.Printf("origin: %v, trim: %v, right: %v\n", a, b, right)

	//	output:
	//	origin: +-+Beautiful-+-, trim: +-, right: +-+Beautiful
}

// 自定义cut条件，cut掉字符串右侧匹配到的字符串，返回左侧的字符串
func TestTrimRightFunc(t *testing.T) {
	a := "+-+Beautiful-+-"
	right := strings.TrimRightFunc(a, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})
	fmt.Printf("origin: %v, right: %v\n", a, right)

	//	output:
	//	origin: +-+Beautiful-+-, right: +-+Beautiful
}

// cut掉字符串两边的空格，返回trim后的字符串
func TestTrimSpace(t *testing.T) {
	a := "    Beautiful    "
	trim := strings.TrimSpace(a)
	fmt.Printf("origin: [%v], trim: [%v]\n", a, trim)

	//	output:
	//	origin: [    Beautiful    ], trim: [Beautiful]
}

// cut掉指定的后缀字符串，保留剩余前缀字符串
func TestTrimSuffix(t *testing.T) {
	a := "+-+Beautiful-+-"
	b := "ful-+-"
	suffix := strings.TrimSuffix(a, b)
	fmt.Printf("origin: %v, trim: %v, suffix: %v\n", a, b, suffix)

	//	output:
	//	origin: +-+Beautiful-+-, trim: ful-+-, suffix: +-+Beauti
}
