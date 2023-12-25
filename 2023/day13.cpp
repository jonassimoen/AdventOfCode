#include <string>
#include <vector>
#include <input.h>
#include <grid.h>
#include <queue>
using namespace std;

namespace Day13 {

	vector<util::Grid<char>> convertToGrid(const vector<string>& lines) {
		return lines | views::split("") | views::transform([](const auto&& sub) {
			vector<string> gridLines = sub | ranges::to<vector>();
			auto grid = util::Grid<char>{ static_cast<int>(gridLines[0].size()), static_cast<int>(gridLines.size()) };
			for (int y = 0; y < grid.getHeight(); ++y)
			{
				const string& line = gridLines[y];
				for (int x = 0; x < grid.getWidth(); ++x)
				{
					grid.at(x, y) = line[x];
				}
			}
			return grid;
			}) | ranges::to<vector>();
	}

	int64_t part1(const vector<string>& lines)
	{
		int64_t result = 0;
		vector<util::Grid<char>> grids = convertToGrid(lines);

		int gridIdx = 0;
		
		for (const auto& grid : grids) {

			for (int row = 1; row < grid.getHeight(); ++row) 
			{
				int size = min(row, grid.getHeight() - row);
				bool reflectionFound = true;

				for (int s = 0; s < size; ++s) {
					for (int col = 0; col < grid.getWidth(); ++col) {
						if (grid.at(col, row - s -1) != grid.at(col, row + s))
						{
							reflectionFound = false;
							break;
						}
					}
				}
				if (reflectionFound)
				{
					result += (100 * row);
				}
			}

			for (int col = 1; col < grid.getWidth(); ++col) 
			{
				int size = min(col, grid.getWidth() - col);
				bool reflectionFound = true;

				for (int s = 0; s < size; ++s) {
					for (int row = 0; row < grid.getHeight(); ++row) {
						if (grid.at(col - s - 1, row) != grid.at(col + s, row))
						{
							reflectionFound = false;
							break;
						}
					}
				}
				if (reflectionFound)
				{
					result += col;
				}
			}
		}

		return result;
	}

	int64_t part2(const vector<string>& lines)
	{
		int64_t result = 0;
		vector<util::Grid<char>> grids = convertToGrid(lines);

		int gridIdx = 0;

		for (const auto& grid : grids) {
			gridIdx++;
			vector<int> smudgesRows(grid.getHeight(), 0);
			vector<int> smudgesCols(grid.getWidth(), 0);

			for (int row = 1; row < grid.getHeight(); ++row)
			{
				int size = min(row, grid.getHeight() - row);
				bool reflectionFound = true;

				for (int s = 0; s < size; ++s) {
					for (int col = 0; col < grid.getWidth(); ++col) {
						if (grid.at(col, row - s - 1) != grid.at(col, row + s))
						{
							++smudgesRows[row];
						}
					}
				}
			}

			for (int col = 1; col < grid.getWidth(); ++col)
			{
				int size = min(col, grid.getWidth() - col);
				bool reflectionFound = true;

				for (int s = 0; s < size; ++s) {
					for (int row = 0; row < grid.getHeight(); ++row) {
						if (grid.at(col - s - 1, row) != grid.at(col + s, row))
						{
							++smudgesCols[col];
						}
					}
				}
			}

			bool foundSmudge = false;
			for (int row = 0; row < smudgesRows.size(); ++row)
			{
				if (smudgesRows[row] == 1) {
					result += 100 * row;
					foundSmudge = true;
				}
			}


			if (!foundSmudge) {
				for (int col = 0; col < smudgesCols.size(); ++col)
				{
					if (smudgesCols[col] == 1) {
						result += col;
						foundSmudge = true;
					}
				}

				if (!foundSmudge) {
					cout << "Grid " << gridIdx << ": no smudge found" << endl;
				}
			}

		}

		return result;
	}
}
