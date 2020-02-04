package basic

import (
	"fmt"
	"sync"
)

type myMutex sync.Mutex

type myLocker sync.Locker

func sync2() {
	/*
		var mtx myMutex
		mtx.Lock()   //error 新的类型定义不会给他分配实际的空间
		mtx.Unlock() // error
	 */

	var mtx sync.Mutex
	mtx.Lock()   //error
	mtx.Unlock() // error

	var lock myLocker = new(sync.Mutex)  //TODO 待理解
	lock.Lock() //ok
	lock.Unlock() //ok
}

func sync1() {
	wg := sync.WaitGroup{}

	const synNum int = 5
	wg.Add(synNum)
	for i := 0; i < synNum; i++ {
		go func(i int) {
			fmt.Println(i, " sync1 done-1")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(" sync1 end")
}

func Sync() {
	sync1()
	sync2()
}
