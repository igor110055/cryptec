package kraken

// Endpoints
type Endpoint int64

const (
	UNDEFINED Endpoint = iota
	BALANCES
	LEDGERS
)

func (e Endpoint) String() string {
	switch e {
	case BALANCES:
		return "/private/Balance"
	case LEDGERS:
		return "/private/Ledgers"
	}
	return "unknown"
}
