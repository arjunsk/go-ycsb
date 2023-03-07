# Workload A
Description: Scan the latest insert (90% read, 10% insert)

We are using the properties
- operation count = 5M
- record count = 5M
- data integrity check true
- thread count = 10

```shell
cd ../..
./bin/go-ycsb load memtable -P workloads/workloada -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true
```

```shell
cd ../..
./bin/go-ycsb run memtable -P workloads/workloada -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true
```