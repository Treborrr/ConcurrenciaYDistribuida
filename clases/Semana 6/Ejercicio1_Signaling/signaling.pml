/* Ejercicio 1 - Patron de Senalizacion (Signaling)
   Verificar: Proceso_B nunca ejecuta su tarea sin la senal de Proceso_A.
   Ejecutar verificacion:
     spin -a signaling.pml && gcc -o pan pan.c && ./pan
*/

#define wait(s)   atomic { s > 0 -> s-- }
#define signal(s) s++

byte sem = 0;          /* semaforo inicializado en 0 */
bool aTermino = false; /* bandera de verificacion    */

active proctype Proceso_A() {
    /* Descarga del paquete de datos */
    printf("A: descargando paquete...\n");
    aTermino = true;   /* marca que A termino ANTES de hacer signal */
    signal(sem);
    printf("A: senal enviada.\n");
}

active proctype Proceso_B() {
    wait(sem);         /* bloquea hasta recibir senal de A */
    /* Invariante de seguridad: A debe haber terminado antes de llegar aqui */
    assert(aTermino == true);
    printf("B: procesando y visualizando datos.\n");
}
