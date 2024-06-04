/**
 * 1672-richest-customer-wealth.cpp
 *
 * https://leetcode.com/problems/richest-customer-wealth/description/
 */

class Solution {
public:
    int maximumWealth(vector<vector<int>>& accounts) {
        int max_wealth = 0;
        for (const auto& account : accounts) {
            const int wealth = std::accumulate(account.begin(), account.end(), 0);
            max_wealth = (wealth > max_wealth) ? wealth : max_wealth;
        }

        return max_wealth;
    }
};
