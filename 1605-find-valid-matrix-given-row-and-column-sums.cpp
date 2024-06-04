class Solution {
public:
    vector<vector<int>> restoreMatrix(vector<int>& row_sum,
                                      vector<int>& col_sum) {
        vector<vector<int>> res(row_sum.size(),
                                vector<int>(col_sum.size(), {}));
        for (size_t i = 0; i < row_sum.size(); ++i) {
            for (size_t j = 0; j < col_sum.size(); ++j) {
                res[i][j] = min(row_sum[i], col_sum[j]);
                row_sum[i] -= res[i][j];
                col_sum[j] -= res[i][j];
            }
        }
        return res;
    }
};
