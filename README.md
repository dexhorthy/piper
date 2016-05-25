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
- source:
    user: postgres
    database: postgres
    password: postgres
    host: localhost
    port: 5432
  query: SELECT count(*) FROM users
  dest: users.count
  graphite:
    host: localhost
    port: 2003
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
2016/05/24 20:23:04 Executing query: SELECT count(*) FROM users
2016/05/24 20:23:04 Loaded value: 11
2016/05/24 20:23:04 Connected to Graphite at localhost:2003
2016/05/24 20:23:04 sent 11 to users.count
```


Configuration
-------------

Piper uses a yaml file for configuration.
Contents are a list of `PipeConfig`s.

At the top level each `PipeConfig` has four fields:

- `source` -- details for connecting to a database
- `query` -- query to use to extract a single statistic from a database
- `dest` -- destination path in graphite
- `graphite` -- details for connecting to graphite


You can specify any number of source/query/destination/graphite combinations in a single
yaml file, they will all be executed by a call to piper.


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
