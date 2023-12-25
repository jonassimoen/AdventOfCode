#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day14 {
	util::Grid<char> parseGrid(const vector<string>& lines)
	{
		util::Grid<char> grid{ static_cast<int>(lines[0].size()), static_cast<int>(lines.size()) };
		for (int y = 0; y < lines.size(); ++y) {
			for (int x = 0; x < lines[y].size(); ++x) {
				grid.at(x, y) = lines[y][x];
			}
		}
		return grid;
	}

	int64_t part1(const vector<string>& lines)
	{
		int64_t result = 0;
		util::Grid<char> grid = parseGrid(lines);
		for (int x = 0; x < grid.getWidth(); ++x)
		{
			int ySouthestBlocked = -1;
			//cout << "Column " << x << ": " << endl;
			for (int y = 0; y < lines.size(); ++y)
			{
				if (grid.at(x, y) == '#')
				{
					ySouthestBlocked = y;
				}
				else if (grid.at(x, y) == 'O')
				{
					++ySouthestBlocked;
					result += (lines.size() - ySouthestBlocked);
					//cout << "\tO [" << x << ", " << y << "] => [" << x << ", " << ySouthestBlocked << "] => " << lines.size() - ySouthestBlocked << endl;
				}
			}
		}
		return result;
	}

	util::Grid<char> cycleOnce(util::Grid<char>& grid)
	{
		// NORTH
		for (int x = 0; x < grid.getWidth(); ++x)
		{
			int ySouthestBlocked = -1;
			for (int y = 0; y < grid.getHeight(); ++y)
			{
				if (grid.at(x, y) == '#')
				{
					ySouthestBlocked = y;
				}
				else if (grid.at(x, y) == 'O')
				{
					grid.at(x, y) = '.';
					grid.at(x, ++ySouthestBlocked) = 'O';
				}
			}
		}
		// WEST
		for (int y = 0; y < grid.getHeight(); ++y)
		{
			int xSouthestBlocked = -1;
			for (int x = 0; x < grid.getWidth(); ++x)
			{
				if (grid.at(x, y) == '#')
				{
					xSouthestBlocked = x;
				}
				else if (grid.at(x, y) == 'O')
				{
					grid.at(x, y) = '.';
					grid.at(++xSouthestBlocked, y) = 'O';
				}
			}
		}
		// SOUTH
		for (int x = 0; x < grid.getWidth(); ++x)
		{
			int ySouthestBlocked = grid.getHeight();
			for (int y = grid.getHeight() - 1; y >= 0; --y)
			{
				if (grid.at(x, y) == '#')
				{
					ySouthestBlocked = y;
				}
				else if (grid.at(x, y) == 'O')
				{
					grid.at(x, y) = '.';
					grid.at(x, --ySouthestBlocked) = 'O';
				}
			}
		}
		// EAST
		for (int y = 0; y < grid.getHeight(); ++y)
		{
			int xSouthestBlocked = grid.getWidth();
			for (int x = grid.getWidth()-1; x >=0 ; --x)
			{
				if (grid.at(x, y) == '#')
				{
					xSouthestBlocked = x;
				}
				else if (grid.at(x, y) == 'O')
				{
					grid.at(x, y) = '.';
					grid.at(--xSouthestBlocked, y) = 'O';
				}
			}
		}
		return grid;
	}

	int64_t score(const util::Grid<char>& grid)
	{
		int64_t result = 0;
		for (int y = 0; y < grid.getHeight(); ++y)
		{
			for (int x = 0; x < grid.getWidth(); ++x)
			{
				if (grid.at(x, y) == 'O')
				{
					result += grid.getHeight() - y;
				}
			}
		}
		return result;
	}

	int64_t part2(const vector<string>& lines)
	{
		util::Grid<char> grid = parseGrid(lines);

		unordered_map<string, int> iterationPerCycle; 

		bool foundCycle = false;
		static int MAX_ITERATIONS = 1000000000;
		for (int i = 0; i < MAX_ITERATIONS; ++i) {
			grid = cycleOnce(grid);

			string gridsSerialized = "";
			for (char c : grid) {
				gridsSerialized += c;
			}

			if (iterationPerCycle.contains(gridsSerialized)) {
				int lengthOfCycle = i - iterationPerCycle[gridsSerialized];
				i += (MAX_ITERATIONS - i - 1) / lengthOfCycle * lengthOfCycle;

			}
			iterationPerCycle[gridsSerialized] = i;
		}
		return score(grid);
	}
}
