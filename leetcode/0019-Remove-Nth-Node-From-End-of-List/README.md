# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

<div><p>Given the <code>head</code> of a linked list, remove the <code>n<sup>th</sup></code> node from the end of the list and return its head.</p>

<p><strong>Follow up:</strong>&nbsp;Could you do this in one pass?</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg" style="width: 542px; height: 222px;">
<pre><strong>Input:</strong> head = [1,2,3,4,5], n = 2
<strong>Output:</strong> [1,2,3,5]
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> head = [1], n = 1
<strong>Output:</strong> []
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> head = [1,2], n = 1
<strong>Output:</strong> [1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the list is <code>sz</code>.</li>
	<li><code>1 &lt;= sz &lt;= 30</code></li>
	<li><code>0 &lt;= Node.val &lt;= 100</code></li>
	<li><code>1 &lt;= n &lt;= sz</code></li>
</ul>
</div>

<p>&nbsp;</p>


## My Solutions
### Solution 1: Two Pointer
#### C++ Code
```cpp
// Problem: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Author: Araceae
// Date: 2021/4/18

// Solution: Two pointer

// Time Complexity: O(N)
// Space Complexity: O(1)

// Performance: 
// Runtime: 0 ms, faster than 100.00% of C++ online submissions for Remove Nth Node From End of List.
// Memory Usage: 10.6 MB, less than 92.61% of C++ online submissions for Remove Nth Node From End of List.

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
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* res = new ListNode(-1, head);
        ListNode* left = res;
        ListNode* right = left;
        
        for (int i = 0; i < n; ++i)
            right = right->next;
        
        while (right->next){
            left = left->next;
            right = right->next;
        }
        
        left->next = left->next->next;
        
        return res->next;
    }
};
```