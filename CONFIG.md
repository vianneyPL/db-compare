# Json config
## qdb-benchmark:
A path providing a specific version of the qdb-benchmark tool.
If left empty, db-compare will automatically download the last version of the software.

## databases:
An array of databases names that will be tested.
To get a list of available databases you can run `./db-compare -list-databases`

## tests:
An array of tests names that will run.
To get a list of available tests you can run `./db-compare -list-tests`

## tests-config:
###    threads:
    An array describing the number of threads running simultaneously against the cluster.
###    sizes:
    An array describing the size of the elements.
    Ex: A size of `1M` (bytes) on blob_put means each blob will have a size of 1MB
###    packs:
    An array describing the size of a pack.
    Ex: a packs value of `100` on batch_blob_put means each batch will have 100 blob_put
###    pause:
    A number describing the pause between two tests.
###    duration:
    A number describing the duration of a test.

## clusters:
An array of description of clusters.
###    location:
    The location of the cluster.
###    system:
    The operating system of the cluster
###    nodes:
    An array describing the number of nodes per cluster
###    threads:
    An array describing the number of threads per node
To get a list of available location and system you can run `./db-compare -list-clusters`


# Example:
```
{
    "qdb-benchmark": "",
    "databases": [
        "qdb"
    ],
    "tests": [
        "blob_put",
        "blob_get",
        "blob_update"
    ],
    "tests-config": {
        "threads": [
            1,
            2,
            4,
            8,
            16
        ],
        "sizes": [
            "8",
            "2k",
            "64k",
            "2M"
        ],
        "packs": [
            "10",
            "100"
        ],
        "pause": "20",
        "duration": "20"
    },
    "clusters": [
        {
            "location": "belgium",
            "system": "linux",
            "nodes": [
                1,
                3,
                5
            ],
            "threads": [
                8
            ]
        }
    ],
    "transient": true
}```