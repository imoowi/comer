package components

import (
	"log"
	"testing"

	"github.com/bmizerany/assert"
)

type TestStudentModel struct {
	Name   string
	Gender string
	Class  string
}

func TestCache(t *testing.T) {
	cache := NewMemCacheT[*TestStudentModel]()
	lilei := &TestStudentModel{
		Name:   `李雷`,
		Gender: `男`,
		Class:  `一(1)班`,
	}
	cache.SetOne(`lilei`, lilei)
	_lilei, _ := cache.GetOne(`lilei`)
	assert.Equal(t, `李雷`, _lilei.Name)
	assert.Equal(t, `男`, _lilei.Gender)
	assert.Equal(t, `一(1)班`, _lilei.Class)
	hanmeimei := &TestStudentModel{
		Name:   `韩梅梅`,
		Gender: `女`,
		Class:  `二(1)班`,
	}
	lily := &TestStudentModel{
		Name:   `lily`,
		Gender: `女`,
		Class:  `一(3)班`,
	}
	lucy := &TestStudentModel{
		Name:   `lucy`,
		Gender: `女`,
		Class:  `一(2)班`,
	}
	students := make([]*TestStudentModel, 0)
	students = append(students, lilei, hanmeimei, lily, lucy)
	cache.SetArray(`students`, students)
	if _students, ok := cache.GetArray(`students`); ok {
		assert.Equal(t, 4, len(_students))
		for k, v := range _students {
			log.Println(`---range-students---`, k, v.Name, v.Class, v.Gender)
		}
	} else {
		t.Errorf(`取所有同学的缓存失败`)
	}

	cache2 := NewMemCacheT[[]*TestStudentModel]()
	cache2.SetOne(`students_by_arr`, students)
	if __students, ok := cache2.GetOne(`students_by_arr`); ok {
		assert.Equal(t, 4, len(__students))
		for k, v := range __students {
			log.Println(`---range-students---`, k, v.Name, v.Class, v.Gender)
		}
	} else {
		t.Errorf(`取所有同学的缓存失败`)
	}
}
