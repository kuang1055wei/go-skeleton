package utils

import (
	"reflect"
	"sort"
)

// Int64Slice attaches the methods of Interface to []int64, sorting a increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// SortInt64s sorts []int64s in increasing order.
func SortInt64s(a []int64) {
	sort.Sort(Int64Slice(a))
}

//SearchInt64s 在已排序的 int64s 切片中搜索 x 并返回 Search 指定的索引。
//如果 x 不存在，则返回值是插入 x 的索引（可能是 len(a)）。
//切片必须按升序排序。
func SearchInt64s(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// InInts 检查 []int 的切片中是否存在 int 值。
func InInts(needle int, haystack []int) bool {
	if len(haystack) == 0 {
		return false
	}

	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

// InInt64s 检查 []int64 的切片中是否存在 int64 的值。
func InInt64s(needle int64, haystack []int64) bool {
	if len(haystack) == 0 {
		return false
	}

	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

// InFloat64s 检查 []float64 的切片中是否存在 float64 的值。
func InFloat64s(needle float64, haystack []float64) bool {
	if len(haystack) == 0 {
		return false
	}

	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

// InStrings checks if a value of string exists in a slice of []string.
func InStrings(needle string, haystack []string) bool {
	if len(haystack) == 0 {
		return false
	}

	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

// InArray 检查 []interface{} 的切片中是否存在 interface{} 的值。
func InArray(needle interface{}, haystack []interface{}) bool {
	if len(haystack) == 0 {
		return false
	}

	for _, v := range haystack {
		if reflect.DeepEqual(needle, v) {
			return true
		}
	}

	return false
}

// IntsUnique 接受一个整数输入切片并返回一个没有重复值的新整数切片。
func IntsUnique(a []int) []int {
	l := len(a)

	if l <= 1 {
		return a
	}

	m := make(map[int]byte, l)
	r := make([]int, 0, l)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = 0
			r = append(r, v)
		}
	}

	return r
}

// Int64sUnique takes an input slice of int64s and
// returns a new slice of int64s without duplicate values.
func Int64sUnique(a []int64) []int64 {
	l := len(a)

	if l <= 1 {
		return a
	}

	m := make(map[int64]byte, l)
	r := make([]int64, 0, l)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = 0
			r = append(r, v)
		}
	}

	return r
}

//Float64sUnique 接受一个 float64s 的输入切片，并返回一个没有重复值的新的 float64s 切片。
func Float64sUnique(a []float64) []float64 {
	l := len(a)

	if l <= 1 {
		return a
	}

	m := make(map[float64]byte, l)
	r := make([]float64, 0, l)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = 0
			r = append(r, v)
		}
	}

	return r
}

//StringsUnique 接受输入的字符串切片并返回没有重复值的新字符串切片。
func StringsUnique(a []string) []string {
	l := len(a)

	if l <= 1 {
		return a
	}

	m := make(map[string]byte, l)
	r := make([]string, 0, l)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = 0
			r = append(r, v)
		}
	}

	return r
}
