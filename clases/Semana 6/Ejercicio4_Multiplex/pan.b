	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC Tecnico */

	case 3: // STATE 2
		;
		now.multiplex = trpt->bup.oval;
		;
		goto R999;

	case 4: // STATE 5
		;
		now.cnt_mutex = trpt->bup.oval;
		;
		goto R999;

	case 5: // STATE 7
		;
		now.en_sala = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 7: // STATE 9
		;
		now.cnt_mutex = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 9: // STATE 12
		;
		now.cnt_mutex = trpt->bup.oval;
		;
		goto R999;

	case 10: // STATE 14
		;
		now.en_sala = trpt->bup.oval;
		;
		goto R999;

	case 11: // STATE 15
		;
		now.cnt_mutex = trpt->bup.oval;
		;
		goto R999;

	case 12: // STATE 16
		;
		now.multiplex = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 14: // STATE 18
		;
		p_restor(II);
		;
		;
		goto R999;
	}

