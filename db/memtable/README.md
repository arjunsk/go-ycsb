# GO-YCSB Benchmarking commands

## Workload E (Uniform Distribution)
Description: Scan the latest insert (90% read, 10% insert)

We are using the properties
- operation count = 5M
- record count = 5M
- scan width distribution = uniform distribution between 1 to 1000
- data integrity check true
- thread count = 10
- request distribution = `uniform`
```shell
cd ../..
./bin/go-ycsb load memtable -P workloads/workloade -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true -p minscanlength=1 -p maxscanlength=1000 -p scanlengthdistribution=uniform
```

```shell
cd ../..
./bin/go-ycsb run memtable -P workloads/workloade -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true -p minscanlength=1 -p maxscanlength=1000 -p scanlengthdistribution=uniform
```

## Workload E (Sequential Distribution)
Description: Scan the latest insert (90% read, 10% insert)

We are using the properties
- operation count = 5M
- record count = 5M
- scan width distribution = uniform distribution between 1 to 1000
- data integrity check true
- thread count = 10
- request distribution = `sequential`
```shell
cd ../..
./bin/go-ycsb load memtable -P workloads/workloade -p requestdistribution=sequential -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true -p minscanlength=1 -p maxscanlength=1000 -p scanlengthdistribution=uniform
```

```shell
cd ../..
./bin/go-ycsb run memtable -P workloads/workloade -p requestdistribution=sequential -p operationcount=5000000 -p recordcount=5000000 -p threadcount=10 -p dataintegrity=true -p minscanlength=1 -p maxscanlength=1000 -p scanlengthdistribution=uniform
```



## Workload A