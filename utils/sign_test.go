package utils

import (
	"fmt"
	"testing"
)

func TestSigin(t *testing.T) {
	fmt.Printf(Sign("wx12sscwweewwqwwqqs", "1471443563", "1297948553"))
}

//func TestMsgSigin(t *testing.T) {
//	fmt.Printf(MsgSign("HEYIHSI", "hesss", "askasdkas", "ssss"))
//}
