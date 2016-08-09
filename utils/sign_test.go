package utils

import (
	"fmt"
	"testing"
)

func TestSigin(t *testing.T) {
	fmt.Printf(Sign("HEYIHSI", "hesss", "askasdkas"))
}
func TestMsgSigin(t *testing.T) {
	fmt.Printf(MsgSign("HEYIHSI", "hesss", "askasdkas", "ssss"))
}
