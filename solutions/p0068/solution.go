package p0068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	if len(words) == 0 {
		return words
	}
	// 相当于双指针
	left, right := 0, 0
	for left < len(words) {
		length := 0
		// 初步计算    一个word + 一个空格
		for right < len(words) && length+len(words[right])+1 <= maxWidth+1 {
			length += len(words[right]) + 1
			right++
		}
		space, extra := 1, 0
		if right != len(words) && right != left+1 {
			// word 之间的空格数
			space = (maxWidth-length+1)/(right-left-1) + 1
			// 最后剩下的 长度 (空格补齐)
			extra = (maxWidth - length + 1) % (right - left - 1)
		}
		sb := &strings.Builder{}
		sb.WriteString(words[left])
		left++
		for left < right {
			for i := 0; i < space; i++ {
				sb.WriteByte(' ')
			}
			if extra > 0 {
				sb.WriteByte(' ')
				extra--
			}
			sb.WriteString(words[left])
			left++
		}
		remaining := maxWidth - sb.Len()
		for remaining > 0 {
			sb.WriteByte(' ')
			remaining--
		}
		result = append(result, sb.String())
	}
	return result
}
