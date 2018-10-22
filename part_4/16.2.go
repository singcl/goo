// 字符串正确的操作方式

package main

import (
	"bytes"
)

var b bytes.Buffer

for condition {
	b.WriteString(str)  // 将字符串str写入缓存buffer
}

return b.String()