# 1161. Maximum Level Sum of a Binary Tree

<div><p>Given the <code>root</code> of a binary tree, the level of its root is <code>1</code>, the level of its children is <code>2</code>, and so on.</p>

<p>Return the <strong>smallest</strong> level <code>x</code> such that the sum of all the values of nodes at level <code>x</code> is <strong>maximal</strong>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG" style="width: 200px; height: 175px;">
<pre><strong>Input:</strong> root = [1,7,0,7,-8,null,null]
<strong>Output:</strong> 2
<strong>Explanation: </strong>
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> root = [989,null,10250,98693,-89388,null,null,null,-32127]
<strong>Output:</strong> 2
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the tree is in the range <code>[1, 10<sup>4</sup>]</code>.</li>
	<li><code>-10<sup>5</sup> &lt;= Node.val &lt;= 10<sup>5</sup></code></li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1: Recursion
#### C++ Code
- Runtime is slow, but the memory usage is less.
```cpp
// Problem: https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Author: Araceae
// Date: 2021/4/18

// Solution: Recursion

// Time Complexity: O(N)
// Space Complexity: O(height(tree))

// Performance:
// Runtime: 224 ms, faster than 6.27% of C++ online submissions for Maximum Level Sum of a Binary Tree.
// Memory Usage: 104.6 MB, less than 98.67% of C++ online submissions for Maximum Level Sum of a Binary Tree.

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
    vector<int> maxValVec;
    
public:
    int maxLevelSum(TreeNode* root, int level = 0) {
        
        if (!root)  return 0;
        
        if (maxValVec.size() <= level)
            maxValVec.push_back(root->val);
        else
            maxValVec[level] += root->val;
        
        maxLevelSum(root->left, level+1);
        maxLevelSum(root->right, level+1);
        
        
        int maxIdx = -1;
        for (int i = 0; i < maxValVec.size(); ++i)
            if (maxIdx == -1 || maxValVec[i] > maxValVec[maxIdx])   maxIdx = i;
        
        return maxIdx + 1;
    }
};
```

### Solution 2: BFS queue
#### C++ code
```cpp
// Problem: https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Author: Araceae
// Date: 2021/4/18

// Solution: BFS queue

// Time Complexity: O(N)
// Space Complexity: O(N)

// Performance:
// Runtime: 172 ms, faster than 64.27% of C++ online submissions for Maximum Level Sum of a Binary Tree.
// Memory Usage: 107.4 MB, less than 34.07% of C++ online submissions for Maximum Level Sum of a Binary Tree.

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int maxLevelSum(TreeNode* root) {
        
        int maxSum = INT_MIN;
        int res_idx = -1;
        
        queue<TreeNode*> q;
        q.push(root);
        
        int idx = -1;
        while (!q.empty()){
            idx++;
            int l = q.size();
            int s = 0;
            for (int i = 0; i < l; ++i){
                TreeNode* curr = q.front();
                q.pop();
                s += curr->val;
                if (curr->left) q.push(curr->left);
                if (curr->right) q.push(curr->right);
            }
            if (s > maxSum){
                res_idx = idx;
                maxSum = s;
            }
        }
        
        return res_idx + 1;
    }
};
```