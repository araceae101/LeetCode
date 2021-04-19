# [377. Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/)

<div><p>Given an array of <strong>distinct</strong> integers <code>nums</code> and a target integer <code>target</code>, return <em>the number of possible combinations that add up to</em>&nbsp;<code>target</code>.</p>

<p>The answer is <strong>guaranteed</strong> to fit in a <strong>32-bit</strong> integer.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> nums = [1,2,3], target = 4
<strong>Output:</strong> 7
<strong>Explanation:</strong>
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> nums = [9], target = 3
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 200</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
	<li>All the elements of <code>nums</code> are <strong>unique</strong>.</li>
	<li><code>1 &lt;= target &lt;= 1000</code></li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow up:</strong> What if negative numbers are allowed in the given array? How does it change the problem? What limitation we need to add to the question to allow negative numbers?</p>
</div>

<p>&nbsp;</p>

## My Solutions

### Solution 1: DP
- Idea: 
    - Check the combination count starting from target number `1` to `target`.
    - For each target number, check every nums if it can be a member.
    [Schematic diagram](./0377-example.png)
- Time Complexity: O(N*target)
- Space Complexity: O(target)

```cpp
// Problem: https://leetcode.com/problems/combination-sum-iv/
// Author: Araceae
// Date: 2021/4/19

// Solution: DP

// Time Complexity: O(N*target)
// Space Complexity: O(target)

// Performance:
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum IV.
// Memory Usage: 6.6 MB, less than 94.77% of Go online submissions for Combination Sum IV.
class Solution {
public:
    int combinationSum4(vector<int>& nums, int target) {
        vector<unsigned int> dp(target+1, 0);
        dp[0] = 1;
        
        for (int t = 1; t <= target; ++t){
            for (auto &n : nums){
                if (t - n >= 0)
                    dp[t] += dp[t-n];
            }
        }
        
        return dp[target];
    }
};
```