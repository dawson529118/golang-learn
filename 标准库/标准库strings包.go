package main

import (
	"fmt"
	"strings"
)

/*
对字符串的处理
- 检索字符串
- 格式化
- 比较大小
- 裁剪
- 炸碎
- 拼接
*/

/*检索字串*/
func mainCheck() {
	fmt.Println(strings.Contains("hello", "e"))          // true
	fmt.Println(strings.Contains("hello", "b"))          // false
	fmt.Println(strings.ContainsAny("hello", "asshole")) // true
	fmt.Println(strings.ContainsAny("hello", "ass"))     // false

	// 以字符集【序号】和【字符】形式显示字符
	fmt.Printf("%U\n", 'h')
	fmt.Printf("%c\n", 0x6211)
	// 判断是否包含任意字符
	fmt.Println(strings.ContainsRune("hello", 'h'))    // true
	fmt.Println(strings.ContainsRune("hello", 0x0068)) // true

	fmt.Println(strings.Index("hello", "e"))           // 1
	fmt.Println(strings.Index("hello", "a"))           // -1
	fmt.Println(strings.IndexAny("hello", "he"))       // 0
	fmt.Println(strings.IndexRune("fuckoff", 0x6211))  // -1
	fmt.Println(strings.IndexRune("fuck我off", 0x6211)) // 4
}

/*格式化*/
func mainFormat() {
	fmt.Println(strings.ToUpper("heLLo")) // HELLO
	fmt.Println(strings.ToLower("HeLLo")) // hello
	fmt.Println(strings.Title("hello"))   // Hello
}

/*比较大小*/
func mainCompare() {
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("c", "b")) // 1
	fmt.Println(strings.Compare("b", "b")) // 0
}

/*裁剪（只裁头尾，不裁中间）*/
func mainTailor() {
	// 干掉头尾的空格
	fmt.Println(strings.TrimSpace("  abc  "))
	// 去掉头
	fmt.Println(strings.TrimPrefix("和：18", "和："))
	// 去掉尾
	fmt.Println(strings.TrimSuffix("asbcd:asd", ":asd"))

	// 干掉开头和结尾的f, u, c, k
	fmt.Println(strings.Trim("fuckyoufuckmeuc", "fuck"))      // youfuckme
	fmt.Println(strings.TrimLeft("fuckyoufuckmeuc", "fuck"))  // youfuckmeuc
	fmt.Println(strings.TrimRight("fuckyoufuckmeuc", "fuck")) // fuckyoufuckme

	// 通过函数自定义复杂的裁剪规则
	fmt.Println(strings.TrimFunc("fuckyoufuckmeuc", filter)) // ckyoufuckmeuc
	fmt.Println(strings.TrimFunc("fuckyoufuckmeuu", func(char rune) bool {
		if char == 'f' || char == 'u' {
			return true
		} else {
			return false
		}
	}))
}

/*哪个字符返回true，哪个字符就上黑名单---去掉*/
func filter(char rune) bool {
	if char == 'f' || char == 'u' {
		return true
	} else {
		return false
	}
}

/*炸碎（分割）*/
func mainSeg() {
	strs := strings.Split("fuck you fuck me", " ")
	fmt.Println(strs, len(strs))

	// 分割为N段，N=-1时，最大分割数；不等于-1时，输入几，就分割几份
	/*strs = strings.SplitN("fuck you fuck me", " ", 2)
	fmt.Println(strs, len(strs))
	for _, value := range strs {
		fmt.Println(value)
	}*/

	// 将分隔符本身包含进去
	strs = strings.SplitAfter("fuck,you,fuck,me", ",")
	fmt.Println(strs, len(strs))
	strs = strings.SplitAfterN("fuck,you,fuck,me", ",", 2)
	fmt.Println(strs, len(strs))
}

/*拼接*/
func mainJoin() {
	bigStr := strings.Join([]string{"fuck", "shit", "you"}, "❤")
	fmt.Println(bigStr)
}
