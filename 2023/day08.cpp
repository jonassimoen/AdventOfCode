#include <string>
#include <vector>
#include <map>
#include <ranges>
#include <numeric>
#include <input.h>
using namespace std;

namespace Day08 {

	struct Node {
		string l, r;
	};

	map<string, Node> parseNodes(const vector<string>& lines) {
		map<string, Node> nodes;
		for (const std::string& line : lines | std::views::drop(2)) {
			string nodeStr = line.substr(0, 3);
			string l = line.substr(7, 3);
			string r = line.substr(12, 3);
			nodes[nodeStr] = { l, r };
		}
		return nodes;
	}

	int64_t part1(const vector<string>& lines)
	{
		string path = lines[0];
		map<string, Node> m = parseNodes(lines);

		bool found = false;
		size_t i = 0;
		auto currentNodeRanges = (m | views::keys | views::filter([](const string& s) {return s == "AAA"; }));

		if (!currentNodeRanges.empty()) {
			string currentNode = currentNodeRanges.front();
			int steps = 0;
			while (i < path.size() && !found) {

				if (path[i] == 'L')
					currentNode = m[currentNode].l;
				else if (path[i] == 'R')
					currentNode = m[currentNode].r;
				else
					break;
				++steps;

				if (currentNode == "ZZZ") {
					found = true;
					break;
				}

				if (i == path.size() - 1)
					i = 0;
				else
					i++;
			}
			return steps;
		}
		else {
			return -1;
		}
	}

	int64_t part2(const vector<string>& lines)
	{
		string path = lines[0];
		map<string, Node> m = parseNodes(lines);

		auto nodesStartingWithA = (m | views::keys | views::filter([](const string& s) { return s[2] == 'A'; }));
		vector<string> currentNodes = vector<string>({ nodesStartingWithA.begin(), nodesStartingWithA.end() });
		vector<int64_t> steps = vector<int64_t>(currentNodes.size(), 0);

		bool found = false;
		int64_t stepsTaken = 0;
		int i = 0;
		while (i < path.size() && !found) {
			++stepsTaken;

			for (int j = 0; j < currentNodes.size(); ++j) {
				if (steps[j] > 0) {
					continue;
				}

				string& currentNode = currentNodes[j];
				currentNode = (path[i]=='L')? m[currentNode].l : m[currentNode].r;

				if(currentNode.back() == 'Z')
				{
					steps[j] = stepsTaken;
				}
			}
			i = (i == path.size() - 1) ? 0 : i + 1;

			found = ranges::all_of(steps, [](int64_t i) { return i > 0; });
		}

		return accumulate(steps.begin(), steps.end(), 1LL, lcm<int64_t, int64_t>);
	}
}
