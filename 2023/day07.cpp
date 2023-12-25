#include <string>
#include <vector>
#include <ranges>
#include <input.h>

using namespace std;

namespace Day07 {
	enum class Type {
		FiveKind,
		FourKind,
		FullHouse,
		ThreeKind,
		TwoPair,
		OnePair,
		HighCard,
	};

	string typeToString(Type t) {
		switch (t)
		{
		case Day07::Type::FiveKind:
			return "FiveKind";
			break;
		case Day07::Type::FourKind:
			return "FourKind";
			break;
		case Day07::Type::FullHouse:
			return "FullHouse";
			break;
		case Day07::Type::ThreeKind:
			return "ThreeKind";
			break;
		case Day07::Type::TwoPair:
			return "TwoPair";
			break;
		case Day07::Type::OnePair:
			return "OnePair";
			break;
		default:
			return "HighCard";
		}
	}

	struct Hand {
		string raw;
		Type type = Type::HighCard;
		int bid = 0;
	};

	Type parseHandType(string input, bool joker = false) {
		unordered_map<char, int> cardCounts;
		for (char c : input) {
				cardCounts[c]++;
		}
		
		int jCount = cardCounts['J'];

		if(joker) 
			cardCounts['J'] = 0; // Remove jokers from the count
			
		unordered_map<int, int> cardCountsCount;
		for (const auto& [card, amount] : cardCounts) {
				cardCountsCount[amount]++;
		}

		if (joker) {
			int maximalCount = max_element(cardCounts.begin(), cardCounts.end(), [](const pair<char, int>& a, const pair<char, int>& b) {return a.second < b.second; })->second;
			cardCountsCount[maximalCount] -= 1; // This entry of MAX does not exist 
			cardCountsCount[maximalCount + jCount] += 1; // But is replaced by count MAX + #J
		}

		if (cardCountsCount[5] >= 1)
			return Type::FiveKind;
		else if (cardCountsCount[4] >= 1)
			return Type::FourKind;
		else if (cardCountsCount[3] >= 1 && cardCountsCount[2] >= 1)
			return Type::FullHouse;
		else if (cardCountsCount[3] >= 1)
			return Type::ThreeKind;
		else if (cardCountsCount[2] >= 2)
			return Type::TwoPair;
		else if (cardCountsCount[2] >= 1)
			return Type::OnePair;
		else
			return Type::HighCard;
	}

	int64_t part1(const vector<string>& lines)
	{
		vector<Hand> hands;
		for (const string& line : lines) {
			vector<string> split = util::SplitString(line, " ");
			hands.emplace_back(split[0], parseHandType(split[0]), stoi(split[1]));
		}
		ranges::sort(
			hands,
			[](const Hand& hl, const Hand& hr)
			{
				if (hl.type == hr.type)
				{
					static const vector<char> strength = { 'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2' };
					for (int i = 0; i < hl.raw.size(); ++i) {
						if (hl.raw[i] != hr.raw[i]) {
							return ranges::find(strength, hl.raw[i]) > ranges::find(strength, hr.raw[i]);
						}
					}
				}
				return hl.type > hr.type;
			}
		);

		int64_t sum = 0;
		for (int r = 0; r < hands.size(); ++r) {
			//cout << (r + 1) <<". " << hands[r].raw <<" " << hands[r].bid << endl;
			sum += hands[r].bid * (r+1);
		}
		return sum;
	}

	int part2(const vector<string>& lines)
	{
		vector<Hand> hands;
		for (const string& line : lines) {
			vector<string> split = util::SplitString(line, " ");
			Type t = parseHandType(split[0], true);
			hands.emplace_back(split[0],t, stoi(split[1]));
		}
		ranges::sort(
			hands,
			[](const Hand& hl, const Hand& hr)
			{
				if (hl.type == hr.type)
				{
					static const vector<char> strength = { 'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J' };
					for (int i = 0; i < hl.raw.size(); ++i) {
						if (hl.raw[i] != hr.raw[i]) {
							return ranges::find(strength, hl.raw[i]) > ranges::find(strength, hr.raw[i]);
						}
					}
				}
				return hl.type > hr.type;
			}
		);

		int sum = 0;
		for (int r = 0; r < hands.size(); ++r) {
			//cout << (r + 1) << ".\t" << hands[r].raw << " " << hands[r].bid << "\t\t(" << typeToString(hands[r].type) << ")" << endl;
			sum += hands[r].bid * (r + 1);
		}
		return sum;
	}
}