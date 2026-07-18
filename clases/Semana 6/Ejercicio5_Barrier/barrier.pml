/* Ejercicio 5 - Barrera Reutilizable (Two-Turnstile Barrier)
   Verificar: ausencia de deadlocks en un modelo de dos fases consecutivas.
   Un proceso mas rapido no puede adelantarse a la siguiente fase sin que
   todos los demas hayan llegado a la barrera de la fase actual.
   Ejecutar verificacion (con busqueda de deadlocks):
     spin -a barrier.pml && gcc -o pan pan.c && ./pan
*/

#define wait(s)   atomic { s > 0 -> s-- }
#define signal(s) s++
#define TOTAL 3    /* numero de trabajadores */
#define FASES 2    /* fases a sincronizar    */

byte count     = 0;
byte mut       = 1;
byte molinete1 = 0; /* inicia cerrado */
byte molinete2 = 1; /* inicia abierto */

proctype Trabajador(byte id) {
    byte f = 0;
    do
    :: f < FASES ->
        printf("Trabajador %d: completando Fase %d\n", id, f + 1);

        /* === Barrera: fase de llegada === */
        wait(mut);
            count++;
            if
            :: count == TOTAL ->
                wait(molinete2);   /* cierra molinete2 para la proxima ronda */
                signal(molinete1); /* abre molinete1 */
            :: else -> skip
            fi;
        signal(mut);

        /* Molinete 1: pass-through */
        wait(molinete1);
        signal(molinete1);

        /* === Barrera: fase de salida === */
        wait(mut);
            count--;
            if
            :: count == 0 ->
                wait(molinete1);   /* cierra molinete1 para la proxima ronda */
                signal(molinete2); /* abre molinete2 */
            :: else -> skip
            fi;
        signal(mut);

        /* Molinete 2: pass-through */
        wait(molinete2);
        signal(molinete2);

        printf("Trabajador %d: paso barrera de Fase %d\n", id, f + 1);
        f++;
    :: else -> break
    od
}

init {
    atomic {
        run Trabajador(1);
        run Trabajador(2);
        run Trabajador(3);
    }
}
