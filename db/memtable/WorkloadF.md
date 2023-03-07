# Workload F
Description: Read Modify Write - 50% read, 50% update

We are using the properties
- operation count = 5M
- record count = 5M
- data integrity check true
- thread count = 10

```shell
cd ../..
./bin/go-ycsb load memtable -P workloads/workloadf -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10
./bin/go-ycsb run memtable -P workloads/workloadf -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10
```

```shell
cd ../..
```