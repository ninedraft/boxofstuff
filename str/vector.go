package str

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Vector []string

func (vector Vector) FirstNonEmpty() string {
	for _, str := range vector {
		if str != "" {
			return str
		}
	}
	return ""
}

func (vector Vector) Len() int {
	return len(vector)
}

func (vector Vector) New() Vector {
	return make(Vector, 0, vector.Len())
}

func (vector Vector) Copy() Vector {
	return append(vector.New(), vector...)
}

func (vector Vector) Filter(pred func(str string) bool) Vector {
	var filtered = vector.New()
	for _, str := range vector {
		if pred(str) {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func (vector Vector) Count() map[string]int {
	var count = make(map[string]int, vector.Len())
	for _, str := range vector {
		count[str]++
	}
	return count
}

func (vector Vector) Slice() []string {
	return []string(vector.Copy())
}

func (vector Vector) Join(delim string) string {
	return strings.Join(vector, delim)
}

func (vector Vector) WriteToChan(ch chan<- string) {
	for _, str := range vector {
		ch <- str
	}
}

func (vector Vector) Append(strs ...string) Vector {
	return append(vector.Copy(), strs...)
}

func (vector Vector) Map(op func(str string) string) Vector {
	var mapped = vector.New()
	for _, str := range vector {
		mapped = append(mapped, op(str))
	}
	return mapped
}

func (vector Vector) Populate(op func(str string) []string) Vector {
	var populated = vector.New()
	for _, str := range vector {
		populated = append(populated, op(str)...)
	}
	return populated
}

func (vector Vector) Fold(start string, op func(acc, str string) string) string {
	var acc = start
	for _, str := range vector {
		acc = op(acc, str)
	}
	return acc
}

func (vector Vector) Delete(i int) Vector {
	vector = vector.Copy()
	return append(vector[:i], vector[i+1:]...)
}

func (vector Vector) FirstFiltered(pred func(str string) bool) (string, bool) {
	for _, str := range vector {
		if pred(str) {
			return str, true
		}
	}
	return "", false
}

func (vector Vector) SortByKey(key func(str string) int) Vector {
	vector = vector.Copy()
	sort.Slice(vector, func(i, j int) bool {
		return key(vector[i]) < key(vector[j])
	})
	return vector
}

func (vector Vector) SortByLess(less func(a, b string) bool) Vector {
	vector = vector.Copy()
	sort.Slice(vector, func(i, j int) bool {
		return less(vector[i], vector[j])
	})
	return vector
}

func (vector Vector) Unique() Vector {
	var set = map[string]struct{}{}
	var unique = vector.New()
	for _, str := range vector {
		if _, ok := set[str]; !ok {
			unique = append(unique, str)
			set[str] = struct{}{}
		}
	}
	return unique
}

func (vector Vector) Shuffle() Vector {
	vector = vector.Copy()
	rnd.Shuffle(vector.Len(), func(i, j int) {
		vector[i], vector[j] = vector[j], vector[i]
	})
	return vector
}

func (vector Vector) Sample(n uint) Vector {
	var sample = make(Vector, 0, n)
	var vLen = vector.Len()
	for i := uint(0); i < n; i++ {
		var ind = rnd.Intn(vLen)
		sample = append(sample, vector[ind])
	}
	return sample
}

func (vector Vector) Top(n uint) Vector {
	var count = vector.Count()
	var sorted = vector.Unique().SortByKey(func(str string) int {
		return -count[str]
	})
	if uint(sorted.Len()) < n {
		return sorted
	}
	return sorted[:n].Copy()
}

func (vector Vector) Classify(key func(str string) string) map[string][]string {
	vector = vector.Unique()
	var classes = make(map[string][]string, len(vector))
	for _, str := range vector {
		var k = key(str)
		classes[k] = append(classes[k], str)
	}
	return classes
}
