/* Ejercicio 3 - Cita (Rendezvous)
   Verificar: es imposible que un jugador esté en Fase 2
              mientras el otro no ha llegado al punto de control.
   Ejecutar verificacion:
     spin -a rendezvous.pml && gcc -o pan pan.c && ./pan
*/

#define wait(s)   atomic { s > 0 -> s-- }
#define signal(s) s++

byte j1Llego = 0;   /* semaforo de llegada Jugador 1 */
byte j2Llego = 0;   /* semaforo de llegada Jugador 2 */

/* banderas para rastrear en que fase esta cada jugador */
bool j1_en_fase2 = false;
bool j2_en_fase2 = false;
bool j1_llego_control = false;
bool j2_llego_control = false;

active proctype Jugador_1() {
    /* Fase 1 */
    printf("J1: en Fase 1\n");
    j1_llego_control = true;
    signal(j1Llego);   /* aviso que llegue al punto de control */
    wait(j2Llego);     /* espero a Jugador 2 */

    /* Invariante: J2 debe haber llegado antes de que J1 entre a Fase 2 */
    assert(j2_llego_control == true);
    j1_en_fase2 = true;
    printf("J1: en Fase 2 - mision iniciada\n");
}

active proctype Jugador_2() {
    /* Fase 1 */
    printf("J2: en Fase 1\n");
    j2_llego_control = true;
    signal(j2Llego);   /* aviso que llegue */
    wait(j1Llego);     /* espero a Jugador 1 */

    /* Invariante: J1 debe haber llegado antes de que J2 entre a Fase 2 */
    assert(j1_llego_control == true);
    j2_en_fase2 = true;
    printf("J2: en Fase 2 - mision iniciada\n");
}
