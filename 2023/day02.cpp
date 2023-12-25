using namespace std;

namespace Day02 {
	int64_t part1(const vector<string>& lines)
	{
		int64_t sum = 0;
		for (const string& line : lines) {
			int gid = stoi(line.substr(5, line.find(':')));

			string lineS = line.substr(line.find(':') + 2);
			char* d = lineS.data();
			char* read_to = nullptr;

			bool lineOK = true;
			while (*d != '\0' && lineOK)
			{
				int cnt = strtol(d, &read_to, 10);
				d = read_to;

				switch (d[1])
				{
				case 'r':
					d += 4;
					if (cnt > 12) {
						lineOK = false;
					}
					break;
				case 'g':
					d += 6;
					if (cnt > 13) {
						lineOK = false;
					}
					break;
				case 'b':
					d += 5;
					if (cnt > 14) {
						lineOK = false;
					}
					break;
				}

				if (*d == ',' || *d == ';')
					d++;
			}
			if (lineOK) 
				sum += gid;
		}
		return sum;
	}

	typedef struct ColorCount {
		int r;
		int g;
		int b;
	} ColorCount;

	int64_t part2(const vector<string>& lines)
	{
		int64_t sum = 0;
		for (const string& line : lines) {
			int gid = stoi(line.substr(5, line.find(':')));

			string lineS = line.substr(line.find(':') + 2);
			char* d = lineS.data();
			char* read_to = nullptr;

			ColorCount c = ColorCount{ 0,0,0 };

			// const char* gamedata = line.substr(line.find(':') + 2).c_str();
			while (*d != '\0')
			{
				int cnt = strtol(d, &read_to, 10);
				d = read_to;

				switch (d[1])
				{
				case 'r':
					d += 4;
					if (cnt > c.r) {
						c.r = cnt;
					}
					break;
				case 'g':
					d += 6;
					if (cnt > c.g) {
						c.g = cnt;
					}
					break;
				case 'b':
					d += 5;
					if (cnt > c.b) {
						c.b = cnt;
					}
					break;
				}

				if (*d == ',' || *d == ';')
					d++;
			}

			sum+= c.r * c.b * c.g;
		}
		return sum;
	}
}
