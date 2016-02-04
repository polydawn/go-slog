go-slog
=======

A streamish log, for when things really slog along.

Bunch of slow processes?  In parallel?
You need some status streams and progress bars so this doesn't feel like a slow slog.

Designed to play nicely with your terminal.
Doesn't screw with your scrollback, take over the full window,
or use any of the heavy weaponry for stateful terminal modification used in applications like e.g. vim.
Full-screen terminal apps are fine; this just isn't one of them.
It's optimized to fit into your flow instead.
You get updating, animated status info at the bottom, and regular logs trailing up like plain old "\n"-delimited logs should.


