package builder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	sql, args, err := Insert(Eq{"c": 1, "d": 2}).Into("table1").ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestInsert2(t *testing.T) {
	sql, args, err := Insert(Eq{"a": Expr("Select b FROM table2 where id=?", 3)}).Into("table1").ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestUpdate(t *testing.T) {
	sql, args, err := Update(Eq{"a": 2}).From("table1").Where(Eq{"a": 1}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestDelete(t *testing.T) {
	sql, args, err := Delete(Eq{"a": 1, "b": 2}).From("table1").ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestSelect(t *testing.T) {
	sql, args, err := Select("c, d").From("table1").Where(Eq{"a": 1, "b": 3}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestLeftJoin(t *testing.T) {
	sql, args, err := Select("t1.c, t2.d").From("table1 t1").LeftJoin("table2 t2", "t1.id=t2.id").Where(Eq{"t1.a": 5}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestRightJoin(t *testing.T) {
	sql, args, err := Select("c, d").From("table1").RightJoin("table3", "table1.id=table3.id").Where(Eq{"a": 1}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestInnerJoin(t *testing.T) {
	sql, args, err := Select("t1.a,t1.b,t2.c,t2.d").From("table1 AS t1").InnerJoin("table2 AS t2", "t1.id = t2.id").Where(Eq{"t1.name": "abc"}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestCrossJoin(t *testing.T) {
	sql, args, err := Select("t1.a,t1.b,t2.c,t2.d").From("table1 AS t1").CrossJoin("table2 AS t2", "t1.id = t2.id").Where(Eq{"t1.name": "abc"}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestFullJoin(t *testing.T) {
	sql, args, err := Select("t1.a,t1.b,t2.c,t2.d").From("table1 AS t1").FullJoin("table2 AS t2", "t1.id = t2.id").Where(Eq{"t1.name": "abc"}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestToSQL(t *testing.T) {
	sql, args, err := ToSQL(Eq{"a": 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)

	sql, args, err = ToSQL(Eq{"b": "c"}.And(Eq{"c": 0}))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestEq(t *testing.T) {
	sql, args, err := ToSQL(Eq{"b": []string{"c", "d"}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Eq{"b": 1, "c": []int{2, 3}})
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Eq{"b": "c"}.Or(Eq{"b": "d"}))
	fmt.Println(sql, args)
}

func TestNeq(t *testing.T) {
	sql, args, _ := ToSQL(Neq{"a": 1})
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Neq{"b": "c"}.And(Neq{"c": 0}))
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Neq{"b": "c"}.Or(Neq{"b": "d"}))
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Neq{"a": []int{1, 2, 3, 4}})
	fmt.Println(sql, args)
}

func TestNotIn(t *testing.T) {
	sql, args, _ := ToSQL(NotIn("a", 1, 2, 3, 4, 5, 6))
	fmt.Println(sql, args)
}

func TestNull(t *testing.T) {
	sql, args, _ := ToSQL(IsNull{"a"})
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(NotNull{"b"})
	fmt.Println(sql, args)
}

func TestLike(t *testing.T) {
	sql, args, _ := ToSQL(Like{"a", "%b"})
	fmt.Println(sql, args)
}

func TestBetween(t *testing.T) {
	sql, args, _ := ToSQL(Between{"a", 1, 2})
	fmt.Println(sql, args)
}

func TestLt(t *testing.T) {
	sql, args, _ := ToSQL(Lt{"a": 1})
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Lte{"b": 2})
	fmt.Println(sql, args)
}

func TestGt(t *testing.T) {
	sql, args, _ := ToSQL(Gt{"a": 1})
	fmt.Println(sql, args)

	sql, args, _ = ToSQL(Gte{"b": 2})
	fmt.Println(sql, args)
}

func TestNot_n(t *testing.T) {
	sql, args, _ := ToSQL(Not{Eq{"a": 1}})
	fmt.Println(sql, args)
}

func TestAnd(t *testing.T) {
	sql, args, err := Select("c, d").From("table1").Where(Eq{"a": 1, "b": 3}).And(NotIn("c", []int{1, 2, 3})).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestOr(t *testing.T) {
	sql, args, err := Select("c, d").From("table1").Where(Eq{"a": 1, "b": 3}).Or(NotIn("c", []int{1, 2, 3})).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)
}

func TestOrderBy(t *testing.T) {
	sql, args, err := Select("c,d").From("table1").Where(Eq{"a": 1}).Order(OrderBy{"b DESC"}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)

	assert.Equal(t, sql, "SELECT c,d FROM table1 WHERE a=? ORDER BY b DESC")
}

func TestLimit(t *testing.T) {
	sql, args, err := Select("c,d").From("table1").Where(Eq{"a": 1}).Order(OrderBy{"b DESC"}).Limit(Limit{0, 20}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)

	assert.Equal(t, "SELECT c,d FROM table1 WHERE a=? ORDER BY b DESC LIMIT 0, 20", sql)
}

func TestGroupBy(t *testing.T) {
	sql, args, err := Select("c,d").From("table1").Where(Eq{"a": 1}).Group(Group{"f"}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)

	assert.Equal(t, "SELECT c,d FROM table1 WHERE a=? GROUP BY f", sql)
}

func TestHaving(t *testing.T) {
	sql, args, err := Select("c,d,sum(num) AS nums").From("table1").Where(Eq{"a": 1}).Group(Group{"f"}).Having(Having{"nums > 10"}).ToSQL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql, args)

	assert.Equal(t, "SELECT c,d,sum(num) AS nums FROM table1 WHERE a=? GROUP BY f HAVING nums > 10", sql)
}
