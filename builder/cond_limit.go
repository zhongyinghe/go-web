package builder

import "fmt"

type Limit struct {
	offset int
	size   int
}

var _ Cond = Limit{}

func (limit Limit) WriteTo(w Writer) error {
	if limit.IsValid() {
		_, err := fmt.Fprintf(w, " LIMIT %d, %d", limit.offset, limit.size)
		return err
	}

	return nil
}

func (limit Limit) And(conds ...Cond) Cond {
	return nil
}

func (limit Limit) Or(conds ...Cond) Cond {
	return nil
}

func (limit Limit) IsValid() bool {
	return limit.size >= 1
}
