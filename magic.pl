:- use_module(library("clpfd")).

%% square(?Table, ?N)
% True when Table is an NxN square matrix.
square(Table, N) :-
	length(Table, N),
	square_(Table, N).
square_([], _).
square_([H|T], N) :-
	length(H, N),
	square_(T, N).

%% magic(?Table, ?N)
% True when Table is an NxN magic-square.
magic(Table, N) :-
	N in 3..sup,
	Sum #= (N * (N * N + 1)) / 2,
	Max #= N*N,
	square(Table, N),
	transpose(Table, Trans),
	magic_rows(Table, Sum),
	magic_rows(Trans, Sum),
	magic_diag(Table, Sum),
	magic_diag(Trans, Sum),
	append(Table, Flat),
	all_distinct(Flat),
	Flat ins 1..Max,
	labeling([ff], Flat).

%% magic_rows(+Table, ?Sum)
% Applies the constraint that every row of Table sums to Sum.
magic_rows([], _).
magic_rows([Row|Rows], Sum) :-
	sum(Row, #=, Sum),
	magic_rows(Rows, Sum).

%% magic_diag(+Table, ?Sum)
% Applies the constraint that the main diagonal of Table sums to Sum.
magic_diag(Table, Sum) :-
	magic_diag_(Table, 1, Diag),
	sum(Diag, #=, Sum).
magic_diag_([], _, []).
magic_diag_([Row|Table], N, [H|Diag]) :-
	element(N, Row, H),
	N1 #= N + 1,
	magic_diag_(Table, N1, Diag).
