package web

type RequestParam struct {
	Key   string
	Value string
}

type RequestParams struct {
	params []RequestParam
}

func (r *RequestParam) Get(name string) (string, error) {
	return "", nil
}
