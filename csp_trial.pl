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
		[opt(count), shortflags([n]), type(integer)],
		[opt(timeout), shortflags([t]), type(integer)]
	], Opts, _),
	(member(gen(Gen), Opts);     Gen = siamese1),
	(member(order(Order), Opts); Order = 3),
	(member(dif(Dif), Opts);     Dif = easy),
	(member(count(Count), Opts); Count=10),
	(member(timeout(Timeout), Opts); Timeout=30),

	% run the experiment
	message_queue_create(Queue),
	findall(Time, (
		between(1, Count, _),
		thread_create((
			trial(Order, Gen, Dif, Square),
			magic(Square, Order),
			statistics(cputime, Time),
			thread_send_message(Queue, Time, [timeout(Timeout)]),
			thread_exit(Square)
		), Thread, []),
		(thread_get_message(Queue, Time, [timeout(Timeout)])
		->
			thread_join(Thread, _)
		;
			write("timeout"),
			nl,
			halt
		)
	), Times),
	sum_list(Times, Total),
	Avg is Total / Count,
	write(Avg).
