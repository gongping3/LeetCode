// 定义一个栈结构
type Stack struct {
	data []byte // 使用切片存储数据
}

// 压栈操作
func (s *Stack) Push(value byte) {
	s.data = append(s.data, value) // 将元素添加到切片末尾
}

// 弹栈操作
func (s *Stack) Pop() (byte, bool) {
	if len(s.data) == 0 {
		return 0, false // 如果栈为空，返回false
	}
	value := s.data[len(s.data)-1]  // 获取最后一个元素
	s.data = s.data[:len(s.data)-1] // 移除最后一个元素
	return value, true
}

// 查看栈顶元素（不弹出）
func (s *Stack) Peek() (byte, bool) {
	if len(s.data) == 0 {
		return 0, false // 栈为空
	}
	return s.data[len(s.data)-1], true
}

// 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}

func isValid(s string) bool {
	stack := &Stack{}
	couple := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		// 如果是左括号，压入栈
		if ch == '(' || ch == '{' || ch == '[' {
			stack.Push(ch)
		} else {
			// 栈顶元素与当前括号匹配
			top, ok := stack.Pop()
			if !ok || couple[top] != ch {
				return false
			}
		}
	}

	// 栈为空时表示所有括号匹配
	return stack.IsEmpty()
}
