package numbers


type number interface {
    add() number
    subtract() number
    multiply() number
    divide() number
}

// ZZ is an arbitrarily large integer.
type ZZ struct {
    sign uint8
    base uint64
    coefficients []uint64
}

// NewZZ creates a new arbitrarily large integer.
func NewZZ() *ZZ {
    n :=  new(ZZ)
    return n
}
