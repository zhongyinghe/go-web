package regmatch

import (
	"fmt"
	"reflect"
	"regexp"
)

//匹配数字
func IsNumber(number interface{}) bool {
	str := fmt.Sprintf("%v", number)
	if m, _ := regexp.MatchString("^[0-9]+$", str); !m {
		return false
	}
	return true
}

//匹配ip
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

//匹配中文
func IsChinese(str string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", str); !m {
		return false
	}
	return true
}

//匹配英文
func IsEnglish(str string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		return false
	}
	return true
}

//匹配邮件地址
func IsEmail(email string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}

//匹配手机号码
func IsPhone(phoneNum string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phoneNum); !m {
		return false
	}
	return true
}

//匹配身份证
func IsUserCard(usercard string) bool {
	if len(usercard) == 15 {
		if m, _ := regexp.MatchString(`^(\d{15})$`, usercard); !m {
			return false
		}
		return true
	}

	if len(usercard) == 18 {
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, usercard); !m {
			return false
		}
		return true
	}

	return false
}

//类似下拉菜单匹配
func InRange(value interface{}, items interface{}) bool {
	itemsValue := reflect.ValueOf(items)
	kind := reflect.TypeOf(items).Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		for i := 0; i < itemsValue.Len(); i++ {
			if itemsValue.Index(i).Interface() == value {
				return true
			}
		}
		return false
	}
	return false
}

//类似多选按钮匹配
func RangeInRange(slice1, slice2 interface{}) bool {
	kind1, kind2 := reflect.TypeOf(slice1).Kind(), reflect.TypeOf(slice2).Kind()
	if (kind1 == reflect.Slice || kind1 == reflect.Array) && (kind2 == reflect.Slice || kind2 == reflect.Array) {
		itemsValue := reflect.ValueOf(slice1)
		for i := 0; i < itemsValue.Len(); i++ {
			if !InRange(itemsValue.Index(i).Interface(), slice2) {
				return false
			}
		}
		return true
	}
	return false
}
