package utils

import "github.com/enorith/supports/collection"

type TreeData[T any] struct {
	Data     T   `json:"data"`
	Children []T `json:"children"`
}

func MapTree[T any](items []T, parent TreeData[T], isParent, compare func(T) bool) (data []TreeData[T]) {
	parents := collection.Filter(items, isParent)

	for _, parent := range parents {
		data = append(data, TreeData[T]{Data: parent, Children: findNodes(items, TreeData[T]{Data: parent}, compare)})
	}

	return
}

func findNodes[T any](items []T, parent TreeData[T], compare func(T) bool) []T {
	children := collection.Filter(items, compare)
	return collection.Map(children, func(co T) T {
		parent.Children = findNodes(items, TreeData[T]{Data: co}, compare)
		return co
	})
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
