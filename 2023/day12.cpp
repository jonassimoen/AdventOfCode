#include <string>
#include <vector>
#include <input.h>
#include <unordered_map>
#include <map>
using namespace std;

namespace Day12 {

	void printTree(const string_view& conditionRecord)
	{
		ofstream outf(format("{}/day12.tgf", ASSETS_FOLDER));
		map<string, int> mapping;
		vector<vector<string>> tree;
		tree.emplace_back(vector<string>{});
		tree[0].emplace_back(conditionRecord);
		int nodeNr = 0;
		mapping.insert({ string(conditionRecord), nodeNr++ });
		for (int i = 1; i <= ranges::count(conditionRecord, '?'); ++i)
		{
			tree.emplace_back(vector<string>{});
			for (auto& s : tree[i - 1])
			{
				string edited = s;
				int idx = edited.find_first_of('?');
				edited[idx] = '#';
				tree[i].emplace_back(edited);
				mapping.insert({ edited, nodeNr++ });
				//cout << edited << ": " << nodeNr << endl;
				edited[idx] = '.';
				tree[i].emplace_back(edited);
				mapping.insert({ edited, nodeNr++ });
				//cout << edited  << ": " << nodeNr << endl;
			}
		}
		cout << mapping.size() << endl;
		for (auto& [s, i] : mapping) {
			outf << i << " " << s << endl;
		}
		outf << "#" << endl;
		for (int l = 0; l < tree.size() - 1; l++) {
			for (int sl = 0; sl < tree[l].size(); sl++) {
				outf << mapping[tree[l][sl]] << " " << mapping[tree[l + 1][sl * 2]] << endl;
				outf << mapping[tree[l][sl]] << " " << mapping[tree[l + 1][sl * 2 + 1]] << endl;
			}
		}
	}

	bool isValidRecord(const string_view& record, const vector<int>& chunkSizes, bool allowWildCards = false)
	{
		int64_t chunkIdx = 0;
		int64_t currentChunkSize = 0;
		if (record.empty()) { return true; }
		if (allowWildCards || record.contains('?')) { return false; }
		if (chunkSizes.size() == 0 && !record.contains('#')) { return true; }

		auto hashChunks = record | views::chunk_by([](char a, char b) { return a == '#' && b == '#'; }) | views::filter([](auto chunk) {return string_view(chunk).contains('#'); });

		if (ranges::distance(hashChunks) != chunkSizes.size()) { return false; }

		for (auto const subrange : hashChunks)
		{
			if (ranges::distance(subrange) != chunkSizes[chunkIdx++]) { return false; }
		}

		for (auto chunk : record | views::chunk_by([](char a, char b) { return a == '#' && b == '#'; }))
		{
			cout << string_view(chunk) << " ";
		}
		cout << "\t";
		for (auto chunkLen : chunkSizes)
		{
			cout << chunkLen << ",";
		}
		cout << "\t" << "VALID";
		cout << endl;
		return true;
	}

	struct Record
	{
		int currentIdx;
		int currentChunk;
		int currentHashChunkSize;
	};
	struct RecordHash {
		auto operator()(const Record& record) const {
			return record.currentIdx * 10000 + record.currentChunk * 100 + record.currentHashChunkSize;
		}
	};

	struct RecordEqual {
		auto operator()(const Record& rec1, const Record& rec2) const {
			return rec1.currentIdx == rec2.currentIdx && rec1.currentChunk == rec2.currentChunk && rec1.currentHashChunkSize == rec2.currentHashChunkSize;
		}
	};

	static unordered_map<Record, int64_t, RecordHash, RecordEqual> cache = {};
	int64_t processRecord(const string_view& record, const vector<int>& chunkSizes, int64_t currentSpring = 0, int64_t currentChunk = 0, int64_t currentHashChunkSize = 0)
	{
		Record r{ currentSpring, currentChunk, currentHashChunkSize };

		if (cache.contains(r)) { return cache[r]; }
		if (currentSpring == record.size()) {
			if (currentChunk == chunkSizes.size() && currentHashChunkSize == 0) { return 1; }
			if (currentChunk == chunkSizes.size() - 1 && currentHashChunkSize == chunkSizes[currentChunk]) { return 1; }
			else { return 0; }
		}


		int64_t res = 0;
		for (char c : vector<char>{ '.', '#' }) {
			if (record[currentSpring] == c or record[currentSpring] == '?') {
				if (c == '.' and currentHashChunkSize == 0)
					res += processRecord(record, chunkSizes, currentSpring + 1, currentChunk, 0);
				else if (c == '.' and currentHashChunkSize > 0 and currentChunk < chunkSizes.size() and currentHashChunkSize == chunkSizes[currentChunk])
					res += processRecord(record, chunkSizes, currentSpring + 1, currentChunk + 1, 0);
				else if (c == '#')
					res += processRecord(record, chunkSizes, currentSpring + 1, currentChunk, currentHashChunkSize + 1);
			}
		}
		cache.insert({ r, res });
		return res;
	}

