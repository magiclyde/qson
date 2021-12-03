/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/12/3 下午4:36
 * @note:
 */

package qson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	assert.Equal(t, "helloqson", Concat("hello", "qson"))
}

func TestSB(t *testing.T) {
	str := "teststr"
	byt := S2B(str)
	assert.Equal(t, str, B2S(byt))
}
