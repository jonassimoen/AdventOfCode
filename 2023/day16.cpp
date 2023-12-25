#include <string>
#include <vector>
#include <stack>
#include <set>
#include <input.h>
using namespace std;

namespace Day16 {
	constexpr util::Point<int64_t> UP{ 0, -1 };
	constexpr util::Point<int64_t> DOWN{ 0, 1 };
	constexpr util::Point<int64_t> LEFT{ -1, 0 };
	constexpr util::Point<int64_t> RIGHT{ 1, 0 };

	util::Grid<char> parseGrid(const vector<string>& lines)
	{
		util::Grid<char> grid{ static_cast<int>(lines[0].size()), static_cast<int>(lines.size()) };
		for (int y = 0; y < lines.size(); ++y)
		{
			for (int x = 0; x < lines[y].size(); ++x)
			{
				grid.at(x, y) = lines[y][x];
			}
		}
		return grid;
	}

	struct RayEntry {
		util::Point<int64_t> currentPosition;
		util::Point<int64_t> direction;
		auto operator<=>(const RayEntry&) const = default;
	};

	int64_t calcEnergizedTiles(const util::Grid<char>& grid, const RayEntry& starting = { {-1,0}, RIGHT })
	{
		util::Grid<int> energyGrid{ grid.getWidth(), grid.getHeight() };
		ranges::fill(energyGrid, 0);

		stack<RayEntry> s = {};
		// keep track of which cell we've visited so we don't get stuck in a loop (same location, same direction)
		set<RayEntry> visitedRayEntry = {};
		s.push(starting);
		int rounds = 0;
		while (!s.empty())
		{
			RayEntry& e = s.top();
			s.pop();

			if (!visitedRayEntry.contains(e)) {
				visitedRayEntry.insert(e);
				util::Point<int64_t> nextPosition = e.currentPosition + e.direction;
				if (grid.inBounds(nextPosition))
				{
					energyGrid.at(nextPosition) = 1;

					char currentChar = grid.at(nextPosition);
					bool hitSplitterPointyEdge = ((currentChar == '|') && (e.direction == UP || e.direction == DOWN)) ||
						((currentChar == '-') && (e.direction == LEFT || e.direction == RIGHT));
					bool hitSplitterFlatEdge = ((currentChar == '|') && (e.direction == LEFT || e.direction == RIGHT)) ||
						((currentChar == '-') && (e.direction == UP || e.direction == DOWN));

					if ((currentChar == '.') || hitSplitterPointyEdge)
					{
						s.push({ nextPosition, e.direction });
					}
					else if (hitSplitterFlatEdge) {
						if (currentChar == '|') {
							s.push({ nextPosition, UP });
							s.push({ nextPosition, DOWN });
						}
						else if (currentChar == '-') {
							s.push({ nextPosition, LEFT });
							s.push({ nextPosition, RIGHT });
						}

					}
					else if (currentChar == '\\') {
						if (e.direction == DOWN)
							s.push({ nextPosition, RIGHT });
						else if (e.direction == UP)
							s.push({ nextPosition, LEFT });
						else if (e.direction == RIGHT)
							s.push({ nextPosition, DOWN });
						else if (e.direction == LEFT)
							s.push({ nextPosition, UP });
					}
					else if (currentChar == '/') {
						if (e.direction == DOWN)
							s.push({ nextPosition, LEFT });
						else if (e.direction == UP)
							s.push({ nextPosition, RIGHT });
						else if (e.direction == RIGHT)
							s.push({ nextPosition, UP });
						else if (e.direction == LEFT)
							s.push({ nextPosition, DOWN });
					}
				}
				//cout << endl;
			}

		}

		return accumulate(energyGrid.begin(), energyGrid.end(), 0);
	}

	int64_t part1(const vector<string>& lines)
	{
		util::Grid<char> grid = parseGrid(lines);
		return calcEnergizedTiles(grid);
	}

	int64_t part2(const vector<string>& lines)
	{
		util::Grid<char> grid = parseGrid(lines);
		vector<int> energizedTiles = {};
		for (int x = 0; x < grid.getWidth(); ++x) {
			energizedTiles.emplace_back(calcEnergizedTiles(grid, { {x, -1}, DOWN }));
			energizedTiles.emplace_back(calcEnergizedTiles(grid, { {x, grid.getHeight()}, UP }));
		}

		for (int y = 0; y< grid.getHeight(); ++y) {
			energizedTiles.emplace_back(calcEnergizedTiles(grid, { {-1, y}, RIGHT }));
			energizedTiles.emplace_back(calcEnergizedTiles(grid, { {grid.getWidth(), y}, LEFT}));
		}
		return *max_element(energizedTiles.begin(), energizedTiles.end());
	}
}
