package basic

import (
	"fmt"
	"runtime"
	"sync"
)

func SyncPool() {
	/*
	sync.Pool 的定位不是做类似连接池的东西，
	它的用途仅仅是增加对象重用的几率，
	减少 gc 的负担，而开销方面也不是很便宜的。
	 */
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println("syncpool get a=", a, ", b=", b)

	b = p.Get().(int)
	fmt.Println("syncpool get again b=", b)

	b = p.Get().(int)
	p.Put(3)
	/*
		runtime.GC() 会清掉 sync.Pool 里的所有缓存
	*/
	runtime.GC()

	b= p.Get().(int)
	fmt.Println("syncpool get b=", b)
}


