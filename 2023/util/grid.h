#ifndef UTIL_GRID_H
#define UTIL_GRID_H

namespace util {
	template<typename T>
	struct Point {
		T x = 0;
		T y = 0;

		Point operator+(const Point& other) const {
			return { x + other.x, y + other.y };
		}

		Point operator+=(const Point& other) {
			*this = *this + other;
			return *this;
		}

		Point operator-(const Point& other) const {
			return { x - other.x, y - other.y };
		}

		Point operator-=(const Point& other) {
			*this = *this - other;
			return *this;
		}

		Point operator*(const int times) const {
			return { x * times, y * times };
		}

		auto operator<=>(const Point&) const = default;
		friend std::ostream& operator<<(std::ostream& os, const Point& pt) {
			os << "(" << pt.x << ", " << pt.y << ")";
			return os;
		}
	};

	template<typename T>
	class Grid {
	public:
		Grid(int width, int height) : width(width), height(height), grid(width* height)
		{}
		Grid(int width, int height, T startvalue) : width(width), height(height), grid(width* height, startvalue)
		{}
		Grid() : Grid(0, 0)
		{}
		Grid(const Grid& g) : width(g.width), height(g.height), grid(g.grid)
		{}
		virtual ~Grid() = default;
		size_t size() const {
			return grid.size();
		}
		int getWidth() const {
			return width;
		}
		int getHeight() const {
			return height;
		}
		virtual bool inBounds(int x, int y) const {
			return x >= 0 && x < width && y >= 0 && y < height;
		}
		virtual bool inBounds(const Point<int64_t>& p) const {
			return inBounds(p.x, p.y);
		}
		virtual T& at(int x, int y) {
			return grid[y * width + x];
		}
		virtual T& at(const Point<int64_t>& p) {
			return at(p.x, p.y);
		}
		virtual const T& at(int x, int y) const {
			return grid[y * width + x];
		}
		virtual const T& at(const Point<int64_t>& p) const {
			return at(p.x, p.y);
		}
		void swap(Grid<T>& other) {
			std::swap(width, other.width);
			std::swap(height, other.height);
			std::swap(grid, other.grid);
		}

		std::vector<T>::iterator begin() noexcept {
			return grid.begin();
		}
		std::vector<T>::iterator end() noexcept {
			return grid.end();
		}
		std::vector<T>::const_iterator begin() const noexcept {
			return grid.begin();
		}
		std::vector<T>::const_iterator end() const noexcept {
			return grid.end();
		}
		std::vector<T>::const_iterator cbegin() const noexcept {
			return grid.cbegin();
		}
		std::vector<T>::const_iterator cend() const noexcept {
			return grid.cend();
		}
		std::vector<T>::reverse_iterator rbegin() noexcept {
			return grid.rbegin();
		}
		std::vector<T>::reverse_iterator rend() noexcept {
			return grid.rend();
		}
		std::vector<T>::const_reverse_iterator rbegin() const noexcept {
			return grid.rbegin();
		}
		std::vector<T>::const_reverse_iterator rend() const noexcept {
			return grid.rend();
		}
		std::vector<T>::const_reverse_iterator crbegin() const noexcept {
			return grid.crbegin();
		}
		std::vector<T>::const_reverse_iterator crend() const noexcept {
			return grid.crend();
		}

		friend std::ostream& operator<<(std::ostream& os, const Grid<T>& grid) {
			for (int y = 0; y < grid.height; ++y) {
				for (int x = 0; x < grid.width; ++x) {
					os << std::setw(1) << grid.at(x, y);
				}
				os << std::endl;
			}
			return os;
		}

	private:
		int width;
		int height;
		std::vector<T> grid;
	};

	template<typename T>
	struct Point3D {
		T x = 0;
		T y = 0;
		T z = 0;

		Point3D operator+(const Point3D& other) const {
			return { x + other.x, y + other.y, z + other.z };
		}

