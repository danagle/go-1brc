# CPU Profile Results

(pprof) top
Showing nodes accounting for 79.19s, 57.01% of 138.90s total
Dropped 275 nodes (cum <= 0.69s)
Showing top 10 nodes out of 57
      flat  flat%   sum%        cum   cum%
    18.13s 13.05% 13.05%     18.25s 13.14%  strconv.readFloat
     9.45s  6.80% 19.86%     30.59s 22.02%  runtime.mapaccess1_faststr
     8.38s  6.03% 25.89%    133.80s 96.33%  github.com/danagle/go-1brc/v01.Run
     7.77s  5.59% 31.48%      7.77s  5.59%  indexbytebody
     6.87s  4.95% 36.43%      6.87s  4.95%  aeshashbody
     6.35s  4.57% 41.00%      6.35s  4.57%  strconv.atof32exact
     6.13s  4.41% 45.41%     12.04s  8.67%  runtime.mallocgcTiny
     5.51s  3.97% 49.38%      5.51s  3.97%  internal/runtime/syscall.Syscall6
     5.34s  3.84% 53.23%     31.84s 22.92%  strconv.atof32
     5.26s  3.79% 57.01%     20.30s 14.61%  bufio.(*Scanner).Scan
