package main
import "testing"

type testpair1 struct {
radius float64 
car float64
cper float64
}
type testpair2 struct {
height float64 
breadth float64
rar float64
rper float64
}
var tests1 = []testpair1{
{ 1,3.141592653589793,6.283185307179586 },
}

var tests2 = []testpair2{
{1,1,1,4},
}
func TestFib(t *testing.T){
for _, pair :=range tests2{
	v:=Arear(pair.height,pair.breadth)
	if v!= pair.rar {
	t.Error(
	"expected rectangle area output is ", pair.rar, " got",v,
	)
	ve:=Perimeterr(pair.height,pair.breadth)
	if ve!= pair.rper {
	t.Error(
	"expected rectangle perimeter output is ", pair.rper, " got",ve,
	)
	}
	
	}
}
for _, pair :=range tests1{
	v:=Areac(pair.radius)
	if v!= pair.car {
	t.Error(
	"expected rectangle area output is ", pair.car, " got",v,
	)
	ve:=Perimeterc(pair.radius)
	if ve!= pair.cper {
	t.Error(
	"expected rectangle perimeter output is ", pair.cper, " got",ve,
	)
	}
	
	}
}}