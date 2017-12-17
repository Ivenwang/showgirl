package utils

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func Int64rand() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}

func Int32rand() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31()
}

// Rand16s : 创建16位随机字符串, [0-9][a-z][A-Z]{16}
func Rand16s() string {
	// 48-57 - '0'-'9' | 65-90 - 'A'-'Z' | 97-122 - 'a'-'z'
	// 10 [0, 10), 26 [10, 36), 26 [36, 62)
	// x - 0 + '0' <- x - 10 + 'A' <- x - 36 + 'a'
	var s16 [16]byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s16); i++ {
		idx := rand.Intn(10 + 26 + 26)
		if idx >= 10+26 {
			s16[i] = byte(idx - 10 - 26 + 'a')
		} else if idx >= 10 {
			s16[i] = byte(idx - 10 + 'A')
		} else {
			s16[i] = byte(idx + '0')
		}
	}
	return string(s16[:])
}
