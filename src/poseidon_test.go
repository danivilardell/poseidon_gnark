package poseidon

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

func TestPoseidon1(t *testing.T) {

	witness := poseidonCircuit{}
	witness.In = [2]frontend.Variable{1, 2}

	out := new(big.Int)
	out, ok := out.SetString("7853200120776062878684798364095072458815029376092732009249414926327459813530", 10)
	if !ok {
		t.Fatal("could not parse big int")
	}
	witness.Out = frontend.Variable(out)

	assert := test.NewAssert(t)
	assert.ProverSucceeded(&poseidonCircuit{}, &witness, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.NoFuzzing(), test.NoSerialization())
}
