package generate

import (
	"math/rand"
	"time"
)

func GeneratePassword(length int, includeSpecial bool) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	digits := "0123456789"
	specials := "~=+%^*/()[]{}!@#"
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var allChars string
	allChars = lower + upper + digits

	buf := make([]byte, length)
	buf[0] = digits[random.Intn(len(digits))]
	buf[1] = lower[random.Intn(len(lower))]
	buf[2] = upper[random.Intn(len(upper))]

	if includeSpecial {
		allChars = allChars + specials
	}

	for i := 3; i < length; i++ {
		buf[i] = allChars[random.Intn(len(allChars))]
	}

	random.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}
