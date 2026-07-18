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

		 /* PROC Tecnico */
	case 3: // STATE 1 - multiplex.pml:18 - [((multiplex>0))] (6:0:1 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		if (!((((int)now.multiplex)>0)))
			continue;
		/* merge: multiplex = (multiplex-1)(0, 2, 6) */
		reached[0][2] = 1;
		(trpt+1)->bup.oval = ((int)now.multiplex);
		now.multiplex = (((int)now.multiplex)-1);
#ifdef VAR_RANGES
		logval("multiplex", ((int)now.multiplex));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 4: // STATE 4 - multiplex.pml:21 - [((cnt_mutex>0))] (7:0:1 - 1)
		IfNotBlocked
		reached[0][4] = 1;
		if (!((((int)now.cnt_mutex)>0)))
			continue;
		/* merge: cnt_mutex = (cnt_mutex-1)(0, 5, 7) */
		reached[0][5] = 1;
		(trpt+1)->bup.oval = ((int)now.cnt_mutex);
		now.cnt_mutex = (((int)now.cnt_mutex)-1);
#ifdef VAR_RANGES
		logval("cnt_mutex", ((int)now.cnt_mutex));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 5: // STATE 7 - multiplex.pml:22 - [en_sala = (en_sala+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		(trpt+1)->bup.oval = ((int)now.en_sala);
		now.en_sala = (((int)now.en_sala)+1);
#ifdef VAR_RANGES
		logval("en_sala", ((int)now.en_sala));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 8 - multiplex.pml:23 - [assert((en_sala<=3))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][8] = 1;
		spin_assert((((int)now.en_sala)<=3), "(en_sala<=3)", II, tt, t);
		_m = 3; goto P999; /* 0 */
	case 7: // STATE 9 - multiplex.pml:24 - [cnt_mutex = (cnt_mutex+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][9] = 1;
		(trpt+1)->bup.oval = ((int)now.cnt_mutex);
		now.cnt_mutex = (((int)now.cnt_mutex)+1);
#ifdef VAR_RANGES
		logval("cnt_mutex", ((int)now.cnt_mutex));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 8: // STATE 10 - multiplex.pml:26 - [printf('Tecnico %d entro. En sala: %d/%d\\n',_pid,en_sala,3)] (0:0:0 - 1)
		IfNotBlocked
		reached[0][10] = 1;
		Printf("Tecnico %d entro. En sala: %d/%d\n", ((int)((P0 *)_this)->_pid), ((int)now.en_sala), 3);
		_m = 3; goto P999; /* 0 */
	case 9: // STATE 11 - multiplex.pml:31 - [((cnt_mutex>0))] (14:0:1 - 1)
		IfNotBlocked
		reached[0][11] = 1;
		if (!((((int)now.cnt_mutex)>0)))
			continue;
		/* merge: cnt_mutex = (cnt_mutex-1)(0, 12, 14) */
		reached[0][12] = 1;
		(trpt+1)->bup.oval = ((int)now.cnt_mutex);
		now.cnt_mutex = (((int)now.cnt_mutex)-1);
#ifdef VAR_RANGES
		logval("cnt_mutex", ((int)now.cnt_mutex));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 10: // STATE 14 - multiplex.pml:32 - [en_sala = (en_sala-1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][14] = 1;
		(trpt+1)->bup.oval = ((int)now.en_sala);
		now.en_sala = (((int)now.en_sala)-1);
#ifdef VAR_RANGES
		logval("en_sala", ((int)now.en_sala));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 11: // STATE 15 - multiplex.pml:33 - [cnt_mutex = (cnt_mutex+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][15] = 1;
		(trpt+1)->bup.oval = ((int)now.cnt_mutex);
		now.cnt_mutex = (((int)now.cnt_mutex)+1);
#ifdef VAR_RANGES
		logval("cnt_mutex", ((int)now.cnt_mutex));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 12: // STATE 16 - multiplex.pml:35 - [multiplex = (multiplex+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][16] = 1;
		(trpt+1)->bup.oval = ((int)now.multiplex);
		now.multiplex = (((int)now.multiplex)+1);
#ifdef VAR_RANGES
		logval("multiplex", ((int)now.multiplex));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 13: // STATE 17 - multiplex.pml:36 - [printf('Tecnico %d salio. En sala: %d/%d\\n',_pid,en_sala,3)] (0:0:0 - 1)
		IfNotBlocked
		reached[0][17] = 1;
		Printf("Tecnico %d salio. En sala: %d/%d\n", ((int)((P0 *)_this)->_pid), ((int)now.en_sala), 3);
		_m = 3; goto P999; /* 0 */
	case 14: // STATE 18 - multiplex.pml:37 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][18] = 1;
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

