#include "input.h"

namespace util {

	std::vector<std::string> ReadLinesInFile(const std::filesystem::path& path) {
		auto all_lines = std::vector<std::string>{};

		if (!std::filesystem::exists(path)) {
			std::cout << "File does not exist" << std::endl;
			return all_lines;
		}

		auto fstream = std::ifstream(path);
		auto curr = std::string{};

		while (std::getline(fstream, curr)) {
			all_lines.emplace_back(curr);
		}
		return all_lines;
	}

	std::vector<std::string> SplitString(std::string_view input, std::string_view delimiter) {
		return SplitStringTransform<std::string>(input, delimiter, [](std::string_view token) { return std::string{token}; });
	}

	std::vector<int> ExtractInts32(std::string_view input) {
		return util::ExtractInts<int32_t>(input);
	}

	std::vector<int64_t> SplitToInt64(std::string_view input, std::string_view delimiter) {
		return util::SplitStringTransform<int64_t>(input, delimiter, [](std::string_view t) { return std::stoll(std::string{ t }); });
	}
}