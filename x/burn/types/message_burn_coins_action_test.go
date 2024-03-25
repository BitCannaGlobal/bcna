package types

import (
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
			name: "valid address",
			msg: MsgBurnCoinsAction{
				Creator: sample.AccAddress(),
				Amount:  sdk.NewInt64Coin("testcoin", 1000),
			},
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
