# topk
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/topk)](https://pkg.go.dev/github.com/hslam/topk)
[![Build Status](https://github.com/hslam/topk/workflows/build/badge.svg)](https://github.com/hslam/topk/actions)
[![codecov](https://codecov.io/gh/hslam/topk/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/topk)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/topk)](https://goreportcard.com/report/github.com/hslam/topk)
[![LICENSE](https://img.shields.io/github/license/hslam/topk.svg?style=flat-square)](https://github.com/hslam/topk/blob/master/LICENSE)

Package topk finds the top k elements in the collection.

## Feature
* Min Heap
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
	Top(list{5, 10, 2, 7, 1}, 3, false)
	Top(list{5, 10, 2, 7, 1}, 3, true)
}

func Top(l list, k int, sortK bool) {
	fmt.Printf("list:%v,k:%d,sortK:%t\t", l, k, sortK)
	topk.Top(l, k, sortK)
	fmt.Printf("==>\ttop%d:%v\n", k, l[:k])
}
```

#### Output
```
list:[5 10 2 7 1],k:3,sortK:false	==>	top3:[5 10 7]
list:[5 10 2 7 1],k:3,sortK:true	==>	top3:[10 7 5]
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)


### Author
topk was written by Meng Huang.


