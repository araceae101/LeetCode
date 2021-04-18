# [1151. Minimum Swaps to Group All 1's Together](https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together/)

<div><p>Given a&nbsp;binary array <code>data</code>, return&nbsp;the minimum number of swaps required to group all <code>1</code>’s present in the array together in <strong>any place</strong> in the array.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> data = [1,0,1,0,1]
<strong>Output:</strong> 1
<strong>Explanation: </strong>
There are 3 ways to group all 1's together:
[1,1,1,0,0] using 1 swap.
[0,1,1,1,0] using 2 swaps.
[0,0,1,1,1] using 1 swap.
The minimum is 1.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> data = [0,0,0,1,0]
<strong>Output:</strong> 0
<strong>Explanation: </strong>
Since there is only one 1 in the array, no swaps needed.
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> data = [1,0,1,0,1,0,0,1,1,0,1]
<strong>Output:</strong> 3
<strong>Explanation: </strong>
One possible solution that uses 3 swaps is [0,0,0,0,0,1,1,1,1,1,1].
</pre>

<p><strong>Example 4:</strong></p>

<pre><strong>Input:</strong> data = [1,0,1,0,1,0,1,1,1,0,1,0,0,1,1,1,0,0,1,1,1,0,1,0,1,1,0,0,0,1,1,1,1,0,0,1]
<strong>Output:</strong> 8
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= data.length &lt;= 10<sup>5</sup></code></li>
	<li><code>data[i]</code>&nbsp;is <code>0</code> or <code>1</code>.</li>
</ul>
</div>

<p>&nbsp;</p>


## My Solutions
### Solution 1: Sliding window
- Idea: 
    1. Count how many `1` in the `data` and set this number as the size of sliding window.
    2. And we just need to count how many `0` in the sliding window.
    3. Shift the sliding window and choose the sliding window with the minimum `0` count.
- Performance
    - Time Complexity: O(N)
    - Space Complexity: O(1)
#### C++ Code
```cpp
// Problem: https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together
// Author: Araceae
// Date: 2021/4/18

// Solution: Sliding window

// Time Complexity: O(N)
// Space Complexity: O(1)

// Performance: 
// Runtime: 56 ms, faster than 99.33% of C++ online submissions for Minimum Swaps to Group All 1's Together.
// Memory Usage: 68.6 MB, less than 95.78% of C++ online submissions for Minimum Swaps to Group All 1's Together.

class Solution {
public:
    int minSwaps(vector<int>& data) {
        int n = data.size();
        int window = accumulate(data.begin(), data.end(), 0);
        
        if (window == 1 || window == n) return 0;
        
        int count_One = accumulate(data.begin(), data.begin()+window, 0);
        int count_Zero = window - count_One;
        
        int left = 0, right = window;
        int res = count_Zero;
        
        while (right < n){
            count_Zero -= !data[left++];
            count_Zero += !data[right++];
            if (count_Zero < res)   res = count_Zero;
        }
        
        return res;
    }
};
```