# [51. N-Queens](https://leetcode.com/problems/n-queens/)
<div><p>The <strong>n-queens</strong> puzzle is the problem of placing <code>n</code> queens on an <code>n x n</code> chessboard such that no two queens attack each other.</p>

<p>Given an integer <code>n</code>, return <em>all distinct solutions to the <strong>n-queens puzzle</strong></em>.</p>

<p>Each solution contains a distinct board configuration of the n-queens' placement, where <code>'Q'</code> and <code>'.'</code> both indicate a queen and an empty space, respectively.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/queens.jpg" style="width: 600px; height: 268px;">
<pre><strong>Input:</strong> n = 4
<strong>Output:</strong> [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
<strong>Explanation:</strong> There exist two distinct solutions to the 4-queens puzzle as shown above
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> n = 1
<strong>Output:</strong> [["Q"]]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 9</code></li>
</ul>
</div>

<p>&nbsp;</p>

## My Solution
### Solution-1: DFS
#### C++ Code:
```cpp
class Solution {
    bool isValid(vector<string>& board, int row, int col) {
        for (int i = row; i >= 0; --i)
            if (board[i][col] == 'Q') return false;
        for (int i = row, j = col; i >= 0 && j >= 0; --i, --j)
            if (board[i][j] == 'Q') return false;
        for (int i = row, j = col; i >= 0 && j < board.size(); --i, ++j)
            if (board[i][j] == 'Q') return false;
        return true;
    }
    void dfs(vector<string>& board, int row, vector<vector<string>>& res) {
        if (row == board.size()) {
            res.push_back(board);
            return;
        }
        for (int i = 0; i < board.size(); ++i) {
            if (isValid(board, row, i)) {
                board[row][i] = 'Q';
                dfs(board, row+1, res);
                board[row][i] = '.';
            }
        }
    }
public:
    vector<vector<string>> solveNQueens(int n) {
        if (n <= 0) return {{}};
        vector<string> board(n, string(n, '.'));
        vector<vector<string>> res;
        dfs(board, 0, res);
        return res;
    }
};
```