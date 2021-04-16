# [542. 01 Matrix](https://leetcode.com/problems/01-matrix/)

<div><p>Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.</p>

<p>The distance between two adjacent cells is 1.</p>

<p>&nbsp;</p>

<p><b>Example 1: </b></p>

<pre><strong>Input:</strong>
[[0,0,0],
 [0,1,0],
 [0,0,0]]

<strong>Output:</strong>
[[0,0,0],
&nbsp;[0,1,0],
&nbsp;[0,0,0]]
</pre>

<p><b>Example 2: </b></p>

<pre><b>Input:</b>
[[0,0,0],
 [0,1,0],
 [1,1,1]]

<strong>Output:</strong>
[[0,0,0],
 [0,1,0],
 [1,2,1]]
</pre>

<p>&nbsp;</p>

<p><b>Note:</b></p>

<ol>
	<li>The number of elements of the given matrix will not exceed 10,000.</li>
	<li>There are at least one 0 in the given matrix.</li>
	<li>The cells are adjacent in only four directions: up, down, left and right.</li>
</ol>
</div>

<p>&nbsp;</p>

## My Solutions

### Solution 1. DP
#### C++
```cpp
// Problem: https://leetcode.com/problems/01-matrix/
// Author: Araceae
// Date: 2021/4/17

// Solution: DP solution
// DP from 4 quadrant and to ask the neighbors in front
// of their nearest zero distance and plus 1: 
//     (1) Upper Left  --to--> Lower Right
//     (2) Upper Right --to--> Lower Left
//     (3) Lower Right --to--> Upper Left
//     (4) Lower Left  --to--> Upper Right
// Time Complexity: O(MN)
// Space Complexity: O(MN)

// Performance: 
// Runtime: 68 ms, faster than 66.03% of C++ online submissions
// Memory Usage: 26 MB, less than 92.59% of C++ online submissions

class Solution {
public:
    vector<vector<int>> updateMatrix(vector<vector<int>>& matrix) {
        int m = matrix.size(); if (m < 1)   return {{}};
        int n = matrix[0].size();   if (n < 1)  return {{}};
        
        vector<vector<int>> res(m, vector<int>(n, INT_MAX));
        
        for (int i = 0; i < m; ++i){
            for (int j = 0; j < n; ++j){
                
                // Upper left
                if (matrix[i][j] == 0)  res[i][j] = 0;
                else {
                    if (i-1 >= 0 && res[i-1][j] != INT_MAX) res[i][j] = min(res[i][j], res[i-1][j] + 1);
                    if (j-1 >= 0 && res[i][j-1] != INT_MAX) res[i][j] = min(res[i][j], res[i][j-1] + 1);
                }
                
                // Upper Right
                if (matrix[i][n-1-j] == 0)  res[i][n-1-j] = 0;
                else {
                    if (i-1 >= 0 && res[i-1][n-1-j] != INT_MAX) res[i][n-1-j] = min(res[i][n-1-j], res[i-1][n-1-j] + 1);
                    if (n-1-j+1 < n && res[i][n-1-j+1] != INT_MAX) res[i][n-1-j] = min(res[i][n-1-j], res[i][n-1-j+1] + 1);
                }

                // Lower Right
                if (matrix[m-1-i][n-1-j] == 0)  res[m-1-i][n-1-j] = 0;
                else {
                    if (m-1-i+1 < m && res[m-1-i+1][n-1-j] != INT_MAX) res[m-1-i][n-1-j] = min(res[m-1-i][n-1-j], res[m-1-i+1][n-1-j] + 1);
                    if (n-1-j+1 < n && res[m-1-i][n-1-j+1] != INT_MAX) res[m-1-i][n-1-j] = min(res[m-1-i][n-1-j], res[m-1-i][n-1-j+1] + 1);
                }
                
                // Lower Left
                if (matrix[m-1-i][j] == 0)  res[m-1-i][j] = 0;
                else {
                    if (m-1-i+1 < m && res[m-1-i+1][j] != INT_MAX) res[m-1-i][j] = min(res[m-1-i][j], res[m-1-i+1][j] + 1);
                    if (j-1 >= 0 && res[m-1-i][j-1] != INT_MAX) res[m-1-i][j] = min(res[m-1-i][j], res[m-1-i][j-1] + 1);
                }
            }
        }
        
        return res;
    }
};
```