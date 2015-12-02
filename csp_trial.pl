#!/usr/bin/env swipl -g main -t halt

:- use_module(library(optparse)).

:- consult('./csp/magic.pl').
:- consult('./trials.pl').

main :-
	OptsSpec = [
		[opt(gen), shortflags([g]), type(atom)],
		[opt(order), shortflags([o]), type(integer)],
		[opt(dif), shortflags([d]), type(atom)],
		[opt(count), shortflags([n]), type(integer)]
	],
	opt_arguments(OptsSpec, Opts, _),
	member(gen(Gen), Opts),
	member(order(Order), Opts),
	member(dif(Dif), Opts),
	(member(count(Count), Opts); Count=10),
	nonvar(Gen),
	nonvar(Order),
	nonvar(Dif),
	statistics(walltime, _),
	forall(between(0, Count, _), (
		trial(Order, Gen, Dif, Schema),
		magic(Schema, Order)
	)),
	statistics(walltime, [_, Walltime]),
	Duration is Walltime / 1000 / Count,
	write(Duration),
	nl.
