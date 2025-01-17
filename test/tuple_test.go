package test

import (
	"testing"

	"github.com/algorand/go-algorand/data/transactions/logic"
	"github.com/stretchr/testify/require"

	"github.com/pzbitskiy/tealang/compiler"
	"github.com/pzbitskiy/tealang/dryrun"
)

func performTest(t *testing.T, source string) {
	t.Helper()
	a := require.New(t)
	result, errors := compiler.Parse(source)
	a.NotEmpty(result, errors)
	a.Empty(errors)
	teal := compiler.Codegen(result)
	op, err := logic.AssembleString(teal)
	a.NoError(err)

	pass, err := dryrun.Run(op.Program, "", nil)
	a.NoError(err)
	a.True(pass)
}

func TestAddw(t *testing.T) {
	source := `
function logic() {
	let carry, sum = addw(10, 20)
	assert(sum == 30)
	assert(carry == 0)
	return 1
}`
	performTest(t, source)
}

func TestMulw(t *testing.T) {
	source := `
function logic() {
	let high, low = mulw(10, 20)
	assert(low == 200)
	assert(high == 0)
	return 1
}`
	performTest(t, source)
}

func TestExpw(t *testing.T) {
	source := `
function logic() {
	let high, low = expw(2, 3)
	assert(low == 8)
	assert(high == 0)
	return 1
}`
	performTest(t, source)
}

func TestDivmodw(t *testing.T) {
	source := `
function logic() {
	let qhigh, qlow, rhigh, rlow = divmodw(2, 0, 0, 1)
	assert(qlow == 2)
	assert(qhigh == 0)
	assert(rlow == 0)
	assert(rhigh == 0)
	
return 1
}`
	performTest(t, source)
}
