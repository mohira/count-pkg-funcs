package a

import (
	_ "fmt"
)

func foo() {}
func bar() {}

func init() {}
func init() {}
func init() {}
func init() {}
func init() {}

// func init(x int) {}

var x int

// ブランク識別子は2つ以上宣言できる
func _() {}
func _() {}

// シグニチャ変えても許される！
func _(x int)    {}
func _(s string) {}

var _ int = 1
var _ = 1

func hoge(_ int) {}

// 識別子がおけるところにはおける。が、参照はだめだよ
var _ int // ← これはOK
// a := _ + // <- 参照しているのでだめ！

type _ int

type Foo struct {
	_ int
}

func (f *Foo) name() {}

func (_ *Foo) age() {}

func (*Foo) address() {}
