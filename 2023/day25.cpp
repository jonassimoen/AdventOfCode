#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day25 {
	map<string, set<string>> parseInput(const vector<string>& lines)
	{
		map<string, set<string>> result;
		for (auto& line : lines) {
			auto connections = util::SplitString(line, ": ");
			auto connectedParts = util::SplitString(connections[1], " ");
			for (auto& connectedPart : connectedParts)
			{
				result[connections[0]].insert(connectedPart);
				result[connectedPart].insert(connections[0]);
			}
		}
		return result;
	}

	bool hasDisconnectedGroups(const map<string, set<string>>& parts)
	{
		set<string> visited;
		queue<string> toVisit;
		toVisit.push(parts.begin()->first);
		while (!toVisit.empty())
		{
			string current = toVisit.front();
			toVisit.pop();
			if (visited.find(current) != visited.end())
			{
				continue;
			}
			visited.insert(current);
			for (auto& connectedPart : parts.at(current))
			{
				toVisit.push(connectedPart);
			}
		}
		return visited.size() != parts.size();
	}

	int64_t multiplyGroupSizes(const map<string, set<string>>& parts)
	{
		set<string> visited;
		queue<string> toVisit;
		toVisit.push(parts.begin()->first);
		while (!toVisit.empty())
		{
			string current = toVisit.front();
			toVisit.pop();
			if (visited.find(current) != visited.end())
			{
				continue;
			}
			visited.insert(current);
			for (auto& connectedPart : parts.at(current))
			{
				toVisit.push(connectedPart);
			}
		}
		return visited.size() * (parts.size() - visited.size());
	}

	void printGraph(const map<string, set<string>>& parts) {
		set<pair<string, string>> alreadyDrawnEdges;
		ofstream outf(format("{}/day25.dot", ASSETS_FOLDER));
		outf << "graph Day25 {" << endl;
		for (auto& [part, nbs] : parts)
		{
			for (auto& nb : nbs)
			{
				if (!alreadyDrawnEdges.contains({ part, nb }) && !alreadyDrawnEdges.contains({ nb, part })) {
					outf << part << " -- " << nb << endl;
					alreadyDrawnEdges.insert({ part, nb });
				}
			}
		}
		outf << "}" << endl;
	}

	void removeEdges(map<string, set<string>>& parts, const array<pair<string, string>, 3> edges) {
		for (const auto& edge : edges) {
			parts[edge.first].erase(edge.second);
			parts[edge.second].erase(edge.first);
		}
	}

	int64_t part1(const vector<string>& lines)
	{
		map<string, set<string>> parts = parseInput(lines);
		printGraph(parts);

		// By visualizing the graph, u can clearly see the 3 edges to be removed
		system(format("neato -Tpng {}/day25.dot -o {}/day25graph.png", ASSETS_FOLDER, ASSETS_FOLDER).c_str());

		/* SAMPLE */
		//removeEdges(parts, {
		//	make_pair("bvb","cmg"),
		//	make_pair("hfx","pzl"),
		//	make_pair("nvd","jqt")
		//	});

		/* INPUT */
		removeEdges(parts, {
			make_pair("rsm","bvc"),
			make_pair("bkm","ldk"),
			make_pair("zmq","pgh")
			});
		
		if (!hasDisconnectedGroups(parts)) {
			return -1;
		}
		return multiplyGroupSizes(parts);
	}

	int64_t part2(const vector<string>& lines)
	{
		return 0;
	}
}
