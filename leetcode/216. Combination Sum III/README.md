# [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)

<p>Find all valid combinations of <code>k</code> numbers that sum up to <code>n</code> such that the following conditions are true:</p>

<ul>
	<li>Only numbers <code>1</code> through <code>9</code> are used.</li>
	<li>Each number is used <strong>at most once</strong>.</li>
</ul>

<p>Return <em>a list of all possible valid combinations</em>. The list must not contain the same combination twice, and the combinations may be returned in any order.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> k = 3, n = 7
<strong>Output:</strong> [[1,2,4]]
<strong>Explanation:</strong>
1 + 2 + 4 = 7
There are no other valid combinations.</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> k = 3, n = 9
<strong>Output:</strong> [[1,2,6],[1,3,5],[2,3,4]]
<strong>Explanation:</strong>
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> k = 4, n = 1
<strong>Output:</strong> []
<strong>Explanation:</strong> There are no valid combinations. [1,2,1] is not valid because 1 is used twice.
</pre>

<p><strong>Example 4:</strong></p>

<pre><strong>Input:</strong> k = 3, n = 2
<strong>Output:</strong> []
<strong>Explanation:</strong> There are no valid combinations.
</pre>

<p><strong>Example 5:</strong></p>

<pre><strong>Input:</strong> k = 9, n = 45
<strong>Output:</strong> [[1,2,3,4,5,6,7,8,9]]
<strong>Explanation:</strong>
1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45
​​​​​​​There are no other valid combinations.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= k &lt;= 9</code></li>
	<li><code>1 &lt;= n &lt;= 60</code></li>
</ul>


## My Solutions
### Solution 1. Recurcive
#### C++
- Time complexity: Best is O(1). Best case is `1` when n > 45; Worst case is `C(9, 4) = 126` when k = 4.
- Space complexity: Best is O(1)

```cpp
class Solution {
    int nums[9] = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int len = 9;
    int sum = 45;
    
    void findAllCombination(int k, int n, int start, vector<vector<int>>& res, vector<int> tmp){
        if (k == 0 && n == 0){
            if (!tmp.empty())   res.push_back(tmp);
            return;
        }
        for (int i = start; i < len; ++i){
            if (n < nums[i])    break;
            else if (n >= nums[i]){
                tmp.push_back(nums[i]);
                findAllCombination(k-1, n-nums[i], i+1, res, tmp);
                tmp.pop_back();
            }
        }
    }
    
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        if (n > sum) return {};
        
        vector<vector<int>> res;
        findAllCombination(k, n, 0, res, {});
        return res;
    }
};
```