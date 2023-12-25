#include <string>
#include <vector>
#include <input.h>
#include <grid.h>
#include <queue>
using namespace std;

namespace Day10 {
	constexpr util::Point<int64_t> NORTH{ 0, -1 }, SOUTH{ 0,1 }, EAST{ 1,0 }, WEST{ -1,0 };

	util::Grid<char> convertToGrid(const vector<string>& lines) {
		util::Grid<char> grid{ static_cast<int>(lines[0].size()), static_cast<int>(lines.size()) };
		for (int y = 0; y < lines.size(); ++y)
		{
			const string& line = lines[y];
			for (int x = 0; x < line.size(); ++x)
			{
				grid.at(x, y) = line[x];
			}
		}
		return grid;
	}

	vector<util::Point<int64_t>> getNeighbors(const util::Grid<char>& grid, const util::Point<int64_t>& p)
	{
		vector<util::Point<int64_t>> possibleNeighbors;
		const char c = grid.at(p);

		switch (c)
		{
		case '-':
			possibleNeighbors.emplace_back(p + WEST);
			possibleNeighbors.emplace_back(p + EAST);
			break;
		case '|':
			possibleNeighbors.emplace_back(p + NORTH);
			possibleNeighbors.emplace_back(p + SOUTH);
			break;
		case 'L':
			possibleNeighbors.emplace_back(p + NORTH);
			possibleNeighbors.emplace_back(p + EAST);
			break;
		case 'J':
			possibleNeighbors.emplace_back(p + NORTH);
			possibleNeighbors.emplace_back(p + WEST);
			break;
		case 'F':
			possibleNeighbors.emplace_back(p + SOUTH);
			possibleNeighbors.emplace_back(p + EAST);
			break;
		case '7':
			possibleNeighbors.emplace_back(p + SOUTH);
			possibleNeighbors.emplace_back(p + WEST);
			break;
		default:
			break;
		}

		for (auto it = possibleNeighbors.begin(); it != possibleNeighbors.end();)
		{
			if (!grid.inBounds(*it))
			{
				it = possibleNeighbors.erase(it);
			}
			else
			{
				++it;
			}
		}


		return possibleNeighbors;
	}

	util::Point<int64_t> replaceStartingPointIteratorWithCorrectPipe(util::Grid<char>& grid)
	{
		auto startingPointIterator = ranges::find(grid, 'S');
		util::Point<int64_t> startingPosition = { static_cast<int64_t>((startingPointIterator - grid.begin()) % grid.getWidth()), static_cast<int64_t>((startingPointIterator - grid.begin()) / grid.getWidth()) };

		for (char pipePossibilities : {'|', '-', 'L', 'J', '7', 'F'})
		{
			// What if S is this pipe?
			grid.at(startingPosition) = pipePossibilities;
			vector<util::Point<int64_t>> neighbors = getNeighbors(grid, startingPosition);

			bool closedLoopWithCurrentPipe = neighbors.size() == 2;
			if (!closedLoopWithCurrentPipe)
			{
				continue;
			}
			for (util::Point<int64_t> neighbor : neighbors)
			{
				char c = grid.at(neighbor);
				switch (c)
				{
				case '|':
					closedLoopWithCurrentPipe &= (neighbor + NORTH == startingPosition || neighbor + SOUTH == startingPosition);
					break;
				case '-':
					closedLoopWithCurrentPipe &= (neighbor + EAST == startingPosition || neighbor + WEST == startingPosition);
					break;
				case 'L':
					closedLoopWithCurrentPipe &= (neighbor + NORTH == startingPosition || neighbor + EAST == startingPosition);
					break;
				case 'J':
					closedLoopWithCurrentPipe &= (neighbor + NORTH == startingPosition || neighbor + WEST == startingPosition);
					break;
				case 'F':
					closedLoopWithCurrentPipe &= (neighbor + SOUTH == startingPosition || neighbor + EAST == startingPosition);
					break;
				case '7':
					closedLoopWithCurrentPipe &= (neighbor + SOUTH == startingPosition || neighbor + WEST == startingPosition);
					break;
				case '.':
					closedLoopWithCurrentPipe = false;
				}
			}
			if (closedLoopWithCurrentPipe)
			{
				return startingPosition;
			}
		}
		return util::Point<int64_t>{-1, -1};
	}

	int64_t part1(const vector<string>& lines)
	{
		util::Grid<char> grid = convertToGrid(lines);
		util::Point<int64_t> startingPosition = replaceStartingPointIteratorWithCorrectPipe(grid);
		util::Grid<int64_t> distancesFromStart{ grid.getWidth(), grid.getHeight() };
		ranges::fill(distancesFromStart, -1);
		distancesFromStart.at(startingPosition) = 0;

		queue<pair<util::Point<int64_t>, int64_t>> q;
		for (const util::Point<int64_t>& nb : getNeighbors(grid, startingPosition))
		{
			q.emplace(nb, 1);
		}

		while (!q.empty())
		{
			auto [p, distance] = q.front();
			q.pop();

			if (distancesFromStart.at(p) != -1)
			{
				// at the end of the loop, two nodes will have same neighbor & add this one to the queue twice
				// so, check if we already visited this node
				continue;
			}
			distancesFromStart.at(p) = distance;
			for (const util::Point<int64_t>& nb : getNeighbors(grid, p))
			{
				if (distancesFromStart.at(nb) == -1)
				{
					q.emplace(nb, distance + 1);
				}
			}
		}
		return *max_element(distancesFromStart.begin(), distancesFromStart.end());
	}

