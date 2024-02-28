# readseekr-bench

Just a benchmark repo to compare two approaches of reading zip files.

```shell
$ benchstat old.txt new.txt                                
goos: darwin
goarch: arm64
pkg: github.com/simar7/readseekr-bench
          │   old.txt   │              new.txt               │
          │   sec/op    │   sec/op     vs base               │
Normal-10   8.700µ ± 3%   8.311µ ± 5%  -4.47% (p=0.009 n=10)

          │   old.txt    │               new.txt                │
          │     B/op     │     B/op      vs base                │
Normal-10   9.039Ki ± 0%   6.256Ki ± 0%  -30.79% (p=0.000 n=10)

          │  old.txt   │              new.txt               │
          │ allocs/op  │ allocs/op   vs base                │
Normal-10   25.00 ± 0%   21.00 ± 0%  -16.00% (p=0.000 n=10)
```
