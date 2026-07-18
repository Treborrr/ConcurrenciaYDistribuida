	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC Proceso_B */

	case 3: // STATE 2
		;
		now.sem = trpt->bup.oval;
		;
		goto R999;
;
		;
		;
		;
		
	case 6: // STATE 6
		;
		p_restor(II);
		;
		;
		goto R999;

		 /* PROC Proceso_A */
;
		;
		
	case 8: // STATE 2
		;
		now.aTermino = trpt->bup.oval;
		;
		goto R999;

	case 9: // STATE 3
		;
		now.sem = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 11: // STATE 5
		;
		p_restor(II);
		;
		;
		goto R999;
	}

