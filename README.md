# 概要
最近在网上看到一篇[有趣的文章](https://medium.com/@jayaganesh1997/let-us-sort-25e41a4ba854)。
说的是在归并排序中，可以将每个子数组的排序工作交给n个携程，再用n个携程将这些子数归并来降低时间复杂度。
据说“可以将运行时间显著降为 O(N)。”

这种说法是否可行呢？

带着疑问我做了下面的实验。

# 实验
## 电脑配置
* MacBook Pro (13-inch, 2016, Four Thunderbolt 3 Ports)
* 2.9 GHz Intel Core i5
* 8 GB 2133 MHz LPDDR3

## 算法比较
1. 快速排序（QuickSort）
2. 归并排序（MergeSortS）
3. 并行归并排序（用6个携程）（MergeSortSSC）
4. 并行归并排序（用n个携程）（MergeSortJ）

## Benchmark

![CPU](https://raw.githubusercontent.com/shabaoma/Golang-Parallel-Merge-Sort/master/result/cpu.png)
![MEM](https://raw.githubusercontent.com/shabaoma/Golang-Parallel-Merge-Sort/master/result/mem.png)

### 计算时间
数组过小时 MergeSortJ > MergeSortSSC > QuickSort ≒ MergeSortS  
数组够大时 MergeSortJ > QuickSort ≒ MergeSortS > MergeSortSSC

### 占用内存
数组过小时 MergeSortJ > MergeSortSSC > MergeSortS > QuickSort  
数组够大时 MergeSortJ > MergeSortSSC ≒ MergeSortS > QuickSort

# 分析
从结果上看，文章中所提到的方法并没有实现O(n)的计算时间，反而是最耗时的一种方法。  
通过pprof我们可以看到MergeSortJ在每一次拆分和合并切片时都要创建一个携程，从而耗费了大量时间和内存。
```
      50ms       50ms      7:func Sort(arr []int, ch chan int) {
      60ms       90ms      8:	defer close(ch)
         .          .      9:	if len(arr) == 1 {
         .      1.92s     10:		ch <- arr[0]
         .      100ms     11:		return
         .          .     12:	}
         .          .     13:
         .          .     14:	mid := len(arr) >> 1
      20ms      3.41s     15:	s1 := make(chan int, mid)
         .      3.84s     16:	s2 := make(chan int, len(arr)-mid)
         .          .     17:
         .          .     18:	// Concurrency established
      10ms      800ms     19:	go Sort(arr[:mid], s1)
         .      790ms     20:	go Sort(arr[mid:], s2)
         .          .     21:	// The sorting of arr[mid:] & arr[:mid] occurs Concurrently now.
         .          .     22:
         .          .     23:	// Merging happens simultaneously and is not blocked on individual sorting.
         .        18s     24:	merge(s1, s2, ch)
         .      170ms     25:}

```

MergeSortS和MergeSortSSC的内存占用量几乎没有区别。都是用在了生成切片上，MergeSortSSC比MergeSortS仅多生成了6个携程。  
但是MergeSortSSC的执行效率在数组足够长的情况下比普通快速排序和归并排序快了1.5倍。
```
      10ms       10ms     11:func merge(s1, s2 []int) []int {
         .      610ms     12:	s := make([]int, len(s1)+len(s2))
         .          .     13:	i, j, k := 0, 0, 0
     110ms      110ms     14:	for i < len(s1) && j < len(s2) {
      10ms       10ms     15:		if s1[i] < s2[j] {
     210ms      210ms     16:			s[k] = s1[i]
      80ms       80ms     17:			i++
         .          .     18:		} else {
     250ms      250ms     19:			s[k] = s2[j]
      70ms       70ms     20:			j++
         .          .     21:		}
         .          .     22:		k++
         .          .     23:	}
      10ms       10ms     24:	for i < len(s1) {
      10ms       10ms     25:		s[k] = s1[i]
         .          .     26:		i++
         .          .     27:		k++
         .          .     28:	}
         .          .     29:	for j < len(s2) {
         .          .     30:		s[k] = s2[j]
         .          .     31:		j++
         .          .     32:		k++
         .          .     33:	}
      10ms       10ms     34:	return s
         .          .     35:}
```

快速排序并没有消耗很多内存。生成初始数组用了O(n)的空间。  
快速排序的时间主要用在了sort.doPivot函数上。
```
         .       30ms    183:func quickSort(data Interface, a, b, maxDepth int) {
         .          .    184:	for b-a > 12 { // Use ShellSort for slices <= 12 elements
         .          .    185:		if maxDepth == 0 {
         .          .    186:			heapSort(data, a, b)
         .          .    187:			return
         .          .    188:		}
         .          .    189:		maxDepth--
         .      1.96s    190:		mlo, mhi := doPivot(data, a, b)
         .          .    191:		// Avoiding recursion on the larger subproblem guarantees
         .          .    192:		// a stack depth of at most lg(b-a).
         .          .    193:		if mlo-a < b-mhi {
         .      1.44s    194:			quickSort(data, a, mlo, maxDepth)
         .          .    195:			a = mhi // i.e., quickSort(data, mhi, b)
         .          .    196:		} else {
         .      1.62s    197:			quickSort(data, mhi, b, maxDepth)
         .          .    198:			b = mlo // i.e., quickSort(data, a, mlo)
         .          .    199:		}
         .          .    200:	}
         .          .    201:	if b-a > 1 {
         .          .    202:		// Do ShellSort pass with gap 6
         .          .    203:		// It could be written in this simplified form cause b-a <= 12
         .          .    204:		for i := a + 6; i < b; i++ {
         .          .    205:			if data.Less(i, i-6) {
         .          .    206:				data.Swap(i, i-6)
         .          .    207:			}
         .          .    208:		}
         .       10ms    209:		insertionSort(data, a, b)
         .          .    210:	}
         .          .    211:}
```

# 结论
文章所说的计算时间并不成立。因为创建携程需要消耗时间和内存。  
如果数组不是很大时，并行归并排序并不能提高排序效率。  
而当数组足够大时，同时排序子数组可以提高整体的排序速度。但是归并排序需要足够的内存。  
快速排序在相同排序效率下用内存空间最少。但是很难通过并行的方法加快排序效率。

# 参考
[A Comparison of Parallel Sorting Algorithms on Different Architectures](https://parasol.tamu.edu/publications/download.php?file_id=191)
