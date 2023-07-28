# Golang Benchmark Tests


## Sample Results


| Package | 1                      | 2                      | 3                      | 4                       | 5                      |
|---------|------------------------|------------------------|------------------------|-------------------------|------------------------|
| glog    | 526244   - 2329 ns/op  | 8479018  - 144.2 ns/op | 8221225 - 149.6 ns/op  | 8437315  - 137.7 ns/op  | 8019350 - 133.3 ns/op  |
| logrus  | 433173   - 2740 ns/op  | 430206   - 2776 ns/op  | 434234  - 2735 ns/op   | 435856   -  2736 ns/op  | 433822  - 2718 ns/op   |
| zerolog | 14175481 - 84.53 ns/op | 512898   - 2304 ns/op  | 509218  - 2335 ns/op   | 521764   -  2312 ns/op  | 525550 -  2316 ns/op   |
| zap     | 7420369  - 161.7 ns/op | 14149322 - 83.98 ns/op | 13737259 - 85.55 ns/op | 13894208 -  85.04 ns/op | 13850374 - 85.41 ns/op |
| slog    | 132628   - 8381 ns/op  | 136209   - 8307 ns/op  | 138687 -   7256 ns/op  | 149724   -  7046 ns/op  | 124480 -   8832 ns/op  |