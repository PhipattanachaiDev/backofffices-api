package convertDataFormatService

import (
	"strconv"
	"strings"
)

// Helper function to convert []int to a CSV string
func ConvertSliceToCSV(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, num := range slice {
		strSlice[i] = strconv.Itoa(num)
	}
	return strings.Join(strSlice, ",")
}
