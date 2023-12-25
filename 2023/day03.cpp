using namespace std;

namespace Day03 {
	int has_punct(string line, int from, int to)
	{
		if (!line.empty())
		{
			int _to = to < line.length() ? to : line.length() - 1;
			int _from = from >= 0 ? from : 0;

			for (int i = _from; i <= _to; i++)
			{
				char c = line[i];
				if ((c < '0' || c > '9') && c != '.' && c != '\0')
				{
					return 1;
				}
			}
		}
		return 0;
	}

	int64_t part1(const vector<string>& lines)
	{
		int64_t sum = 0;
		for (int i = 0; i < lines.size(); i++)
		{
			string line = '.' + lines[i] + '.';
			int start = -1, number = 0;
			for (int j = 0; j < line.length(); j++)
			{
				if (isdigit(line[j]))
				{
					if (start == -1)
						start = j;
					number = number * 10 + (line[j] - '0');
				}
				else if (number != 0)
				{
					int add = 0;
					if (i > 0)
						add |= has_punct(lines[i - 1], start - 1, j);
					if (i < lines.size() - 1)
						add |= has_punct(lines[i + 1], start - 1, j);
					add |= has_punct(line, start - 1, j);

					if (add)
						sum += number;

					start = -1;
					number = 0;
				}
			}
		}
		return sum;
	}

	void sum_numbers_line(const string& line, int place_gear, int* n_num, int* multipl)
	{
		int j = place_gear;
		int nr = -1;
		while (j > 0 && line[j - 1] >= '0' && line[j - 1] <= '9')
			j--;
		while ((j <= place_gear + 1) || nr != -1)
		{
			// cout << "Checking " << line[j] << endl;
			if (line[j] >= '0' && line[j] <= '9')
			{
				if (nr == -1)
					nr = 0;
				nr = nr * 10 + (line[j] - '0');
			}
			else if (nr != -1)
			{
				// cout << "Found number " << nr << endl;
				*multipl *= nr;
				*n_num += 1;
				nr = -1;
			}
			j++;
		}
	}

	int64_t part2(const vector<string>& lines)
	{
		int64_t sum = 0;
		for (size_t i = 0; i < lines.size(); i++)
		{
			for (size_t j = 0; j < lines[i].size(); j++)
			{
				if (lines[i][j] == '*')
				{
					// cout << "Gear found at " << i << ", " << j << endl;
					int n_num = 0, multipl = 1;
					if (i > 0)
					{
						sum_numbers_line(lines[i - 1], j, &n_num, &multipl);
					}
					if (i < lines.size() - 1)
					{
						sum_numbers_line(lines[i + 1], j, &n_num, &multipl);
					}
					sum_numbers_line(lines[i], j, &n_num, &multipl);

					if (n_num == 2)
					{
						// cout << "Adding " << multipl << endl;
						sum += multipl;
					}
				}
			}
		}
		return sum;
	}
}