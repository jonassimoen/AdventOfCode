#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day18 {
	// UP - RIGHT - DOWN - LEFT
	util::Point<int64_t> direction(char c) {
		switch (c) {
		case 'U':
			return util::Point<int64_t>{0, -1};
		case 'R':
			return util::Point<int64_t>{1, 0};
		case 'D':
			return util::Point<int64_t>{0, 1};
		case 'L':
			return util::Point<int64_t>{-1, 0};
		default:
			return util::Point<int64_t>{0, 0};
		}
	}
	// RIGHT - DOWN - LEFT - UP
	const vector<util::Point<int64_t>> DIRECTIONS = {
		util::Point<int64_t>{1, 0},
		util::Point<int64_t>{0, 1},
		util::Point<int64_t>{-1, 0},
		util::Point<int64_t>{0, -1},
	};

	int64_t part1(const vector<string>& lines)
	{
		int64_t perimeter = 0;
		vector<util::Point<int64_t>> corners;
		corners.emplace_back(0, 0);
		util::Point<int64_t> pos = util::Point<int64_t>{ 0, 0 };
		for (const string& line : lines) {
			vector<string>  tokens = util::SplitString(line, " ");
			int64_t number = util::ExtractInts<int64_t>(tokens[1])[0];
			corners.emplace_back(corners.back() + direction(tokens[0][0]) * number);
			perimeter += number;
		}
		int64_t area = 0;
		for (int i = 0; i < corners.size() - 1;++i) {
			area += ((corners[i + 1].y * corners[i].x) - (corners[i + 1].x * corners[i].y));
		}
		return area / 2 + perimeter / 2 + 1;
	}

	int64_t part2(const vector<string>& lines)
	{
		int64_t perimeter = 0;
		vector<util::Point<int64_t>> corners;
		corners.emplace_back(0, 0);
		util::Point<int64_t> pos = util::Point<int64_t>{ 0, 0 };
		for (const string& line : lines) {
			vector<string>  tokens = util::SplitString(line, " ");
			string hexa = tokens[2].substr(2, 5);
			int directionInt = tokens[2][7] - '0';
			int number = stoi(hexa, 0, 16);
			corners.emplace_back(corners.back() + DIRECTIONS[directionInt] * number);
			perimeter += number;
		}
		int64_t area = 0;
		for (int i = 0; i < corners.size() - 1; ++i) {
			area += ((corners[i + 1].y * corners[i].x) - (corners[i + 1].x * corners[i].y));
		}
		area /= 2;
		return area + perimeter / 2 + 1;
	}
}
