package state

import (
	"os"
	"testing"
	"time"

	"github.com/aergoio/aergo-lib/db"
	"github.com/aergoio/aergo/types"
	"github.com/stretchr/testify/assert"
)

var chainStateDB *ChainStateDB
var stateDB *StateDB

func initTest(t assert.TestingT) {
	chainStateDB = NewChainStateDB()
	_ = chainStateDB.Init(string(db.BadgerImpl), "test", nil, false)
	stateDB = chainStateDB.GetStateDB()
	genesis := types.GetTestGenesis()

	err := chainStateDB.SetGenesis(genesis, nil)
	assert.NoError(t, err, "failed init")
}
func deinitTest() {
	_ = chainStateDB.Close()
	_ = os.RemoveAll("test")
}
func TestContractStateCode(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")
	testBytes := []byte("test_bytes")

	// open contract state
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// set code
	err = contractState.SetCode(testBytes)
	assert.NoError(t, err, "set code to contract state")

	// get code
	res, err := contractState.GetCode()
	assert.NoError(t, err, "get code from contract state")
	assert.Equal(t, testBytes, res, "different code detected")
}

func TestContractStateData(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")
	testBytes := []byte("test_bytes")
	testKey := []byte("test_key")

	// open contract state
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// set data
	err = contractState.SetData(testKey, testBytes)
	assert.NoError(t, err, "set data to contract state")

	// get data
	res, err := contractState.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.Equal(t, testBytes, res, "different data detected")

	// stage contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")
}

func TestContractStateInitialData(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")
	testBytes := []byte("test_bytes")
	testKey := []byte("test_key")

	// open contract state
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// get initial data
	res, err := contractState.GetInitialData(testKey)
	assert.NoError(t, err, "get initial data from contract state")
	assert.Nil(t, res, "get initial data from contract state")

	// set data
	err = contractState.SetData(testKey, testBytes)
	assert.NoError(t, err, "set data to contract state")

	// get data
	res, err = contractState.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.Equal(t, testBytes, res, "different data detected")

	// get initial data
	res, err = contractState.GetInitialData(testKey)
	assert.NoError(t, err, "get initial data from contract state")
	assert.Nil(t, res, "get initial data from contract state")

	// stage contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")

	// update and commit statedb
	err = stateDB.Update()
	assert.NoError(t, err, "update statedb")
	err = stateDB.Commit()
	assert.NoError(t, err, "commit statedb")

	// reopen contract state
	contractState, err = stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// get initial data
	res, err = contractState.GetInitialData(testKey)
	assert.NoError(t, err, "get initial data from contract state")
	assert.Equal(t, testBytes, res, "get initial data from contract state")
}

func TestContractStateDataDelete(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")
	testBytes := []byte("test_bytes")
	testKey := []byte("test_key")

	// open contract state and set test data
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")
	err = contractState.SetData(testKey, testBytes)
	assert.NoError(t, err, "set data to contract state")

	// stage and re-open contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")
	contractState, err = stateDB.OpenContractState(types.ToAccountID(testAddress), contractState.State)
	assert.NoError(t, err, "could not open contract state")

	// get and delete test data
	res, err := contractState.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.Equal(t, testBytes, res, "different data detected")
	err = contractState.DeleteData(testKey)
	assert.NoError(t, err, "delete data from contract state")

	// stage and re-open contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")
	contractState, err = stateDB.OpenContractState(types.ToAccountID(testAddress), contractState.State)
	assert.NoError(t, err, "could not open contract state")

	// get test data
	res, err = contractState.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.Nil(t, res, "garbage data detected")

	// stage contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")
}

func TestContractStateHasKey(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")
	testBytes := []byte("test_bytes")
	testKey := []byte("test_key")

	// open contract state and set test data
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")
	assert.False(t, contractState.HasKey(testKey))

	err = contractState.SetData(testKey, testBytes)
	assert.NoError(t, err, "set data to contract state")
	assert.True(t, contractState.HasKey(testKey))

	// get test data
	_, err = contractState.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.True(t, contractState.HasKey(testKey))

	// delete test data
	err = contractState.DeleteData(testKey)
	assert.NoError(t, err, "delete data from contract state")
	assert.True(t, contractState.HasKey(testKey))

	// stage contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")

	// update and commit
	err = stateDB.Update()
	assert.NoError(t, err, "failed to update stateDB")
	err = stateDB.Commit()
	assert.NoError(t, err, "failed to commit stateDB")

	// re-open contract state
	contractState, err = stateDB.OpenContractState(types.ToAccountID(testAddress), contractState.State)
	assert.NoError(t, err, "could not open contract state")

	// check key existence
	assert.False(t, contractState.HasKey(testKey))
}

