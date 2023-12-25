#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day22 {
	struct Brick {
		int64_t c = 0;
		util::Point3D<int64_t> startPos;
		util::Point3D<int64_t> size;

		bool operator==(const Brick& other) const {
			return startPos == other.startPos && size == other.size;
		}

		auto operator<=>(const Brick& other) const {
			return startPos <=> other.startPos;
		}

		struct Hash {
			size_t operator()(const Brick& brick) const {
				return std::hash<int64_t>()(brick.startPos.x) ^ (std::hash<int64_t>()(brick.startPos.y) << 1) ^ (std::hash<int64_t>()(brick.startPos.z) << 2)
					^ (std::hash<int64_t>()(brick.size.x) << 3) ^ (std::hash<int64_t>()(brick.size.y) << 4) ^ (std::hash<int64_t>()(brick.size.z) << 5);
				//return (brick.startPos.x * 10000 + brick.startPos.y * 100 + brick.startPos.z) * 100 + (brick.startPos.x * 10000 + brick.startPos.y * 100 + brick.startPos.z);
			}
		};
	};

	vector<Brick> parseBricks(const vector<string>& lines)
	{
		vector<Brick> bricks;
		int c = 0;
		for (const auto& line : lines) {
			auto parts = util::SplitString(line, "~");
			auto startPos = util::SplitToInt64(parts[0], ",");
			auto endPos = util::SplitToInt64(parts[1], ",");
			bricks.push_back({ c++, {startPos[0], startPos[1], startPos[2]}, {endPos[0] - startPos[0], endPos[1] - startPos[1], endPos[2] - startPos[2]} });
		}
		return bricks;
	}

	unordered_map<Brick, set<Brick>, Brick::Hash> drop_bricks(vector<Brick>& bricks)
	{
		map<util::Point<int64_t>, int> heightAtPos2D;
		map<util::Point3D<int64_t>, Brick> containsBrick;
		unordered_map<Brick, set<Brick>, Brick::Hash> restingOn;
		sort(bricks.begin(), bricks.end(), [](const Brick& a, const Brick& b) {
			return (a.startPos.z != b.startPos.z) ? a.startPos.z < b.startPos.z : a.startPos.z + a.size.z < b.startPos.z + b.size.z;
			});
		for (auto& brick : bricks) {
			//cout << "Dropping block " << brick.c << endl;
			int64_t lowestPossibleZ = 0;

			for (int x = brick.startPos.x; x <= brick.startPos.x + brick.size.x; ++x) {
				for (int y = brick.startPos.y; y <= brick.startPos.y + brick.size.y; ++y) {
					int64_t heightAtCurrentPos = 1;
					if (heightAtPos2D.contains({ x, y }))
						heightAtCurrentPos = heightAtPos2D[{x, y}] + 1;
					lowestPossibleZ = max(heightAtCurrentPos, lowestPossibleZ);
				}
			}
			//cout << "\tfrom " << brick.startPos << " to " << brick.endPos << " lowest possible z is " << lowestPossibleZ << endl;

			brick.startPos.z = lowestPossibleZ;
			//cout << "\t\tDropped: " << brick.startPos << " to " << brick.endPos << endl;

			for (int x = brick.startPos.x; x <= brick.startPos.x + brick.size.x; ++x) {
				for (int y = brick.startPos.y; y <= brick.startPos.y + brick.size.y; ++y) {
					heightAtPos2D[{x, y}] = brick.startPos.z + brick.size.z;
					for (int z = brick.startPos.z; z <= brick.startPos.z + brick.size.z; ++z)
						containsBrick[{x, y, z}] = brick;
				}
			}
			for (int x = brick.startPos.x; x <= brick.startPos.x + brick.size.x; ++x) {
				for (int y = brick.startPos.y; y <= brick.startPos.y + brick.size.y; ++y) {
					if (containsBrick.contains({ x, y, brick.startPos.z - 1 })) {
						//cout << "\t\t\tResting on: " << containsBrick.at({x,y, brick.startPos.z-1}).c << endl;
						restingOn[brick].emplace(containsBrick.at({ x, y, brick.startPos.z - 1 }));
					}
				}
			}

		}
		return restingOn;
	}

	unordered_map<Brick, set<Brick>, Brick::Hash> restingOnToSupporting(const unordered_map<Brick, set<Brick>, Brick::Hash>& restingOn) {
		unordered_map<Brick, set<Brick>, Brick::Hash> supporting;
		for (const auto& [brick, restingOnBricks] : restingOn) {
			for (const auto& restingOnBrick : restingOnBricks) {
				supporting[restingOnBrick].emplace(brick);
			}
		}
		return supporting;
	}

	int64_t part1(const vector<string>& lines)
	{
		vector<Brick> bricks = parseBricks(lines);
		unordered_map<Brick, set<Brick>, Brick::Hash> restingOn = drop_bricks(bricks);
		unordered_map<Brick, set<Brick>, Brick::Hash> supporting = restingOnToSupporting(restingOn);

		int64_t totalDisintegratable = 0;
		ofstream outf(format("{}/log2", ASSETS_FOLDER));

		for (const Brick& brick : bricks)
		{
			outf << "Brick " << brick.c << " [ " << brick.startPos << " ] [ " << brick.startPos + brick.size << " supporting " << supporting[brick].size() << " bricks" << endl;
			bool isDisintegratable = true;
			// Check all bricks for which the current brick is a supporting brick
			for (const Brick& supportedBrick : supporting[brick]) {
				outf << "\tSupported brick " << supportedBrick.c << " [ " << brick.startPos << " ] [ " << brick.startPos + brick.size << " is resting on " << restingOn[supportedBrick].size() << " bricks" << endl;
				// Does this supported brick have multiple bricks on top of it? ==> disintegratable
				isDisintegratable &= restingOn[supportedBrick].size() > 1;
			}
			totalDisintegratable += isDisintegratable ? 1 : 0;
		}
		return totalDisintegratable;
	}

	int64_t part2(const vector<string>& lines)
	{
		vector<Brick> bricks = parseBricks(lines);
		unordered_map<Brick, set<Brick>, Brick::Hash> restingOn = drop_bricks(bricks);

		int64_t totalFall = 0;

		for (const Brick& brick : bricks)
		{
			set<Brick> bricksWouldFall{ brick };
			for (const Brick& brick2 : bricks) {
				if (brick2.startPos.z == 0)
					continue;

				
				if (restingOn[brick2].size() != 0 && 
					restingOn[brick2].size() == ranges::distance(restingOn[brick2] | views::filter([&bricksWouldFall](const Brick& b) { return bricksWouldFall.contains(b); })))
					bricksWouldFall.emplace(brick2);
				
			}
			totalFall += bricksWouldFall.size() - 1;
		}
		// Efficiënter: van boven naar onderen, met cache van aantal bricks die vallen als die brick valt
		return totalFall;
	}
}
