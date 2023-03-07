# Workload B
Description: Mostly read - 5% Update, 95% read

We are using the properties
- operation count = 5M
- record count = 5M
- data integrity check true
- thread count = 10

```shell
cd ../..
./bin/go-ycsb load memtable -P workloads/workloadb -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true
```

```shell
cd ../..
./bin/go-ycsb run memtable -P workloads/workloadb -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true
```