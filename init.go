package gojieba

import (
	"github.com/jizizr/jieba/deps/cppjieba"
	"github.com/jizizr/jieba/deps/limonp"
	"github.com/jizizr/jieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