func TestContractStateEmpty(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")

	// open contract state
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// stage contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")
}

func TestContractStateReOpenData(t *testing.T) {
	initTest(t)
	defer deinitTest()
	testAddress := []byte("test_address")
	testBytes := []byte("test_bytes")
	testKey := []byte("test_key")

	// open contract state
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// set data
	err = contractState.SetData(testKey, testBytes)
	assert.NoError(t, err, "set data to contract state")

	// get data
	res, err := contractState.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.Equal(t, testBytes, res, "different data detected")

	// stage contract state
	err = stateDB.StageContractState(contractState)
	assert.NoError(t, err, "stage contract state")

	// re-open contract state
	//contractState2, err := chainStateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	contractState2, err := stateDB.OpenContractState(types.ToAccountID(testAddress), contractState.State)
	assert.NoError(t, err, "could not open contract state")

	// get data
	res2, err := contractState2.GetData(testKey)
	assert.NoError(t, err, "get data from contract state")
	assert.Equal(t, testBytes, res2, "different data detected")
}

func TestContractStateRollback(t *testing.T) {
	initTest(t)
	defer deinitTest()

	testAddress := []byte("test_address")
	testKey := []byte("test_key")

	// open contract state
	contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
	assert.NoError(t, err, "could not open contract state")

	// test data
	_ = contractState.SetData(testKey, []byte("1")) // rev 1
	_ = contractState.SetData(testKey, []byte("2")) // rev 2
	res, _ := contractState.GetData(testKey)
	assert.Equal(t, []byte("2"), res)

	// snapshot: rev 2
	revision := contractState.Snapshot()
	t.Log("revision", revision)
	assert.Equal(t, 2, int(revision))

	// test data
	_ = contractState.SetData(testKey, []byte("3")) // rev 3
	_ = contractState.SetData(testKey, []byte("4")) // rev 4
	_ = contractState.SetData(testKey, []byte("5")) // rev 5
	res, _ = contractState.GetData(testKey)
	assert.Equal(t, []byte("5"), res)

	// rollback: rev 2
	contractState.Rollback(revision)
	res, _ = contractState.GetData(testKey)
	assert.Equal(t, []byte("2"), res)

	// rollback to empty: rev 0
	contractState.Rollback(Snapshot(0))
	res, _ = contractState.GetData(testKey)
	assert.Nil(t, res)
}

func BenchmarkContractStateKeyCheck(b *testing.B) {
	count := 10000
	benchmarks := []struct {
		name     string
		count    int
		keyCheck bool
		repeat   int
	}{
		{"True.0", count, true, 0},
		{"True.1", count, true, 1},
		{"True.2", count, true, 2},
		{"True.3", count, true, 3},
		{"False.0", count, false, 0},
		{"False.1", count, false, 1},
		{"False.2", count, false, 2},
		{"False.3", count, false, 3},
	}

	testAddress := []byte("test_address")
	testValue := []byte("test_value")
	for _, bm := range benchmarks {
		initTest(b)
		b.Run(bm.name, func(b *testing.B) {
			b.N = bm.count
			repeat := bm.repeat
			if repeat == 0 {
				repeat = 1
			}
			laptime := []time.Time{}
			for x := 0; x < repeat; x++ {
				laptime = append(laptime, time.Now())
				contractState, err := stateDB.OpenContractStateAccount(types.ToAccountID(testAddress))
				assert.NoError(b, err, "could not open contract state")
				laptime = append(laptime, time.Now())
				for i := 0; i < bm.count; i++ {
					key := []byte("key_" + string(i))
					if bm.keyCheck {
						_ = contractState.HasKey(key)
					}
					_ = contractState.SetData(key, testValue)
				}
				laptime = append(laptime, time.Now())
				if bm.repeat > 0 {
					stateDB.Update()
					stateDB.Commit()
				}
				laptime = append(laptime, time.Now())
				b.Logf("(time) open: %v ns / data: %v ns / commit: %v ns",
					laptime[1].Sub(laptime[0]).Nanoseconds(),
					laptime[2].Sub(laptime[1]).Nanoseconds(),
					laptime[3].Sub(laptime[2]).Nanoseconds())
				stateDB.StageContractState(contractState)
			}
		})
		deinitTest()
	}

}
