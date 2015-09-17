package main
import "fmt"
var numbers,i int
func fib(numbers int) int{
if numbers==0{
	return 0
}else if numbers ==1{
	return 1
	}else{
	return fib(numbers-1) + fib(numbers-2)}
}
func main(){
fmt.Println("enter a number")
fmt.Scanf("%d", &numbers)
i=0
fmt.Println("the fibonacci numbers are the following")
for i<numbers{
fmt.Println(fib(i))
i=i+1	
}
}

/////////////////////////////////////////////////////////
 
Defer is basically used when we need to delay a function from execution for a while till the completion of another function. It is mainly used in cases where we need some resources to be freed before its execution.
Panic on the other hand helps in error handling by causing a run time failure. This can be prevented by recover. It stops the panic from occuring by returning a value, this was passed on to panic when it had stopped its execution
This usually occurs in cases of error from the programmers side and it can lead it to huge unrecoverable errors later.
package main
import "fmt"
func main(){
	defer func(){
	example:=recover()
	fmt.Println(example)
	}()
panic("oh no there's a problem")
}

////////////////////////////////////////////////////////
package main
import("fmt";"math")
type rectangle struct{
h,b float64
}

type circle struct{
r float64}

type Shape interface {
area() float64
perimeter() float64
}

func (re *rectangle) area() float64 {
return re.h*re.b
}

func (c *circle) area() float64{
return math.Pi *c.r * c.r
}

func (re *rectangle) perimeter() float64{
return 2*(re.h+re.b)
}

func (c*circle) perimeter() float64{
return 2*math.Pi*c.r
}
func main(){
var h,b,r float64
fmt.Println("enter the radius of the circle and sides of the rect")
fmt.Scanf("%f%f%f",&r,&h,&b)
c:=circle{r}
fmt.Println("circle area is ",c.area())
fmt.Println("cicle perimter is ",c.perimeter())
re:=rectangle{h,b}
fmt.Println("rectangle area is ",re.area())
fmt.Println("rectangle perimter is ",re.perimeter())
}

/////////////////////////////////////////////////////

package main
import ("fmt"
"time")

type Duration int64
const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)

func main() {
seconds:=10
fmt.Println("Starting !")
select {
case <- time.After(time.Duration(seconds)*time.Second):
	fmt.Println("Done !")
}

/////////////////////////////////////////////////////////
}