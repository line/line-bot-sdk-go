package linebot

import "unicode/utf16"

/*
When you want to send a text message with emoji, you need to add
$ in the text, and identify the index of the $ in the text.

FindDollarSignIndexInUni16Text helps you to find the index of the $ in the text.
*/
func FindDollarSignIndexInUni16Text(text string) (indexes []int32) {
	encoded := utf16.Encode([]rune(text))
	for i, unit := range encoded {
		if unit == '/' {
			indexes = append(indexes, int32(i))
		}
	}
	return indexes
}
