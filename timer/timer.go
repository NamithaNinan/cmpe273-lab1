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
func Times (seconds int){
 <- time.After(time.Duration(seconds)*time.Second)
	fmt.Println("Done !")
}

func main() {
seconds:=10
fmt.Println("Starting !")
Times(seconds)
}
