package convert

import "strconv"

type StrTo string

// 返回字串
func (s StrTo) String() string {
	return string(s)
}

// 返回Int
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

//Int 但不會返回錯誤
func (s StrTo) MustInt() int {
	v, _ := strconv.Atoi(s.String())
	return v
}

//Uint32
func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

//Uint32 但不會返回錯誤

func (s StrTo) MustUint32() uint32 {
	v, _ := strconv.Atoi(s.String())
	return uint32(v)
}
