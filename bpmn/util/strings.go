package bpmn_util

// 删除数组内元素
func RemoveString(strings []string, s string) []string {
	for i, aString := range strings {
		if aString == s {
			strings[i] = strings[0]
			strings = strings[1:]
			break
		}
	}
	return strings
}

func ContainsString(strings []string, s string) bool {
	for _, aString := range strings {
		if aString == s {
			return true
		}
	}
	return false
}
