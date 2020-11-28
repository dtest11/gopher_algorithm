package sort

/***
451. Sort Characters By Frequency
Medium

1960

138

Add to List

Share
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/
type Record struct {
	Key     rune
	Counter int
}

func frequencySort(s string) string {
	// a  Ascall=97, z=112

	var allKeys []rune
	var exist = make(map[rune]int)
	for _, v := range s {
		count, ex := exist[v]
		if ex {
			exist[v] = count + 1
		} else {
			exist[v] = 1
			allKeys = append(allKeys, v)
		}
	}

	//sort record accord ascall code value
	for i := 0; i < len(allKeys); i++ {
		for j := i + 1; j < len(allKeys) && j-1 >= 0; j-- {
			if allKeys[j-1] < allKeys[j] {
				break // sorted
			} else {
				allKeys[j-1], allKeys[j] = allKeys[j], allKeys[j-1]
			}
		}
	}
	var record []Record
	for _, v := range allKeys {
		t := Record{
			Key:     v,
			Counter: exist[v],
		}
		record = append(record, t)
	}
	//sort record
	for i := 0; i < len(record); i++ {
		for j := i + 1; j < len(record) && j-1 >= 0; j-- {
			if record[j-1].Counter >= record[j].Counter {
				break // sorted
			} else {
				record[j-1], record[j] = record[j], record[j-1]
			}
		}
	}
	var result string
	for _, v := range record {
		counter := v.Counter
		tmp := ""
		for i := 0; i < counter; i++ {
			tmp += string(v.Key)
		}
		result += tmp

	}
	return result
}
