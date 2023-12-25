#include <string>
#include <vector>
#include <regex>
#include <input.h>

using namespace std;

namespace Day24 {
	/*
	From input we have points : (x, y) and velocities : (vx, vy)
	Points on the line: Px, Py: (x,y) + t * (vx, vy)
	These points comply with: (Px - x) / vx = (Py - y) / vy
	Refactoring gives: Px * vy - x * vy = Py * vx - y * vx => (vy) * x + (- vx) * y = (Px * vy - Py * vx)
	*/

	struct HailStone {
		HailStone(util::Point3D<int64_t> start, util::Point3D<int64_t> velocity) : start(start), velocity(velocity) {
			a = velocity.y;
			b = -velocity.x;
			c = start.x * velocity.y - start.y * velocity.x;
		}
		double a, b, c;
		util::Point3D<int64_t> start;
		util::Point3D<int64_t> velocity;
	};

	vector<HailStone> parse(const vector<string>& lines) {
		vector<HailStone> result;
		for (auto& line : lines) {
			auto parts = util::SplitString(line, "@");
			auto start = util::SplitString(parts[0], ",");
			auto velocity = util::SplitString(parts[1], ",");
			result.push_back(HailStone({ stoll(start[0]), stoll(start[1]),stoll(start[2]) }, { stoi(velocity[0]), stoi(velocity[1]),stoi(velocity[2]) }));
		}
		return result;
	}

	int64_t part1(const vector<string>& lines)
	{
		const double LOWER_LIMIT = 200000000000000;
		const double UPPER_LIMIT = 400000000000000;

		int64_t result = 0;
		vector<HailStone> hailstones = parse(lines);

		// Two 3D lines (ax+by=c) are parallel if (a1*b2 - a2*b1 = 0) or (a1/a2 = b1/b2)
		// Set of two equations (a1*x + b1*y = c1) and (a2*x + b2*y = c2) ==> solving that equation gives both x and y

		vector<string> eq = {};

		for (int i = 0; i < hailstones.size(); ++i) {
			const auto& [a, b, c, start, velo] = hailstones[i];
			string eq = format(
				"(x-{})*({}-b) - (y-{})*({}-a)",
				start.x, velo.y, start.y, velo.x
			);
			string eq2 = format(
				"(y-{})*({}-c) - (z-{})*({}-b)",
				start.y, velo.z, start.z, velo.y
			);

			eq = regex_replace(eq, regex("--"), "+");
			eq2 = regex_replace(eq2, regex("--"), "+");
			cout << eq << endl;
			cout << eq2 << endl;
		}
		return result;
	}

	int64_t part2(const vector<string>& lines)
	{
		return 0;
	}
}
