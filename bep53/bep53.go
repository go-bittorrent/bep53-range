package bep53

import (
	"errors"
	"strconv"
	"strings"
)

// ErrParseFailed ...
var ErrParseFailed = errors.New("failed to parse bep53 range")

// Parse range into []int.
func Parse(rawRange []string) ([]int, error) {
	var r []int

	for _, v := range rawRange {
		rangeSplit := strings.Split(v, "-")
		switch len(rangeSplit) {
		case 1:
			val, err := strconv.Atoi(rangeSplit[0])
			if err != nil {
				return nil, err
			}

			r = append(r, val)
		case 2:
			start, err := strconv.Atoi(rangeSplit[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(rangeSplit[1])
			if err != nil {
				return nil, err
			}

			for i := start; i < end+1; i++ {
				r = append(r, i)
			}
		default:
			return nil, ErrParseFailed
		}
	}

	return r, nil
}

// Compose parsed range into []string.
func Compose(numbers []int) []string {
	var r []string

	if len(numbers) > 1 {
		for i := 1; i < len(numbers); i++ {
			if numbers[i] == numbers[i-1]+1 {
				start := numbers[i-1]
				stop := numbers[i-1]

				for j := i; j < len(numbers); j++ {
					if numbers[j-1]+1 == numbers[j] {
						stop++
					} else {
						break
					}
				}

				i += stop - start
				r = append(r, strconv.Itoa(start)+"-"+strconv.Itoa(stop))
			} else {
				r = append(r, strconv.Itoa(numbers[i-1]))
			}
		}
	} else if len(numbers) > 0 {
		r = append(r, strconv.Itoa(numbers[0]))
	}

	return r
}
