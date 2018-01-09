package builder

import "fmt"

type Between struct {
	Col     string
	LessVal interface{}
	MoreVal interface{}
}

var _ Cond = Between{}

func (between Between) WriteTo(w Writer) error {
	if _, err := fmt.Fprintf(w, "%s BETWEEN ? AND ?", between.Col); err != nil {
		return err
	}
	w.Append(between.LessVal, between.MoreVal)
	return nil
}

// And implments And with other conditions
func (between Between) And(conds ...Cond) Cond {
	return And(between, And(conds...))
}

// Or implments Or with other conditions
func (between Between) Or(conds ...Cond) Cond {
	return Or(between, Or(conds...))
}

// IsValid tests if the condition is valid
func (between Between) IsValid() bool {
	return len(between.Col) > 0
}
