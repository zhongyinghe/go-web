package orm

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
	"time"
)

/*type Entity struct {
	Id   int       `table:"table_name" column:"id"`
	Time time.Time `column:"time"`
	Name string    `column:"day"`
	Age  string    `column:"age"`
}*/

type UserInfo struct {
	Uid        int    `table:"userinfo" column:"uid"`
	UserName   string `column:"username"`
	Departname string `column:"departname"`
	Created    int64  `column:"created"`
}

func GetDao() *BaseDao {
	dao := &BaseDao{EntityType: reflect.TypeOf(new(UserInfo)).Elem()}
	dao.Init()
	return dao
}

/*func TestInit(t *testing.T) {
	dao := GetDao()
	dao.PrintInfo()
}

func TestInsertPrepareSQL(t *testing.T) {
	dao := GetDao()
	fieldNames, sql := dao.insertPrepareSQL()
	for e := fieldNames.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(sql)
}*/

func TestPrepareValues(t *testing.T) {
	user := &UserInfo{
		UserName:   "abc",
		Departname: "golang",
		Created:    time.Now().Unix(),
	}
	/*user := new(UserInfo)
	user.Uid = 12
	user.UserName = "abc"
	user.Departname = "golang"
	user.Created = time.Now().Unix()*/

	dao := GetDao()
	columns, sql := dao.insertPrepareSQL()
	fmt.Println(columns)
	fmt.Println(sql)
	values := dao.prepareValues(user, columns)
	fmt.Println(values)
}

func TestSave(t *testing.T) {
	user := &UserInfo{
		UserName:   "abc_kk",
		Departname: "golang_kk",
		Created:    time.Now().Unix(),
	}

	dao := GetDao()
	dao.Save(user)
}

func TestSaveAll(t *testing.T) {
	var datas list.List
	for i := 0; i < 3; i++ {
		user := &UserInfo{
			UserName:   "efg",
			Departname: "研发",
			Created:    time.Now().Unix(),
		}
		datas.PushBack(user)
	}

	dao := GetDao()
	dao.SaveAll(datas)
}

func TestUpdate(t *testing.T) {
	u := &UserInfo{
		Uid:        11,
		UserName:   "siscall",
		Departname: "AI",
		Created:    time.Now().Unix(),
	}

	dao := GetDao()
	dao.Update(u)
}

func TestSaveOrUpdate(t *testing.T) {
	u := &UserInfo{
		UserName:   "abc123",
		Departname: "AI",
		Created:    time.Now().Unix(),
	}

	dao := GetDao()
	dao.SaveOrUpdate(u)
}

func TestDelete(t *testing.T) {
	u := &UserInfo{
		Uid: 13,
	}

	dao := GetDao()
	dao.Delete(u)
}

func TestFind(t *testing.T) {
	sql := "Select * from userinfo"
	dao := GetDao()
	datas, _ := dao.Find(sql)
	//fmt.Println(datas)
	for e := datas.Front(); e != nil; e = e.Next() {
		i := e.Value
		u := *(i.(*UserInfo))
		info(u)
	}
}

func TestFindOne(t *testing.T) {
	sql := "Select * FROM userinfo WHERE uid=12"
	dao := GetDao()
	data, _ := dao.FindOne(sql)
	p := data.(*UserInfo)
	info(*p)
}

func info(u UserInfo) {
	fmt.Println(u.Uid, u.UserName, u.Departname, u.Created)
}
