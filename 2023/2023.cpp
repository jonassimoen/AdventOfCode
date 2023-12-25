#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <algorithm>
#include <functional>
#include <sstream>
#include <cstring>
#include <string_view>
#include <ranges>
#include <set>
#include <stack>
#include <assert.h>

#include "util/input.h"
#include "util/grid.h"

#include "day01.cpp"
#include "day02.cpp"
#include "day03.cpp"
#include "day04.cpp"
#include "day05.cpp"
#include "day06.cpp"
#include "day07.cpp"
#include "day08.cpp"
#include "day09.cpp"
#include "day10.cpp"
#include "day11.cpp"
#include "day12.cpp"
#include "day13.cpp"
#include "day14.cpp"
#include "day15.cpp"
#include "day16.cpp"
#include "day17.cpp"
#include "day18.cpp"
#include "day19.cpp"
#include "day20.cpp"
#include "day21.cpp"
#include "day22.cpp"
#include "day23.cpp"
#include "day24.cpp"
#include "day25.cpp"

std::pair<int, int> getStartEndDay() {
	time_t now = time(0);
	tm* ltm = localtime(&now);
	if (ltm->tm_mon == 11 && ltm->tm_mday < 25)
		return { ltm->tm_mday, ltm->tm_mday };
	return { 1, 25 };
}

const std::vector<std::function<int64_t(const std::vector<std::string>&)>> Part1 = {
	Day01::part1, Day02::part1, Day03::part1, Day04::part1, Day05::part1, Day06::part1, Day07::part1, Day08::part1, Day09::part1, Day10::part1, Day11::part1, Day12::part1, Day13::part1, Day14::part1, Day15::part1, Day16::part1, Day17::part1, Day18::part1, Day19::part1, Day20::part1, Day21::part1, Day22::part1, Day23::part1, Day24::part1, Day25::part1
};
const std::vector<std::function<int64_t(const std::vector<std::string>&)>> Part2 = {
	Day01::part2, Day02::part2, Day03::part2, Day04::part2, Day05::part2, Day06::part2, Day07::part2, Day08::part2, Day09::part2, Day10::part2, Day11::part2, Day12::part2, Day13::part2, Day14::part2, Day15::part2, Day16::part2, Day17::part2, Day18::part2, Day19::part2, Day20::part2, Day21::part2, Day22::part2, Day23::part2, Day24::part2, Day25::part2
};

//		day			sample/test  # files (often a single one)
const std::vector < std::vector <std::vector<std::string>>> input = {
	{{"day01_sample.txt"}, {"day01_input.txt"}},
	{{"day02_sample.txt"}, {"day02_input.txt"}},
	{{"day03_sample.txt"}, {"day03_input.txt"}},
	{{"day04_sample.txt"}, {"day04_input.txt"}},
	{{"day05_sample.txt"}, {"day05_input.txt"}},
	{{"day06_sample.txt"}, {"day06_input.txt"}},
	{{"day07_sample.txt"}, {"day07_input.txt"}},
	{{"day08_sample.txt","day08_sample2.txt"}, {"day08_input.txt"}},
	{{"day09_sample.txt"}, {"day09_input.txt"}},
	{{"day10_sample.txt","day10_sample2.txt","day10_sample3.txt","day10_sample4.txt","day10_sample5.txt"}, {"day10_input.txt"}},
	{{"day11_sample.txt"}, {"day11_input.txt"}},
	{{"day12_sample.txt"}, {"day12_input.txt"}},
	{{"day13_sample.txt"}, {"day13_input.txt"}},
	{{"day14_sample.txt"}, {"day14_input.txt"}},
	{{"day15_sample.txt"}, {"day15_input.txt"}},
	{{"day16_sample.txt"}, {"day16_input.txt"}},
	{{"day17_sample.txt"}, {"day17_input.txt"}},
	{{"day18_sample.txt"}, {"day18_input.txt"}},
	{{"day19_sample.txt"}, {"day19_input.txt"}},
	{{"day20_sample1.txt","day20_sample2.txt"}, {"day20_input.txt"}},
	{{/*"day21_sample.txt",*/ "day21_sample2.txt"}, {"day21_input.txt"}},
	{{"day22_sample.txt"}, {"day22_input.txt"}},
	{{"day23_sample.txt"}, {"day23_input.txt"}},
	{{"day24_sample.txt"}, {"day24_input.txt"}},
	{{"day25_sample.txt"}, {"day25_input.txt"}},
};

int64_t solution(string filename, std::function<int64_t(const std::vector<std::string>&)> func)
{
	filesystem::path input = filesystem::path(INPUT_FOLDER) / filename;
	vector<string> lines = util::ReadLinesInFile(input);

	if (func != nullptr)
	{
		return func(lines);
	}
	else
	{
		std::cout << "\x1B[31mFunction not found." << std::endl;
		return -1;
	}
}

int main(int argc, char** argv) {
	if (argc < 2) {
		return -1;
	}

	bool useTestInput = strcmp(argv[1], "-sample");
	bool useSampleInput = strcmp(argv[1], "-test");
	auto [START_DAY, END_DAY] = (argc >= 3) ? std::pair<int, int>(atoi(argv[2]), atoi(argv[2])) : getStartEndDay();

	std::cout <<
		left << setw(4) << "DAY" << " | " <<
		left << setw(20) << "FILE" << " | " <<
		left << setw(20) << "RESULT P1" << " | " <<
		left << setw(10) << "Tp1 [ms]" << " | " <<
		left << setw(20) << "RESULT P2" << " | " <<
		left << setw(10) << "Tp2 [ms]" << std::endl;
	std::cout << std::string(100, '-') << std::endl;

	for (int i = START_DAY; i <= END_DAY; ++i) {

		if (useSampleInput) {
			for (const string& sampleInput : input[i - 1][0]) {
				auto start_time = std::chrono::steady_clock::now();

				int64_t res1 = solution(sampleInput, Part1[i - 1]);
				auto stop_time1 = std::chrono::steady_clock::now();
				int64_t res2 = solution(sampleInput, Part2[i - 1]);
				auto stop_time2 = std::chrono::steady_clock::now();

				std::cout
					<< "\x1B[32m"
					<< left << setw(4) << i << " | "
					<< left << setw(20) << sampleInput << " | "
					<< left << setw(20) << res1 << " | " << left << setw(10) << std::chrono::duration_cast<std::chrono::milliseconds>(stop_time1 - start_time).count() << " | "
					<< left << setw(20) << res2 << " | " << left << setw(10) << std::chrono::duration_cast<std::chrono::milliseconds>(stop_time2 - stop_time1).count()
					<< std::endl;
			}
		}

		if (useTestInput) {
			for (const string& testInput : input[i - 1][1]) {
				auto start_time = std::chrono::steady_clock::now();

				int64_t res1 = solution(testInput, Part1[i - 1]);
				auto stop_time1 = std::chrono::steady_clock::now();
				int64_t res2 = solution(testInput, Part2[i - 1]);
				auto stop_time2 = std::chrono::steady_clock::now();

				std::cout
					<< "\x1B[36m"
					<< left << setw(4) << i << " | "
					<< left << setw(20) << testInput << " | "
					<< left << setw(20) << res1 << " | " << left << setw(10) << std::chrono::duration_cast<std::chrono::milliseconds>(stop_time1 - start_time).count() << " | "
					<< left << setw(20) << res2 << " | " << left << setw(10) << std::chrono::duration_cast<std::chrono::milliseconds>(stop_time2 - stop_time1).count()
					<< std::endl;
			}
		}
		std::cout << "\x1B[0m";
	}
	return 0;
}