package data

import (
	stdJSON "encoding/json"

	"github.com/QinPengLin/starknet-go/pkg/abi"
	"github.com/goccy/go-json"
)

// Class -
type Class struct {
	EntryPointsByType EntrypointsByType  `json:"entry_points_by_type"`
	Abi               abi.Abi            `json:"-"`
	RawAbi            stdJSON.RawMessage `json:"abi"`
}

// GetAbi -
func (c *Class) GetAbi() (abi.Abi, error) {
	var abi abi.Abi
	err := json.Unmarshal(c.RawAbi, &abi)
	return abi, err
}

// Handler -
type Handler struct {
	Offset   string `json:"offset"`
	Selector string `json:"selector"`
}

// EntrypointsByType -
type EntrypointsByType struct {
	CONSTRUCTOR []Handler `json:"CONSTRUCTOR"`
	EXTERNAL    []Handler `json:"EXTERNAL"`
	L1HANDLER   []Handler `json:"L1_HANDLER"`
}