		Point3D operator+=(const Point3D& other) {
			*this = *this + other;
			return *this;
		}

		Point3D operator-(const Point3D& other) const {
			return { x - other.x, y - other.y, z - other.z };
		}

		Point3D operator-=(const Point3D& other) {
			*this = *this - other;
			return *this;
		}

		Point3D operator*(const int times) const {
			return { x * times, y * times, z * times };
		}

		auto operator<=>(const Point3D&) const = default;
		friend std::ostream& operator<<(std::ostream& os, const Point3D& pt) {
			os << "(" << pt.x << ", " << pt.y << ", " << pt.z << ")";
			return os;
		}
	};

	template<typename T>
	class Grid3D {
	public:
		Grid3D(int width, int height, int depth) : width(width), height(height), depth(depth), grid(width* height*depth)
		{}
		Grid3D(int width, int height, int depth, T startvalue) : width(width), height(height), depth(depth), grid(width * height * depth, startvalue)
		{}
		Grid3D() : Grid3D(0, 0, 0)
		{}
		Grid3D(const Grid3D& g) : width(g.width), height(g.height), depth(g.depth), grid(g.grid)
		{}
		virtual ~Grid3D() = default;
		size_t size() const {
			return grid.size();
		}
		int getWidth() const {
			return width;
		}
		int getHeight() const {
			return height;
		}
		int getDepth() const {
			return depth;
		}
		virtual bool inBounds(int x, int y, int z) const {
			return x >= 0 && x < width && y >= 0 && y < height && z >= 0 && z < depth;
		}
		virtual bool inBounds(const Point3D<int64_t>& p) const {
			return inBounds(p.x, p.y, p.z);
		}
		virtual T& at(int x, int y, int z) {
			return grid[(z * height + y) * width + x];
		}
		virtual T& at(const Point3D<int64_t>& p) {
			return at(p.x, p.y, p.z);
		}
		virtual const T& at(int x, int y, int z) const {
			return grid[(z * height + y) * width + x];
		}
		virtual const T& at(const Point3D<int64_t>& p) const {
			return at(p.x, p.y, p.z);
		}
		void swap(Grid3D<T>& other) {
			std::swap(width, other.width);
			std::swap(height, other.height);
			std::swap(depth, other.depth);
			std::swap(grid, other.grid);
		}

		std::vector<T>::iterator begin() noexcept {
			return grid.begin();
		}
		std::vector<T>::iterator end() noexcept {
			return grid.end();
		}
		std::vector<T>::const_iterator begin() const noexcept {
			return grid.begin();
		}
		std::vector<T>::const_iterator end() const noexcept {
			return grid.end();
		}
		std::vector<T>::const_iterator cbegin() const noexcept {
			return grid.cbegin();
		}
		std::vector<T>::const_iterator cend() const noexcept {
			return grid.cend();
		}
		std::vector<T>::reverse_iterator rbegin() noexcept {
			return grid.rbegin();
		}
		std::vector<T>::reverse_iterator rend() noexcept {
			return grid.rend();
		}
		std::vector<T>::const_reverse_iterator rbegin() const noexcept {
			return grid.rbegin();
		}
		std::vector<T>::const_reverse_iterator rend() const noexcept {
			return grid.rend();
		}
		std::vector<T>::const_reverse_iterator crbegin() const noexcept {
			return grid.crbegin();
		}
		std::vector<T>::const_reverse_iterator crend() const noexcept {
			return grid.crend();
		}

		friend std::ostream& operator<<(std::ostream& os, const Grid3D<T>& grid) {
			for (int z = 0; z < grid.depth; ++z) {
				for (int y = 0; y < grid.height; ++y) {
					for (int x = 0; x < grid.width; ++x) {
						os << std::setw(1) << grid.at(x, y, z);
					}
					os << std::endl;
				}
			}
			return os;
		}

	private:
		int width;
		int height;
		int depth;
		std::vector<T> grid;
	};
}

#endif