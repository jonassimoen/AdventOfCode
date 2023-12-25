#include <string>
#include <vector>
#include <input.h>
#include <queue>
using namespace std;

namespace Day20 {
	enum class ModuleType { FlipFlop = '%', Conjunction = '&', Broadcaster = 'b', Output = 'o' };
	enum class Pulse { High, Low };

	class Module;

	struct QueueItem {
		Module* module;
		Module* source;
		Pulse pulse;
	};

	class Module {
	public:
		string name;
		Module() : name(""), type(ModuleType::Output), destinations() {}
		Module(const Module&) = default;
		Module(string name, ModuleType type) : name(name), type(type), destinations() {}
		virtual void sendPulse(Module* source, Pulse pulse, queue<QueueItem>& q) {
			//string pulseStr = (pulse == Pulse::High ? "high" : "low");
			//cout << source->name << " -" << pulseStr << "-> " << name << endl;
		};
		void addDestination(Module* destination) {
			destinations.push_back(destination);
		}    
		Module& operator=(const Module&) = default;
		bool operator==(const Module& other) const {
			return name == other.name;
		}
		ModuleType getType() const {
			return type;
		}
		bool hasDestination(string destination) const {
			return ranges::any_of(destinations, [destination](Module* mod) { return mod->name == destination; });
		}
	protected:
		vector<Module*> destinations;
	private:
		ModuleType type;
	};

	class FlipFlopModule : public Module{
	public:
		FlipFlopModule(string name) : Module(name, ModuleType::FlipFlop) {}
		void sendPulse(Module* source, Pulse pulse, queue<QueueItem>& q) override {
			Module::sendPulse(source, pulse, q);
			if (pulse == Pulse::Low) {
				if (state) {
					for(auto& destination : destinations)
						q.emplace(destination, this, Pulse::Low);
				}
				else {
					for (auto& destination : destinations)
						q.emplace(destination, this, Pulse::High);
				}
				state = !state;
			}
		}
	private:
		bool state = false;
	};

	class ConjunctionModule : public Module {
	public:
		ConjunctionModule(string name) : Module(name, ModuleType::Conjunction) {}
		void setInputs(vector<string> inputs) {
			for(int i=0; i<inputs.size(); ++i)
				history[inputs[i]] = Pulse::Low;
		}
		void sendPulse(Module* source, Pulse pulse, queue<QueueItem>& q) override {
			Module::sendPulse(source, pulse, q);
			history[source->name] = pulse;
			Pulse result = Pulse::High;
			if (ranges::all_of(history, [](auto& pair) { return pair.second == Pulse::High; })) {
				result = Pulse::Low;
			}
			for (auto& destination : destinations)
				q.emplace(destination, this, result);
		}
	private:
		unordered_map<string, Pulse> history;
	};

	class BroadcasterModule : public Module {
	public:
		BroadcasterModule(string name) : Module(name, ModuleType::Broadcaster) {}
		void sendPulse(Module* source, Pulse pulse, queue<QueueItem>& q) override {
			Module::sendPulse(source, pulse, q);
			for (auto& destination : destinations) 
				q.emplace(destination, this, pulse);
		}
	};

	unordered_map<string, Module*> parseModules(const vector<string>& lines)
	{
		unordered_map<string, Module*> res;
		unordered_map<string, vector<string>> neighborsPerModule;
		unordered_map<string, vector<string>> inputsPerModule;
		ofstream outf(format("{}/day20.dot", ASSETS_FOLDER));
		for (const string& line : lines)
		{
			vector<string> tokens = util::SplitString(line, " -> ");
			string name = tokens[0].substr(1);
			if ((ModuleType)tokens[0][0] == ModuleType::FlipFlop) 
				res.emplace(name, new FlipFlopModule(name));
			else if ((ModuleType)tokens[0][0] == ModuleType::Conjunction)
				res.emplace(name, new ConjunctionModule(name));
			else if ((ModuleType)tokens[0][0] == ModuleType::Broadcaster) {
				name = tokens[0];
				res.emplace(name, new BroadcasterModule(name));
			}
			neighborsPerModule.emplace(name, util::SplitString(tokens[1], ", "));
		}

		for (const auto& [name, neighbors] : neighborsPerModule) {
			for (const string& neighbor : neighbors) {
				outf << name << " -> " << neighbor << endl;
				if (res[neighbor] == nullptr)
					res[neighbor] = new Module(neighbor, ModuleType::Output);
				res[name]->addDestination(res[neighbor]);
				inputsPerModule[neighbor].push_back(name);
			}
		}		
		for (const auto& [name, module] : res) {
			if(module != nullptr && module->getType() == ModuleType::Conjunction)
				dynamic_cast<ConjunctionModule*>(module)->setInputs(inputsPerModule[name]);
		}

		return res;
	}

