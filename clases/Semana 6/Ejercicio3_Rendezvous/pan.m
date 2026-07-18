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

		 /* PROC Jugador_2 */
	case 3: // STATE 1 - rendezvous.pml:35 - [printf('J2: en Fase 1\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[1][1] = 1;
		Printf("J2: en Fase 1\n");
		_m = 3; goto P999; /* 0 */
	case 4: // STATE 2 - rendezvous.pml:36 - [j2_llego_control = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[1][2] = 1;
		(trpt+1)->bup.oval = ((int)now.j2_llego_control);
		now.j2_llego_control = 1;
#ifdef VAR_RANGES
		logval("j2_llego_control", ((int)now.j2_llego_control));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 3 - rendezvous.pml:37 - [j2Llego = (j2Llego+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[1][3] = 1;
		(trpt+1)->bup.oval = ((int)now.j2Llego);
		now.j2Llego = (((int)now.j2Llego)+1);
#ifdef VAR_RANGES
		logval("j2Llego", ((int)now.j2Llego));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 4 - rendezvous.pml:38 - [((j1Llego>0))] (7:0:1 - 1)
		IfNotBlocked
		reached[1][4] = 1;
		if (!((((int)now.j1Llego)>0)))
			continue;
		/* merge: j1Llego = (j1Llego-1)(0, 5, 7) */
		reached[1][5] = 1;
		(trpt+1)->bup.oval = ((int)now.j1Llego);
		now.j1Llego = (((int)now.j1Llego)-1);
#ifdef VAR_RANGES
		logval("j1Llego", ((int)now.j1Llego));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 7: // STATE 7 - rendezvous.pml:41 - [assert((j1_llego_control==1))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][7] = 1;
		spin_assert((((int)now.j1_llego_control)==1), "(j1_llego_control==1)", II, tt, t);
		_m = 3; goto P999; /* 0 */
	case 8: // STATE 8 - rendezvous.pml:42 - [j2_en_fase2 = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[1][8] = 1;
		(trpt+1)->bup.oval = ((int)j2_en_fase2);
		j2_en_fase2 = 1;
#ifdef VAR_RANGES
		logval("j2_en_fase2", ((int)j2_en_fase2));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 9: // STATE 9 - rendezvous.pml:43 - [printf('J2: en Fase 2 - mision iniciada\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[1][9] = 1;
		Printf("J2: en Fase 2 - mision iniciada\n");
		_m = 3; goto P999; /* 0 */
	case 10: // STATE 10 - rendezvous.pml:44 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[1][10] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */

		 /* PROC Jugador_1 */
	case 11: // STATE 1 - rendezvous.pml:22 - [printf('J1: en Fase 1\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		Printf("J1: en Fase 1\n");
		_m = 3; goto P999; /* 0 */
	case 12: // STATE 2 - rendezvous.pml:23 - [j1_llego_control = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[0][2] = 1;
		(trpt+1)->bup.oval = ((int)now.j1_llego_control);
		now.j1_llego_control = 1;
#ifdef VAR_RANGES
		logval("j1_llego_control", ((int)now.j1_llego_control));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 13: // STATE 3 - rendezvous.pml:24 - [j1Llego = (j1Llego+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][3] = 1;
		(trpt+1)->bup.oval = ((int)now.j1Llego);
		now.j1Llego = (((int)now.j1Llego)+1);
#ifdef VAR_RANGES
		logval("j1Llego", ((int)now.j1Llego));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 14: // STATE 4 - rendezvous.pml:25 - [((j2Llego>0))] (7:0:1 - 1)
		IfNotBlocked
		reached[0][4] = 1;
		if (!((((int)now.j2Llego)>0)))
			continue;
		/* merge: j2Llego = (j2Llego-1)(0, 5, 7) */
		reached[0][5] = 1;
		(trpt+1)->bup.oval = ((int)now.j2Llego);
		now.j2Llego = (((int)now.j2Llego)-1);
#ifdef VAR_RANGES
		logval("j2Llego", ((int)now.j2Llego));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 15: // STATE 7 - rendezvous.pml:28 - [assert((j2_llego_control==1))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		spin_assert((((int)now.j2_llego_control)==1), "(j2_llego_control==1)", II, tt, t);
		_m = 3; goto P999; /* 0 */
	case 16: // STATE 8 - rendezvous.pml:29 - [j1_en_fase2 = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[0][8] = 1;
		(trpt+1)->bup.oval = ((int)j1_en_fase2);
		j1_en_fase2 = 1;
#ifdef VAR_RANGES
		logval("j1_en_fase2", ((int)j1_en_fase2));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 17: // STATE 9 - rendezvous.pml:30 - [printf('J1: en Fase 2 - mision iniciada\\n')] (0:0:0 - 1)
		IfNotBlocked
		reached[0][9] = 1;
		Printf("J1: en Fase 2 - mision iniciada\n");
		_m = 3; goto P999; /* 0 */
	case 18: // STATE 10 - rendezvous.pml:31 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][10] = 1;
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

