#!/bin/bash

for i in {01..25}; do
	echo "using namespace std;

namespace Day${i} {
	int64_t part1(const vector<string>& lines)
	{
		return 0;
	}

	int64_t part2(const vector<string>& lines)
	{
		return 0;
	}
}" > day${i}.cpp
	touch input/day${i}_sample.txt
	touch input/day${i}_input.txt
done



