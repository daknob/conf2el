# conf2el

`conf2el` is a service that uses `ssh` to connect to remote hosts, and then run
arbitrary commands, such as to retrieve a configuration file, and then after it
fetches the command output, it pushes the file to
[eldim](https://github.com/daknob/eldim).

The main usecase is for equipment where an `eldim` client cannot run natively,
such as for example network equipment (routers, switches, etc.). By configuring
an SSH account that can login with either a username and a password or an SSH
key and has read-only access to the configuration file (or a normal file that
it can `cat`) it can export and store information securely.

Moreover, for compliance or audit reasons, equipment state at a given point in
time may be necessary. By running the necessary commands (such as showing the
firewall rules, interface statistics, installed packages and versions, ...) and
then storing them in `eldim`, these requirements can be addressed.
