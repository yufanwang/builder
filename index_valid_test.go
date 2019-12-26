package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	ID       int64  `validx:"index,col:id"`
	Name     string `validx:"index,col:name"`
	Age      int
	NickName string `validx:"index,col:nick_name"`
	Country  int32
	Email    string
}

func TestIndexValidSingle(t *testing.T) {
	cond1 := If(1 > 0, Eq{"a": 1}, Eq{"b": 1})
	idx1, err := IdxValid(&Data{}, cond1)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx1)

	cond2 := Eq{"id": 1}
	idx2, err := IdxValid(&Data{}, cond2)
	assert.NoError(t, err)
	assert.EqualValues(t, true, idx2)

	cond3 := Gt{"age": 1}
	idx3, err := IdxValid(&Data{}, cond3)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx3)
}

func TestIndexValidAnd(t *testing.T) {
	cond1 := And(Eq{"a": 1}, Eq{"b": 1})
	idx1, err := IdxValid(&Data{}, cond1)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx1)

	cond2 := And(Eq{"id": 1}, Like{"email", "qqqq"})
	idx2, err := IdxValid(&Data{}, cond2)
	assert.NoError(t, err)
	assert.EqualValues(t, true, idx2)

	cond3 := And(Eq{"a": 1}, Like{"nick_name", "qqqq"})
	idx3, err := IdxValid(&Data{}, cond3)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx3)

	cond4 := And(Eq{"a": 1}, Like{"nick_name", "%qqqq"})
	idx4, err := IdxValid(&Data{}, cond4)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx4)

	cond5 := And(Eq{"a": 1}, Eq{"b": "qqqq"}, Eq{"id": 1})
	idx5, err := IdxValid(&Data{}, cond5)
	assert.NoError(t, err)
	assert.EqualValues(t, true, idx5)

	cond6 := And(Eq{"a": 1}, Eq{"b": "qqqq"}, Eq{"id": "1"})
	idx6, err := IdxValid(&Data{}, cond6)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx6)
}
