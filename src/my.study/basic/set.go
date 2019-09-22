package basic

import (
	"github.com/cihub/seelog"
	"sync"
)

/*
	set集合只有键，没有值，

*/
type Set struct {
	/*
		对set的所有操作， 都可以通过值为bool型的map来完成
		因为这里的map不需要值，所有值得类型设为bool型
		一般set是字符串集合， 这里为操作方便， 用整数集合
	*/
	m map[int]bool
	/*
		读写锁
		一般世界写父类， 不会在为父类写成员变量名，
		这样子类对象调用父类的成员，和调用子类成员的写法完全一样
	*/
	sync.RWMutex
}

/*
函数返回声明为结构体指针， 返回的必须是结构体的地址
*/
func New() *Set {
	/*
		里面的锁不用初始化
	*/
	return &Set{
		m: map[int]bool{},
	}
}

/*
对集合类的标准函数：
增， 删， 集合元素全部清楚， 长度计算，判断有无
*/

/*
	因为需要改变结构体成员里的值， 所以要用结构体指针
*/
func (s *Set) Add(item int) {
	/*
		简单变量锁的基本用法， 这个函数只用于变量赋值，
		进来先上锁
		出去就解锁
		s.Lock()
		defer s.Unlock()
		是函数开始的标准写法
	*/
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	/*
		map是键值对集合， 通过键取到值， 返回两个数据，
		一个是取到的值，一个表示有没有取到， 这里用第二个数据
	*/
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	/*
		s.List()放回一个切片， len可以计算出切片的长度
	*/
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	/*
		map可以随意增加长度， 也可以随意清空，清空的写法就是空的大括号即可
	*/
	s.m = map[int]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

/*
返回set集合里面所有的元素
*/
func (s *Set) List() []int {
	/*
		只对读的动作加锁， 就用Rlock()
	*/
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func Set1() {
	seelog.Info("<-------------------- Set begin ----------------->")
	s := New()

	s.Add(1)
	s.Add(1)
	s.Add(2)

	s.Clear()
	if s.IsEmpty() {
		seelog.Info("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		seelog.Info("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	seelog.Info("list of all items", s.List())

	seelog.Info("<-------------------- Set end ----------------->")
}
