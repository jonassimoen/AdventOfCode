#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day17 {
	void printTree(util::Grid<int> grid)
	{
		ofstream outf(format("{}/day17.tgf", ASSETS_FOLDER));
		for (int y = 0; y < grid.getHeight(); y++) {
			for (int x = 0; x < grid.getWidth(); x++) {
				outf << y * grid.getWidth() + x << " (" << x << "," << y << ")" << endl;
			}
		}
		outf << "#" << endl;
		for (int y = 0; y < grid.getHeight(); y++) {
			for (int x = 1; x < grid.getWidth(); x++) {
				outf << y * grid.getWidth() + (x - 1) << " " << y * grid.getWidth() + x << " " << grid.at(x, y) << endl;
			}
		}
		for (int x = 0; x < grid.getWidth(); x++) {
			for (int y = 1; y < grid.getWidth(); y++) {
				outf << (y - 1) * grid.getWidth() + x << " " << y * grid.getWidth() + x << " " << grid.at(x, y) << endl;
			}
		}
	}

	util::Grid<int> parseGrid(const vector<string>& lines)
	{
		util::Grid<int> grid(lines[0].size(), lines.size());
		for (int y = 0; y < lines.size(); y++) {
			for (int x = 0; x < lines[y].size(); x++) {
				grid.at(x, y) = lines[y][x] - '0';
			}
		}
		return grid;
	}

	// UP - RIGHT - DOWN - LEFT
	const vector<util::Point<int64_t>> DIRECTIONS = {
		util::Point<int64_t>{0, -1},
		util::Point<int64_t>{1, 0},
		util::Point<int64_t>{0, 1},
		util::Point<int64_t>{-1, 0}
	};

	const vector<char> STRING_DIRECTIONS = { '^','>','v','<' };

	struct State {
		util::Point<int64_t> position = util::Point<int64_t>{ 0,0 };
		int direction = 0;
		int heatLoss = 0;
		int blocksInSameDirection = 0;
		bool operator == (const State& state) const {
			return state.position == position && state.blocksInSameDirection == blocksInSameDirection && state.direction == direction;
			// heat loss should not need a comparison as using a priority queue
		}
		friend std::ostream& operator << (std::ostream& os, const State& s);
	};

	std::ostream& operator << (std::ostream& os, const State& s) {
		os << s.position << " [" << s.direction << "] [" << s.blocksInSameDirection << "] [" << s.heatLoss << "]" << endl;
		return os;
	}
	struct StateComparator {
		bool operator()(const State& s1, const State& s2) const {
			return s1.heatLoss > s2.heatLoss;
		}
	};
	struct StateHasher {
		size_t operator()(const State& state) const {
			return state.position.y * 100000 + state.position.x;
		}
	};

	int64_t part1A(const vector<string>& lines)
	{
		util::Grid<int> grid = parseGrid(lines);
		priority_queue<State, vector<State>, StateComparator> q;
		unordered_set<State, StateHasher> visited;

		util::Point<int64_t> destination{ grid.getWidth() - 1, grid.getHeight() - 1 };

		State startState{ util::Point<int64_t>{0,0}, 1, 0, 0 };
		q.push(startState);
		startState.direction = 2;
		q.push(startState);

		State s;

		while (!q.empty()) {
			s = q.top();
			q.pop();
			if (visited.contains(s))
			{
				continue;
			}
			visited.insert(s);

			if (s.position == destination) {
				return s.heatLoss;
			}
			// Move left
			int direction = (s.direction - 1 + 4) % 4;
			util::Point<int64_t> newPosition = s.position + DIRECTIONS[direction];
			if (grid.inBounds(newPosition)) {
				State newState{ newPosition, direction, s.heatLoss + grid.at(newPosition), 0 };
				q.push(newState);
			}

			// Move right
			direction = (s.direction + 1 + 4) % 4;
			newPosition = s.position + DIRECTIONS[direction];
			if (grid.inBounds(newPosition)) {
				State newState{ newPosition, direction, s.heatLoss + grid.at(newPosition), 0 };
				q.push(newState);
			}

			// Move forward if possible
			if (s.blocksInSameDirection < 2)
			{
				util::Point<int64_t> newPosition = s.position + DIRECTIONS[s.direction];
				if (grid.inBounds(newPosition)) {
					State newState{ newPosition, s.direction, s.heatLoss + grid.at(newPosition), s.blocksInSameDirection + 1 };
					q.push(newState);
				}
			}
		}
		return 0;
	}

	int64_t part1B(const vector<string>& lines)
	{
		util::Grid<int> grid = parseGrid(lines);
		priority_queue<State, vector<State>, StateComparator> q;
		unordered_set<State, StateHasher> visited;
		util::Point<int64_t> start{ 0,0 };

		q.push({ util::Point<int64_t>{ 0,0 },-1,0,-1 });

		while (!q.empty()) {
			State s = q.top();
			q.pop();
			if (visited.contains(s))
			{
				continue;
			}
			if (s.position == util::Point<int64_t>{ grid.getWidth() - 1, grid.getHeight() - 1 })
			{
				return s.heatLoss;
			}
			visited.insert(s);

			for (int i = 0; i < DIRECTIONS.size(); ++i)
			{
				util::Point<int64_t> newPosition = s.position + DIRECTIONS[i];
				int newDirection = i;
				int blocksInSameDirection = s.direction == newDirection ? s.blocksInSameDirection + 1 : 1;
				bool notReversed = s.direction != (newDirection + 2) % 4;

				if (grid.inBounds(newPosition) && notReversed && blocksInSameDirection <= 3)
				{
					int cost = grid.at(newPosition);
					State newState{ newPosition, newDirection, s.heatLoss + cost, blocksInSameDirection };
					if (visited.contains(newState))
					{
						continue;
					}
					q.push(newState);
				}
			}
		}
		return 0;
	}

	int64_t part1(const vector<string>& lines)
	{
		//int64_t resultA = part1A(lines);
		int64_t resultB = part1B(lines);
		return resultB;
	}

	int64_t part2(const vector<string>& lines)
	{
		util::Grid<int> grid = parseGrid(lines);
		priority_queue<State, vector<State>, StateComparator> q;
		unordered_set<State, StateHasher> visited;
		util::Point<int64_t> start{ 0,0 };

		q.push({ util::Point<int64_t>{ 0,0 },-1,0,-1 });

		while (!q.empty()) {
			State s = q.top();
			q.pop();
			if (visited.contains(s))
			{
				continue;
			}
			if (s.position == util::Point<int64_t>{ grid.getWidth() - 1, grid.getHeight() - 1 })
			{
				return s.heatLoss;
			}
			visited.insert(s);

			for (int i = 0; i < DIRECTIONS.size(); ++i)
			{
				util::Point<int64_t> newPosition = s.position + DIRECTIONS[i];
				int newDirection = i;
				int blocksInSameDirection = s.direction == newDirection ? s.blocksInSameDirection + 1 : 1;
				bool notReversed = s.direction != (newDirection + 2) % 4;

				if (grid.inBounds(newPosition) && notReversed && blocksInSameDirection <= 10 && (newDirection == s.direction || s.blocksInSameDirection >= 4 || s.blocksInSameDirection == -1))
				{
					int cost = grid.at(newPosition);
					State newState{ newPosition, newDirection, s.heatLoss + cost, blocksInSameDirection };
					if (visited.contains(newState))
					{
						continue;
					}
					q.push(newState);
				}
			}
		}
		return 0;
	}
}
