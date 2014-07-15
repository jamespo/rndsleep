rndsleep
========

A command-line utility written in Go to sleep for a random period, then run another command line executable.

Typical usage would be to run puppet agent in onetime mode in cron to save resources - the random sleep will help to avoid thundering herd issues at the puppetmaster.

`rndsleep --command="puppet agent --no-daemonize --onetime"`

would run puppet agent with a random timeout of 0-30s (default).

Can be compiled on all platforms that Go is available for.

## Command-line arguments ##

With default arguments shown

       rndsleep
         -command="": command to run
         -port=0: localhost TCP port to lock on
         -randmax=0: maximum delay (seconds)
         -verbose=false: enable verbose output

### Specifying a port ###

Port allows you to specify a TCP port on localhost that rndsleep will listen on. Nothing will be done with input on this socket but rndsleep will fail and the command will not be run if this socket can't be opened - hence this stops duplicate invocations. This could be useful if you want to invoke a puppet run immediately (perhaps remotely with ansible) and prevent the cronned instance running.

So (1) ensure this doesn't collide with other local ports on your system (eg 3306 for mysql) and (2) if you want to use this feature choose a unique port per job executed (eg port 14355 for puppet runs, port 14356 for database maintenance jobs).

By default this argument is set to 0 and disabled.

