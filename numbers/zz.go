package numbers
import (
    //"strconv"
)

type ZZ struct {
    base uint64
    coefficients []uint32
    sign uint8
}

func NewZZ(val string) *ZZ {
    z := new(ZZ)
    
    z.coefficients = make([]uint32, len(val))

    if val[0] == '-' {
        z.sign = 1 
        val = val[1:]
    } else {
        z.sign = 0
    }

    for pos, char := range val {
        z.coefficients[pos] = uint32(char - '0')
    }
    return z
}

func (a *ZZ, b *ZZ) Add *ZZ {
    carry := 0

}
