#include <string>
#include <vector>
#include <unordered_set>
#include <input.h>
using namespace std;

namespace Day11 {
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

	int64_t part1(const vector<string>& lines)
	{
		util::Grid<char> grid = convertToGrid(lines);

		unordered_set<int> nonGalaxyRows;
		unordered_set<int> nonGalaxyCols;

		for (int i = 0; i < grid.getHeight(); ++i)
		{
			nonGalaxyRows.insert(i);
		}
		for (int i = 0; i < grid.getWidth(); ++i)
		{
			nonGalaxyCols.insert(i);
		}

		vector<util::Point<int>> galaxyPoints;
		for (int y = 0; y < grid.getHeight(); ++y)
		{
			for (int x = 0; x < grid.getWidth(); ++x)
			{
				if (grid.at(x, y) == '#')
				{
					galaxyPoints.emplace_back(x, y);
					nonGalaxyCols.erase(x);
					nonGalaxyRows.erase(y);
				}
			}
		}
		int64_t totalDistance = 0;
		for (int i = 0; i < galaxyPoints.size(); ++i)
		{
			const util::Point<int>& pA = galaxyPoints[i];
			for (int j = i + 1; j < galaxyPoints.size(); ++j) {
				const util::Point<int>& pB = galaxyPoints[j];

				int manhattan = abs(pA.x - pB.x) + abs(pA.y - pB.y);
				int emptyColsInPath = ranges::count_if(nonGalaxyCols, [pA, pB](int x) {
					return x > min(pA.x, pB.x) && x < max(pA.x, pB.x);
				});
				int emptyRowsInPath = ranges::count_if(nonGalaxyRows, [pA, pB](int y) {
					return y > min(pA.y, pB.y) && y < max(pA.y, pB.y);
				});

				totalDistance += (manhattan + emptyColsInPath + emptyRowsInPath);
			}
			//cout << "Point " << i << ": " << p.x << "," << p.y << endl;
		}

		return totalDistance;
	}

	int64_t part2(const vector<string>& lines)
	{
		util::Grid<char> grid = convertToGrid(lines);

		unordered_set<int> nonGalaxyRows;
		unordered_set<int> nonGalaxyCols;

		for (int i = 0; i < grid.getHeight(); ++i)
		{
			nonGalaxyRows.insert(i);
		}
		for (int i = 0; i < grid.getWidth(); ++i)
		{
			nonGalaxyCols.insert(i);
		}

		vector<util::Point<int>> galaxyPoints;
		for (int y = 0; y < grid.getHeight(); ++y)
		{
			for (int x = 0; x < grid.getWidth(); ++x)
			{
				if (grid.at(x, y) == '#')
				{
					galaxyPoints.emplace_back(x, y);
					nonGalaxyCols.erase(x);
					nonGalaxyRows.erase(y);
				}
			}
		}
		int64_t totalDistance = 0;
		for (int i = 0; i < galaxyPoints.size(); ++i)
		{
			const util::Point<int>& pA = galaxyPoints[i];
			for (int j = i + 1; j < galaxyPoints.size(); ++j) {
				const util::Point<int>& pB = galaxyPoints[j];

				int manhattan = abs(pA.x - pB.x) + abs(pA.y - pB.y);
				int emptyColsInPath = ranges::count_if(nonGalaxyCols, [pA, pB](int x) {
					return x > min(pA.x, pB.x) && x < max(pA.x, pB.x);
					});
				int emptyRowsInPath = ranges::count_if(nonGalaxyRows, [pA, pB](int y) {
					return y > min(pA.y, pB.y) && y < max(pA.y, pB.y);
					});
				
				constexpr int64_t emptySpaceCost = 1000000-1;
				totalDistance += (manhattan + emptyColsInPath* emptySpaceCost + emptyRowsInPath* emptySpaceCost);
			}
			//cout << "Point " << i << ": " << p.x << "," << p.y << endl;
		}

		return totalDistance;
	}
}
