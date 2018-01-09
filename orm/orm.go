package orm

import (
	"bytes"
	"container/list"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	timeFormate = "2006-01-02 15:04:05"
)

var sqlDB *sql.DB

func Open() *sql.DB {
	if sqlDB == nil {
		db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
		if err != nil {
			fmt.Println(err.Error())
		}
		sqlDB = db
	}

	return sqlDB
}

type IBaseDao interface {
	Init()
	Save(data interface{}) error
	Update(data interface{}) error
	SaveOrUpdate(data interface{}) error
	SaveAll(datas list.List) error
	Delete(data interface{}) error
	Find(sql string) (*list.List, error)
	FindOne(sql string) (interface{}, error)
}

type BaseDao struct {
	EntityType    reflect.Type
	sqlDB         *sql.DB
	tableName     string //表名
	pk            string //主键
	columnToField map[string]string
	fieldToColumn map[string]string
}

//初始化
func (this *BaseDao) Init() {
	this.columnToField = make(map[string]string)
	this.fieldToColumn = make(map[string]string)

	types := this.EntityType

	for i := 0; i < types.NumField(); i++ {
		typ := types.Field(i)
		tag := typ.Tag
		if len(tag) > 0 {
			column := tag.Get("column")
			name := typ.Name

			this.columnToField[column] = name
			this.fieldToColumn[name] = column
			if len(tag.Get("table")) > 0 {
				this.tableName = tag.Get("table")
				this.pk = column
			}
		}
	}
}

func (this *BaseDao) PrintInfo() {
	fmt.Println(this.tableName)
	fmt.Println(this.pk)
	fmt.Println(this.columnToField)
	fmt.Println(this.fieldToColumn)
}

//预处理insert语句
func (this *BaseDao) insertPrepareSQL() (fieldNames list.List, sql string) {
	names := new(bytes.Buffer)
	values := new(bytes.Buffer)

	i := 0
	for column, fieldName := range this.columnToField {
		if i != 0 {
			names.WriteString(",")
			values.WriteString(",")
		}
		fieldNames.PushBack(fieldName)
		names.WriteString(column)
		values.WriteString("?")
		i++
	}
	sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", this.tableName, names.String(), values.String())
	return
}

//预处理占位符的数据
func (this *BaseDao) prepareValues(data interface{}, fieldNames list.List) []interface{} {
	values := make([]interface{}, len(this.columnToField))
	object := reflect.ValueOf(data).Elem()

	i := 0
	for e := fieldNames.Front(); e != nil; e = e.Next() {
		name := e.Value.(string)
		field := object.FieldByName(name)
		values[i] = this.fieldValue(field)
		i++
	}
	return values
}

func (this *BaseDao) fieldValue(v reflect.Value) interface{} {
	if !v.IsValid() {
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Struct:
		switch v.Type().String() {
		case "time.Time":
			m := v.MethodByName("Format")
			rets := m.Call([]reflect.Value{reflect.ValueOf(timeFormate)})
			t := rets[0].String()
			return t
		default:
			return nil
		}
	default:
		return nil
	}
}

//增加单个记录
func (this *BaseDao) Save(data interface{}) error {
	columns, sql := this.insertPrepareSQL()
	stmt, err := Open().Prepare(sql)
	args := this.prepareValues(data, columns)
	fmt.Println(sql, " ", args)

	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return err
}

//增加多个记录
func (this *BaseDao) SaveAll(datas list.List) error {
	if datas.Len() == 0 {
		return nil
	}

	columns, sql := this.insertPrepareSQL()
	stmt, err := Open().Prepare(sql)
	if err != nil {
		panic(err.Error())
	}

	for e := datas.Front(); e != nil; e = e.Next() {
		args := this.prepareValues(e.Value, columns)
		fmt.Println(sql, " ", args)
		_, err = stmt.Exec(args...)
		if err != nil {
			panic(err.Error())
		}
	}
	return err
}

//更新预处理
func (this *BaseDao) updatePrepareSQL() (fieldNames list.List, sql string) {
	sets := new(bytes.Buffer)

	i := 0
	for column, fieldName := range this.columnToField {
		if strings.EqualFold(column, this.pk) {
			continue
		}
		if i != 0 {
			sets.WriteString(",")
		}
		fieldNames.PushBack(fieldName)
		sets.WriteString(column)
		sets.WriteString("=?")
		i++
	}

	fieldNames.PushBack(this.columnToField[this.pk])
	sql = fmt.Sprintf("UPDATE %s SET %s WHERE %s=?", this.tableName, sets.String(), this.pk)
	return
}

