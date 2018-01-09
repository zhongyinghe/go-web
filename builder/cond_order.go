package builder

import "fmt"

type OrderBy [1]string

var _ Cond = OrderBy{""}

func (orderby OrderBy) WriteTo(w Writer) error {
	if orderby.IsValid() {
		_, err := fmt.Fprintf(w, " ORDER BY %s", orderby[0])
		return err
	}
	return nil
}

func (orderby OrderBy) And(conds ...Cond) Cond {
	return nil
}

func (orderby OrderBy) Or(conds ...Cond) Cond {
	return nil
}

func (orderby OrderBy) IsValid() bool {
	return len(orderby[0]) > 0
}
