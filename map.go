package rf

func GetMapValue[R any](m map[string]any, key string) (r R) {
	if len(m) == 0 {
		return r
	}
	v, ok := m[key]
	if !ok {
		return r
	}
	r, _ = v.(R)
	return r
}
