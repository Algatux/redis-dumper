# redis-dumper
Wanna dump some values from a redis database ?

The idea behind this redis-dumper will provide the extracted values in a format that can be easily dumped back directly using the redis-cli!

Cmd output (dump)

    HSET alga.algatux.dev RDS_HOST "database.algatux.dev"
    HSET alga.algatux.dev RDS_USER "algatux"
    HSET alga.algatux.dev RDS_REPLICA "replica.algatux.dev"
    HSET alga.algatux.dev RDS_PASSWORD "********"

to be placed inside a dump file

Restore

    cat dump.dat | redis-cli -h localhost -n 0

