# PSQLCompare

PSQLCompare is a golang-based nagios check which works like [Bucardo's check_postgresql](https://bucardo.org/wiki/Check_postgres) check and will take 2
or more (PostgreSQL) strings containing a pair of a connection string and a query to run on for example:

    "postgres://user:password@hostOrIP:Port/DBname#Query"

which can be something like this

    "postgres://myuser:secretpassword@localhost:5432/myapp#SELECT id FROM randomtable WHERE columnx = 23"

## Installation
Just clone the repo and built it with

    go build check_psqlcompare.go

Alternatively you can also use the latest binaries pre-built for different OS and architectures using [gox](https://github.com/mitchellh/gox).
You can find them in the Releases tab.

## Usage
    ./check_psqlcompare --compare "<querystring1>" "<querystring2>" "<querystring...>"

It will split up the querystrings for the connection and query parts and get the output of the query parts in a new slice. After that
it compares all slice values with the first querystring output as a reference. If it finds differences in the output values the check
will exit CRITICAL and tell you which querystring does not match the reference. Otherwise the check will return OK.

## Known issues

The check does not validate neither the connection string nor the query.

## Contributing
This is an open source project and your contribution is very much appreciated.

1. Check for open issues or open a fresh issue to start a discussion around a feature idea or a bug.
2. Fork the repository on Github and make your changes on the **develop** branch (or branch off of it).
3. Send a pull request (with the **develop** branch as the target).


## Changelog
See [CHANGELOG.md](changelog.md)

## License
PSQLCompare is available under the MIT license. See the [LICENSE](LICENSE) file for more info.