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

func TestIndexValid(t *testing.T) {
	cond1 := If(1 > 0, Eq{"a": 1}, Eq{"b": 1})
	idx, err := IdxValid(&Data{}, cond1)
	assert.NoError(t, err)
	assert.EqualValues(t, false, idx)

	cond2 := Eq{"id": 1}
	idx1, err := IdxValid(&Data{}, cond2)
	assert.NoError(t, err)
	assert.EqualValues(t, true, idx1)

}
