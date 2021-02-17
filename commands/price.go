package commands

import (
	"fmt"

	"github.com/leogsouza/bitcoin-telegram-bot/utils"
)

func GetPrice() (string, error) {
	p, err := utils.GetAPICall()

	return fmt.Sprintf("%.2f", p.Last), err
}
