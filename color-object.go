package main


import (
    "fmt"
    )
    
const(
	WHITE = iota
	BLACK
    BLUE
    RED
    YELLOW
)
    
type Color byte

type Box struct {
     width,height,depth float64
     color Color
}

func (bl BoxList) BiggestsColor() Color {
    v := 0.00
    k := Color(WHITE)
    for _, b := range bl {
        if bv := b.Volume(); bv > v {
            v = bv
            k = b.color
        }
    }
    return k
}


type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height *b.depth
}

func (b *Box) SetColor(c Color){
	b.color = c
}

func (bl BoxList) PaintItBlack(){
    for i,_:=range bl {
    	bl[i].SetColor(BLACK)
    }
}

func (c Color) String() string {
    strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
    return strings[c]
}
    
func main(){
    boxes := BoxList {
    	Box{4,4,4,RED},
    	Box{4,42,4,YELLOW},
    	Box{71,4,4,RED},
    	Box{4,8,4,WHITE},
    }
    fmt.Printf("We have %d boxes in our set\n", len(boxes))
    fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
    fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
    fmt.Println("The biggest one is", boxes.BiggestsColor().String())

    fmt.Println("Let's paint them all black")
    boxes.PaintItBlack()
    fmt.Println("The color of the second one is", boxes[1].color.String())

    fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
}

