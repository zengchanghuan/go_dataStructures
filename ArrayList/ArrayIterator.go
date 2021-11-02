package ArrayList

type Iterator interface {
	//是否有下一个
	HaxNext() bool
	//下一个
	Next() (interface{},error)
	Remove()
	GetIndex()int
}

type Iterabel interface {
	//构造初始化接口
    Iterator() Iterator
}

//构造指针访问数组
type ArrayIterator struct {
	//数组指针
	array *ArrayList
	//当前索引
	currentIndex int
}
