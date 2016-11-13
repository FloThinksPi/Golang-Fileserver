package Utils

// 0 = All , 3 = Only Errors
const VERBOSITY = 0

/*
Error 	means that the execution of some task could not be completed
	an email couldn't be sent, a page couldn't be rendered, some data couldn't be stored to a database,
	something like that. Something has definitively gone wrong.

Warning means that something unexpected happened, but that execution can continue, perhaps in a degraded mode
	a configuration file was missing but defaults were used, a price was calculated as negative,
	so it was clamped to zero, etc. Something is not right, but it hasn't gone properly wrong yet
	- warnings are often a sign that there will be an error very soon.

Info 	means that something normal but significant happened
	the system started, the system stopped, the daily inventory update job ran, etc.
	There shouldn't be a continual torrent of these, otherwise there's just too much to read.

Debug 	means that something normal and insignificant happened
	new user came to the site, a page was rendered,	an order was taken, a price was updated.
	This is the stuff excluded from info because there would be too much of it.
*/