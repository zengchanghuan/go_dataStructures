package ArrayList

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
	array *ArrayList
	//最大范围
	capsize int
}

func NewStack() *Stack  {
	stack := new(Stack)
	stack.array = NewArray()
	stack.capsize = 10
	return stack
}

func (stack *Stack)Clear()  {
	stack.array.Clear()
	stack.capsize = 10
}

func (stack *Stack)Size() int {
	return stack.array.TheSize
}

func (stack *Stack)Pop()interface{}  {
	if !stack.IsEmpty() {
		last := stack.array.dataStore[stack.array.TheSize - 1]
		stack.array.Delete(stack.array.TheSize - 1)
		return last
	}
	return nil
}

func (stack *Stack)Push(data interface{})  {
	if !stack.IsFull() {
		stack.array.Append(data)
	}
}

func (stack *Stack)Peek()(data interface{})  {
	if !stack.IsEmpty() {
		last, _ := stack.array.Get(stack.array.TheSize - 1)
		return last
	}
	return nil
}

func (stack *Stack)IsFull() bool {
	if stack.array.TheSize >= stack.capsize {
		return true
	} else {
		return false
	}
}

func (stack *Stack)IsEmpty() bool  {
	if stack.array.TheSize == 0 {
		return true
	} else {
		return false
	}
}
