package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StrToInt64(data string) (int64, error) {
	return strconv.ParseInt(data, 10, 64)
}

func StrsToInt64s(datas []string) ([]int64, error) {
	i64s := make([]int64, 0, len(datas))
	for _, data := range datas {
		i64, err := StrToInt64(data)
		if err != nil {
			return nil, err
		}
		i64s = append(i64s, i64)
	}

	return i64s, nil
}

// XxYy to xx_yy , XxYY to xx_yy
func SnakeStrs(changeStrs []string) []string {
	newStrs := make([]string, 0, len(changeStrs))
	for _, changeStr := range changeStrs {
		newStr := SnakeStr(changeStr)
		newStrs = append(newStrs, newStr)
	}

	return newStrs
}

// XxYy to xx_yy , XxYY to xx_yy
func SnakeStr(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelStrs(changeStrs []string) []string {
	newStrs := make([]string, 0, len(changeStrs))
	for _, changeStr := range changeStrs {
		newStr := CamelStr(changeStr)
		newStrs = append(newStrs, newStr)
	}

	return newStrs
}

// camel string, xx_yy to XxYy
func CamelStr(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func ReplaceAllStrs(oldStrs []string, old string, new string) []string {
	newStrs := make([]string, 0, len(oldStrs))
	for _, oldStr := range oldStrs {
		newStr := strings.ReplaceAll(oldStr, old, new)
		newStrs = append(newStrs, newStr)
	}

	return newStrs
}

func GetExistsStrs(from []string, to []string) []string {
	fromMap := StrsToMap(from)
	toStrs := make([]string, 0, len(to))
	for _, toStr := range to {
		if _, ok := fromMap[toStr]; ok {
			toStrs = append(toStrs, toStr)
		}
	}

	return toStrs
}

func GetExistsStrsPrefix(from []string, to []string, toPrefix string) []string {
	fromMap := StrsToMap(from)
	toStrs := make([]string, 0, len(to))
	for _, toStr := range to {
		prefixToStr := fmt.Sprintf("%s%s", toPrefix, toStr)
		if _, ok := fromMap[prefixToStr]; ok {
			toStrs = append(toStrs, toStr)
		}
	}

	return toStrs
}

func StrsToMap(strs []string) map[string]struct{} {
	strsMap := make(map[string]struct{})
	for _, str := range strs {
		strsMap[str] = struct{}{}
	}

	return strsMap
}

// 获取host port, 是否是host port
func AddrToHostPort(addr string) (string, int64, bool) {
	hostPort := strings.Split(strings.TrimSpace(addr), ":")
	if len(hostPort) == 2 { // 数组长度不是两个
		host := hostPort[0]
		port, err := strconv.ParseInt(hostPort[1], 10, 64)
		if err != nil { // port 数字
			return addr, 0, false
		}

		return host, port, true
	} else {
		return addr, 0, false

	}
}

func AddrI64(host string, port int64) string {
	return fmt.Sprintf("%v:%v", host, port)
}

func DeduplicationStrs(strs []string) []string {
	strMap := make(map[string]struct{})
	for _, str := range strs {
		strMap[str] = struct{}{}
	}

	newStrs := make([]string, 0, len(strMap))
	for key, _ := range strMap {
		newStrs = append(newStrs, key)
	}

	return newStrs
}

func SplitStringAndFilterEmpty(data string, sep string) []string {
	items := strings.Split(strings.TrimSpace(data), sep)
	newItems := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		newItems = append(newItems, item)
	}

	return newItems
}

func StrsToUnique(strs []string) []string {
	strMap := make(map[string]struct{})
	for _, str := range strs {
		strMap[str] = struct{}{}
	}

	newStrs := make([]string, 0, len(strMap))
	for key, _ := range strMap {
		newStrs = append(newStrs, key)
	}

	return newStrs
}

// 随机获取一个字符串
func GetRandStr(strs []string) string {
	return strs[RandN(len(strs))]
}

func HasPrefixAll(strs []string, prefix string) bool {
	for _, str := range strs {
		if !strings.HasPrefix(str, prefix) {
			return false
		}
	}

	return true
}

func SuffixIsNumAll(strs []string, suffixSep string) bool {
	for _, str := range strs {
		items := strings.Split(str, suffixSep)
		if len(items) < 2 {
			return false
		}

		// 看看最后一个是否是数字
		numStr := items[len(items)-1]
		if _, err := strconv.ParseInt(numStr, 10, 64); err != nil {
			return false
		}
	}

	return true
}

func GetSuffix(str string, suffixSep string) string {
	items := strings.Split(str, suffixSep)
	return items[len(items)-1]
}
