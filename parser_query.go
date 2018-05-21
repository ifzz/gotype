package gotype

func (r *Parser) Search(name string) Type {

	v := r.Values.Search(name)
	if v != nil {
		return v
	}
	v = r.Funcs.Search(name)
	if v != nil {
		return v
	}
	v = r.Types.Search(name)
	if v != nil {
		return v
	}
	return nil
}

func (r *Parser) Child(i int) Type {
	vl := r.Values.Len()
	if i < vl {
		return r.Values.Index(i)
	}
	i -= vl

	fl := r.Funcs.Len()
	if i < fl {
		return r.Funcs.Index(i)
	}
	i -= fl

	tl := r.Types.Len()
	if i < tl {
		return r.Types.Index(i)
	}
	// i -= tl

	return nil
}

func (r *Parser) NumChild() int {
	return len(r.Values) + len(r.Funcs) + len(r.Types)
}

func (r *Parser) Range(f func(k string, v Type) bool) {
	r.RangeValues(f)
	r.RangeFuncs(f)
	r.RangeTypes(f)
}

func (r *Parser) RangeValues(f func(k string, v Type) bool) {
	for _, v := range r.Values {
		if !f(v.Name(), v) {
			return
		}
	}
}

func (r *Parser) RangeFuncs(f func(k string, v Type) bool) {
	for _, v := range r.Funcs {
		if !f(v.Name(), v) {
			return
		}
	}
}

func (r *Parser) RangeTypes(f func(k string, v Type) bool) {
	for _, v := range r.Types {
		if !f(v.Name(), v) {
			return
		}
	}
}