package kata

func Stati(strg string) string {
	if strg == "" {
		return ""
	}
	var times int
	var total int
	var min int
	var max int
	var arr []int
	for _, v := range strg {
		if v == ',' {
			arr = append(arr, total)
			if total < min {
				min = total
			}
			if total > max {
				max = total
			}
			times++
			total = 0
		} else if v == '|' {
			arr = append(arr, total)
			if total < min {
				min = total
			}
			if total > max {
				max = total
			}
			times++
			break
		} else if v == ' ' {
			continue
		} else {
			total = total*10 + int(v-'0')
		}
	}
	if times == 0 {
		return ""
	}
	var sum int
	for _, v := range arr {
		sum += v
	}
	avg := sum / times
	med := arr[times/2]
	return "Range: " + format(max-min) + " Average: " + format(avg) + " Median: " + format(med)
}
