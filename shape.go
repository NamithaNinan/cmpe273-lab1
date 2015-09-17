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
