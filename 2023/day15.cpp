#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day15 {
	int64_t HASH(const string& str) {
		int64_t step_value = 0;
		for (char c : str)
		{
			step_value += static_cast<int>(c);
			step_value *= 17;
			step_value %= 256;
		}
		return step_value;
	}
	int64_t part1(const vector<string>& lines)
	{
		int64_t result = 0;
		vector<string> steps = util::SplitString(lines[0], ",");
		for (const string& step : steps)
		{
			result += HASH(step);
		}

		return result;
	}

	struct Lens {
		string label;
		int focal_length;
	};
	struct Box {
		vector<Lens> lenses;
	};

	int64_t part2(const vector<string>& lines)
	{
		map<int, Box> boxes; // Boxes, each containing a map of labels to focal lengths
		vector<string> steps = util::SplitString(lines[0], ",");
		for (const string& step : steps)
		{
			int idxDelimit = step.find_first_of("-=");
			string label = step.substr(0, idxDelimit);
			int box = HASH(label);

			if (step.at(idxDelimit) == '=') {
				// Check if lens already in box
				int focal_length = stoi(step.substr(idxDelimit + 1));
				auto it = find_if(boxes[box].lenses.begin(), boxes[box].lenses.end(), [&label](const Lens& lens) { return lens.label == label; });
				if (it != boxes[box].lenses.end()) {
					it->focal_length = focal_length;
				}
				else {
					boxes[box].lenses.push_back({ label, focal_length });
				}
			}
			else if (step.at(idxDelimit) == '-') {
				// Remove lens from box if it exists
				auto it = find_if(boxes[box].lenses.begin(), boxes[box].lenses.end(), [&label](const Lens& lens) { return lens.label == label; });
				if (it != boxes[box].lenses.end()) {
					boxes[box].lenses.erase(it);
				}
			}
		}

		int64_t result = 0;
		for (int i = 0; i < boxes.size(); ++i)
		{
			if (boxes[i].lenses.size() > 0) {
				for (int j = 0; j < boxes[i].lenses.size(); ++j) {
					Lens& l = boxes[i].lenses[j];
					result += (i + 1) * (j + 1) * l.focal_length;
				}
			}
		}
		return result;
	}
}
