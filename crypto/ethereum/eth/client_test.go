//go:build integration
// +build integration

package eth_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hiromaily/go-crypto-wallet/pkg/testutil"
)

type clientTest struct {
	testutil.ETHTestSuite
}

// TestBalanceAt is test for BalanceAt
func (ct *clientTest) TestBalanceAt() {
	type args struct {
		addr string
	}
	type want struct {
		balance uint64
		isErr   bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "happy path",
			args: args{"0x967B50a5E4d1D35Fa9aAf7DB8A391b0546209fD2"},
			want: want{100, false},
		},
		{
			name: "happy path",
			args: args{"0x16e486ED0148A59C58400232D5b4AF6bE4dC1ef0"},
			want: want{100, false},
		},
		{
			name: "happy path",
			args: args{"0x57033fC5434F3271e83B4695dAd726F348b854c2"},
			want: want{100, false},
		},
		{
			name: "happy path",
			args: args{"0x3727eE9FA88B21a0703946f9afEE3930f5980c15"},
			want: want{100, false},
		},
		{
			name: "happy path",
			args: args{"0xe933a3318C3f5D94c2A3B2BEAEF772F67a45311c"},
			want: want{100, false},
		},
		{
			name: "happy path",
			args: args{"0xe933a3318C3f5D94c2A3B2BEAEF772F67a45311c"},
			want: want{100, false},
		},
		{
			name: "address is random string",
			args: args{"0xe933a3318C3f5D94c2A3B2BEAEF772F67a45314d"},
			want: want{100, false},
		},
		{
			name: "address has no 0x",
			args: args{"e933a3318C3f5D94c2A3B2BEAEF772F67a45311c"},
			want: want{100, false},
		},
		{
			name: "address is btc address",
			args: args{"2N4TcHSCteXwiF2dj8SQijj3w2HieR4x6r5"},
			want: want{100, true},
		},
	}

	for _, tt := range tests {
		ct.T().Run(tt.name, func(t *testing.T) {
			balance, err := ct.ETH.BalanceAt(tt.args.addr)
			ct.Equal(tt.want.isErr, err != nil)
			if err == nil {
				t.Log(balance)
			}
		})
	}
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(clientTest))
}
