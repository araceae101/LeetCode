# [1780. Check if Number is a Sum of Powers of Three](https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/)

<div><p>Given an integer <code>n</code>, return <code>true</code> <em>if it is possible to represent </em><code>n</code><em> as the sum of distinct powers of three.</em> Otherwise, return <code>false</code>.</p>

<p>An integer <code>y</code> is a power of three if there exists an integer <code>x</code> such that <code>y == 3<sup>x</sup></code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> n = 12
<strong>Output:</strong> true
<strong>Explanation:</strong> 12 = 3<sup>1</sup> + 3<sup>2</sup>
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> n = 91
<strong>Output:</strong> true
<strong>Explanation:</strong> 91 = 3<sup>0</sup> + 3<sup>2</sup> + 3<sup>4</sup>
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> n = 21
<strong>Output:</strong> false
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>7</sup></code></li>
</ul>
</div>

<p>&nbsp;</p>


## My Solutions
### **Idea**
The question can be converted to whether there is a `2` digit  after converting the input number to base-3 number.
For example: 
- The base-3 number for `12` is `110` , and there is no `2` digit, so the answer is `true`.
- The base-3 number for `91` is `10101`, and there is no `2` digit, so the answer is `true`.
- The base-3 number for `21` is `210`, there is a `2` digit, so the answer is `false`.

### **C++ Code**
#### Recursive:
```cpp
class Solution {
public:
    bool checkPowersOfThree(int n, int r = 0) {
        if (r == 2) return false;
        if (n == 0) return true;
        
        return checkPowersOfThree(n/3, n%3);
    }
};

// Time Complexity: O(log3 N)
// Space Complexity: O(log3 N)
```

#### Iterative:
```cpp
class Solution {
public:
    bool checkPowersOfThree(int n) {
        int r = 0;
        while (n > 0) {
            r = n % 3;
            n /= 3;
            if (r == 2) return false;
        }
        return true;
    }
};

// Time Complexity: O(log3 N)
// Space Complexity: O(1)
```