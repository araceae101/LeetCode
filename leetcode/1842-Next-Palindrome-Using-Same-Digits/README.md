# [1842. Next Palindrome Using Same Digits](https://leetcode.com/problems/next-palindrome-using-same-digits/)

<div><p>You are given a numeric string <code>num</code>, representing a very large <strong>palindrome</strong>.</p>

<p>Return<em> the <strong>smallest palindrome larger than </strong></em><code>num</code><em> that can be created by rearranging its digits. If no such palindrome exists, return an empty string </em><code>""</code>.</p>

<p>A <strong>palindrome</strong> is a number that reads the same backward as forward.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> num = "1221"
<strong>Output:</strong> "2112"
<strong>Explanation:</strong>&nbsp;The next palindrome larger than "1221" is "2112".
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> num = "32123"
<strong>Output:</strong> ""
<strong>Explanation:</strong>&nbsp;No palindromes larger than "32123" can be made by rearranging the digits.
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> num = "45544554"
<strong>Output:</strong> "54455445"
<strong>Explanation:</strong> The next palindrome larger than "45544554" is "54455445".
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= num.length &lt;= 10<sup>5</sup></code></li>
	<li><code>num</code> is a <strong>palindrome</strong>.</li>
</ul></div>

<p>&nbsp;</p>

## My Solutions
### Solution 1: Next permutation
The idea is the same as [31. Next Permutation](https://leetcode.com/problems/next-permutation/).
We can first find the next permutation for the first half of array and then fill the remaining half according the first half.
- Time complexity: O(N)
- Space complexity: O(1)

#### C++ Code:
```cpp
class Solution {
public:
    string nextPalindrome(string num) {
        int n = num.size();
		
		// find the index k, l and swap them to get the next permutaion
        int k;
        for (k = n/2 - 2; k >= 0; --k)
            if (num[k] < num[k+1]) break;
        
        if (k < 0)  return "";
        
        int l;
        for (l = n/2 - 1; l >= 0; --l)
            if (num[k] < num[l])    break;
        
        swap(num[k], num[l]);
        reverse(num.begin() + k+1, num.begin() + n/2);
        
		// generate the data for remaining half 
        for (int i = 0; i < n/2; ++i)
             num[num.size()-1-i] = num[i];
        
        return num;
    }
};
```