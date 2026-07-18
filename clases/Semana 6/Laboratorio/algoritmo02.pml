//patrón productor-consumidor
//semáforo
//1.- binario
//2.- contador
#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

#define L 20 //Finito
//modelo buffer finito
byte buffer[L]
byte top = 0

//control - semáforo
byte noVacio = 0
byte noLleno = L
byte mutex=1

active proctype Productor(){
    byte d
    do 
    ::  
        d++
        wait(noLleno) //hasta q punto está lleno
        wait(mutex) //bloquea
        buffer[top]=d
        top++
        printf("Proceso P genera valor=%d\n",d)
        signal(mutex) //libera
        signal(noVacio)
    od
}

active proctype Consumidor(){ 
    byte d
    do 
    ::  
        wait(noVacio)
        wait(mutex)
        top--
        d=buffer[top]
        printf("Proceso C consume valor=%d\n",d)
        signal(mutex)
        signal(noLleno)
    od
}