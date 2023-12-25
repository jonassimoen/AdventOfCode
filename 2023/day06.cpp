using namespace std;

namespace Day06 {
	int part1(const vector<string>& lines)
	{
		vector<int> time = util::ExtractInts32(lines[0].substr(lines[0].find(':')));
		vector<int> distance = util::ExtractInts32(lines[1].substr(lines[1].find(':')));

		int multipl = 1;
		for (int i = 0; i < time.size(); ++i)
		{
			int t = time[i];
			int maximalDistance = distance[i] + 1;

			float a = -1.0f;
			float b = static_cast<float>(t);
			float c = static_cast<float>(-maximalDistance);
			float d = std::sqrt(b * b - 4 * a * c);
			float minX = (b - d) / 2.0f;
			float maxX = (b + d) / 2.0f;

			multipl *= static_cast<int>(floor(maxX)) - static_cast<int>(ceil(minX)) + 1;
		}

		return multipl;
	}

	int part2(const vector<string>& lines)
	{
		int64_t time = 0;
		for (char c : lines[0]) {
			if (isdigit(c)) 
				time = time * 10 + (c - '0');
		}
		int64_t distance = 0;
		for (char c : lines[1]) {
			if (isdigit(c))
				distance = distance * 10 + (c - '0');
		}
		int64_t multipl = 1;


		float a = -1.0f;
		float b = static_cast<float>(time);
		float c = static_cast<float>(-(distance+1));
		float d = std::sqrt(b * b - 4 * a * c);
		float minX = (b - d) / 2.0f;
		float maxX = (b + d) / 2.0f;

		return static_cast<int>(floor(maxX)) - static_cast<int>(ceil(minX)) + 1;
		
	}
}