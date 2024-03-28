package types

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/BitCannaGlobal/bcna/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgBurnCoinsAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBurnCoinsAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgBurnCoinsAction{
				Creator: "invalid_address",
				Amount:  sdk.NewInt64Coin("testcoin", 1000),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address and valid amount",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.NewInt64Coin("testcoin", 1000),
			},
		}, {
			name: "negative amount",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.Coin{Denom: "testcoin", Amount: sdk.NewInt(-1000)},
			},
			err: fmt.Errorf("amount must be positive or valid"),
		}, {
			name: "zero amount",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.NewInt64Coin("testcoin", 0),
			},
			err: fmt.Errorf("amount must be positive or valid"),
		}, {
			name: "invalid denom",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			err: fmt.Errorf("amount must be positive or valid"),
		},
		{
			name: "excessive token amount length",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.Coin{Denom: "ubcna", Amount: sdk.NewInt(math.MaxInt64)}, //9223372036854775807
			},
			err: fmt.Errorf("token amount exceeds maximum length of 18 digits"),
		},
		{
			name: "excessive denom length",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.Coin{Denom: strings.Repeat("a", 129), Amount: sdk.NewInt(1000)},
			},
			err: fmt.Errorf("amount must be positive or valid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.Contains(t, err.Error(), tt.err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
