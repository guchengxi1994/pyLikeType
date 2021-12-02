/*
	following https://github.com/smbanaie/floyd-warshall-python
		and   https://blog.csdn.net/yuewenyao/article/details/81021319
*/

package pyLikeType

import (
	"errors"
	"fmt"
)

var INF = 999999

type Route struct {
	From     int
	To       int
	Distance int
	Path     []int
}

type Node struct {
	Index    int
	Next     int
	Distance int
}

func (r *Route) ToString() string {

	if r.From == r.To {
		return fmt.Sprintf("%v -> %v, is the same", r.From, r.To)
	}

	if len(r.Path) == 0 {
		return fmt.Sprintf("%v -> %v, has no path", r.From, r.To)
	}

	pathStr := ""
	for i, v := range r.Path {
		if i == 0 {
			pathStr = fmt.Sprint(v)
		} else {
			pathStr += " ---|> " + fmt.Sprint(v)
		}
	}

	return fmt.Sprintf("%v -> %v, distance: %v, path: %s", r.From, r.To, r.Distance, pathStr)
}

func ArrayToNode(array []int) (Node, error) {
	node := Node{}
	if len(array) != 3 {
		return node, errors.New("array length error")
	}
	node.Distance = array[2]
	node.Index = array[0]
	node.Next = array[1]
	return node, nil
}

func FloydWarshall(n int, edge []Node) {
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	nxt := make([][]int, n)
	for i := range dist {
		nxt[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			nxt[i][j] = 0
		}
	}

	for _, e := range edge {
		dist[e.Index-1][e.Next-1] = e.Distance
		nxt[e.Index-1][e.Next-1] = e.Next - 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					nxt[i][j] = nxt[i][k]
				}

			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			route := Route{}
			route.From = i + 1
			route.To = j + 1
			if i != j {
				var path []int
				path = append(path, i)
				for path[len(path)-1] != j {
					path = append(path, nxt[path[len(path)-1]][j])
				}

				for _i, _v := range path {
					path[_i] = _v + 1
				}
				// fmt.Printf("path: %v\n", path)

				route.Distance = dist[i][j]
				route.Path = path
			}
			fmt.Printf("route.ToString(): %v\n", route.ToString())
		}
	}
}
