#include <string>
#include <vector>
#include <algorithm>
#include <input.h>
using namespace std;

template<typename T>
void printVector(const string& prefix, const vector<T>& v)
{
	cout << prefix << ": ";
	for (const T& val : v)
	{
		cout << val << " ";
	}
	cout << endl;
}

namespace Day09 {
	int64_t part1(const vector<string>& lines)
	{
		auto linesSplittedInts = views::transform(lines, [](const string& line) { return util::SplitToInt64(line, " "); });
		vector<vector<int64_t>> sequences = vector<vector<int64_t>>{linesSplittedInts.begin(), linesSplittedInts.end()};

		vector<int64_t> nextNumberPerSequence;

		for (const vector<int64_t>& seq : sequences) 
		{
			vector<vector<int64_t>> layers;
			layers.push_back(seq);

			while (!ranges::all_of(layers.back(), [](const int64_t& val) { return val == 0; }))
			{
				// add new layer
				vector<int64_t> nextLayer;
				for (int i = 0; i < layers.back().size() - 1; ++i)
				{
					nextLayer.push_back(layers.back()[i + 1] - layers.back()[i]);
				}
				layers.push_back(move(nextLayer));
			}

			vector<int64_t> predictionPerLayer{ 0 };
			for (int layer = layers.size()-2; layer >= 0; --layer)
			{
				predictionPerLayer.emplace_back(layers[layer].back() + predictionPerLayer.back());
			}
			nextNumberPerSequence.emplace_back(predictionPerLayer.back());
		}
		return accumulate(nextNumberPerSequence.begin(), nextNumberPerSequence.end(), 0);
	}

	int64_t part2(const vector<string>& lines)
	{
		auto linesSplittedInts = views::transform(lines, [](const string& line) { return util::SplitToInt64(line, " "); });
		vector<vector<int64_t>> sequences = vector<vector<int64_t>>{ linesSplittedInts.begin(), linesSplittedInts.end() };

		vector<int64_t> nextNumberPerSequence;

		for (const vector<int64_t>& seq : sequences)
		{
			vector<vector<int64_t>> layers;
			layers.push_back(seq);

			while (!ranges::all_of(layers.back(), [](const int64_t& val) { return val == 0; }))
			{
				// add new layer
				vector<int64_t> nextLayer;
				for (int i = 0; i < layers.back().size() - 1; ++i)
				{
					nextLayer.push_back(layers.back()[i + 1] - layers.back()[i]);
				}

				layers.push_back(move(nextLayer));
			}

			vector<int64_t> predictionPerLayer{ 0 };
			for (int layer = layers.size() - 2; layer >= 0; --layer)
			{
				predictionPerLayer.emplace_back(layers[layer].front() - predictionPerLayer.back());
			}
			nextNumberPerSequence.emplace_back(predictionPerLayer.back());
		}
		return accumulate(nextNumberPerSequence.begin(), nextNumberPerSequence.end(), 0);
	}
}
