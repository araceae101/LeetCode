# [1820. Maximum Number of Accepted Invitations](https://leetcode.com/problems/maximum-number-of-accepted-invitations/)

<div><p>There are <code>m</code> boys and <code>n</code> girls in a class attending an upcoming party.</p>

<p>You are given an <code>m x n</code> integer matrix <code>grid</code>, where <code>grid[i][j]</code> equals <code>0</code> or <code>1</code>. If <code>grid[i][j] == 1</code>, then that means the <code>i<sup>th</sup></code> boy can invite the <code>j<sup>th</sup></code> girl to the party. A boy can invite at most<strong> one girl</strong>, and a girl can accept at most <strong>one invitation</strong> from a boy.</p>

<p>Return <em>the <strong>maximum</strong> possible number of accepted invitations.</em></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> grid = [[1,1,1],
               [1,0,1],
               [0,0,1]]
<strong>Output:</strong> 3<strong>
Explanation:</strong> The invitations are sent as follows:
- The 1<sup>st</sup> boy invites the 2<sup>nd</sup> girl.
- The 2<sup>nd</sup> boy invites the 1<sup>st</sup> girl.
- The 3<sup>rd</sup> boy invites the 3<sup>rd</sup> girl.</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> grid = [[1,0,1,0],
               [1,0,0,0],
               [0,0,1,0],
               [1,1,1,0]]
<strong>Output:</strong> 3
<strong>Explanation:</strong> The invitations are sent as follows:
-The 1<sup>st</sup> boy invites the 3<sup>rd</sup> girl.
-The 2<sup>nd</sup> boy invites the 1<sup>st</sup> girl.
-The 3<sup>rd</sup> boy invites no one.
-The 4<sup>th</sup> boy invites the 2<sup>nd</sup> girl.</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>grid.length == m</code></li>
	<li><code>grid[i].length == n</code></li>
	<li><code>1 &lt;= m, n &lt;= 200</code></li>
	<li><code>grid[i][j]</code> is either <code>0</code> or <code>1</code>.</li>
</ul>

<p>&nbsp;</p>

## My Solutions
### Solution 1. Hungrian Algorithm
- Description: Hungrian Algorithm uses DFS to recursively update the best pairing combination
- Time complexity: O(MNN)
- Space complexity: O(N)

#### C++
```cpp
class Solution {
    bool bpm(vector<vector<int>>& grid, vector<int>& match, vector<int>& visited, int u){
        int m = grid.size(), n = grid[0].size();
        
        for (int v = 0; v < n; ++v){
            // if the girl-v has checked or was not invited by the boy-u,
            // then continue to check the next girl
            if (visited[v] || !grid[u][v])
                continue;

            visited[v] = true;
            // if we have not yet found the partner for girl-v, 
            // or the girl-v already has the partner but we can find another partner for her,
            // then update the match list
            if (match[v] < 0 || bpm(grid, match, visited, match[v])){
                match[v] = u;
                return true;
            }
        }
        return false;
    }
public:
    int maximumInvitations(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<int> match(n, -1);
        int res = 0;
        
        for (int u = 0; u < m; ++u){
            vector<int> visited(n, 0);
            if (bpm(grid, match, visited, u))
                res++;
        }
        return res;
    }
};
```