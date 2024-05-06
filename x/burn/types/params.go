package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Keys for parameter access
var (
	KeyBurnDenom = []byte("BurnDenom")
)

// // Params defines the parameters for the module.
// type Params struct {
// 	BurnDenom string `json:"burn_denom" yaml:"burn_denom"`
// }

// NewParams creates a new Params instance
func NewParams(burnDenom string) Params {
	return Params{
		BurnDenom: burnDenom,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		BurnDenom: "ubcna",
	}
}

// ParamKeyTable the param key table for the module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBurnDenom, &p.BurnDenom, validateBurnDenom),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBurnDenom(p.BurnDenom); err != nil {
		return err
	}
	return nil
}

// ValidateBurnDenom validates the BurnDenom parameter
func validateBurnDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v == "" {
		return fmt.Errorf("burn denom cannot be empty")
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
