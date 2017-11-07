mdocker
==============

<pre><code>
[root@mhc beegoTest]# ./mdocker
mdocker is a Fast and Flexible tool for managing your DB cluster.

USAGE
    mdocker [command] [subcommand] [arguments]

AVAILABLE COMMANDS

  o cluster              cluster operation

  o snapshot             no short description

  o database             no short description

  o operator             no short description

  o log                  no short description

  o ops                  no short description


Use mdocker help for more information about all commands.

Use mdocker [command] help for more information about a command.

Use mdocker [command] [subcommand] help for more information about a subcommand.

Required env:
    export APPID=test
    export APPSCERET=ee1d7336-73e3-452b-9e63-fdeaf2dccde6
Optional env:
    export OPSSERVER=http://127.0.0.1:8181
    export MGMTSERVER=http://127.0.0.1:8080
------------------------------------------------------------------------------------
[root@mhc beegoTest]# ./mdocker help
mdocker is a Fast and Flexible tool for managing your DB cluster.

USAGE
    mdocker [command] [subcommand] [arguments]

AVAILABLE COMMANDS

  o cluster                         cluster operation

        o info                      get cluster info
        o start                     start cluster
        o stop                      stop cluster
        o restart                   restart cluster
        o metrics                   get metrics

  o snapshot                        no short description

        o start                     start backup
        o status                    get backup status
        o start-scheduled-backup    start scheduled backup
        o stop-scheduled-backup     stop scheduled backup
        o list                      list all backup
        o restore                   start restore
        o delete-backup             delete a backup
        o delete-all                delete all backup

  o database                        no short description

        o create                    create a database
        o delete                    delete a database from master
        o create-with-user          create a database and a bind user for it
        o delete-with-user          delete a database and it's bind user
        o get-with-user             get all database and it's bind user

  o operator                        no short description

        o add                       add a user
        o delete                    delete a user by name

  o log                             no short description

        o cluster                   get cluster logs
        o node                      get a node log by containerId

  o ops                             no short description

        o create                    create a cluster
        o scale                     scale a cluster
        o upgrade                   upgrade a service
        o delete                    delete cluster
        o packages                  get packages from triton
        o resize                    resize a service


Use mdocker [command] help for more information about a command.

Use mdocker [command] [subcommand] help for more information about a subcommand.

Required env:
    export APPID=test
    export APPSCERET=ee1d7336-73e3-452b-9e63-fdeaf2dccde6
Optional env:
    export OPSSERVER=http://127.0.0.1:8181
    export MGMTSERVER=http://127.0.0.1:8080

----------------------------------------------------------------------------------------------
[root@mhc beegoTest]# ./mdocker cluster help
DESCRIPTION
  cluster operation

USAGE
    mdocker cluster [subcommand] [arguments]

AVAILABLE SUBCOMMANDS:

     o info                      get cluster info
     o start                     start cluster
     o stop                      stop cluster
     o restart                   restart cluster
     o metrics                   get metrics

Use mdocker cluster  [subcommand] help for more information about a subcommand.
---------------------------------------------------------------------------------------------
[root@mhc beegoTest]# ./mdocker cluster metrics help
DESCRIPTION
  get metrics

USAGE
  mdocker cluster metrics [-keyword='kw1&kw2'] [-ip=***] [-container-id=***]

OPTIONS
  -container-id
      Specify a containerID to query.

  -ip
      Specify a ip to query.

  -keyword
      Specify a keyword to filter metrics.



</code></pre>