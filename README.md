# json-vs-gob

Mac
```shell
go test -bench=./... -test.benchmem
```


Windows
```shell
go test -bench . -benchmem
```

Results:
```
go test -bench=./... -test.benchmem
goos: darwin
goarch: arm64
pkg: json_vs_gob
BenchmarkEncodeMap/algo=Gob:key=10-8      884458              1346 ns/op            1264 B/op         36 allocs/op
BenchmarkEncodeMap/algo=JSON:key=10-8     853940              1385 ns/op             640 B/op         22 allocs/op
BenchmarkEncodeMap/algo=Gob:key=50-8      333781              3623 ns/op            2528 B/op        119 allocs/op
BenchmarkEncodeMap/algo=JSON:key=50-8     136562              8680 ns/op            3233 B/op        102 allocs/op
BenchmarkEncodeMap/algo=Gob:key=100-8     175381              6751 ns/op            4112 B/op        220 allocs/op
BenchmarkEncodeMap/algo=JSON:key=100-8             63613             18792 ns/op            6594 B/op        202 allocs/op
BenchmarkEncodeMap/algo=Gob:key=1000-8             19288             61955 ns/op           48592 B/op       2027 allocs/op
BenchmarkEncodeMap/algo=JSON:key=1000-8             4657            257062 ns/op           74431 B/op       2902 allocs/op
BenchmarkDecodeMap/algo=Gob:key=10-8              117212              8912 ns/op            7084 B/op        177 allocs/op
BenchmarkDecodeMap/algo=JSON:key=10-8             568738              2070 ns/op             731 B/op         17 allocs/op
BenchmarkDecodeMap/algo=Gob:key=50-8              102765             11600 ns/op            8335 B/op        180 allocs/op
BenchmarkDecodeMap/algo=JSON:key=50-8             111764             10719 ns/op            3155 B/op         65 allocs/op
BenchmarkDecodeMap/algo=Gob:key=100-8              79242             15172 ns/op           10032 B/op        181 allocs/op
BenchmarkDecodeMap/algo=JSON:key=100-8             56529             21398 ns/op            6382 B/op        120 allocs/op
BenchmarkDecodeMap/algo=Gob:key=1000-8             17832             67383 ns/op           54513 B/op        178 allocs/op
BenchmarkDecodeMap/algo=JSON:key=1000-8             5457            221765 ns/op           90766 B/op       1028 allocs/op
BenchmarkEncodeSliceMap/algo=Gob:key=10:len=10-8                  168139              7130 ns/op            4104 B/op        222 allocs/op
BenchmarkEncodeSliceMap/algo=JSON:key=10:len=10-8                  85863             13707 ns/op            6426 B/op        212 allocs/op
BenchmarkEncodeSliceMap/algo=Gob:key=10:len=1000-8                  1987            599874 ns/op          356421 B/op      20036 allocs/op
BenchmarkEncodeSliceMap/algo=JSON:key=10:len=1000-8                  868           1381298 ns/op          642155 B/op      21004 allocs/op
BenchmarkEncodeSliceMap/algo=Gob:key=100:len=10-8                   1984            608119 ns/op          421931 B/op      20036 allocs/op
BenchmarkEncodeSliceMap/algo=JSON:key=100:len=10-8                   622           1894042 ns/op          652193 B/op      20102 allocs/op
BenchmarkEncodeSliceMap/algo=Gob:key=1000:len=10-8                    18          59859426 ns/op        55961732 B/op    2000111 allocs/op
BenchmarkEncodeSliceMap/algo=JSON:key=1000:len=10-8                    4         256632688 ns/op        82345686 B/op    2901265 allocs/op
BenchmarkDecodeSliceMap/algo=Gob:key=10:len=10-8                   72087             16608 ns/op           12319 B/op        243 allocs/op
BenchmarkDecodeSliceMap/algo=JSON:key=10:len=10-8                  58644             20473 ns/op            6158 B/op        149 allocs/op
BenchmarkDecodeSliceMap/algo=Gob:key=10:len=1000-8                  1530            780583 ns/op          540021 B/op       6260 allocs/op
BenchmarkDecodeSliceMap/algo=JSON:key=10:len=1000-8                  601           1998121 ns/op          589524 B/op      14091 allocs/op
BenchmarkDecodeSliceMap/algo=Gob:key=100:len=10-8                   1692            704198 ns/op          357017 B/op       1191 allocs/op
BenchmarkDecodeSliceMap/algo=JSON:key=100:len=10-8                   566           2137873 ns/op          624335 B/op      11754 allocs/op
BenchmarkDecodeSliceMap/algo=Gob:key=1000:len=10-8                    19          59695877 ns/op        47892279 B/op       7183 allocs/op
BenchmarkDecodeSliceMap/algo=JSON:key=1000:len=10-8                    5         219293992 ns/op        90623942 B/op    1025813 allocs/op
BenchmarkEncodeSingleStruct/algo=Gob:struct_size=small-8         1000000              1010 ns/op            1200 B/op         21 allocs/op
BenchmarkEncodeSingleStruct/algo=JSON:struct_size=small-8               12506340                95.83 ns/op           24 B/op          1 allocs/op
BenchmarkEncodeSingleStruct/algo=Gob:struct_size=medium-8                 381122              3152 ns/op            2248 B/op         47 allocs/op
BenchmarkEncodeSingleStruct/algo=JSON:struct_size=medium-8               1962462               616.9 ns/op           288 B/op          3 allocs/op
BenchmarkEncodeSingleStruct/algo=Gob:struct_size=big-8                    221665              5385 ns/op            3288 B/op         71 allocs/op
BenchmarkEncodeSingleStruct/algo=JSON:struct_size=big-8                   522942              2287 ns/op            1152 B/op          9 allocs/op
BenchmarkDecodeSingleStruct/algo=Gob:struct_size=small-8                  143022              8281 ns/op            6744 B/op        178 allocs/op
BenchmarkDecodeSingleStruct/algo=JSON:struct_size=small-8                2921467               411.7 ns/op           248 B/op          6 allocs/op
BenchmarkDecodeSingleStruct/algo=Gob:struct_size=medium-8                  99013             12032 ns/op            9116 B/op        245 allocs/op
BenchmarkDecodeSingleStruct/algo=JSON:struct_size=medium-8                651886              1856 ns/op             448 B/op         11 allocs/op
BenchmarkDecodeSingleStruct/algo=Gob:struct_size=big-8                     69404             20420 ns/op           12165 B/op        323 allocs/op
BenchmarkDecodeSingleStruct/algo=JSON:struct_size=big-8                   169856              7771 ns/op            1024 B/op         21 allocs/op
BenchmarkEncodeSliceStruct/algo=Gob:struct_size=small:len=100-8           216348              5526 ns/op            5488 B/op         28 allocs/op
BenchmarkEncodeSliceStruct/algo=JSON:struct_size=small:len=100-8          246421              4893 ns/op            2689 B/op          1 allocs/op
BenchmarkEncodeSliceStruct/algo=Gob:struct_size=medium:len=100-8           66609             18380 ns/op           13432 B/op         55 allocs/op
BenchmarkEncodeSliceStruct/algo=JSON:struct_size=medium:len=100-8          21175             56455 ns/op           28691 B/op        201 allocs/op
BenchmarkEncodeSliceStruct/algo=Gob:struct_size=big:len=100-8              20096             60154 ns/op           44808 B/op         83 allocs/op
BenchmarkEncodeSliceStruct/algo=JSON:struct_size=big:len=100-8              5294            226120 ns/op          120404 B/op        801 allocs/op
BenchmarkDecodeSliceStruct/algo=Gob:struct_size=small:len=100-8            84081             14441 ns/op           11548 B/op        296 allocs/op
BenchmarkDecodeSliceStruct/algo=JSON:struct_size=small:len=100-8           37430             32321 ns/op            8184 B/op        114 allocs/op
BenchmarkDecodeSliceStruct/algo=Gob:struct_size=medium:len=100-8           41070             29365 ns/op           30245 B/op        660 allocs/op
BenchmarkDecodeSliceStruct/algo=JSON:struct_size=medium:len=100-8           7052            171479 ns/op           43064 B/op        416 allocs/op
BenchmarkDecodeSliceStruct/algo=Gob:struct_size=big:len=100-8              16440             73522 ns/op           85350 B/op       1531 allocs/op
BenchmarkDecodeSliceStruct/algo=JSON:struct_size=big:len=100-8              1748            685140 ns/op          135401 B/op       1217 allocs/op
BenchmarkEncodeSliceStructPointer/algo=Gob:struct_size=small:len=100-8            198888              6049 ns/op            5488 B/op         28 allocs/op
BenchmarkEncodeSliceStructPointer/algo=JSON:struct_size=small:len=100-8           216258              5972 ns/op            2689 B/op          1 allocs/op
BenchmarkEncodeSliceStructPointer/algo=Gob:struct_size=medium:len=100-8            62888             19008 ns/op           13432 B/op         55 allocs/op
BenchmarkEncodeSliceStructPointer/algo=JSON:struct_size=medium:len=100-8           20928             57497 ns/op           28695 B/op        201 allocs/op
BenchmarkEncodeSliceStructPointer/algo=Gob:struct_size=big:len=100-8               19639             60542 ns/op           44808 B/op         83 allocs/op
BenchmarkEncodeSliceStructPointer/algo=JSON:struct_size=big:len=100-8               5274            227377 ns/op          120454 B/op        801 allocs/op
BenchmarkDecodeSliceStructPointer/algo=Gob:struct_size=small:len=100-8             64630             18512 ns/op           12156 B/op        396 allocs/op
BenchmarkDecodeSliceStructPointer/algo=JSON:struct_size=small:len=100-8            33022             34665 ns/op            5224 B/op        214 allocs/op
BenchmarkDecodeSliceStructPointer/algo=Gob:struct_size=medium:len=100-8            34756             34455 ns/op           31206 B/op        760 allocs/op
BenchmarkDecodeSliceStructPointer/algo=JSON:struct_size=medium:len=100-8            6996            173398 ns/op           20096 B/op        516 allocs/op
BenchmarkDecodeSliceStructPointer/algo=Gob:struct_size=big:len=100-8               15123             79818 ns/op           86887 B/op       1631 allocs/op
BenchmarkDecodeSliceStructPointer/algo=JSON:struct_size=big:len=100-8               1755            682780 ns/op           68160 B/op       1317 allocs/op
PASS
ok      json_vs_gob     97.866s
```