package leetcode

type UnionFindSet struct {
	fathers []int
}

func MakeUnionFindSet(n int) *UnionFindSet {
	s := &UnionFindSet{fathers: make([]int, n)}
	for i := 0; i < n; i++ {
		s.fathers[i] = -1
	}
	return s
}

func (s *UnionFindSet) SetFather(x int) {
	s.fathers[x] = x
}

func (s *UnionFindSet) Find(x int) int {
	if s.fathers[x] != x {
		s.fathers[x] = s.Find(s.fathers[x])
	}
	return s.fathers[x]
}

func (s *UnionFindSet) Union(x, y int) {
	x = s.Find(x)
	y = s.Find(y)
	if x != y {
		s.fathers[y] = x
	}
}

func (s *UnionFindSet) Judge(x, y int) bool {
	x = s.Find(x)
	y = s.Find(y)
	return x == y
}

func (s *UnionFindSet) UnionCount() int {
	cnt := 0
	for i := 0; i < len(s.fathers); i++ {
		if s.fathers[i] == i {
			cnt++
		}
	}
	return cnt
}

func (s *UnionFindSet) UnionSet() [][]int {
	m := make(map[int][]int)
	for i := 0; i < len(s.fathers); i++ {
		f := s.Find(i)
		m[f] = append(m[f], i)
	}
	sets := make([][]int, 0, len(m))
	for _, v := range m {
		sets = append(sets, v)
	}
	return sets
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n <= 0 {
		return 0
	}
	m := len(grid[0])
	s := MakeUnionFindSet(n * m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				s.SetFather(i*m + j)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				if i-1 >= 0 {
					if grid[i-1][j] == '1' {
						x := (i-1)*m + j
						y := i*m + j
						s.Union(x, y)
					}
				}
				if i+1 < n {
					if grid[i+1][j] == '1' {
						x := (i+1)*m + j
						y := i*m + j
						s.Union(x, y)
					}
				}
				if j-1 >= 0 {
					if grid[i][j-1] == '1' {
						x := i*m + (j - 1)
						y := i*m + j
						s.Union(x, y)
					}
				}
				if j+1 < m {
					if grid[i][j+1] == '1' {
						x := i*m + (j + 1)
						y := i*m + j
						s.Union(x, y)
					}
				}
			}
		}
	}
	return s.UnionCount()
}
