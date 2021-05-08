# [906. Super Palindromes](https://leetcode.com/problems/super-palindromes/)

<div><p>Let's say a positive integer is a <strong>super-palindrome</strong> if it is a palindrome, and it is also the square of a palindrome.</p>

<p>Given two positive integers <code>left</code> and <code>right</code> represented as strings, return <em>the number of <strong>super-palindromes</strong> integers in the inclusive range</em> <code>[left, right]</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> left = "4", right = "1000"
<strong>Output:</strong> 4
<strong>Explanation</strong>: 4, 9, 121, and 484 are superpalindromes.
Note that 676 is not a superpalindrome: 26 * 26 = 676, but 26 is not a palindrome.
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> left = "1", right = "2"
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= left.length, right.length &lt;= 18</code></li>
	<li><code>left</code> and <code>right</code> consist of only digits.</li>
	<li><code>left</code> and <code>right</code> cannot have leading zeros.</li>
	<li><code>left</code> and <code>right</code> represent integers in the range <code>[1, 10<sup>18</sup>]</code>.</li>
	<li><code>left</code> is less than or equal to <code>right</code>.</li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution 1 : Force Brute
- Idea:
    1. To generate all possible palindrome strings
    2. For each palindrome string, check if its square is still a palindrome string.
- But this approach is slow... need to improve the performance.
#### **C++ Code**
```cpp
class Solution {
    vector<string> parlindrom;
    vector<char> num {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};
    
    void gernParlindrome(int n, string tmp) {
        // finish generating parlindrome string with n characters
        if (tmp.size() == n) {
            parlindrom.push_back(tmp);
            return;
        }
        
        // finish generating half of parlindrome string
        // and need to complete full string
        bool isOdd = n % 2;
        // for even number of string:
        if (!isOdd && tmp.size() == n/2) {
            for (int i = tmp.size()-1; i >= 0; --i)
                tmp.push_back(tmp[i]);
            parlindrom.push_back(tmp);
            return;
        }
        // for odd number of string:
        if (isOdd && tmp.size() > n/2) {
            for (int i = tmp.size()-2; i >= 0; --i)
                tmp.push_back(tmp[i]);
            parlindrom.push_back(tmp);
            return;
        }
        
        // generate for all possible characters
        for (int i = 0; i < 10; ++i) {
            if (tmp.size() == 0 && i == 0)  continue;
            tmp.push_back(num[i]);
            gernParlindrome(n, tmp);
            tmp.pop_back();
        }
    }
    
    bool isPalindrome(long long int num) {
        string str;
        while (num) {
            str.push_back(num % 10);
            num /= 10;
        }
        int l = 0, r = str.size()-1;
        while (l < r)
            if (str[l++] != str[r--])
                return false;
        
        return true;
    }
    
public:
    int superpalindromesInRange(string lStr, string rStr) {
        
        for (int i = 1; i <= 10; ++i)
            gernParlindrome(i, "");
        
        long long left = stoll(lStr);
        long long right = stoll(rStr);
        
        int res = 0;
        for (auto &e : parlindrom) {
            long long currPalim = stoll(e);
            long long squareCurrPalim = currPalim * currPalim;
            if (squareCurrPalim < left) continue;
            if (squareCurrPalim > right) break;
            if (isPalindrome(squareCurrPalim)) ++res;
        }
        
        return res;
    }
};
```