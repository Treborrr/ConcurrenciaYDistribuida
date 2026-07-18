package main

import (
	"fmt"
	"time"
)

func P(ch1 chan string){
	
}

func Q(ch2 chan string, ch3 chan string){
	time.Sleep(100*time.Millisecond)
	ch2 <- <- ch3 
	

}


func main(){
// canales sincronos
ch1:=make(chan string)
ch2:=make(chan string)
// canal asincrono
ch3:=make(chan string,1)

go P(ch1)
go Q(ch2, ch3)

//usa el selector para recepcionar la inf de los canales 1 y 2
// 1 intento de lectura
for i := 0; i < 2; i++ {
	select {
	case msg1 := <-ch1:
		fmt.Println("mensaje de ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("mensaje de ch2:", msg2)
	}
}

}