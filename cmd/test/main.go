package main

import (
	"fmt"
	"github.com/gardener/controller-manager-library/cmd/test/certs"
	"github.com/gardener/controller-manager-library/cmd/test/field"
	"github.com/gardener/controller-manager-library/cmd/test/match"
	"github.com/gardener/controller-manager-library/cmd/test/scheme"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller"
)

var values = map[controller.ResourceKey]int{}

func main() {

	if false {
		field.FieldMain()
	}

	if false {
		scheme.SchemeMain()
	}
	if false {
		match.MatchMain()
	}
	if true {
		certs.CertsMain()
	}
}

type Interface interface {
	Func()
}

type A struct{}
type B struct{}

type Common struct {
	Interface
	*B
}

func (*A) Func() {
	fmt.Printf("A.Func\n")
}
func (*B) Func() {
	fmt.Printf("B.Func\n")
}

func main0() {
	//c := &Common{&A{}, &B{}}
	//c.Func()
}

func main1() {
	k1 := controller.NewResourceKey("a", "b")
	k2 := controller.NewResourceKey("a", "b")
	values[k1] = 1
	fmt.Printf("k1: %d\n", values[k1])
	fmt.Printf("k2: %d\n", values[k2])

	fmt.Printf("cluster mapping: %s", set)
}

type C struct {
	name string
}

type S struct {
	m map[string]*C
}

var set = &S{map[string]*C{"a": &C{"A"}}}

func (c *C) String() string {
	return c.name
}
func (c *S) String() string {
	return fmt.Sprintf("%v", c.m)[3:]
}
