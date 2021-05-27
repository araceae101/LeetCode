# [318. Maximum Product of Word Lengths](https://leetcode.com/problems/maximum-product-of-word-lengths/)

<div><p>Given a string array <code>words</code>, return <em>the maximum value of</em> <code>length(word[i]) * length(word[j])</code> <em>where the two words do not share common letters</em>. If no such two words exist, return <code>0</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> words = ["abcw","baz","foo","bar","xtfn","abcdef"]
<strong>Output:</strong> 16
<strong>Explanation:</strong> The two words can be "abcw", "xtfn".
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> words = ["a","ab","abc","d","cd","bcd","abcd"]
<strong>Output:</strong> 4
<strong>Explanation:</strong> The two words can be "ab", "cd".
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> words = ["a","aa","aaa","aaaa"]
<strong>Output:</strong> 0
<strong>Explanation:</strong> No such pair of words.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= words.length &lt;= 1000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 1000</code></li>
	<li><code>words[i]</code> consists only of lowercase English letters.</li>
</ul>
</div>

<p>&nbsp;</p>

## My Solutions
### Solution-1: bruce-force
- Time Complexity: O(N^2)
- Space Complexity: O(N)
- Performance:
  - Runtime: 92 ms, faster than 29.24% of C++ online submissions for Maximum Product of Word Lengths.
  - Memory Usage: 17.9 MB, less than 26.37% of C++ online submissions for Maximum Product of Word Lengths.
#### C++ Code:
```cpp
class Solution {
public:
    int maxProduct(vector<string>& words) {
        int n = words.size();
        vector<vector<char>> m(n, vector<char>(26, 0));
        for (int i = 0; i < n; ++i)
            for (int j = 0;j < words[i].size(); ++j)
                m[i][words[i][j] - 'a']++;
        
        int res = 0;
        for (int i = 0; i < n; ++i) {
            for (int j = i; j < n; ++j) {
                int k = 0;
                for (; k < 26; ++k)
                    if (m[i][k] != 0 && m[j][k] != 0)   break;
                if (k == 26)
                    res = max(res, int(words[i].size() * words[j].size()));
            }
        }
        return res;
    }
};
```

### Solution-2: bit manipulation
- Time Complexity: O(N^2)
- Space Complexity: O(N)
- Performance: 
  - Runtime: 32 ms, faster than 97.49% of C++ online submissions for Maximum Product of Word Lengths.
  - Memory Usage: 15.5 MB, less than 96.42% of C++ online submissions for Maximum Product of Word Lengths.
#### C++ Code:
```cpp
class Solution {
public:
    int maxProduct(vector<string>& words) {
        int n = words.size();
        vector<int> mask(n, 0);
        for (int i = 0; i < n; ++i)
            for (auto &c : words[i])
                mask[i] |= 1 << (c - 'a');
        int res = 0;
        for (int i = 0; i < n; ++i)
            for (int j = i; j < n; ++j)
                if (!(mask[i] & mask[j]))
                    res = max(res, int(words[i].size() * words[j].size()));
        return res;
    }
};
```