	int64_t part2(const vector<string>& lines)
	{
		util::Grid<char> grid = convertToGrid(lines);
		util::Point<int64_t> startingPosition = replaceStartingPointIteratorWithCorrectPipe(grid);
		util::Grid<int64_t> distancesFromStart{ grid.getWidth(), grid.getHeight() };
		ranges::fill(distancesFromStart, -1);
		distancesFromStart.at(startingPosition) = 0;

		queue<pair<util::Point<int64_t>, int64_t>> q;
		for (const util::Point<int64_t>& nb : getNeighbors(grid, startingPosition))
		{
			q.emplace(nb, 1);
		}

		while (!q.empty())
		{
			auto [p, distance] = q.front();
			q.pop();

			if (distancesFromStart.at(p) != -1)
			{
				// at the end of the loop, two nodes will have same neighbor & add this one to the queue twice
				// so, check if we already visited this node
				continue;
			}
			distancesFromStart.at(p) = distance;
			for (const util::Point<int64_t>& nb : getNeighbors(grid, p))
			{
				if (distancesFromStart.at(nb) == -1)
				{
					q.emplace(nb, distance + 1);
				}
			}
		}

		for (int i = 0; i < grid.getHeight(); ++i)
		{
			for (int j = 0; j < grid.getWidth(); ++j)
			{
				if (grid.at(j, i) != '.' && distancesFromStart.at(j, i) == -1)
				{
					grid.at(j, i) = '.';
				}
			}
		}

		util::Grid<char> gridEnlarged{ grid.getWidth() * 2, grid.getHeight() * 2 };
		ranges::fill(gridEnlarged, ' ');
		for (int i = 0; i < gridEnlarged.getWidth(); ++i)
		{
			gridEnlarged.at(i, 0) = 'O';
			gridEnlarged.at(i, gridEnlarged.getHeight() - 1) = 'O';
		}
		for (int i = 0; i < gridEnlarged.getHeight(); ++i)
		{
			gridEnlarged.at(0, i) = 'O';
			gridEnlarged.at(gridEnlarged.getWidth() - 1, i) = 'O';
		}

		for (int i = 0; i < grid.getHeight(); ++i)
		{
			for (int j = 0; j < grid.getWidth(); ++j)
			{
				if (grid.at(j, i) != '.' && distancesFromStart.at(j, i) != -1)
				{
					gridEnlarged.at(j * 2, i * 2) = grid.at(j, i);
					continue;
				}
				if (i == 0 || i == grid.getHeight() - 1 || j == 0 || j == grid.getWidth() - 1)
				{
					continue;
				}
				gridEnlarged.at(j * 2, i * 2) = 'I';
			}
		}

		for (int i = 0; i < grid.getHeight(); ++i) // y
		{
			for (int j = 0; j < grid.getWidth(); ++j) // x
			{
				if (j < grid.getWidth() - 1)
				{
					char westPipe = grid.at(j, i);
					char eastPipe = grid.at(j + 1, i);
					switch (westPipe)
					{
					case '-':
					case 'L':
					case 'F':
						switch (eastPipe)
						{
						case '-':
						case 'J':
						case '7':
							gridEnlarged.at(j * 2 + 1, i * 2) = 'X';
							break;
						}
						break;
					}
				}
				if (i < grid.getHeight() - 1)
				{
					char northPipe = grid.at(j, i);
					char southPipe = grid.at(j, i + 1);
					switch (northPipe)
					{
					case '|':
					case '7':
					case 'F':
						switch (southPipe)
						{
						case '|':
						case 'J':
						case 'L':
							gridEnlarged.at(j * 2, i * 2 + 1) = 'X';
							break;
						}
						break;
					}
				}
			}
		}

		for (int i = 0; i < gridEnlarged.getHeight(); ++i) // y
		{
			for (int j = 0; j < gridEnlarged.getWidth(); ++j) // x
			{
				if (gridEnlarged.at(j, i) != 'O')
				{
					continue;
				}
				util::Point<int64_t> p{ j, i };
				stack<util::Point<int64_t>> stack;
				stack.emplace(p + NORTH);
				stack.emplace(p + SOUTH);
				stack.emplace(p + EAST);
				stack.emplace(p + WEST);
				while (!stack.empty())
				{
					util::Point<int64_t> neighbor = stack.top();
					stack.pop();

					if (!gridEnlarged.inBounds(neighbor) || (gridEnlarged.at(neighbor) != ' ' && gridEnlarged.at(neighbor) != 'I'))
					{
						continue;
					}
					gridEnlarged.at(neighbor) = 'O';
					stack.emplace(neighbor + NORTH);
					stack.emplace(neighbor + SOUTH);
					stack.emplace(neighbor + EAST);
					stack.emplace(neighbor + WEST);
				}
			}
		}
		return count(gridEnlarged.begin(), gridEnlarged.end(), 'I');
	}
}
