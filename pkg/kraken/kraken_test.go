package kraken

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLedgersInfo(t *testing.T) {
	kraken := NewKraken("Ayhkz0MxatgBpe4IqA4xUEnDqBYQeVQTqEms31V7XmUVurcMX9gH6VtE", "gjFY842MYmsujg1UUTz2gXPTOIxV7Adg537cRxaR7UUG0c0/XIR4GR+8pknIWqoUkP8OsVCzdTru/lwquxUlvQ==")
	ledgersResponse, err := kraken.getLedgersInfo(getLedgersArgs{
		Asset:  "",
		Aclass: "",
		Type_:  "",
		Start:  time.Now().UnixMilli() - 60*1000,
		End:    time.Now().UnixMilli(),
		Offset: 0,
	})
	assert.Empty(t, err)
	assert.NotEmpty(t, ledgersResponse)
}

func TestGetBalance(t *testing.T) {
	kraken := NewKraken("Ayhkz0MxatgBpe4IqA4xUEnDqBYQeVQTqEms31V7XmUVurcMX9gH6VtE", "gjFY842MYmsujg1UUTz2gXPTOIxV7Adg537cRxaR7UUG0c0/XIR4GR+8pknIWqoUkP8OsVCzdTru/lwquxUlvQ==")
	r, err := kraken.getAccountBalance()
	assert.Empty(t, err)
	fmt.Printf("%s\n", *r)
}
