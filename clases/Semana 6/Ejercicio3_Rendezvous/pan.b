	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC Jugador_2 */
;
		;
		
	case 4: // STATE 2
		;
		now.j2_llego_control = trpt->bup.oval;
		;
		goto R999;

	case 5: // STATE 3
		;
		now.j2Llego = trpt->bup.oval;
		;
		goto R999;

	case 6: // STATE 5
		;
		now.j1Llego = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 8: // STATE 8
		;
		j2_en_fase2 = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 10: // STATE 10
		;
		p_restor(II);
		;
		;
		goto R999;

		 /* PROC Jugador_1 */
;
		;
		
	case 12: // STATE 2
		;
		now.j1_llego_control = trpt->bup.oval;
		;
		goto R999;

	case 13: // STATE 3
		;
		now.j1Llego = trpt->bup.oval;
		;
		goto R999;

	case 14: // STATE 5
		;
		now.j2Llego = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 16: // STATE 8
		;
		j1_en_fase2 = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 18: // STATE 10
		;
		p_restor(II);
		;
		;
		goto R999;
	}

