# [120. Triangle](https://leetcode.com/problems/triangle/)

<div><p>Given a <code>triangle</code> array, return <em>the minimum path sum from top to bottom</em>.</p>

<p>For each step, you may move to an adjacent number of the row below. More formally, if you are on index <code>i</code> on the current row, you may move to either index <code>i</code> or index <code>i + 1</code> on the next row.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
<strong>Output:</strong> 11
<strong>Explanation:</strong> The triangle looks like:
   <u>2</u>
  <u>3</u> 4
 6 <u>5</u> 7
4 <u>1</u> 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> triangle = [[-10]]
<strong>Output:</strong> -10
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= triangle.length &lt;= 200</code></li>
	<li><code>triangle[0].length == 1</code></li>
	<li><code>triangle[i].length == triangle[i - 1].length + 1</code></li>
	<li><code>-10<sup>4</sup> &lt;= triangle[i][j] &lt;= 10<sup>4</sup></code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you&nbsp;do this using only <code>O(n)</code> extra space, where <code>n</code> is the total number of rows in the triangle?</div>

<p>&nbsp;</p>

## My Solutions

### Solution 1: DP

#### C++ Code (version-1. Up to Down)
```cpp
// Problem: https://leetcode.com/problems/triangle/
// Author: Araceae
// Date: 2021/4/21

// Solution: DP (Up to Down)

// Time Complexity: O(MN)
// Space Complexity: O(1)

// Performance: 
// Runtime: 4 ms, faster than 90.75% of C++ online submissions for Triangle.
// Memory Usage: 8.3 MB, less than 93.93% of C++ online submissions for Triangle.
class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int m = triangle.size();
        for (int i = 0; i < m; ++i){
            int n = triangle[i].size();
            
            for (int j = 0; j < n; ++j){
                int x1 = (i-1 >= 0 && j-1 >= 0) ? triangle[i-1][j-1] : INT_MAX;
                int x2 = (i-1 >= 0 && j < n-1) ? triangle[i-1][j] : INT_MAX;
                
                if (x1 != INT_MAX && x2 != INT_MAX) triangle[i][j] += min(x1, x2);
                else if (x1 != INT_MAX) triangle[i][j] += x1;
                else if (x2 != INT_MAX) triangle[i][j] += x2;
            }
        }
        int res = INT_MAX;
        for (int i = 0; i < triangle[m-1].size(); ++i)
            res = triangle[m-1][i] < res ? triangle[m-1][i] : res;
        return res;
    }
};
```

#### C++ Code (version-2. Down to Up)
```cpp
// Problem: https://leetcode.com/problems/triangle/
// Author: Araceae
// Date: 2021/4/21

// Solution: DP (Down to Up)

// Time Complexity: O(MN)
// Space Complexity: O(1)

// Performance: 
// Runtime: 0 ms, faster than 100.00% of C++ online submissions for Triangle.
// Memory Usage: 8.4 MB, less than 71.00% of C++ online submissions for Triangle.

class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int m = triangle.size();
        
        for (int i = m-1; i > 0; --i){
            int n = triangle[i].size();
            for (int j = 0; j < n-1; ++j){
                triangle[i-1][j] += min(triangle[i][j], triangle[i][j+1]);
            }
        }
        
        return triangle[0][0];
    }
};
```