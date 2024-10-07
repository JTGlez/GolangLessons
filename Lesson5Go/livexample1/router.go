package main

type Router struct {
	handlers map[string]func(a, b int)
}

func (r Router) PatternMatching(key string, a, b int) {
	h, ok := r.handlers[key]
	if !ok {
		return
	}

	a += 1
	b -= 1

	h(a, b)
}
