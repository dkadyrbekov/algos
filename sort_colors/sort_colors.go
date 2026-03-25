package sort_colors

func sortColors(colors []int) []int {
	left := 0
	right := len(colors) - 1

	for i := 0; i <= right; {
		switch colors[i] {
		case 0:
			tmp := colors[left]
			colors[left] = colors[i]
			colors[i] = tmp
			left++
			i++
		case 1:
			i++
		case 2:
			tmp := colors[right]
			colors[right] = colors[i]
			colors[i] = tmp
			right--
		default:
			return []int{}
		}
	}

	return colors
}
