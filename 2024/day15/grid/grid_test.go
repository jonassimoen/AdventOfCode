package grid

import (
	"fmt"
	"testing"
)

func MainMovement(t *testing.T, g *Grid, p1, p2 Direction, double, shouldBeAbleToMove bool) {
	if double {
		g.DoubleGrid()
	}
	if g.robCurPos != p1 {
		t.Logf("Expected robCurPos to be at position %d, got %d", p1, g.robCurPos)
		t.Fail()
	}
	able := g.isAbleToMove(p1, g.robotMoves[0])
	if shouldBeAbleToMove != able {
		t.Logf("Expected isAbleToMove %t, got %t", shouldBeAbleToMove, able)
		t.Fail()
	}
	g.Print()
	fmt.Println(g.robCurPos)
	if double {
		g.StartMovingBis()
	} else {
		g.StartMoving()
	}
	g.Print()
	fmt.Println(g.robCurPos)
	if g.robCurPos != p2 {
		t.Logf("Expected robCurPos to be at position %d, got %d", p2, g.robCurPos)
		t.Fail()
	}
}

func Test_P2Left(t *testing.T) {
	ls := [][]string{
		{
			"##########",
			"#..O@....#",
			"#........#",
			"#........#",
			"##########",
			"",
			"<",
		},
		{
			"##########",
			"#.OO@....#",
			"#........#",
			"#........#",
			"##########",
			"",
			"<",
		},
		{
			"##########",
			"#OOO@....#",
			"#........#",
			"#........#",
			"##########",
			"",
			"<",
		},
	}
	ps := []Direction{{8, 1}, {8, 1}, {8, 1}}
	pe := []Direction{{7, 1}, {7, 1}, {8, 1}}
	double := true
	shouldBeAbleToMove := []bool{true, true, false}
	for idx, l := range ls {
		g := ParseInput(l)
		MainMovement(t, g, ps[idx], pe[idx], double, shouldBeAbleToMove[idx])
	}
}

func Test_P2Right(t *testing.T) {
	ls := [][]string{
		{
			"##########",
			"#....@O..#",
			"#........#",
			"#........#",
			"##########",
			"",
			">>",
		},
		{
			"##########",
			"#....@OO.#",
			"#........#",
			"#........#",
			"##########",
			"",
			">>",
		},
		{
			"##########",
			"#....@OOO#",
			"#........#",
			"#........#",
			"##########",
			"",
			">>",
		},
	}
	ps := []Direction{{10, 1}, {10, 1}, {10, 1}}
	pe := []Direction{{12, 1}, {12, 1}, {11, 1}}
	double := true
	shouldBeAbleToMove := []bool{true, true, false}
	for idx, l := range ls {
		g := ParseInput(l)
		MainMovement(t, g, ps[idx], pe[idx], double, shouldBeAbleToMove[idx])
		fmt.Println()
	}
}

func Test_P2Up(t *testing.T) {
	ls := [][]string{
		{
			"##########",
			"#.....O..#",
			"#.....@..#",
			"#........#",
			"##########",
			"",
			"^",
		},
		{
			"##########",
			"#........#",
			"#.....O..#",
			"#.....@..#",
			"##########",
			"",
			"^",
		},
		{
			"##########",
			"#....OO..#",
			"#.....O..#",
			"#.....@..#",
			"##########",
			"",
			"^",
		},
		{
			"##########",
			"#........#",
			"#....@OO.#",
			"#.....O..#",
			"#........#",
			"##########",
			"",
			">><vv>^",
		},
		{
			"##########",
			"#........#",
			"#....OO@.#",
			"#....O...#",
			"#........#",
			"##########",
			"",
			"<vv<<^",
		},
		{
			"##########",
			"#.....#..#",
			"#....OO..#",
			"#.....O..#",
			"#.....@..#",
			"##########",
			"",
			"^",
		},
		{
			"##########",
			"#.....#..#",
			"#....OO@.#",
			"#....O...#",
			"#........#",
			"##########",
			"",
			"<vv<<^",
		},
	}
	ps := []Direction{{12, 2}, {12, 3}, {12, 3}, {10, 2}, {14, 2}, {12, 4}, {14, 2}}
	pe := []Direction{{12, 2}, {12, 2}, {12, 3}, {12, 3}, {11, 3}, {12, 4}, {11, 4}}
	double := true
	shouldBeAbleToMove := []bool{false, true, false, true, true, false, true}
	for idx, l := range ls {
		fmt.Printf("Case %d\n", idx+1)
		g := ParseInput(l)
		MainMovement(t, g, ps[idx], pe[idx], double, shouldBeAbleToMove[idx])
		fmt.Println()
	}
}

func Test_P2Down(t *testing.T) {
	ls := [][]string{
		{
			"##########",
			"#.....@..#",
			"#.....O..#",
			"#........#",
			"##########",
			"",
			"v",
		},
		{
			"##########",
			"#........#",
			"#.....@..#",
			"#.....O..#",
			"##########",
			"",
			"v",
		},
		{
			"##########",
			"#....@...#",
			"#....OO..#",
			"#.....O..#",
			"##########",
			"",
			"v",
		},
		{
			"##########",
			"#........#",
			"#....@OO.#",
			"#.....O..#",
			"#........#",
			"##########",
			"",
			">>^>v",
		},
		{
			"##########",
			"#........#",
			"#....@OO.#",
			"#.....O#.#",
			"#........#",
			"##########",
			"",
			">>^>v",
		},
		{
			"##########",
			"#........#",
			"#....@O..#",
			"#.....O..#",
			"##########",
			"",
			">>^>>v",
		},
		{
			"##########",
			"#......O.#",
			"#.......O#",
			"#.....OOO#",
			"#.....O.O#",
			"#.....O.O#",
			"#........#",
			"#....@O.O#",
			"#.....O..#",
			"##########",
			"",
			">><^^^^^>>v",
		},
		{
			"##########",
			"#........#",
			"#....OOO.#",
			"#...@OO..#",
			"#...O....#",
			"##########",
			"",
			">><^^>v",
		},
	}
	ps := []Direction{{12, 1}, {12, 2}, {10, 1}, {10, 2}, {10, 2}, {10, 2}, {10, 7}, {8, 3}}
	pe := []Direction{{12, 2}, {12, 2}, {10, 2}, {13, 2}, {13, 1}, {14, 1}, {13, 3}, {10, 2}}
	double := true
	shouldBeAbleToMove := []bool{true, false, true, true, true, true, true, true}
	for idx, l := range ls {
		fmt.Printf("Case %d\n", idx+1)
		g := ParseInput(l)
		MainMovement(t, g, ps[idx], pe[idx], double, shouldBeAbleToMove[idx])
		fmt.Println()
	}
}

func TestGrid_Moves(t *testing.T) {
	ls := [][]string{
		{
			"",
			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
			"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
			"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
			"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
			"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
			"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
			">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
			"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
			"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
			"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
		},
	}

	for _, l := range ls {
		g := ParseInput(l)
		orig := ""
		for _, li := range l {
			orig = fmt.Sprintf("%s%s", orig, li)
		}
		s := ""
		for _, d := range g.robotMoves {
			s = fmt.Sprintf("%s%c", s, UnparseDirection(d))
		}
		if orig != s {
			t.Errorf("\nExpected:\n%s\nActual:\n%s", orig, s)
			t.Fail()
		}
	}
}
