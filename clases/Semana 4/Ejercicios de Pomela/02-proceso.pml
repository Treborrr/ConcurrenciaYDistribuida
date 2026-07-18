proctype child()
{
	printf("child %d\n", _pid)
}

active proctype parent()
{
	do
	:: (_nr_pr == 1) ->
		run child()
	od
}