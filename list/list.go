package list

import "errors"

func Remove(numbers []int, n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("เลขลบไม่ได้อะ")
	}
	head := numbers[:n]
	end := numbers[n+1:]
	ns := append(head, end...)
	return ns, nil
}
