package utility

type CompatibleArrayTypes interface {
	string | interface{}
}

// RemoveRowColumnsAtIndexes, removes rows and/or columns, at their specified indexes from the 2D `inputArr`
func RemoveRowsColumnsAtIndexes[V CompatibleArrayTypes](rowIndexesToRemove, columnIndexesToRemove []int, inputArr [][]V) [][]V {
	var (
		rowIndexLookup    = map[int]struct{}{}
		columnIndexLookup = map[int]struct{}{}
		outputArr         = make([][]V, len(inputArr))
	)

	for _, ri := range rowIndexesToRemove {
		rowIndexLookup[ri] = struct{}{}
	}
	for _, ci := range columnIndexesToRemove {
		columnIndexLookup[ci] = struct{}{}
	}

	var originalI int
	for i, r := range inputArr {
		if _, ok := rowIndexLookup[i]; ok {
			outputArr = append(outputArr[:originalI], outputArr[originalI+1:]...)
		} else {
			outputArr[originalI] = r

			var originalJ int
			for j, _ := range r {
				if _, ok := columnIndexLookup[j]; ok {
					outputArr[originalI] = append(outputArr[originalI][:originalJ], outputArr[originalI][originalJ+1:]...)
				} else {
					originalJ++
				}
			}
			originalI++
		}
	}

	return outputArr
}

// FilterElementsMatchingCondition, removes any elements from `inputArr` that return "true" from `matchFunc`
func FilterElementsMatchingCondition[V CompatibleArrayTypes](inputArr [][]V, matchFunc func(V) bool) [][]V {
	outputArr := make([][]V, len(inputArr))
	for i, r := range inputArr {
		var (
			matchCount = 0
			newR       = make([]V, len(r))
		)
		for _, c := range r {
			if matchFunc(c) {
				newR[matchCount] = c
				matchCount++
			}
		}
		outputArr[i] = newR[:matchCount]
	}
	return outputArr
}
