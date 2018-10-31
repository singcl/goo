package parse

import (
	"fmt"
	"strings"
	"strconv"
)


// 这是所有自定义包实现者应该遵守的最佳实践：
// 1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
// 2）向包的调用者返回错误值（而不是 panic）。

//在包内部，特别是在非导出函数中有很深层次的嵌套调用时，对主调函数来说用 panic 来表示应该被翻译成错误的错误场景是很有用的（并且提高了代码可读性）


// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int // The index into the space-separated list of words.
	Word string // The word that generated the parse error.
	Err error // The raw error that precipitated this error, if any.
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no worlds to parse")
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
