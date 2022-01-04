package p0049

import "sort"

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	container := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		key := sortString(strs[i])
		container[key] = append(container[key], strs[i])
	}
	for _, item := range container {
		result = append(result, item)
	}
	return result

}

func sortString(s string) string {
	r := []byte(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
