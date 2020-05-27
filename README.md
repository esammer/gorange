# gorange

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/esammer/gorange?label=latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/esammer/gorange/Build)](https://github.com/esammer/gorange/actions?query=workflow%3ABuild)

A continuous value range library for Go.

This library is suitable for representing continuous ranges of values and performing common operations on those
ranges including order comparison, merging, testing if a value is within a range, and so on. Originally it was
developed for representing range-encoded values in a column store. This library has no dependencies (beyond its
test suite) beyond Go 1.13. Earlier versions of Go may work, but aren't tested.

## Usage

Add gorange to your project.

    go get github.com/esammer/gorange

Create a RangeValue implementation.

Or use one of the prebuilt implementations for standard Go types.

    package mypkg
    
    import gr "github.com/esammer/gorange"
    
    type MyValue struct {
        // ....
    }
    
    // LessThan determines whether one value is less than another. If two values are mutually less than each other
    // (e.g. !a.LessThan(b) && !b.LessThan(a)), they are considered equal.
    func (v *MyValue) LessThan(other gr.RangeValue) bool {
        // Some way of determining if v is less than other.
        return false
    }

Perform range operations.

    r1 := Range{
        Begin: IntValue(0),
        End:   IntValue(10),
    }
    r2 := Range{
        Begin: IntValue(5),
        End:   IntValue(15),
    }
    
    r1.LessThan(r2)             // true
    r2.LessThan(r1)             // false
    r3 := r1.Merge(r2)          // r3: Range{Begin: IntValue(0), End: IntValue(15)}
    
    r1.Contains(IntValue(3))    // true
    r1.Intersects(r2)           // true

## Performance

Range is similar to time.Time in that you should almost always use values rather than pointers to values. A
range contains only two interface members. None of the methods on Range allocate memory on the heap. As of
sha 902c17a, the included benchmarks are as follows.

    esammer@C02C86Q6MD6R range % go test -bench '.*' -benchmem       
    goos: darwin
    goarch: amd64
    pkg: github.com/esammer/gorange
    BenchmarkRange/LessThan-16        423571209      2.76 ns/op      0 B/op      0 allocs/op
    BenchmarkRange/Before-16          422689849      2.77 ns/op      0 B/op      0 allocs/op
    BenchmarkRange/After-16           432240738      2.80 ns/op      0 B/op      0 allocs/op
    BenchmarkRange/Intersects-16      171612853      6.89 ns/op      0 B/op      0 allocs/op
    BenchmarkRange/Merge-16           146388944      8.22 ns/op      0 B/op      0 allocs/op
    BenchmarkRange/Contains-16        100000000     11.0 ns/op       0 B/op      0 allocs/op

on a MBP 16" with the following specs:

    Model Name:                 MacBook Pro
    Model Identifier:           MacBookPro16,1
    Processor Name:             8-Core Intel Core i9
    Processor Speed:            2.4 GHz
    Number of Processors:       1
    Total Number of Cores:      8
    L2 Cache (per Core):        256 KB
    L3 Cache:                   16 MB
    Hyper-Threading Technology: Enabled
    Memory:                     32 GB

    System Version:             macOS 10.15.2 (19C57)
    Kernel Version:             Darwin 19.2.0

## Documentation

You can view the gorange API docs at [pkg.go.dev/github.com/esammer/gorange](https://pkg.go.dev/github.com/esammer/gorange).

## Issues

Feel free to file Github issues if you find a bug. PRs are welcome.

## License

This software is licensed under the Apache License 2.0.

Copyright 2020, Eric Sammer.
