package StackArray

type StackArray interface {
	// Clear 清空
    Clear()

	// Size 大小
	Size()int

	// Pop 弹出
	Pop() interface{}

	// Push 压栈
	Push(data interface{})

	// Peek 数据
	Peek(data interface{})

	// IsFull 栈是否满
	IsFull() bool

	// IsEmpty 栈是否空
	IsEmpty() bool
}

type Stack struct {
	dataSource [] interface{}

	//最大范围
	capsize int

	//实际大小
	currentSize int
}

func NewStack() *Stack  {
	stack := new(Stack)
	stack.dataSource = make([]interface{},0,10)
	stack.capsize = 10
	stack.currentSize = 0
	return stack
}

func (stack *Stack)Clear()  {
	stack.dataSource = make([]interface{},0,10)
	stack.capsize = 10
	stack.currentSize = 0
}

func (stack *Stack)Size() int {
	return stack.currentSize
}

func (stack *Stack)Pop()interface{}  {
	if !stack.IsEmpty() {
		last := stack.dataSource[stack.currentSize - 1]
		stack.dataSource = stack.dataSource[:stack.currentSize - 1]
		stack.currentSize--
		return last
	}
	return nil
}

func (stack *Stack)Push(data interface{})  {
	if !stack.IsFull() {
		stack.dataSource = append(stack.dataSource,data)
		stack.currentSize++
	}
}

func (stack *Stack)Peek()(data interface{})  {
	if !stack.IsEmpty() {
		last := stack.dataSource[stack.currentSize - 1]
		return last
	}
	return nil
}

func (stack *Stack)IsFull() bool {
	if stack.currentSize >= stack.capsize {
		return true
	} else {
		return false
	}
}

func (stack *Stack)IsEmpty() bool  {
	if stack.currentSize == 0 {
		return true
	} else {
		return false
	}
}