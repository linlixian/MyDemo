package computer

import (
	// "MyProject/test/design_pattern/strategy/strategy/strategy"
	"MyProject/models/strategy"
	"fmt"
)

type Computer struct {
	Num1, Num2 int
	strate     strategy.Strategier
}

func (p *Computer) SetStrategy(strate strategy.Strategier) {
	p.strate = strate
}

func (p Computer) Do() int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
		}
	}()

	if p.strate == nil {
		panic("Strategier is null")
	}

	return p.strate.Compute(p.Num1, p.Num2)
}
