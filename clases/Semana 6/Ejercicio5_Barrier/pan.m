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

		 /* PROC :init: */
	case 3: // STATE 1 - barrier.pml:63 - [(run Trabajador(1))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][1] = 1;
		if (!(addproc(II, 1, 0, 1)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 4: // STATE 2 - barrier.pml:64 - [(run Trabajador(2))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][2] = 1;
		if (!(addproc(II, 1, 0, 2)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 3 - barrier.pml:65 - [(run Trabajador(3))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][3] = 1;
		if (!(addproc(II, 1, 0, 3)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 5 - barrier.pml:67 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[1][5] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */

		 /* PROC Trabajador */
	case 7: // STATE 1 - barrier.pml:22 - [((f<2))] (5:0:0 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		if (!((((int)((P0 *)_this)->f)<2)))
			continue;
		/* merge: printf('Trabajador %d: completando Fase %d\\n',id,(f+1))(0, 2, 5) */
		reached[0][2] = 1;
		Printf("Trabajador %d: completando Fase %d\n", ((int)((P0 *)_this)->id), (((int)((P0 *)_this)->f)+1));
		_m = 3; goto P999; /* 1 */
	case 8: // STATE 3 - barrier.pml:26 - [((mut>0))] (6:0:1 - 1)
		IfNotBlocked
		reached[0][3] = 1;
		if (!((((int)now.mut)>0)))
			continue;
		/* merge: mut = (mut-1)(0, 4, 6) */
		reached[0][4] = 1;
		(trpt+1)->bup.oval = ((int)now.mut);
		now.mut = (((int)now.mut)-1);
#ifdef VAR_RANGES
		logval("mut", ((int)now.mut));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 9: // STATE 6 - barrier.pml:27 - [count = (count+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][6] = 1;
		(trpt+1)->bup.oval = ((int)now.count);
		now.count = (((int)now.count)+1);
#ifdef VAR_RANGES
		logval("count", ((int)now.count));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 10: // STATE 7 - barrier.pml:29 - [((count==3))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		if (!((((int)now.count)==3)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 11: // STATE 8 - barrier.pml:30 - [((molinete2>0))] (11:0:1 - 1)
		IfNotBlocked
		reached[0][8] = 1;
		if (!((((int)now.molinete2)>0)))
			continue;
		/* merge: molinete2 = (molinete2-1)(0, 9, 11) */
		reached[0][9] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete2);
		now.molinete2 = (((int)now.molinete2)-1);
#ifdef VAR_RANGES
		logval("molinete2", ((int)now.molinete2));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 12: // STATE 11 - barrier.pml:31 - [molinete1 = (molinete1+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][11] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete1);
		now.molinete1 = (((int)now.molinete1)+1);
#ifdef VAR_RANGES
		logval("molinete1", ((int)now.molinete1));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 13: // STATE 16 - barrier.pml:34 - [mut = (mut+1)] (0:0:1 - 3)
		IfNotBlocked
		reached[0][16] = 1;
		(trpt+1)->bup.oval = ((int)now.mut);
		now.mut = (((int)now.mut)+1);
#ifdef VAR_RANGES
		logval("mut", ((int)now.mut));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 14: // STATE 17 - barrier.pml:37 - [((molinete1>0))] (20:0:1 - 1)
		IfNotBlocked
		reached[0][17] = 1;
		if (!((((int)now.molinete1)>0)))
			continue;
		/* merge: molinete1 = (molinete1-1)(0, 18, 20) */
		reached[0][18] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete1);
		now.molinete1 = (((int)now.molinete1)-1);
#ifdef VAR_RANGES
		logval("molinete1", ((int)now.molinete1));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 15: // STATE 20 - barrier.pml:38 - [molinete1 = (molinete1+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][20] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete1);
		now.molinete1 = (((int)now.molinete1)+1);
#ifdef VAR_RANGES
		logval("molinete1", ((int)now.molinete1));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 16: // STATE 21 - barrier.pml:41 - [((mut>0))] (24:0:1 - 1)
		IfNotBlocked
		reached[0][21] = 1;
		if (!((((int)now.mut)>0)))
			continue;
		/* merge: mut = (mut-1)(0, 22, 24) */
		reached[0][22] = 1;
		(trpt+1)->bup.oval = ((int)now.mut);
		now.mut = (((int)now.mut)-1);
#ifdef VAR_RANGES
		logval("mut", ((int)now.mut));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 17: // STATE 24 - barrier.pml:42 - [count = (count-1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][24] = 1;
		(trpt+1)->bup.oval = ((int)now.count);
		now.count = (((int)now.count)-1);
#ifdef VAR_RANGES
		logval("count", ((int)now.count));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 18: // STATE 25 - barrier.pml:44 - [((count==0))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][25] = 1;
		if (!((((int)now.count)==0)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 19: // STATE 26 - barrier.pml:45 - [((molinete1>0))] (29:0:1 - 1)
		IfNotBlocked
		reached[0][26] = 1;
		if (!((((int)now.molinete1)>0)))
			continue;
		/* merge: molinete1 = (molinete1-1)(0, 27, 29) */
		reached[0][27] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete1);
		now.molinete1 = (((int)now.molinete1)-1);
#ifdef VAR_RANGES
		logval("molinete1", ((int)now.molinete1));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 20: // STATE 29 - barrier.pml:46 - [molinete2 = (molinete2+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][29] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete2);
		now.molinete2 = (((int)now.molinete2)+1);
#ifdef VAR_RANGES
		logval("molinete2", ((int)now.molinete2));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 21: // STATE 34 - barrier.pml:49 - [mut = (mut+1)] (0:0:1 - 3)
		IfNotBlocked
		reached[0][34] = 1;
		(trpt+1)->bup.oval = ((int)now.mut);
		now.mut = (((int)now.mut)+1);
#ifdef VAR_RANGES
		logval("mut", ((int)now.mut));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 22: // STATE 35 - barrier.pml:52 - [((molinete2>0))] (38:0:1 - 1)
		IfNotBlocked
		reached[0][35] = 1;
		if (!((((int)now.molinete2)>0)))
			continue;
		/* merge: molinete2 = (molinete2-1)(0, 36, 38) */
		reached[0][36] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete2);
		now.molinete2 = (((int)now.molinete2)-1);
#ifdef VAR_RANGES
		logval("molinete2", ((int)now.molinete2));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 23: // STATE 38 - barrier.pml:53 - [molinete2 = (molinete2+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][38] = 1;
		(trpt+1)->bup.oval = ((int)now.molinete2);
		now.molinete2 = (((int)now.molinete2)+1);
#ifdef VAR_RANGES
		logval("molinete2", ((int)now.molinete2));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 24: // STATE 39 - barrier.pml:55 - [printf('Trabajador %d: paso barrera de Fase %d\\n',id,(f+1))] (0:43:1 - 1)
		IfNotBlocked
		reached[0][39] = 1;
		Printf("Trabajador %d: paso barrera de Fase %d\n", ((int)((P0 *)_this)->id), (((int)((P0 *)_this)->f)+1));
		/* merge: f = (f+1)(43, 40, 43) */
		reached[0][40] = 1;
		(trpt+1)->bup.oval = ((int)((P0 *)_this)->f);
		((P0 *)_this)->f = (((int)((P0 *)_this)->f)+1);
#ifdef VAR_RANGES
		logval("Trabajador:f", ((int)((P0 *)_this)->f));
#endif
		;
		/* merge: .(goto)(0, 44, 43) */
		reached[0][44] = 1;
		;
		_m = 3; goto P999; /* 2 */
	case 25: // STATE 46 - barrier.pml:59 - [-end-] (0:0:0 - 3)
		IfNotBlocked
		reached[0][46] = 1;
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

