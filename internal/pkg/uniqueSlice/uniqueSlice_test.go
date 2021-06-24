package uniqueSlice

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestIsUniqueStrings(t *testing.T) {

	type TestCase struct {
		inputSlice []string
		resultBool bool
	}

	t.Parallel()

	cases := []TestCase{
		{
			inputSlice: []string{"1", "2", "3"},
			resultBool: true,
		},
		{
			inputSlice: []string{"1", "2", "1"},
			resultBool: false,
		},
		{
			inputSlice: []string{},
			resultBool: true,
		},
	}

	for caseNum, item := range cases {

		caseNumLabel := "Case №" + strconv.Itoa(caseNum+1)

		resultUnique := IsUniqueStrings(item.inputSlice)

		require.Equal(t, item.resultBool, resultUnique, caseNumLabel)

	}
}
