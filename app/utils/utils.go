package utils

import (
	"github.com/rs/zerolog/log"
)

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func Spreed(args ...[]interface{}) []interface{} {
	var result []interface{}
	for _, arg := range args {
		result = append(result, arg...)
	}
	return result
}

func FlatMap[T, U any](ts []T, f func(T) []U) []U {
	res := make([]U, 0)
	for i := range ts {
		us := f(ts[i])
		if us != nil {
			res = append(res, us...)
		}
	}
	return res
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func Unique[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}
	return uniqueSlice
}

func Must[T any](val T, err error) T {
	if err != nil {
		log.Err(err).Stack().Msg("Error occurred")
		panic(err)
	}
	return val
}

func Must2[T1 any, T2 any](val1 T1, val2 T2, err error) (T1, T2) {
	if err != nil {
		log.Err(err).Stack().Msg("Error occurred")
		panic(err)
	}
	return val1, val2
}

func Must0(err error) {
	if err != nil {
		log.Err(err).Stack().Msg("Error occurred")
		panic(err)
	}
}

func ChooseFirst[T any](t1 T, t2 any) T {
	return t1
}

func PrintErr0(err error) {
	if err != nil {
		log.Err(err).Stack().Msg("Error occurred")
	}
}

func Pointer[T any](val T) *T {
	t := new(T)
	*t = val
	return t
}

func ToPointerArray[T any](val []T) []*T {
	r := make([]*T, len(val))
	for i := 0; i < len(val); i++ {
		r[i] = &val[i]
	}
	return r
}
func AddNotEnoughItemsToList[Val any, Key comparable](list []Val, total []Key, keyExtractor func(*Val) Key, valueMaker func(Key) Val) []Val {
	m := make(map[Key]byte)

	// Populate the map based on the list
	for i := range list {
		key := keyExtractor(&list[i])
		m[key] = 1
	}

	// Check for missing keys in the map and add default values
	for _, key := range total {
		if _, ok := m[key]; !ok {
			defaultValue := valueMaker(key)
			list = append(list, defaultValue)
		}
	}
	return list
}

func KeysOfMap[K comparable](m map[K]any) []K {
	list := make([]K, 0, len(m))
	for k, _ := range m {
		list = append(list, k)
	}
	return list
}

func ValueOfMap[K comparable, V any](m map[K]V) []V {
	list := make([]V, 0, len(m))
	for _, v := range m {
		list = append(list, v)
	}
	return list
}

func GroupByOfList[L any, K comparable](list []L, keyExt func(L) K) map[K][]L {
	groupMap := make(map[K][]L)
	for _, item := range list {
		key := keyExt(item)
		group := groupMap[key]
		if group == nil {
			group = make([]L, 0)
		}
		group = append(group, item)
		groupMap[key] = group
	}
	return groupMap
}

func ListToSet[T comparable](list []T) map[T]any {
	set := make(map[T]any)
	for _, t := range list {
		set[t] = 1
	}
	return set
}

func ListToMap[L any, K comparable](list []L, keyExt func(L) K) map[K]L {
	groupMap := make(map[K]L)
	for _, item := range list {
		groupMap[keyExt(item)] = item
	}
	return groupMap
}

func IsNotEmptyString(s string) bool {
	return !(len(s) == 0)
}
