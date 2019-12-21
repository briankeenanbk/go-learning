package learning

import "fmt"

type Export struct {
}

//DoMagic
func (c Export) DoMagic() {
	fmt.Println("Magic function was called")
}

func (c Export) String() string {
	return fmt.Sprint("ta da da! \n")
}
