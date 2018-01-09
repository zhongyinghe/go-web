package builder

//import "fmt"

type optype byte

const (
	condType optype = iota
	selectType
	insertType
	updateType
	deleteType
)

type join struct {
	joinType  string
	joinTable string
	joinCond  Cond
}

type Builder struct {
	optype
	tableName string
	cond      Cond
	endConds  []Cond //用来保存order by、 group by、 having和limit的条件
	selects   []string
	joins     []join
	inserts   Eq
	updates   []Eq
}

func Select(cols ...string) *Builder {
	builder := &Builder{cond: NewCond()}
	return builder.Select(cols...)
}

func Insert(eq Eq) *Builder {
	builder := &Builder{cond: NewCond()}
	return builder.Insert(eq)
}

func Update(updates ...Eq) *Builder {
	builder := &Builder{cond: NewCond()}
	return builder.Update(updates...)
}

func Delete(conds ...Cond) *Builder {
	builder := &Builder{cond: NewCond()}
	return builder.Delete(conds...)
}

func (b *Builder) Where(cond Cond) *Builder {
	b.cond = b.cond.And(cond)
	//fmt.Println("cond is ", b.cond.(condAnd))
	return b
}

func (b *Builder) From(tableName string) *Builder {
	b.tableName = tableName
	return b
}

func (b *Builder) Into(tableName string) *Builder {
	b.tableName = tableName
	return b
}

func (b *Builder) Join(joinType, joinTable string, joinCond interface{}) *Builder {
	switch joinCond.(type) {
	case Cond:
		b.joins = append(b.joins, join{joinType, joinTable, joinCond.(Cond)})
	case string:
		b.joins = append(b.joins, join{joinType, joinTable, Expr(joinCond.(string))})
	}

	return b
}

func (b *Builder) InnerJoin(joinTable string, joinCond interface{}) *Builder {
	return b.Join("INNER", joinTable, joinCond)
}

func (b *Builder) LeftJoin(joinTable string, joinCond interface{}) *Builder {
	return b.Join("LEFT", joinTable, joinCond)
}

func (b *Builder) RightJoin(joinTable string, joinCond interface{}) *Builder {
	return b.Join("RIGHT", joinTable, joinCond)
}

func (b *Builder) CrossJoin(joinTable string, joinCond interface{}) *Builder {
	return b.Join("CROSS", joinTable, joinCond)
}

func (b *Builder) FullJoin(joinTable string, joinCond interface{}) *Builder {
	return b.Join("FULL", joinTable, joinCond)
}

func (b *Builder) Select(cols ...string) *Builder {
	b.selects = cols
	b.optype = selectType
	return b
}

func (b *Builder) And(cond Cond) *Builder {
	b.cond = And(b.cond, cond)
	return b
}

func (b *Builder) Or(cond Cond) *Builder {
	b.cond = Or(b.cond, cond)
	return b
}

func (b *Builder) Insert(eq Eq) *Builder {
	//fmt.Println(eq)
	b.inserts = eq
	b.optype = insertType
	return b
}

func (b *Builder) Update(updates ...Eq) *Builder {
	b.updates = updates
	b.optype = updateType
	return b
}

func (b *Builder) Delete(conds ...Cond) *Builder {
	b.cond = b.cond.And(conds...)
	b.optype = deleteType
	return b
}

func (b *Builder) Group(cond Cond) *Builder {
	b.endConds = append(b.endConds, cond)
	return b
}

func (b *Builder) Having(cond Cond) *Builder {
	b.endConds = append(b.endConds, cond)
	return b
}

func (b *Builder) Order(cond Cond) *Builder {
	b.endConds = append(b.endConds, cond)
	return b
}

func (b *Builder) Limit(cond Cond) *Builder {
	b.endConds = append(b.endConds, cond)
	return b
}

func (b *Builder) WriteTo(w Writer) error {
	switch b.optype {
	case condType:
		return b.cond.WriteTo(w)
	case selectType:
		return b.selectWriteTo(w)
	case insertType:
		return b.insertWriteTo(w)
	case updateType:
		return b.updateWriteTo(w)
	case deleteType:
		return b.deleteWriteTo(w)
	}

	return ErrNotSupportType
}

func (b *Builder) ToSQL() (string, []interface{}, error) {
	w := NewWriter()
	//fmt.Println(b)
	if err := b.WriteTo(w); err != nil {
		return "", nil, err
	}

	if len(b.endConds) > 0 {
		for _, cond := range b.endConds {
			cond.WriteTo(w)
		}
	}

	return w.writer.String(), w.args, nil
}

func ToSQL(cond interface{}) (string, []interface{}, error) {
	switch cond.(type) {
	case Cond:
		return condToSQL(cond.(Cond))
	case *Builder:
		return cond.(*Builder).ToSQL()
	}
	return "", nil, ErrNotSupportType
}
