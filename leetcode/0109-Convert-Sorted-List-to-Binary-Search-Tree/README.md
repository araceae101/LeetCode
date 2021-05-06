# [109. Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)

<div><p>Given the <code>head</code> of a singly linked list where elements are <strong>sorted in ascending order</strong>, convert it to a height balanced BST.</p>

<p>For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of <em>every</em> node never differ by more than 1.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/08/17/linked.jpg" style="width: 600px; height: 466px;">
<pre><strong>Input:</strong> head = [-10,-3,0,5,9]
<strong>Output:</strong> [0,-3,9,-10,null,5]
<strong>Explanation:</strong> One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> head = []
<strong>Output:</strong> []
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> head = [0]
<strong>Output:</strong> [0]
</pre>

<p><strong>Example 4:</strong></p>

<pre><strong>Input:</strong> head = [1,3]
<strong>Output:</strong> [3,1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in <code>head</code> is in the range <code>[0, 2 * 10<sup>4</sup>]</code>.</li>
	<li><code>-10^5 &lt;= Node.val &lt;= 10^5</code></li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1 : Binary Search and use extra space
1. Convert the listNode to vector.
1. Recursively find the mid of the element in the vector and add to the tree node.

- Time Complexity: O(N)
- Space Complexity: O(N)
#### C++ Code:
```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
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
    TreeNode* buildTree(vector<int>& list, int start, int end) {
        if (start > end)    return NULL;
        int mid = start + (end - start + 1) / 2;
        TreeNode* node = new TreeNode(list[mid]);
        node->left = buildTree(list, start, mid-1);
        node->right = buildTree(list, mid+1, end);
        return node;
    }
public:
    TreeNode* sortedListToBST(ListNode* head) {
        vector<int> list;
        for (ListNode* curr = head; curr; curr = curr->next)
            list.push_back(curr->val);
        return buildTree(list, 0, list.size()-1);
    }
};
```


### Solution 1 : Binary Search and no extra space used
- Time Complexity: O(N)
- Space Complexity: O(1)
#### C++ Code:
```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
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
    ListNode* currentNode;
    TreeNode* buildTree(int start, int end) {
        if (start > end)    return NULL;
        int mid = start + (end - start + 1) / 2;
        TreeNode* node = new TreeNode();
        node->left = buildTree(start, mid-1);
        node->val = currentNode->val;
        currentNode = currentNode->next;
        node->right = buildTree(mid+1, end);
        return node;
    }
public:
    TreeNode* sortedListToBST(ListNode* head) {
        currentNode = head;
        int len = 0;
        for (ListNode* curr = head; curr; curr = curr->next)
            ++len;
        return buildTree(0, len-1);
    }
};
```