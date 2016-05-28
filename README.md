piper
=====

[![Build Status](https://travis-ci.org/horthy/piper.png)](https://travis-ci.org/horthy/piper)


Send database stats to graphite.


Get Started
-----------

Lets say we have a PostgreSQL table `users` containing user info,
and we'd like to send a time-series count of total users to a graphite stat like `users.count`.

We create a config file named `piper_users.yml`, with contents:

```yaml
# piper_users.yml
---
source:
  driver: postgres
  user: postgres
  database: postgres
  password: postgres
  host: localhost
  port: 5432
graphite:
  host: localhost
  port: 2003
pipes:
  - query: SELECT count(*) FROM users
    dest: users.count
```


This example assumes both postgres and graphite are running on localhost,
but connection details should be updated as appropriate.

To ship the results of the query to graphite, build the module and run `piper`,
specifying our new config file:

```sh
go get github.com/horthy/piper
piper -f piper_users.yml
```

Output should look something like

```
2016/05/24 20:23:04 Loaded config from piper_users.yml
2016/05/24 20:23:04 Sent 11     ->   users.count
```


Configuration
-------------

Piper uses a yaml file for configuration.

At the top, configuration consists of:

- `source` -- details for connecting to a database
- `graphite` -- details for connecting to graphite
- `pipes` -- list of `PipeConfig`

Each `PipeConfig` contains

- `dest` -- destination stat in graphite
- `query` -- query to use to extract a single statistic from a database

You can specify any number of query/destination combinations in a single
yaml file, they will all be executed by a call to piper.

You can specify a configuration file with `piper -f <path to file>`.


TODO
----


Would like to add:

- Support for other databases like mysql, cass, etc
- Support for other time-series data stores like OpenTSDB, CloudWatch, InfluxDB, etc.
- Parallelize shipping multiple stats from a single yaml file
- Cache connections to graphite / postgres / etc


License
---------

MIT
