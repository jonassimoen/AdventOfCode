#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day23 {
	constexpr util::Point<int64_t> NORTH{ 0, -1 }, SOUTH{ 0,1 }, EAST{ 1,0 }, WEST{ -1,0 };
	const vector<util::Point<int64_t>> DIRECTIONS{ NORTH,SOUTH,EAST,WEST };
	const map<char, vector<util::Point<int64_t>>> DIRECTION_MAP{
		{ '^', {NORTH} },
		{ 'v', {SOUTH} },
		{ '>', {EAST} },
		{ '<', {WEST} },
		{ '.', {NORTH, SOUTH, EAST, WEST} }
	};

	util::Grid<char> parseGrid(const vector<string>& lines)
	{
		util::Grid<char> grid(lines[0].size(), lines.size());
		for (int y = 0; y < lines.size(); y++) {
			for (int x = 0; x < lines[y].size(); x++) {
				grid.at(x, y) = lines[y][x];
			}
		}
		return grid;
	}

	pair<util::Point<int64_t>, util::Point<int64_t>> getStartAndEndPoint(const util::Grid<char>& grid)
	{
		util::Point<int64_t> start, end;
		for (int x = 0; x < grid.getWidth(); x++) {
			if (grid.at(x, 0) == '.') {
				start = { x, 0 };
			}
			else if (grid.at(x, grid.getHeight() - 1) == '.') {
				end = { x, grid.getHeight() - 1 };
			}
		}
		return { start, end };
	}

	// Reduce grid: cells with only 1 possible path can be discarded, so we keep only the intersections, where a choice needs to be made
	set<util::Point<int64_t>> getVertices(const util::Grid<char>& grid)
	{
		set<util::Point<int64_t>> vertices;
		for (int y = 0; y < grid.getHeight(); ++y) {
			for (int x = 0; x < grid.getWidth(); ++x) {
				if (x == 3 && y == 5) {
					int a = 0;
				}
				if (grid.at(x, y) == '.') {
					int count = 0;
					for (const util::Point<int64_t>& dir : DIRECTIONS) {
						if (grid.inBounds(x + dir.x, y + dir.y) && grid.at(x + dir.x, y + dir.y) != '#') {
							count++;
						}
					}
					if (count > 2) {
						vertices.insert({ x, y });
					}
				}
			}
		}
		return vertices;
	}

	struct Edge {
		util::Point<int64_t> to;
		int64_t weight;
	};

	map<util::Point<int64_t>, vector<Edge>> getEdges(const util::Grid<char>& grid, const set<util::Point<int64_t>>& vertices, bool allDirections = false)
	{
		map<util::Point<int64_t>, vector<Edge>> edges;
		for (const util::Point<int64_t>& vertex : vertices) {
			set<util::Point<int64_t>> visited;
			queue<Edge> q;
			q.push({ vertex, 0 });
			while (!q.empty()) {
				Edge current = q.front();
				q.pop();
				// If we've already visited this vertex, we don't need to do it again
				if (visited.find(current.to) != visited.end()) {
					continue;
				}
				visited.insert(current.to);

				// Check if we've found a new vertex (not the one we started from) & add the distance from current to that vertex
				if (current.to != vertex && vertices.contains(current.to)) {
					edges[vertex].emplace_back(current.to, current.weight);
					continue;
				}

				// Now proceed with all neighbours of the current vertex
				auto possibleDirections = DIRECTION_MAP.at(grid.at(current.to));
				if (allDirections)
					possibleDirections = DIRECTIONS;
				for (const util::Point<int64_t>& dir : possibleDirections) {
					auto np = current.to + dir;
					if (grid.inBounds(np) && grid.at(np) != '#') {
						q.push({ np, current.weight + 1 });
					}
				}

			}
		}
		return edges;
	}

	int64_t part1(const vector<string>& lines)
	{
		int64_t result = 0;
		util::Grid<char> grid = parseGrid(lines);

		pair<util::Point<int64_t>, util::Point<int64_t>> startEnd = getStartAndEndPoint(grid);
		util::Point<int64_t> start = startEnd.first;
		util::Point<int64_t> end = startEnd.second;

		set<util::Point<int64_t>> vertices = getVertices(grid);
		vertices.insert(start);
		vertices.insert(end);

		map<util::Point<int64_t>, vector<Edge>> edges = getEdges(grid, vertices);

		set<util::Point<int64_t>> visited;

		stack<Edge> dq;
		dq.push({ start, 0 });

		while (!dq.empty()) {
			Edge c = dq.top();
			dq.pop();

			if (visited.find(c.to) != visited.end()) {
				continue;
			}

			visited.insert(c.to);
			if (c.to == end) {
				result = max(result, c.weight);
				visited.erase(c.to);
				continue;
			}

			for (const Edge& e : edges[c.to]) {
				dq.push({ e.to, c.weight + e.weight });
			}
			visited.erase(c.to);
		}
		return result;
	}



	int64_t part2(const vector<string>& lines)
	{
		int64_t result = 0;
		util::Grid<char> grid = parseGrid(lines);

		pair<util::Point<int64_t>, util::Point<int64_t>> startEnd = getStartAndEndPoint(grid);
		util::Point<int64_t> start = startEnd.first;
		util::Point<int64_t> end = startEnd.second;

		set<util::Point<int64_t>> vertices = getVertices(grid);
		vertices.insert(start);
		vertices.insert(end);

		map<util::Point<int64_t>, vector<Edge>> edges = getEdges(grid, vertices, true);

		set<util::Point<int64_t>> visited;
		map<util::Point<int64_t>, bool> seen;

		function<int64_t(util::Point<int64_t>)> DFS;
		DFS = [&DFS, &result, &end, &visited, &edges](util::Point<int64_t> node) -> int64_t {
			if (node == end)
				return 0;
			int64_t dist = 0;
			visited.insert(node);
			for (const Edge& e : edges[node]) 
				if (visited.find(e.to) == visited.end()) 
					dist = max(dist, DFS(e.to) + e.weight);
				
			visited.erase(node);
			return dist;
			};
		return DFS(start);
	}
}