	int64_t part1(const vector<string>& lines)
	{
		unordered_map<string, Module*> modules = parseModules(lines);

		Module* btn = new Module { "btn", ModuleType::FlipFlop };
		queue<QueueItem> queue;

		const int MAX_BTN_PRESSES = 1000; 
		int highFinal = 0, lowFinal = 0;
		for (int i = 0; i < MAX_BTN_PRESSES; ++i) {
			int high = 0, low = 0;
			queue.push({ modules["broadcaster"], btn, Pulse::Low });
			while (!queue.empty()) {
				QueueItem item = queue.front();
				queue.pop();
				if (item.pulse == Pulse::High)
					++high;
				else
					++low;
				item.module->sendPulse(item.source, item.pulse, queue);
			}
			highFinal += high;
			lowFinal += low;
		}
		delete btn;
		for(const auto& [name, module] : modules)
			delete module;
		return highFinal * lowFinal;
	}

	int64_t part2(const vector<string>& lines)
	{
		unordered_map<string, Module*> modules = parseModules(lines);
		if (!modules.contains("rx")) {
			return 0;
		}

		Module* btn = new Module{ "btn", ModuleType::FlipFlop };
		
		// The following nodes we should be watching: parents of 'dt'
		set<string> nodesToWatch = modules
			| views::values
			| views::filter([](Module* m) { return m->hasDestination("dt"); })
			| views::transform([](Module* m) { return m->name; })
			| ranges::to<set<string>>();
		// We keep track of the cycles: when the last time (aka which timestamp) we encountered this parent of 'dt' receiving a LOW pulse
		unordered_map<string, int64_t> history;
		// Keep track of the lengths of the cycles of the parents of 'dt' (should be vector of length 4 in the end)
		vector<int64_t> cycles;

		queue<QueueItem> queue;

		/* 
			After investigation of input:
			- &dt -> rx (Conjunction), so if inputs of 'dt' are all high, output to 'rx' will be low
			- &x, &y -> &dt (Conjunction), so if inputs of 'x' and 'y' are all low, so all inputs of 'dt' will be high
			- We should know when 'x' and 'y' are receiving LOW inputs (should be a cycle) and 'df' should receive HIGH at same moment of all
			- Result is LCM of those cycles
		*/
		for (int i = 0; i < 10000; ++i)
		{
			queue.push({ modules["broadcaster"], btn, Pulse::Low });
			while (!queue.empty()) {
				QueueItem item = queue.front();
				queue.pop();

				if (item.pulse == Pulse::Low) {
					// We encounter parent of 'dt' receiving a LOW pulse, we already have encountered this one twice, so we can calculate the cycle
					if (nodesToWatch.contains(item.module->name) and history.contains(item.module->name)) {
						cycles.push_back(i - history.at(item.module->name));
					}
					history[item.module->name] = i;
				}

				item.module->sendPulse(item.source, item.pulse, queue);
			}
		}

		for (const auto& [name, module] : modules)
			delete module;

		// 'dt' has 4 parents, so we should have 4 cycle lengths
		assert(cycles.size() == 4);

		int64_t lcm_ = cycles[0];
		for(const int64_t& cycle : cycles)
			lcm_ = lcm(lcm_, cycle);
		return lcm_;
	}
}
