package test

import (
	"sort"
	"strings"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/3/4 18:45
 * @Email 1328919715@qq.com
 * @Description:
 **/

func Test_SortKey(t *testing.T) {
	// var  dst = &strings.Builder{}
	var sortedKeys = []string{"1-key", "0-key", "6-key", "90-key", "87-key"}
	sort.Slice(sortedKeys, func(i, j int) bool {
		numI := []byte(strings.Split(sortedKeys[i], "-")[0])
		numJ := []byte(strings.Split(sortedKeys[j], "-")[0])
		if numI[0] > numJ[0] {
			return false
		}
		return true
	})

	t.Log(sortedKeys)
}