	// Currently not working
	int64_t processRecord2(const string_view& record, const vector<int>& chunkSizes, int level = 0, string prefix = "") {
		//cout << level << ". [" << prefix << "_" << record << "] [";
		//for (auto c : chunkSizes) {
		//	cout << c << ",";
		//}
		//cout << "]" << endl;

		int sumChunkSize = std::accumulate(chunkSizes.begin(), chunkSizes.end(), 0);
		int nrChunksToProcess = chunkSizes.size();

		if (record.empty()) {
			// No more strings to process => check if all chunks are processed
			if (nrChunksToProcess == 0) { /*cout << "\tVALID: " << prefix << record << endl;*/ return 1; }
			else { return 0; }
		}
		if (record.size() < sumChunkSize) {
			// Not enough strings to achieve the chunk sizes => invalid
			return 0;
		}
		else if (record.size() == sumChunkSize && !record.contains('.') && nrChunksToProcess == 1) {
			// Only 1 chunk to process and this is possible (no . in the string) => valid
			//cout << "\tVALID: " << prefix << record << endl;
			return 1;
		}
		else if (nrChunksToProcess == 0) {
			// No chunks to process => check if # in the string
			if (record.contains('#')) { return 0; }
			else { /*cout << "\tVALID: " << prefix << record << endl;*/ return 1; }
		} 

		int chunkSize = chunkSizes[0];

		if (record.at(0) == '.') {
			//cout << "\t[.] Proceed" << endl;
			return processRecord2(record.substr(1), chunkSizes, level + 1, prefix + ".");
		}
		else if (record.at(0) == '?') {
			string_view chunk = record.substr(0, chunkSize);
			//cout << "\t[?] Check possible chunk: " << chunk << endl;
			if (record.size() < chunkSize || chunk.contains('.') || record.at(chunkSize) == '#') {
				//cout << "\t[?] Chunk invalid => only check ." << endl;
				return processRecord2(record.substr(1), chunkSizes, level + 1, prefix + ".");
			}
			else {
				//cout << "\t[?] Chunk valid => check . & #" << endl;
				return processRecord2(record.substr(1), chunkSizes, level + 1, prefix + ".") + processRecord2(record.substr(chunkSize + 1), vector<int>(chunkSizes.begin() + 1, chunkSizes.end()), level + 1, prefix + string(chunkSize, '#') + '.');
			}
		}
		else if (record.at(0) == '#') {
			//cout << "\t[#] Check possible chunk" << endl;
			string_view chunk = record.substr(0, chunkSize);
			int nonHashCharIdx = record.find_first_not_of('#');
			if (nonHashCharIdx != string::npos && nonHashCharIdx <= chunkSize && !chunk.contains('.')) {
				//cout << "\t[#] Chunk valid => check after chunk" << endl;
				return processRecord2(record.substr(chunkSize + 1), vector<int>(chunkSizes.begin() + 1, chunkSizes.end()), level + 1, prefix + string(chunkSize, '#') + '.');
			}
		}
		return 0;
	}

	int64_t part1(const vector<string>& lines)
	{
		int64_t result = 0;
		for (const auto& line : lines)
		{
			cache.clear();
			vector<string> parsedLine = util::SplitString(line, " ");
			int64_t res = processRecord(parsedLine[0], util::ExtractInts32(parsedLine[1]));
			result += res;
		}
		return result;
	}

	int64_t part2(const vector<string>& lines)
	{
		int64_t result = 0;
		for (const auto& line : lines)
		{
			cache.clear();
			vector<string> parsedLine = util::SplitString(line, " ");
			string record = parsedLine[0];
			string chunkSizes = parsedLine[1];

			for (int i = 0; i < 4; ++i) {
				record += '?' + parsedLine[0];
				chunkSizes += ',' + parsedLine[1];
			}

			int64_t res = processRecord(record, util::ExtractInts32(chunkSizes));
			result += res;
		}
		return result;
	}
}
