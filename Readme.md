# GoSwift - Embedded cache for golang

High-performance, concurrent embedded caching library for Go applications with support for Hash data type

## Features

- Set & Get command
- Del command
- Update command
- Exists command
- Support for TTL
- Support for Disk Save(Snapshots)
- Support Hash Data type Hset, Hget, HgetAll, HMset
- Safe Locking

## Installation

```shell
go mod init github.com/my/repo
```

Then install goswift:

```shell
go get github.com/leoantony72/goswift
```

## Quickstart

```go

package main

import (
    "fmt"
    "github.com/leoantony72/goswift"
)

func main(){
    cache := goswift.NewCache()

    // Value 0 indicates no expiry
    cache.Set("key", "value", 0)

    val, err := cache.Get("key")
    if err !=nil{
        fmt.Println(err)
        return
    }
    fmt.Println("key", val)
}

```

## Disk Save

### Snapshot

```go
opt := goswift.CacheOptions{
		EnableSnapshots:  true,
		SnapshotInterval: time.Second*5,
	}
c := goswift.NewCache(opt)
```

> **_NOTE:_** If the **EnableSnapshot** is **false**, Data saved in the file will not imported

This will take a snapshot of the Data Every 5sec and saves it into a **_Snapshot.data_** file. By default Snapshots are disabled and if the SnapshotInterval is not provided default value is **5seconds**.

> **_NOTE:_** Don't delete the **_Snapshot.data_** File <br>

## Error Handling

```go
const (
	ErrKeyNotFound   = "key does not Exists"
	ErrFieldNotFound = "field does not Exists"
	ErrNotHashvalue  = "not a Hash value/table"
	ErrHmsetDataType = "invalid data type, Expected Struct/Map"
)
```

These are the common Errors that may occur while writing the code. These Varible provide you a clear and easy **Error** comparison method to determine errors.

```go
data,err := cache.Get("key")
if err != nil {
	if err.Error() == goswift.ErrKeyNotFound {
        //do something
}
}
```

## Benchmarks

The following benchmarks demonstrate the performance of different operations in GoSwift:

```
cpu: AMD EPYC 7763 64-Core Processor
BenchmarkSet-2                  11319936                97.45 ns/op           40 B/op          2 allocs/op
BenchmarkSetWithExpiry-2         3381864               485.6 ns/op           113 B/op          3 allocs/op
BenchmarkGet-2                  75574640                19.07 ns/op            0 B/op          0 allocs/op
BenchmarkExists-2               78714792                19.15 ns/op            0 B/op          0 allocs/op
BenchmarkDel-2                   6834908               216.0 ns/op            40 B/op          2 allocs/op
BenchmarkUpdate-2               10504995               170.5 ns/op            40 B/op          2 allocs/op
BenchmarkHset-2                 11203929               123.6 ns/op            16 B/op          1 allocs/op
BenchmarkHGet-2                 28004112                47.77 ns/op            0 B/op          0 allocs/op
BenchmarkHGetAll-2              79354870                18.93 ns/op            0 B/op          0 allocs/op
BenchmarkHMset-2                 1000000              1030 ns/op             409 B/op          4 allocs/op
BenchmarkAllData-2               2179728               499.2 ns/op           336 B/op          2 allocs/op
```

Key observations from the benchmarks:

- Get/Exists operations are extremely fast (~19 ns/op) with zero allocations
- Basic Set operations are efficient (~97 ns/op)
- Hash operations show good performance, especially HGetAll (~19 ns/op)
- Set with expiry has higher overhead due to TTL management
- HMset has the highest overhead due to handling multiple fields

To run the benchmarks yourself:

```bash
go test -bench=. ./...
```

## Usage

```go
// Set Value with Expiry
// @Set(key string, val interface{}, exp int)
// Here expiry is set to 1sec
cache.Set("key","value",1000)


// Get Value with key
// @Get(key string) (interface{}, error)
val,err := cache.Get("key")
if err != nil{
    fmt.Println(err)
    return
}


// Update value
// @Update(key string, val interface{}) error
err = cache.Update("key","value2")
if err != nil{
    fmt.Println(err)
    return
}


// Delete command
// @Del(key string)
cache.Del("key")


// Hset command
// @Hset(key, field string, value interface{}, exp int)
// in this case the "key" expires in 1sec
cache.Hset("key","name","value",1000)
cache.Hset("key","age",18,1000)


// HMset command
// @HMset(key string, d interface{}, exp int) error
// Set a Hash by passing a Struct/Map
// ---by passing a struct---
type Person struct{
    Name  string
    Age   int
    Place string
}

person1 := &Person{Name:"bob",Age:18,Place:"NYC"}
err = cache.HMset("key",person1)
if err != nil{
    fmt.Println(err)
    return
}

// ---by passing a map---
person2 := map[string]interface{Name:"john",Age:18,Place:"NYC"}
err = cache.HMset("key",person2)
if err != nil{
    fmt.Println(err)
    return
}


// Hget command
// @HGet(key, field string) (interface{}, error)
// get individual fields in Hash
data,err := cache.HGet("key","field")
if err != nil{
    fmt.Println(err)
    return
}
fmt.Println(data)

// HgetAll command
// @HGetAll(key string) (map[string]interface{}, error)
// gets all the fields with value in a hash key
// retuns a map[string]interface{}
data,err = cache.HGetAll("key")
if err != nil{
    fmt.Println(err)
    return
}


// Exist command
// @Exists(key string) bool
// Check if the key exists
value = cache.Exists("key")
fmt.Println(value)



// AllData command
// @AllData() (map[string]interface{}, int)
// returns all the data in the cache with keys, also with no.of keys present
// returns the value as a map[strirng]interface{}
// !It does not return the expiry time of the key
data,counter := cache.AllData()
fmt.Println(data,counter)

```

## Run the Test

```go
go test ./...
```
