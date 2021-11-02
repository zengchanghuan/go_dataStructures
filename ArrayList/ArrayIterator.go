package ArrayList

import "errors"

type Iterator interface {
	// HaxNext 是否有下一个
	HaxNext() bool

	// Next 下一个
	Next() (interface{},error)

	// Remove 删除
	Remove()

	// GetIndex 得到索引
	GetIndex()int
}

type Iterable interface {
	// Iterator 构造初始化接口
    Iterator() Iterator
}

// ArrayIterator 构造指针访问数组
type ArrayIterator struct {
	//数组指针
	array *ArrayList

	//当前索引
	currentIndex int
}

func (iterator *ArrayIterator) HaxNext() bool {
	return iterator.currentIndex < iterator.array.TheSize
}

func (iterator *ArrayIterator) Next() (interface{}, error) {
	if !iterator.HaxNext() {
		return nil,errors.New("没有下一个")
	}
	value,err := iterator.array.Get(iterator.currentIndex)
	iterator.currentIndex++
	return value,err
}

func (iterator *ArrayIterator) Remove() {
	iterator.currentIndex--
	iterator.array.Delete(iterator.currentIndex)
}

func (iterator *ArrayIterator) GetIndex() int {
	return iterator.currentIndex
}

func (array *ArrayList) Iterator() Iterator  {
	iterator := new(ArrayIterator)
	iterator.currentIndex = 0
	iterator.array = array
	return iterator
}