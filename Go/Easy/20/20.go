type Stack []rune

func isValid(s string) bool {
    stack := Stack{}

    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            stack.Push(c);
        } else {
            if stack.IsEmpty() {
                return false
            }
            top, _ := stack.Peek()
            if (c == ')' && top == '(') || (c == ']' && top == '[') || (c == '}' && top == '{') {
                stack.Pop();
            } else {
                return false
            }
        }
    }
    return stack.IsEmpty()
}

func (s *Stack) Push(char rune) {
	*s = append(*s, char)
}


func (s *Stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}


func (s *Stack) Peek() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	return (*s)[index], true
}


func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

