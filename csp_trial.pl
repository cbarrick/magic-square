#!/usr/bin/env swipl -g main -t halt

:- use_module(library(optparse)).

:- consult('./csp/magic.pl').
:- consult('./trials.pl').

main :-
	% parse command line arguments
	opt_arguments([
		[opt(gen), shortflags([g]), type(atom)],
		[opt(order), shortflags([o]), type(integer)],
		[opt(dif), shortflags([d]), type(atom)],
		[opt(count), shortflags([n]), type(integer)]
	], Opts, _),
	(member(gen(Gen), Opts);     Gen = siamese1),
	(member(order(Order), Opts); Order = 3),
	(member(dif(Dif), Opts);     Dif = easy),
	(member(count(Count), Opts); Count=10),

	% run the experiment
	statistics(walltime, _),
	forall(between(0, Count, _), (
		trial(Order, Gen, Dif, Schema),
		magic(Schema, Order)
	)),
	statistics(walltime, [_, Walltime]),
	Duration is Walltime / 1000 / Count,
	write(Duration),
	nl.
