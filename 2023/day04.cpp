using namespace std;

namespace Day04 {
	void construct_trees(const string& line, set<int>& winning, vector<int>& picked)
	{
		int nr = 0;
		int verticalBar = 0;
		for (int j = 0; j < line.size(); j++)
		{
			if (line[j] >= '0' && line[j] <= '9')
			{
				nr = nr * 10 + (line[j] - '0');
			}
			else if (nr != 0)
			{
				if (line[j] != ':')
				{
					if (verticalBar)
						picked.push_back(nr);
					else
						winning.insert(nr);
				}
				nr = 0;
			}
			else if (line[j] == '|')
			{
				verticalBar = 1;
			}
		}
		if (nr != 0)
			picked.push_back(nr);
	}

	int64_t part1(const vector<string>& lines)
	{
		int sum = 0;
		for (int i = 0; i < lines.size(); i++)
		{
			set<int> winning;
			vector<int> picked;
			construct_trees(lines[i], winning, picked);

			int count = 0;
			for (int p : picked)
				count += winning.count(p);
			if (count > 0)
				sum += (1 << count - 1);
		}
		return sum;
	}

	int64_t part2(const vector<string>& lines)
	{
		int64_t sum = 0;
		vector<int> amountOfCards = vector<int>(lines.size(), 1);
		for (int i = 0; i < lines.size(); i++)
		{
			set<int> winning;
			vector<int> picked;
			construct_trees(lines[i], winning, picked);

			int count = 0;
			for (int p : picked)
				count += winning.count(p);

			for (int j = 1; j <= count; j++)
				amountOfCards[i + j] += amountOfCards[i];
		}
		for (int i : amountOfCards)
			sum += i;
		return sum;
	}
}