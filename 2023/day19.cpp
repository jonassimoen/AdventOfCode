#include <string>
#include <vector>
#include <input.h>
using namespace std;

namespace Day19 {
	enum PartAttribute { Cool = 'x', Musical = 'm', Aerodynamic = 'a', Shiny = 's' };
	struct Rule {
		PartAttribute part;
		char op;
		int64_t value;
		string nextWorkflowName;

		friend ostream& operator << (ostream&, const Rule&);
	};

	struct Workflow {
		vector<Rule> rules;
		string fallbackWorkflowName;

		friend ostream& operator << (ostream&, const Workflow&);
	};

	struct Part {
		map<PartAttribute, int64_t> attributes;

		friend ostream& operator << (ostream&, const Part&);
		int64_t sum() const {
			int64_t sum = 0;
			for (const auto& [k, v] : attributes) {
				sum += v;
			}
			return sum;
		}
	};

	ostream& operator << (ostream& os, const Rule& r) {
		os << "[" << (char)r.part << r.op << r.value << "] => " << r.nextWorkflowName;
		return os;
	}

	ostream& operator << (ostream& os, const Workflow& wf) {
		for (const Rule& r : wf.rules) {
			os << "\t" << r << endl;
		}
		os << "\t===> " << wf.fallbackWorkflowName << endl;
		return os;
	}

	ostream& operator << (ostream& os, const Part& p) {
		cout << "Part:" << endl;
		for (const auto& [k, v] : p.attributes) {
			os << "\t" << (char)k << ": " << v << endl;
		}
		return os;
	}

	map<string, Workflow> parseWorkflows(const vector<string>& lines)
	{
		map<string, Workflow> wfs;
		for (const string& l : lines)
		{
			int idx = l.find("{");
			string name = l.substr(0, idx);

			vector<string> rules = util::SplitString(l.substr(idx + 1, l.size() - idx - 2), ",");

			wfs[name] = Workflow{ vector<Rule>{}, rules[rules.size() - 1] };

			for (auto i = 0; i < rules.size() - 1; ++i)
			{
				int condNameSepIdx = rules[i].find(":");
				char part = rules[i][0];
				char op = rules[i][1];
				int64_t value = stoi(rules[i].substr(2, condNameSepIdx - 2));
				string nextWorkflowName = rules[i].substr(condNameSepIdx + 1, rules[i].size() - condNameSepIdx - 1);
				wfs[name].rules.emplace_back((PartAttribute)part, op, value, nextWorkflowName);
			}
		}
		return wfs;
	}

	vector<Part> parseParts(const vector<string>& lines)
	{
		vector<Part> parts;
		for (const string& l : lines)
		{
			int idx = l.find("{");
			string name = l.substr(0, idx);

			vector<int> values = util::ExtractInts32(l.substr(idx + 1, l.size() - idx - 2));
			parts.emplace_back(map<PartAttribute, int64_t>{
				{(PartAttribute)'x', values[0]}, { (PartAttribute)'m', values[1] }, { (PartAttribute)'a', values[2] }, { (PartAttribute)'s', values[3] }
			});
		}
		return parts;
	}

	string evaluateWorkflow(const Workflow& wf, Part p)
	{
		for (int i = 0; i < wf.rules.size(); ++i) {
			Rule r = wf.rules[i];
			int partAttributeValue = p.attributes[r.part];
			if (r.op == '<' && partAttributeValue < r.value)
			{
				return r.nextWorkflowName;
			}
			else if (r.op == '>' && partAttributeValue > r.value)
			{
				return r.nextWorkflowName;
			}
		}
		return wf.fallbackWorkflowName;
	}

	int64_t part1(const vector<string>& lines)
	{
		int empytLineIdx = find(lines.begin(), lines.end(), "") - lines.begin();
		map<string, Workflow> workflows = parseWorkflows(vector<string>{lines.begin(), lines.begin() + empytLineIdx});
		vector<Part> parts = parseParts(vector<string>{lines.begin() + empytLineIdx + 1, lines.end()});

		int64_t sum = 0;
		for (const Part& p : parts)
		{
			string wfName = "in";
			while (wfName != "A" && wfName != "R")
			{
				wfName = evaluateWorkflow(workflows[wfName], p);
			}
			if (wfName == "A")
			{
				sum += p.sum();
			}
		}
		return sum;
	}