//更新数据
func (this *BaseDao) Update(data interface{}) error {
	columns, sql := this.updatePrepareSQL()
	stmt, err := Open().Prepare(sql)

	args := this.prepareValues(data, columns)
	fmt.Println(sql, " ", args)

	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return err
}

func (this *BaseDao) SaveOrUpdate(data interface{}) error {
	if this.isPkValue(data) {
		return this.Update(data)
	} else {
		return this.Save(data)
	}
}

//判断主键是否有值
func (this *BaseDao) isPkValue(data interface{}) bool {
	object := reflect.ValueOf(data).Elem()

	pkName := this.columnToField[this.pk]
	pkValue := object.FieldByName(pkName)
	if !pkValue.IsValid() {
		return false
	}

	switch pkValue.Kind() {
	case reflect.String:
		if len(pkValue.String()) > 0 {
			return true
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if pkValue.Uint() != 0 {
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if pkValue.Int() != 0 {
			return true
		}
	}

	return false
}

//实体转delete sql语句
func (this *BaseDao) deleteSQL(data interface{}) string {
	object := reflect.ValueOf(data).Elem()
	fieldValue := object.FieldByName(this.columnToField[this.pk])
	pkValue := this.valueToString(fieldValue)
	return fmt.Sprintf("DELETE FROM %s WHERE %s=%s", this.tableName, this.pk, pkValue)
}

//reflect.Value转字符串
func (this *BaseDao) valueToString(v reflect.Value) string {
	values := new(bytes.Buffer)

	switch v.Kind() {
	case reflect.String:
		values.WriteString("'")
		values.WriteString(v.String())
		values.WriteString("'")
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		values.WriteString(fmt.Sprintf("%d", v.Uint()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.WriteString(fmt.Sprintf("%d", v.Int()))
	case reflect.Float32, reflect.Float64:
		values.WriteString(fmt.Sprintf("%f", v.Float()))
	case reflect.Struct:
		switch v.Type().String() {
		case "time.Time":
			m := v.MethodByName("Format")
			rets := m.Call([]reflect.Value{reflect.ValueOf(timeFormate)})
			t := rets[0].String()
			values.WriteString("'")
			values.WriteString(t)
			values.WriteString("'")
		default:
			values.WriteString("null")
		}
	default:
		values.WriteString("null")
	}
	return values.String()
}

//删除一个实体
func (this *BaseDao) Delete(data interface{}) error {
	sql := this.deleteSQL(data)
	fmt.Println(sql)
	_, err := Open().Exec(sql)
	return err
}

//查询记录
func (this *BaseDao) Find(sql string) (*list.List, error) {
	return this.query(sql, false)
}

func (this *BaseDao) FindOne(sql string) (interface{}, error) {
	datas, err := this.query(sql, true)
	var data interface{}
	if datas.Len() > 0 {
		data = datas.Front().Value
	}
	return data, err
}

func (this *BaseDao) query(sql string, isOne bool) (*list.List, error) {
	if isOne && !strings.Contains(sql, "limit") {
		sql = sql + " limit 1"
	}

	rows, err := Open().Query(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	datas := list.New()

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		obj := this.parseQuery(columns, values)
		datas.PushBack(obj)
	}

	return datas, err
}

func (this *BaseDao) parseQuery(columns []string, values []interface{}) interface{} {
	obj := reflect.New(this.EntityType).Interface()
	typ := reflect.ValueOf(obj).Elem()

	for i, col := range values {
		if col != nil {
			name := this.columnToField[columns[i]]
			field := typ.FieldByName(name)
			this.pareseQueryColumn(field, string(col.([]byte)))
		}
	}

	return obj
}

//单个属性赋值
func (this *BaseDao) pareseQueryColumn(field reflect.Value, s string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(s)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, _ := strconv.ParseUint(s, 10, 0)
		field.SetUint(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, _ := strconv.ParseInt(s, 10, 0)
		field.SetInt(v)
	case reflect.Float32:
		v, _ := strconv.ParseFloat(s, 32)
		field.SetFloat(v)
	case reflect.Float64:
		v, _ := strconv.ParseFloat(s, 64)
		field.SetFloat(v)
	case reflect.Struct:
		switch field.Type().String() {
		case "time.Time":
			v, _ := time.Parse(timeFormate, s)
			field.Set(reflect.ValueOf(v))
		}
	default:
	}
}
