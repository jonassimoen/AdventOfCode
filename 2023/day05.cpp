using namespace std;

namespace Day05 {
	struct MapData
	{
		int64_t src, dest, len;
	};

	pair<vector<int64_t>, vector<vector<MapData>>> parseInput(const vector<string>&lines) {
		vector<int64_t> seeds = util::SplitToInt64(lines[0].substr(lines[0].find(':') + 2), " ");

		vector<vector<MapData>> maps;
		maps.resize(7);
		int mapNr = 0;

		for (int l = 2; l < lines.size(); l++) {
			if (lines[l].empty()) {
				// mapping done
				++mapNr;
				continue;
			}
			if (!isdigit(lines[l][0])) {
				// mapping title
				continue;
			}

			vector<int64_t> lineDataSplit = util::SplitToInt64(lines[l], " ");
			maps[mapNr].emplace_back(MapData{ lineDataSplit[1], lineDataSplit[0], lineDataSplit[2] });
		}
		return make_pair(seeds, maps);
	}

	int part1(const vector<string>&lines)
	{
		auto [seeds, maps] = parseInput(lines);
		int64_t minimalLocation = numeric_limits<int64_t>::max();
		for (int64_t seed : seeds) {
			for (const vector<MapData>& m : maps) {
				for (const MapData& md : m) {
					if (md.src <= seed && seed < md.src + md.len) {
						seed = md.dest + (seed - md.src);
						break;
					}
				}
			}
			minimalLocation = min(minimalLocation, seed);
		}
		return minimalLocation;
	}

	vector<MapData> merge(vector<vector<MapData>> maps) {
		// make 1 new map covering the whole seed number space
		for (vector<MapData>& map : maps) {
			ranges::sort(map, {}, &MapData::src);
			vector<MapData> newEntries = {};

			int64_t i = 0;
			for (const MapData& md : map) {
				// Numbers keep same between current position & next entry
				if (i < md.src) {
					newEntries.emplace_back(i, i, md.src - i);
					i += md.src - i;
				}
				newEntries.emplace_back(md);
				i += md.len;
			}

			// Other numbers should be filled to.
			newEntries.emplace_back(i, i, std::numeric_limits<int64_t>::max() - i);
			map.swap(newEntries);
		}

		while (maps.size() > 1) {
			vector<MapData> r = maps.back();
			maps.pop_back();
			vector<MapData> l = maps.back();
			maps.pop_back();

			vector<MapData> merged;
			// C++23: stack<MapData> toProcess(l.rbegin(), l.rend());
			stack<MapData> toProcess;
			for (auto it = l.rbegin(); it != l.rend(); ++it) {
				toProcess.push(*it);
			}

			while (!toProcess.empty()) {
				const MapData& md = toProcess.top();
				toProcess.pop();

				// Find first intersection: right src <-> left dest
				int iIS = 0;
				while ((md.dest < r[iIS].src) || r[iIS].src + r[iIS].len <= md.dest) {
					++iIS;
				}

				const MapData& is = r[iIS];
				int64_t offset = md.dest - is.src;

				// Left entry fully inside range Right entry
				if (md.dest + md.len <= is.src + is.len) {
					// Remapping to new destination:
					merged.emplace_back(md.src, is.dest + offset, md.len);
					continue;
				}

				// Left entry partially in range Right entry, chop part of & process it
				int64_t choppedLength = is.len - offset;
				merged.emplace_back(md.src, is.dest + offset, choppedLength);

				// Other part: process later
				toProcess.emplace(md.src + choppedLength, md.dest + choppedLength, md.len - choppedLength);
			}
			maps.push_back(merged);
		}

		return maps[0];
	}

	int part2(const vector<string>&lines)
	{
		auto [seeds, maps] = parseInput(lines);
		vector<MapData> generalMap = merge(maps);
		int64_t minimalLocation = numeric_limits<int64_t>::max();

		for (int seedNumber = 0; seedNumber < seeds.size(); seedNumber += 2) {
			int64_t minLocPerSeedRange = numeric_limits<int64_t>::max();

			for (const MapData& md : generalMap) {
				int64_t seedStart = seeds[seedNumber];
				int64_t seedEnd = seeds[seedNumber] + seeds[seedNumber + 1];
				int64_t mdStart = md.src;
				int64_t mdEnd = md.src+md.len;

				int64_t overlapSeed = -1;
				if (seedStart <= mdStart && mdStart < seedEnd) {
					overlapSeed = mdStart;
				}
				else if (mdStart <= seedStart && seedStart < mdEnd) {
					overlapSeed = seedStart;
				}
				else {
					continue;
				}

				minLocPerSeedRange = min(minLocPerSeedRange, md.dest + (overlapSeed - mdStart));
			}
			minimalLocation = min(minimalLocation, minLocPerSeedRange);
		}
		return minimalLocation;
	}
}