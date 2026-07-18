/* Ejercicio 4 - Multiplex
   Verificar: el numero de tecnicos en la sala NUNCA supera N.
   TECNICOS > N para modelar carga que supera la capacidad.
   Ejecutar verificacion:
     spin -a multiplex.pml && gcc -o pan pan.c && ./pan
*/

#define wait(s)   atomic { s > 0 -> s-- }
#define signal(s) s++
#define N        3   /* capacidad maxima de la sala */
#define TECNICOS 7   /* tecnicos intentando entrar (supera N) */

byte multiplex = N;  /* semaforo inicializado con la capacidad */
byte en_sala   = 0;  /* contador de tecnicos actuales en sala  */
byte cnt_mutex = 1;  /* protege el contador en_sala            */

active [TECNICOS] proctype Tecnico() {
    wait(multiplex);      /* wait: toma un cupo (bloquea si sala llena) */

    /* Entrada: incrementar contador */
    wait(cnt_mutex);
    en_sala++;
    assert(en_sala <= N); /* INVARIANTE: nunca mas de N en sala */
    signal(cnt_mutex);

    printf("Tecnico %d entro. En sala: %d/%d\n", _pid, en_sala, N);

    /* Trabajo en la sala (modelado como sección critica) */

    /* Salida: decrementar contador */
    wait(cnt_mutex);
    en_sala--;
    signal(cnt_mutex);

    signal(multiplex);    /* signal: libera el cupo */
    printf("Tecnico %d salio. En sala: %d/%d\n", _pid, en_sala, N);
}
