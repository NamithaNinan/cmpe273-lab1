package main
import ("testing";"fmt"
"time")
type testpair struct {
values int
}
var tests = []testpair{{1},}

func TestTimes(t *testing.T){
for _, pair :=range tests{
	startime:=time.Now()
	Times(1)
	v:=time.Since(startime)
	a:=int(v)/1000000000
	fmt.Println(v,a,pair.values)
	if a!=pair.values{
	t.Error("sorry")
	}
	
	}
	}