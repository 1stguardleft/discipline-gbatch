package job

type Properties struct {
	name string
}

func (r *Properties) Name() string {
	return r.name
}