	// Parts in [low, high] should be valid for rule r
	pair<int64_t, int64_t> partAttrRangeForRule(pair<int64_t, int64_t> range, Rule r) {
		if (r.op == '<')
		{
			// High should be less then r.value (or just keep the upper bound if it is already lower)
			// E.g. if low = 5 and r.value = 7, then low should be 5
			range.second = min(range.second, r.value - 1);
		}
		else if (r.op == '>')
		{
			// Low should be greater then r.value (or just keep the lower bound if it is already higher)
			// E.g. if high = 5 and r.value = 3, then high should be 5
			range.first = max(range.first, r.value + 1);
		}
		else if (r.op == '/') // represents <=
		{
			// High should be less then or equal to r.value (or just keep the upper bound if it is already lower)
			// E.g. if low = 5 and r.value = 7, then low should be 5
			range.second = min(range.second, r.value);
		}
		else if (r.op == '~') // represents >=
		{
			// Low should be greater then or equal to r.value (or just keep the lower bound if it is already higher)
			// E.g. if high = 5 and r.value = 3, then high should be 5
			range.first = max(range.first, r.value);
		}
		else {
			cout << "Unknown operator: " << r.op << endl;
		}
		return { range.first, range.second };
	}

	struct PartsRange {
		pair<int64_t, int64_t> rangeX = { 1,4000 };
		pair<int64_t, int64_t> rangeM = { 1,4000 };
		pair<int64_t, int64_t> rangeA = { 1,4000 };
		pair<int64_t, int64_t> rangeS = { 1,4000 };
		friend ostream& operator << (ostream&, const PartsRange&);
		int64_t sum() const {
			return (rangeX.second - rangeX.first + 1) * (rangeM.second - rangeM.first + 1) * (rangeA.second - rangeA.first + 1) * (rangeS.second - rangeS.first + 1);
		}
	};

	ostream& operator << (ostream& os, const PartsRange& r) {
		cout << "X: [" << r.rangeX.first << "," << r.rangeX.second << "] (" << (r.rangeX.second - r.rangeX.first + 1) << ") ";
		cout << "M: [" << r.rangeM.first << "," << r.rangeM.second << "] (" << (r.rangeM.second - r.rangeM.first + 1) << ") ";
		cout << "A: [" << r.rangeA.first << "," << r.rangeA.second << "] (" << (r.rangeA.second - r.rangeA.first + 1) << ") ";
		cout << "S: [" << r.rangeS.first << "," << r.rangeS.second << "] (" << (r.rangeS.second - r.rangeS.first + 1) << ") ";
		return os;
	}

	PartsRange partRangesForRule(PartsRange pr, const Rule& r)
	{
		switch (r.part)
		{
		case Cool:
			pr.rangeX = partAttrRangeForRule(pr.rangeX, r);
			break;
		case Musical:
			pr.rangeM = partAttrRangeForRule(pr.rangeM, r);
			break;
		case Aerodynamic:
			pr.rangeA = partAttrRangeForRule(pr.rangeA, r);
			break;
		case Shiny:
			pr.rangeS = partAttrRangeForRule(pr.rangeS, r);
			break;
		default:
			cout << "Unknown part: " << (char)r.part << endl;
			break;
		}
		return pr;
	}

	struct DequeEntry {
		string wfName;
		PartsRange partsRange;
		friend ostream& operator << (ostream&, const DequeEntry&);
	};	
	ostream& operator << (ostream& os, const DequeEntry& de) {
		cout << de.partsRange << " ==> SUM: " << de.partsRange.sum() << endl;
		return os;
	}

	int64_t part2(const vector<string>& lines)
	{
		int64_t sum = 0;
		int empytLineIdx = find(lines.begin(), lines.end(), "") - lines.begin();
		map<string, Workflow> workflows = parseWorkflows(vector<string>{lines.begin(), lines.begin() + empytLineIdx});
		deque<DequeEntry> dq;

		dq.emplace_back("in",PartsRange{});

		while (!dq.empty()) {
			DequeEntry de = dq.back();
			dq.pop_back();

			const auto& [x, m, a, s] = de.partsRange;
			if(x.first > x.second || m.first > m.second || a.first > a.second || s.first > s.second)
				continue;

			if (de.wfName == "A")
			{
				sum += de.partsRange.sum();
				continue;
			}
			else if (de.wfName == "R")
			{
				continue;
			}
			else {
				Workflow wf = workflows[de.wfName];
				PartsRange pr = de.partsRange;

				for (const Rule& r : wf.rules)
				{
					PartsRange rangesValidForCurrentRule = partRangesForRule(pr, r);
					dq.emplace_back(r.nextWorkflowName, rangesValidForCurrentRule);
					// The next rule should NOT be following the current rule
					pr = partRangesForRule(pr, { r.part, (r.op == '>' ? '/' : '~'), r.value, r.nextWorkflowName });

					
				}
				dq.emplace_back(wf.fallbackWorkflowName, pr);
			}
		}

		return sum;
	}
}
