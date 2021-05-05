# [45. Jump Game II](https://leetcode.com/problems/jump-game-ii/)

<div><p>Given an array of non-negative integers <code>nums</code>, you are initially positioned at the first index of the array.</p>

<p>Each element in the array represents your maximum jump length at that position.</p>

<p>Your goal is to reach the last index in the minimum number of jumps.</p>

<p>You can assume that you can always reach the last index.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> nums = [2,3,1,1,4]
<strong>Output:</strong> 2
<strong>Explanation:</strong> The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> nums = [2,3,0,1,4]
<strong>Output:</strong> 2
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 1000</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1: DP
Due to the questions' constraints (nums.length <= 1000), we can simply use for-loop to jump from each position and record the minimum minimum number of jumps.
- Time Complexity: O(N)
#### C++ Code:
```cpp
class Solution {
public:
    int jump(vector<int>& nums) {
        int n = nums.size();
        vector<int> dp(n, 100001);
        dp[0] = 0;
        for (int i = 0; i < n; ++i)
            for (int j = 1; j <= nums[i] && i+j < n; ++j)
                dp[i+j] = min(dp[i+j], dp[i]+1);
        return dp[n-1];
    }
};
```