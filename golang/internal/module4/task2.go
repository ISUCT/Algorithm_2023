package module4

type Stack struct {
	buffer [101]int
	top    int
}

func (s *Stack) push(a int) {
	s.top++
	s.buffer[s.top] = a
}

func (s *Stack) pop() {
	s.top--
}

func (s *Stack) size() int {
	return s.top + 1
}

func (s *Stack) empty() bool {
	return s.top == -1
}

func (s *Stack) clear() {
	s.top = -1
}

func (s *Stack) back() int {
	return s.buffer[s.top]
}

func Task2() {

	// st := Stack{}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	s := scanner.Text()
	// 	if s == "exit" {
	// 		break
	// 	}
	// 	if s == "push" {
	// 		scanner.Scan()
	// 		n, _ := strconv.Atoi(scanner.Text())
	// 		st.push(n)
	// 		fmt.Println("ok")
	// 	} else if s == "pop" {
	// 		if st.empty() {
	// 			fmt.Println("error")
	// 		} else {
	// 			fmt.Println(st.back())
	// 			st.pop()
	// 		}
	// 	} else if s == "back" {
	// 		if st.empty() {
	// 			fmt.Println("error")
	// 		} else {
	// 			fmt.Println(st.back())
	// 		}
	// 	} else if s == "size" {
	// 		fmt.Println(st.size())
	// 	} else if s == "clear" {
	// 		st.clear()
	// 		fmt.Println("ok")
	// 	}
	// }
	// fmt.Println("bye")
}
