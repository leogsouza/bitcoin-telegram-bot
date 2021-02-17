package commands

import (
	"fmt"
	"math"

	"github.com/leogsouza/bitcoin-telegram-bot/utils"
)

func GetSummary() (string, string, error) {
	p, err := utils.GetAPICall()
	l := p.Last
	o := p.Open
	his := ((l - o) / o) * 100

	if !math.Signbit(float64(his)) {

		return fmt.Sprintf("%.2f", p.Last), fmt.Sprintf("%.2f%%", his), err
	} else {

		return fmt.Sprintf("%.2f", p.Last), fmt.Sprintf("-%.2f%%", -1*his), err
	}
}
