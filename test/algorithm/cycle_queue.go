package main

import (
	"errors"
	"fmt"
)

type Cyclequeue struct {
	Queue  []int
	Front  int
	Rear   int
	Length int
}

func NewCyclequeue(len int) *Cyclequeue {
	return &Cyclequeue{
		Queue:  make([]int, len),
		Front:  0,
		Rear:   0,
		Length: len,
	}
}

func (c *Cyclequeue) InQueue(a int) error {
	if (c.Front+1)%c.Length == c.Rear {
		return errors.New("已满")
	}
	c.Queue[c.Front] = a
	c.Front = (c.Front + 1) % c.Length // 不这样会越界
	return nil
}

func (c *Cyclequeue) OutQueue() (int, error) {
	if c.Front == c.Rear {
		return 0, errors.New("已空")
	}
	out := c.Queue[c.Rear]
	c.Queue[c.Rear] = 0
	c.Rear = (c.Rear + 1) % c.Length // 不这样会越界
	return out, nil
}

func (c *Cyclequeue) Len() int {
	return (c.Rear - c.Front + c.Length) % c.Length
}

func main() {
	q := NewCyclequeue(10)
	for i := 0; i < 12; i++ {
		err := q.InQueue(1)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(q.Queue)

	}
	for i := 0; i < 12; i++ {
		_, err := q.OutQueue()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(q.Queue)

	}

	for i := 0; i < 12; i++ {
		err := q.InQueue(1)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(q.Queue)

	}
	for i := 0; i < 12; i++ {
		_, err := q.OutQueue()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(q.Queue)

	}
}
