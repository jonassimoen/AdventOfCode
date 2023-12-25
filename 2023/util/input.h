#ifndef UTIL_INPUT_H
#define UTIL_INPUT_H

#include <vector>
#include <string>
#include <string_view>
#include <filesystem>
#include <fstream>
#include <ranges>
#include <functional>
#include <iostream>

namespace util {
	std::vector<std::string> ReadLinesInFile(const std::filesystem::path&);

	std::vector<int> ExtractInts32(std::string_view);
	std::vector<int64_t> SplitToInt64(std::string_view, std::string_view);
	std::vector<std::string> SplitString(std::string_view, std::string_view );

	template<typename T>
	std::vector<T> ExtractInts(std::string_view input)
	{
		std::vector<T> r;

		bool parsing = false;
		T number = 0;
		for (char c : input) {
			if (isdigit(c)) {
				number = 10 * number + (c - '0');
				parsing = true;
			}
			else if (parsing) {
				r.emplace_back(number);
				number = 0;
				parsing = false;
			}
		}
		if (parsing) {
			r.emplace_back(number);
		}
		return r;
	}

	template<typename T>
	std::vector<T> SplitStringTransform(std::string_view input, std::string_view delimiter, std::function<T(std::string_view)> transformFn) {
		auto r = std::views::split(input, delimiter)
			| std::views::transform([&transformFn](const auto& subrange) { return transformFn(std::string_view{ subrange.begin(), subrange.end() }); });
		return std::vector<T>{r.begin(), r.end() };

	}
}
#endif