# Workload D
Description: Read the latest inserted - 90% read, 10% insert

We are using the properties
- operation count = 5M
- record count = 5M
- data integrity check true
- thread count = 10

```shell
cd ../..
./bin/go-ycsb load badger -P workloads/workloadd -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true
```

```shell
cd ../..
./bin/go-ycsb run badger -P workloads/workloadd -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true
```