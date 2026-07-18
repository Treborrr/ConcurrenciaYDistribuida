//sincronización entre procesos
#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

#define P 5

//recursos disponibles
byte fork[P]={1,1,1,1,1}

active[P] proctype Filosofo(){ 
    pid i=_pid //id del proceso en ejecución
    do 
    :: 
        printf("Filósofo %d pensando!!\n",i)
        wait(fork[i])
        wait(fork[(i+1) % P])
        printf("Filósofo %d comiendo!!\n",i)
        signal(fork[i])
        signal(fork[(i+1) % P])
    od
}
