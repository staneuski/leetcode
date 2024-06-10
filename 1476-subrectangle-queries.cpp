/**
 * subrectangle-queries.cpp
 * 
 * https://leetcode.com/problems/subrectangle-queries
 */

class SubrectangleQueries {
public:
    SubrectangleQueries(vector<vector<int>>& rectangle) : data_(rectangle) {}

    void updateSubrectangle(int row1, int col1, int row2, int col2,
                            int newValue) {
        for (size_t row = row1; row <= row2; ++row)
            for (size_t col = col1; col <= col2; ++col)
                data_[row][col] = newValue;
    }

    int getValue(int row, int col) { return data_[row][col]; }

private:
    std::vector<std::vector<int>> data_;
};
