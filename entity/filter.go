package entity

type FilterItem[T comparable] struct {
	Name  string `json:"name"`
	Op    string `json:"op"`
	Value T      `json:"value"`
}

type FilterFunc interface {
	Filter(m map[string]any) bool
}

func (i *FilterItem[T]) Filter(m map[string]any) bool {
	if _, exists := m[i.Name]; !exists {
		return false
	}
	switch i.Op {
	case "eq":
		return m[i.Name] == i.Value
	case "ne":
		return m[i.Name] == i.Value
	case "gt":
		return m[i.Name] == i.Value
	case "ge":
		return m[i.Name] == i.Value
	case "lt":
		return m[i.Name] == i.Value
	case "le":
		return m[i.Name] == i.Value
	case "in":
		return m[i.Name] == i.Value
	case "nin":
		return m[i.Name] == i.Value
	}
	return false
}

type Filter struct {
	Items []FilterFunc `json:"items"`
}
