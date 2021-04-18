# 2. Add Two Numbers

<div><p>You are given two <strong>non-empty</strong> linked lists representing two non-negative integers. The digits are stored in <strong>reverse order</strong>, and each of their nodes contains a single digit. Add the two numbers and return the sum&nbsp;as a linked list.</p>

<p>You may assume the two numbers do not contain any leading zero, except the number 0 itself.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg" style="width: 483px; height: 342px;">
<pre><strong>Input:</strong> l1 = [2,4,3], l2 = [5,6,4]
<strong>Output:</strong> [7,0,8]
<strong>Explanation:</strong> 342 + 465 = 807.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> l1 = [0], l2 = [0]
<strong>Output:</strong> [0]
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
<strong>Output:</strong> [8,9,9,9,0,0,0,1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in each linked list is in the range <code>[1, 100]</code>.</li>
	<li><code>0 &lt;= Node.val &lt;= 9</code></li>
	<li>It is guaranteed that the list represents a number that does not have leading zeros.</li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1: Recursive
```cpp
// // Problem: https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Author: Araceae
// Date: 2021/4/18

// Solution: Recursively construct linked list

// Performance: 
// Runtime: 24 ms, faster than 80.04%
// Memory Usage: 71.2 MB, less than 50.98%

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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2, int carry = 0) {
        if (l1 == NULL && l2 == NULL && carry == 0)   return NULL;
        
        int n1 = l1 == NULL ? 0 : l1->val;
        int n2 = l2 == NULL ? 0 : l2->val;
        
        ListNode* node = new ListNode((n1 + n2 + carry) % 10);
        carry = (n1 + n2 + carry) / 10;
        node->next = addTwoNumbers(l1 ? l1->next : NULL, l2 ? l2->next : NULL, carry);
        return node;
    }
};
```