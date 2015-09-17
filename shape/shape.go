package main
import("fmt";"math")
var h,b,r float64


func Arear( he float64,be float64) float64 {
return he*be
}

func Areac(r float64) float64{
return math.Pi *r *r
}

func Perimeterr(he float64,be float64) float64{
return 2*(he+be)
}

func Perimeterc(r float64) float64{
return 2*math.Pi*r
}
func main(){

fmt.Println("enter the radius of the circle and sides of the rect")
fmt.Scanf("%f%f%f",&r,&h,&b)

fmt.Println("circle area is ",Areac(r))
fmt.Println("cicle perimter is ",Perimeterc(r))

fmt.Println("rectangle area is ",Arear(h,b))
fmt.Println("rectangle perimter is ",Perimeterr(h,b))
}
