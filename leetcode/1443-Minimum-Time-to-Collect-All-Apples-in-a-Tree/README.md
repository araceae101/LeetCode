# 1443. Minimum Time to Collect All Apples in a Tree

<div><p>Given an undirected tree consisting of <code>n</code> vertices numbered from <code>0</code> to <code>n-1</code>, which has some apples in their vertices. You spend 1 second to walk over one edge of the tree. <em>Return the minimum time in seconds you have to spend to collect all apples in the tree, starting at <strong>vertex 0</strong> and coming back to this vertex.</em></p>

<p>The edges of the undirected tree are given in the array <code>edges</code>, where <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> means that exists an edge connecting the vertices <code>a<sub>i</sub></code> and <code>b<sub>i</sub></code>. Additionally, there is a boolean array <code>hasApple</code>, where <code>hasApple[i] = true</code> means that vertex <code>i</code> has an apple; otherwise, it does not have any apple.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2020/04/23/min_time_collect_apple_1.png" style="width: 300px; height: 212px;"></strong></p>

<pre><strong>Input:</strong> n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,true,true,false]
<strong>Output:</strong> 8 
<strong>Explanation:</strong> The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  
</pre>

<p><strong>Example 2:</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2020/04/23/min_time_collect_apple_2.png" style="width: 300px; height: 212px;"></strong></p>

<pre><strong>Input:</strong> n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,false,true,false]
<strong>Output:</strong> 6
<strong>Explanation:</strong> The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,false,false,false,false,false]
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10^5</code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 &lt;= a<sub>i</sub> &lt; b<sub>i</sub> &lt;= n - 1</code></li>
	<li><code>from<sub>i</sub>&nbsp;&lt; to<sub>i</sub></code></li>
	<li><code>hasApple.length == n</code></li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1. BFS
- Time complexity: O(N)
- Space complexity: O(N)
- Idea: To count how many nodes will pass through (`totalNode`), and the result is `(totalNode - 1) * 2`. I used hash map to store the parent node index, and then used bfs method to find all the parents and mark as needed in the original vector hasApple.
> Note that the edges `(i, j)` only means the connection between two nodes `i, j`, but NOT means `i` is a parent node with a child node `j` .
```
For example, the following edges input is valid:
[[0,1],[0,2],[1,3],[0,4]]
       0 ---
     /  \    \
   1     2    4
  /
 3
```
#### C++ Code
```cpp
// Problem: https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Author: Araceae
// Date: 2021/4/18

// Solution: BFS solution

// Time Complexity: O(N)
// Space Complexity: O(N)

// Performance:
// Runtime: 136 ms, faster than 97.85% of C++ online submissions for Minimum Time to Collect All Apples in a Tree.
// Memory Usage: 57.5 MB, less than 98.09% of C++ online submissions for Minimum Time to Collect All Apples in a Tree.

class Solution {
public:
    int minTime(int n, vector<vector<int>>& edges, vector<bool>& hasApple) {
        
        unordered_map<int, int> nodeParent;
        for (auto &e : edges){
            if (nodeParent.find(e[1]) == nodeParent.end()) nodeParent[e[1]] = e[0];
            else    nodeParent[e[0]] = e[1];
        }
        
        queue<int> q;
        
        for (int i = 0; i < hasApple.size(); ++i)
            if (hasApple[i])  q.push(i);
        
        while (!q.empty()){
            int currIdx = q.front();
            q.pop();
            if (currIdx != 0){
                int nextIdx = nodeParent[currIdx];
				
                // if the node has NOT yet visited, then add to the queue 
                if (!hasApple[nextIdx]){
                    hasApple[nextIdx] = true;
                    q.push(nextIdx);
                }
            }
        }
        
        int countPath = accumulate(hasApple.begin(), hasApple.end(), -1);
        
        return countPath <= 0 ? 0 : 2 * countPath;
    }
};
```