/**
 * 341-flatten-nested-list-iterator.cpp
 *
 * https://leetcode.com/problems/flatten-nested-list-iterator/description/
 */

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *   public:
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool isInteger() const;
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     int getInteger() const;
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     const vector<NestedInteger> &getList() const;
 * };
 */

class NestedIterator {
    typedef vector<NestedInteger>::iterator Iterator;
    typedef vector<NestedInteger>::const_iterator ConstIterator;

    struct Node {
        Iterator it;
        ConstIterator end;

        inline bool hasNext() const {
            return it != end;
        }
    };

public:
    NestedIterator(vector<NestedInteger>& data) {
        nodes_.push({data.begin(), data.end()});
    }

    int next() {
        hasNext();
        return (nodes_.top().it++)->getInteger();
    }

    bool hasNext() {
        while (!nodes_.empty()) {
            if (!nodes_.top().hasNext()) {
                nodes_.pop();
            } else if (const auto node = nodes_.top(); !node.it->isInteger()) {
                ++nodes_.top().it;
                nodes_.push({
                    node.it->getList().begin(),
                    node.it->getList().end()
                });
            } else {
                return true;
            }
        }

        return false;
    }

private:
    stack<Node> nodes_;
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */
