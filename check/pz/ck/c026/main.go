package main

import (
	"fmt"
)

func main() {
	var (
		N, S, p int
		ans     int
		m, s    int
		m_max   int
	)
	fmt.Scan(&N, &S, &p)
	m_max = 0
	for i := 0; i < N; i++ {
		fmt.Scan(&m, &s)
		if s <= S+p && s >= S-p {
			if m > m_max {
				m_max = m
				ans = i + 1
			}
		}
	}
	if ans == 0 {
		fmt.Println("not found")
	} else {
		fmt.Println(ans)
	}
}
