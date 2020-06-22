package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int //一般的な公開鍵のx座標
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}