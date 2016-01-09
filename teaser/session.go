package teaser

import "github.com/cznic/mathutil"

var Sessions = make(map[int]*string)
var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

func CreateSessionId() int {
	return sessionIdGen.Next()
}
