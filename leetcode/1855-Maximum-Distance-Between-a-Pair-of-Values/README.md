# [1855. Maximum Distance Between a Pair of Values](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/)

<div><p>You are given two <strong>non-increasing 0-indexed </strong>integer arrays <code>nums1</code>​​​​​​ and <code>nums2</code>​​​​​​.</p>

<p>A pair of indices <code>(i, j)</code>, where <code>0 &lt;= i &lt; nums1.length</code> and <code>0 &lt;= j &lt; nums2.length</code>, is <strong>valid</strong> if both <code>i &lt;= j</code> and <code>nums1[i] &lt;= nums2[j]</code>. The <strong>distance</strong> of the pair is <code>j - i</code>​​​​.</p>

<p>Return <em>the <strong>maximum distance</strong> of any <strong>valid</strong> pair </em><code>(i, j)</code><em>. If there are no valid pairs, return </em><code>0</code>.</p>

<p>An array <code>arr</code> is <strong>non-increasing</strong> if <code>arr[i-1] &gt;= arr[i]</code> for every <code>1 &lt;= i &lt; arr.length</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
<strong>Output:</strong> 2
<strong>Explanation:</strong> The valid pairs are (0,0), (2,2), (2,3), (2,4), (3,3), (3,4), and (4,4).
The maximum distance is 2 with pair (2,4).
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> nums1 = [2,2,2], nums2 = [10,10,1]
<strong>Output:</strong> 1
<strong>Explanation:</strong> The valid pairs are (0,0), (0,1), and (1,1).
The maximum distance is 1 with pair (0,1).
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> nums1 = [30,29,19,5], nums2 = [25,25,25,25,25]
<strong>Output:</strong> 2
<strong>Explanation:</strong> The valid pairs are (2,2), (2,3), (2,4), (3,3), and (3,4).
The maximum distance is 2 with pair (2,4).
</pre>

<p><strong>Example 4:</strong></p>

<pre><strong>Input:</strong> nums1 = [5,4], nums2 = [3,2]
<strong>Output:</strong> 0
<strong>Explanation:</strong> There are no valid pairs, so return 0.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums2.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums1[i], nums2[j] &lt;= 10<sup>5</sup></code></li>
	<li>Both <code>nums1</code> and <code>nums2</code> are <strong>non-increasing</strong>.</li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1 : Hash Table
Firstly, generate hash table like the following example:
```
nums2: [10,10,1]
key:            1  2  3  4  5  6  7  8  9  10  ...
value (index):  2  1  1  1  1  1  1  1  1  1
```
And then we can iteratively loop for each nums1 and get the largest index`j` in nums2 `(nums2[j])`, which its value is greater than `nums1[i]`
- Time Complexity: O(N1+N2)
- Space Complexity: O(1)

#### C++ Code:
```cpp
class Solution {
public:
    int maxDistance(vector<int>& nums1, vector<int>& nums2) {
        vector<int> h(100001, -1);
        for (int i = 0; i < nums2.size(); ++i)  h[nums2[i]] = i;
        for (int i = h.size()-2; i >= 0; --i)   h[i] = h[i] == -1 ? h[i+1] : h[i];
        
        int res = 0;
        for (int i = 0; i < nums1.size(); ++i) {
            int j = h[nums1[i]];
            res = max(j-i, res);
        }
        return res;
    }
};
```

### Solution 2 : Two Pointer
- Time Complexity: O(N1+N2)
- Space Complexity: O(1)
```cpp
class Solution {
public:
    int maxDistance(vector<int>& nums1, vector<int>& nums2) {
        int i = 0, j = 0, res = 0;
        while (i < nums1.size() && j < nums2.size()) {
            if (nums1[i] > nums2[j])  ++i;
            res = max(res, j++ -i);
        }
        return res;
    }
};
```