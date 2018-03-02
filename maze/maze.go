package main

import (
	"os"
	"fmt"
)

// 读入迷宫数据
func readMaze(filename string) [][]int {

	if file, e := os.Open(filename); e == nil {

		var row, col int
		fmt.Fscanf(file, "%d %d", &row, &col)

		maze := make([][]int, row)
		for i := range maze {
			maze[i] = make([]int, col)
			for j := range maze[i] {
				fmt.Fscanf(file, "%d", &maze[i][j])
			}
		}

		return maze
	} else {
		panic(e)
	}
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{i: p.i + r.i, j: p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i >= 0 && p.i < len(grid) && p.j >= 0 && p.j < len(grid[0]) {
		return grid[p.i][p.j], true
	} else {
		return -1, false
	}
}

// 四个方向
var dirs = [4]point{
	// 上		  左             下           右
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 广度优先走迷宫算法
func walk(maze [][]int, start point, end point) [][]int {

	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 队列
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 只有同时满足maze at next = 0 and steps at next = 0 & next != start，才会去处理它

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			stepVal, _ := cur.at(steps)
			steps[next.i][next.j] = stepVal + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	step, _ := end.at(steps)
	fmt.Printf("共需要多少步：%d\n", step)

	// 从后面遍历就可以找到路径
}
