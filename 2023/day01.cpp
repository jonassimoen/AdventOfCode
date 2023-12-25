using namespace std;

namespace Day01 {
	int64_t part1(const vector<string>& lines)
	{
		int64_t sum = 0;
		for (const string& line : lines) {
			int first = -1, last = -1;

			for (char c : line)
			{
				if (isdigit(c))
				{
					if (first == -1)
					{
						first = c - '0';
					}
					last = c - '0';
				}
			}
			sum += first * 10 + last;
		}
		return sum;
	}

	int64_t part2(const vector<string>& lines)
	{
		string numbers[] = { "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };
		int64_t sum = 0;
		for (const string& line : lines) {
			size_t maxl = line.length();
			int first = -1, last = -1;
			for (int i = 0; i < line.length(); i++)
			{
				for (int j = 0; j < 20; j++)
				{
					if (maxl - i > 0 && numbers[j].compare(line.substr(i, numbers[j].length())) == 0)
					{
						if (first == -1)
						{
							first = j % 10;
						}
						last = j % 10;
					}
				}
			}
			sum += first * 10 + last;
		}
		return sum;
	}
}