:- use_module(library("clpfd")).

make_square(Table, N) :-
	length(Table, N),
	make_square_(Table, N).
make_square_([], _) :- !.
make_square_([H|T], N) :-
	length(H, N),
	make_square_(T, N).

magic(Table, Sum) :-
	length(Table, N),
	Sum #= (N * (N * N + 1)) / 2,
	setup_domain(Table),
	transpose(Table, Trans),
	magic_rows(Table, Sum),
	magic_rows(Trans, Sum),
	magic_diag(Table, Sum),
	magic_diag(Trans, Sum),
	flatten(Table, Flat),
	all_distinct(Flat),
	labeling([ffc], Flat).

magic_rows([], _) :- !.
magic_rows([Row|Rows], Sum) :-
	sum(Row, #=, Sum),
	magic_rows(Rows, Sum).

magic_diag(Table, Sum) :-
	extract_diag(Table, 1, Diag),
	sum(Diag, #=, Sum).

extract_diag([], _, []) :- !.
extract_diag([Row|Table], N, [H|Diag]) :-
	element(N, Row, H),
	N1 #= N + 1,
	extract_diag(Table, N1, Diag).

setup_domain(Table) :-
	length(Table, D),
	Max #= D * D,
	setup_domain_(Table, Max).

setup_domain_([], _) :- !.
setup_domain_([Row|Rows], Max) :-
	Row ins 1..Max,
	setup_domain_(Rows, Max).

flatten([], []) :- !.
flatten([[]|Rows], Flat) :- !,
	flatten(Rows, Flat).
flatten([[H|T]|Rows], [H|Flat]) :-
	flatten([T|Rows], Flat).
