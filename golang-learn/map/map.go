package main

import "fmt"

// 创建： make(map[string]int)
// 获取元素： m[key]
// key不存在时，获得Value类型的初始值
// 用value, ok := m[key] 来判断是否存在key
// 用delete删除一个key

// 使用range 遍历key，或者遍历key，value对
// 不保证遍历顺序，如需顺序，需手动对key排序
// 使用len获得元素个数

// map 使用哈希表，必须可以比较相等
// 除了slice，map，function的内建类型都可以作为key
// Struct类型只要不包含上述字段，也可以作为key

// 无重复字符的最长子串
func lengthOfNonrepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	// map 是无序的
	m := map[string]string{
		"name":    "frank",
		"course":  "golang",
		"site":    "github",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("=== Traversing map ===")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("=== Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	undefined, ok := m["undefined"] // 打印默认空值
	fmt.Println(undefined, ok)

	fmt.Println("=== Deleting values === ")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
