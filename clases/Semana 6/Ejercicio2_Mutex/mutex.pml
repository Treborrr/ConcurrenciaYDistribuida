/* Ejercicio 2 - Exclusion Mutua (Mutex)
   Verificar: saldo == N_HILOS al finalizar todos los procesos.
   N_HILOS reducido a 5 (100 es costoso para model checking).
   Ejecutar verificacion:
     spin -a mutex.pml && gcc -o pan pan.c && ./pan
*/

#define wait(s)   atomic { s > 0 -> s-- }
#define signal(s) s++
#define N_HILOS   5

byte mutex = 1;   /* token unico: solo un proceso en seccion critica */
int  saldo = 0;
byte done  = 0;
byte done_mutex = 1;

active [N_HILOS] proctype Trabajador() {
    wait(mutex);
    /* --- Sección Crítica --- */
    saldo++;
    /* ----------------------- */
    signal(mutex);

    /* registrar finalizacion */
    wait(done_mutex);
    done++;
    signal(done_mutex);
}

/* Verificador: espera a que todos terminen y comprueba el invariante */
active proctype Verificador() {
    (done == N_HILOS); /* guarda: bloquea hasta que done sea N_HILOS */
    assert(saldo == N_HILOS);
    printf("VERIFICACION OK: saldo = %d\n", saldo);
}
