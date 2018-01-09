package builder

import "fmt"

type Group [1]string

var _ Cond = Group{""}

func (group Group) WriteTo(w Writer) error {
	if group.IsValid() {
		_, err := fmt.Fprintf(w, " GROUP BY %s", group[0])
		return err
	}
	return nil
}

func (group Group) And(conds ...Cond) Cond {
	return nil
}

func (group Group) Or(conds ...Cond) Cond {
	return nil
}

func (group Group) IsValid() bool {
	return len(group[0]) > 0
}

type Having [1]string

var _ Cond = Having{""}

func (having Having) WriteTo(w Writer) error {
	if having.IsValid() {
		_, err := fmt.Fprintf(w, " HAVING %s", having[0])
		return err
	}
	return nil
}

func (having Having) And(conds ...Cond) Cond {
	return nil
}

func (having Having) Or(conds ...Cond) Cond {
	return nil
}

func (having Having) IsValid() bool {
	return len(having[0]) > 0
}
