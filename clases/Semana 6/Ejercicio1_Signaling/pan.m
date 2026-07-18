#define rand	pan_rand
#define pthread_equal(a,b)	((a)==(b))
#if defined(HAS_CODE) && defined(VERBOSE)
	#ifdef BFS_PAR
		bfs_printf("Pr: %d Tr: %d\n", II, t->forw);
	#else
		cpu_printf("Pr: %d Tr: %d\n", II, t->forw);
	#endif
#endif
	switch (t->forw) {
	default: Uerror("bad forward move");
	case 0:	/* if without executable clauses */
		continue;
	case 1: /* generic 'goto' or 'skip' */
		IfNotBlocked
		_m = 3; goto P999;
	case 2: /* generic 'else' */
		IfNotBlocked
		if (trpt->o_pm&1) continue;
		_m = 3; goto P999;

		 /* PROC Proceso_B */
	case 3: // STATE 1 - signaling.pml:22 - [((sem>0))] (4:0:1 - 1)
		IfNotBlocked
		reached[1][1] = 1;
		if (!((((int)now.sem)>0)))
			continue;
		/* merge: sem = (sem-1)(0, 2, 4) */
		reached[1][2] = 1;
		(trpt+1)->bup.oval = ((int)now.sem);
		now.sem = (((int)now.sem)-1);
#ifdef VAR_RANGES
		logval("sem", ((int)now.sem));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 4: // STATE 4 - signaling.pml:24 - [assert((aTermino==1))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][4] = 1;
		spin_assert((((int)now.aTermino)==1), "(aTermino==1)", II, tt, t);
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 5 - signaling.pml:25 - [printf('B: procesando y visualizando datos.\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[1][5] = 1;
		Printf("B: procesando y visualizando datos.\n");
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 6 - signaling.pml:26 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[1][6] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */

		 /* PROC Proceso_A */
	case 7: // STATE 1 - signaling.pml:15 - [printf('A: descargando paquete...\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		Printf("A: descargando paquete...\n");
		_m = 3; goto P999; /* 0 */
	case 8: // STATE 2 - signaling.pml:16 - [aTermino = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[0][2] = 1;
		(trpt+1)->bup.oval = ((int)now.aTermino);
		now.aTermino = 1;
#ifdef VAR_RANGES
		logval("aTermino", ((int)now.aTermino));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 9: // STATE 3 - signaling.pml:17 - [sem = (sem+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][3] = 1;
		(trpt+1)->bup.oval = ((int)now.sem);
		now.sem = (((int)now.sem)+1);
#ifdef VAR_RANGES
		logval("sem", ((int)now.sem));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 10: // STATE 4 - signaling.pml:18 - [printf('A: senal enviada.\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[0][4] = 1;
		Printf("A: senal enviada.\n");
		_m = 3; goto P999; /* 0 */
	case 11: // STATE 5 - signaling.pml:19 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][5] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */
	case  _T5:	/* np_ */
		if (!((!(trpt->o_pm&4) && !(trpt->tau&128))))
			continue;
		/* else fall through */
	case  _T2:	/* true */
		_m = 3; goto P999;
#undef rand
	}

