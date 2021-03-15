## Coding challenge: go-permutations
Grouping permutations of integers with a common sum. That is given an array of integers, compute all the permutations with a given count and group all the permutations that have a common sum.

### Requirements

1. Go SDK 1.11 or greater

### Sample run
```sh
# input integer array is: []int{2, 3, 5, 7, 9, 10, 1000}
# combination length is 2 integers
# output column 1 is the sum
# output column 2 is the array of combinations that sum up to the result in column 1

▶ go run main.go
Sorted results keys: [4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 1002 1003 1005 1007 1009 1010 2000]
4   [[2 2]]
5   [[2 3] [3 2]]
6   [[3 3]]
7   [[2 5] [5 2]]
8   [[3 5] [5 3]]
9   [[2 7] [7 2]]
10   [[3 7] [5 5] [7 3]]
11   [[2 9] [9 2]]
12   [[2 10] [3 9] [5 7] [7 5] [9 3] [10 2]]
13   [[3 10] [10 3]]
14   [[5 9] [7 7] [9 5]]
15   [[5 10] [10 5]]
16   [[7 9] [9 7]]
17   [[7 10] [10 7]]
18   [[9 9]]
19   [[9 10] [10 9]]
20   [[10 10]]
1002   [[2 1000] [1000 2]]
1003   [[3 1000] [1000 3]]
1005   [[5 1000] [1000 5]]
1007   [[7 1000] [1000 7]]
1009   [[9 1000] [1000 9]]
1010   [[10 1000] [1000 10]]
2000   [[1000 1000]]
```
### Running unit tests

```sh
▶ go test -v

=== RUN   TestNoInputValue
--- PASS: TestNoInputValue (0.00s)
=== RUN   TestUnSortedArray
--- PASS: TestUnSortedArray (0.00s)
=== RUN   TestSortedWithSmallAndLargeNumbers
--- PASS: TestSortedWithSmallAndLargeNumbers (0.00s)
=== RUN   TestBadResult
    main_test.go:19: sum [10] result contains an invalid combination [[1 2]]
    main_test.go:19: sum [10] result contains an invalid combination [[3 4]]
--- PASS: TestBadResult (0.00s)
=== RUN   TestGoodResult
--- PASS: TestGoodResult (0.00s)
=== RUN   TestBadResultMost7
    main_test.go:19: sum [7] result contains an invalid combination [[1 8]]
--- PASS: TestBadResultMost7 (0.00s)
PASS
ok  	github.com/ltenfield/go-permutations	0.015s
```