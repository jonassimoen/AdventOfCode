#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day21 {
	constexpr util::Point<int64_t> NORTH{ 0, -1 }, SOUTH{ 0,1 }, EAST{ 1,0 }, WEST{ -1,0 };
	const vector<util::Point<int64_t>> DIRECTIONS{ NORTH,SOUTH,EAST,WEST };
	pair<util::Grid<char>, util::Point<int64_t>> parseInput(const vector<string>& lines)
	{
		util::Point<int64_t> sp{ -1, -1 };
		util::Grid<char> grid(lines[0].size(), lines.size());
		for (int y = 0; y < lines.size(); ++y) {
			for (int x = 0; x < lines[y].size(); ++x) {
				grid.at(x, y) = lines[y][x];
				if (grid.at(x, y) == 'S')
					sp = { x,y };
			}
		}
		return { grid,sp };
	}

	struct PointHash
	{
		size_t operator()(const util::Point<int64_t>& point) const
		{
			size_t xHash = std::hash<int>()(point.x);
			size_t yHash = std::hash<int>()(point.y) << 1;
			return xHash ^ yHash;
		}
	};

	int64_t part1(const vector<string>& lines)
	{
		const int MAX_STEPS = 64;

		auto res = parseInput(lines);
		util::Grid<char> grid = res.first;
		util::Point<int64_t> sp = res.second;
		unordered_set<util::Point<int64_t>, PointHash> reachableInCurrentStep;
		unordered_set<util::Point<int64_t>, PointHash> reachableInNextStep;

		reachableInNextStep.insert(sp);

		for (int i = 0; i < MAX_STEPS; ++i) {
			grid = res.first;
			reachableInCurrentStep.clear();
			reachableInCurrentStep = reachableInNextStep;
			reachableInNextStep.clear();
			for (const auto& p : reachableInCurrentStep) {
				for (const auto& d : DIRECTIONS) {
					auto np = p + d;
					if (grid.inBounds(np) && grid.at(np) != '#') {
						grid.at(np) = 'O';
						reachableInNextStep.insert(np);
					}
				}
			}
		}
		return reachableInNextStep.size();
	}

	struct QueueEntry {
		util::Point<int64_t> pt;
		int64_t step;

		bool operator==(const QueueEntry& other) const {
			return pt == other.pt && step == other.step;
		}

		struct Hash {
			size_t operator()(const QueueEntry& qe) const {
				size_t ptHash = PointHash()(qe.pt);
				size_t stepHash = std::hash<int64_t>()(qe.step) << 1;
				return ptHash ^ stepHash;
			}
		};
	};

	int64_t applySteps(const util::Grid<char> grid, util::Point<int64_t> start, int64_t steps)
	{
		unordered_set<util::Point<int64_t>, PointHash> reachableInCurrentStep;
		unordered_set<util::Point<int64_t>, PointHash> reachableInNextStep;
		reachableInNextStep.insert(start);

		for (int step = 0; step < steps; ++step)
		{
			reachableInCurrentStep.clear();
			reachableInCurrentStep = reachableInNextStep;
			reachableInNextStep.clear();
			for (const auto& p : reachableInCurrentStep) {
				for (const auto& d : DIRECTIONS) {
					auto np = p + d;
					if (grid.inBounds(np) && grid.at(np) != '#') {
						reachableInNextStep.insert(np);
					}
				}
			}
		}
		return reachableInNextStep.size();
		/*queue<QueueEntry> q;
		q.push({ start, steps });
		unordered_set< util::Point<int64_t>, PointHash> visited;
		unordered_set<util::Point<int64_t>, PointHash> reachable;

		while (!q.empty()) {
			QueueEntry curr = q.front();
			q.pop();

			if (curr.step % 2 == 0) {
				reachable.insert(curr.pt);
			}
			if(curr.step == 0)
				continue;
			for (const auto& dir : DIRECTIONS) {
				auto np = curr.pt + dir;
				if (grid.inBounds(np) && grid.at(np) != '#' && visited.contains(np)) {
					q.push({ np, curr.step - 1 });
					visited.insert(np);
				}
			}
		}
		return reachable.size();*/
	}

	// https://www.youtube.com/watch?v=9UOMZSL0JTg
	int64_t part2(const vector<string>& lines)
	{
		const int MAX_STEPS = 26501365;

		auto res = parseInput(lines);
		util::Grid<char> grid = res.first;
		util::Point<int64_t> sp = res.second;

		util::Grid<size_t> reachableInSubgrids(5, 5);

		int64_t size = grid.getWidth();

		/*
		Number of grids stacked next to each other:		MAX_STEPS // grid.getWidth()
		Number of grids stacked on top of each other:	MAX_STEPS // grid.getHeight()
		Number of grids we completely pass:				floor(MAX_STEPS // grid.getWidth()) - 1		= MAP_SIZE
		    (not first -> startpoint, not last one -> semi-passing)
		Number of grids we partially pass:				MAX_STEPS % grid.getWidth()
		*/

		int64_t grid_w = MAX_STEPS / size - 1;

		// Odd grid = grids where we have odd number of steps
		int64_t oddGrids = pow(grid_w / 2 * 2 + 1, 2);
		int64_t oddGridsSteps = applySteps(grid, sp, size*2+1); // stepsize is randomly chosen, but large enough

		// Odd grid = grids where we have even number of steps
		int64_t evenGrids = pow((grid_w + 1) / 2 * 2, 2);
		int64_t evenGridsSteps = applySteps(grid, sp, size * 2); // stepsize is randomly chosen, but large enough
		
		// Corners in diamond shape of subgrids
		int64_t corner_top = applySteps(grid, {sp.x, size-1 }, size-1); // size - 1: assume we just have enough steps to reach the top (of the straight line)
		int64_t corner_right = applySteps(grid, { 0, sp.y }, size - 1);
		int64_t corner_bottom = applySteps(grid, { size-1, sp.y }, size - 1);
		int64_t corner_left = applySteps(grid, { sp.x , 0 }, size - 1);

		int64_t corner_small_tr = applySteps(grid, { 0, size-1 }, size / 2 - 1);
		int64_t corner_small_tl = applySteps(grid, { size-1, size - 1 }, size / 2 - 1);
		int64_t corner_small_br = applySteps(grid, { 0, 0 }, size / 2 - 1);
		int64_t corner_small_bl = applySteps(grid, { size-1, 0 }, size / 2 - 1);

		int64_t corner_large_tr = applySteps(grid, { 0, size - 1 },3* size / 2 - 1);
		int64_t corner_large_tl = applySteps(grid, { size - 1, size - 1 }, 3 * size / 2 - 1);
		int64_t corner_large_br = applySteps(grid, { 0, 0 }, 3 * size / 2 - 1);
		int64_t corner_large_bl = applySteps(grid, { size - 1, 0 }, 3 * size / 2 - 1);

		return (
			oddGrids * oddGridsSteps + evenGrids * evenGridsSteps +
			corner_top + corner_right + corner_bottom + corner_left +
			(grid_w + 1) * (corner_small_bl + corner_small_br + corner_small_tl + corner_small_tr) +
			grid_w * (corner_large_bl + corner_large_br + corner_large_tl + corner_large_tr )
			);
	}
}
