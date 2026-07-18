	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC :init: */

	case 3: // STATE 1
		;
		;
		delproc(0, now._nr_pr-1);
		;
		goto R999;

	case 4: // STATE 2
		;
		;
		delproc(0, now._nr_pr-1);
		;
		goto R999;

	case 5: // STATE 3
		;
		;
		delproc(0, now._nr_pr-1);
		;
		goto R999;

	case 6: // STATE 5
		;
		p_restor(II);
		;
		;
		goto R999;

		 /* PROC Trabajador */
;
		
	case 7: // STATE 1
		goto R999;

	case 8: // STATE 4
		;
		now.mut = trpt->bup.oval;
		;
		goto R999;

	case 9: // STATE 6
		;
		now.count = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 11: // STATE 9
		;
		now.molinete2 = trpt->bup.oval;
		;
		goto R999;

	case 12: // STATE 11
		;
		now.molinete1 = trpt->bup.oval;
		;
		goto R999;

	case 13: // STATE 16
		;
		now.mut = trpt->bup.oval;
		;
		goto R999;

	case 14: // STATE 18
		;
		now.molinete1 = trpt->bup.oval;
		;
		goto R999;

	case 15: // STATE 20
		;
		now.molinete1 = trpt->bup.oval;
		;
		goto R999;

	case 16: // STATE 22
		;
		now.mut = trpt->bup.oval;
		;
		goto R999;

	case 17: // STATE 24
		;
		now.count = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 19: // STATE 27
		;
		now.molinete1 = trpt->bup.oval;
		;
		goto R999;

	case 20: // STATE 29
		;
		now.molinete2 = trpt->bup.oval;
		;
		goto R999;

	case 21: // STATE 34
		;
		now.mut = trpt->bup.oval;
		;
		goto R999;

	case 22: // STATE 36
		;
		now.molinete2 = trpt->bup.oval;
		;
		goto R999;

	case 23: // STATE 38
		;
		now.molinete2 = trpt->bup.oval;
		;
		goto R999;

	case 24: // STATE 40
		;
		((P0 *)_this)->f = trpt->bup.oval;
		;
		goto R999;

	case 25: // STATE 46
		;
		p_restor(II);
		;
		;
		goto R999;
	}

