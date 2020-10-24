# topk
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/topk)](https://pkg.go.dev/github.com/hslam/topk)
[![Build Status](https://travis-ci.org/hslam/topk.svg?branch=master)](https://travis-ci.org/hslam/topk)
[![codecov](https://codecov.io/gh/hslam/topk/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/topk)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/topk)](https://goreportcard.com/report/github.com/hslam/topk)
[![LICENSE](https://img.shields.io/github/license/hslam/topk.svg?style=flat-square)](https://github.com/hslam/topk/blob/master/LICENSE)

Package topk implements finding the top k elements in the collection.

## Feature
* Heap Sort
* Top K
* Sort Top K

## Get started

### Install
```
go get github.com/hslam/topk
```
### Import
```
import "github.com/hslam/topk"
```
### Usage
#### Example
```
package main

import (
	"fmt"
	"github.com/hslam/topk"
)

type list []int

func (h list) Len() int           { return len(h) }
func (h list) Less(i, j int) bool { return h[i] < h[j] }
func (h list) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func main() {
	l := list{5, 10, 2, 7, 1}
	k := 3
	topk.Top(l, 3, false)
	fmt.Println(l[:k])
	topk.Top(l, 3, true)
	fmt.Println(l[:k])
}
```

#### Output
```
[5 10 7]
[10 7 5]
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)


### Author
topk was written by Meng Huang.


