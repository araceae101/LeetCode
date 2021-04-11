# [1823. Find the Winner of the Circular Game](https://leetcode.com/problems/find-the-winner-of-the-circular-game/)

There are `n` friends that are playing a game. The friends are sitting in a circle and are numbered from 1 to n in **clockwise order**. More formally, moving clockwise from the ith friend brings you to the `(i+1)^th^` friend for `1 <= i < n`, and moving clockwise from the nth friend brings you to the 1st friend.

The rules of the game are as follows:

1. **Start** at the `1^st^` friend.
1. Count the next `k` friends in the clockwise direction **including** the friend you started at. The counting wraps around the circle and may count some friends more than once.
1. The last friend you counted leaves the circle and loses the game.
1. If there is still more than one friend in the circle, go back to step `2` starting from the friend **immediately clockwise** of the friend who just lost and repeat.
1. Else, the last friend in the circle wins the game.
Given the number of friends, `n`, and an integer `k`, *return the winner of the game*.

<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> n = 5, k = 2
<strong>Output:</strong> 3
<strong>Explanation:</strong> Here are the steps of the game:
1) Start at friend 1.
2) Count 2 friends clockwise, which are friends 1 and 2.
3) Friend 2 leaves the circle. Next start is friend 3.
4) Count 2 friends clockwise, which are friends 3 and 4.
5) Friend 4 leaves the circle. Next start is friend 5.
6) Count 2 friends clockwise, which are friends 5 and 1.
7) Friend 1 leaves the circle. Next start is friend 3.
8) Count 2 friends clockwise, which are friends 3 and 5.
9) Friend 5 leaves the circle. Only friend 3 is left, so they are the winner.</pre>

<p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> n = 6, k = 5
<strong>Output:</strong> 1
<strong>Explanation:</strong> The friends leave in this order: 5, 4, 6, 2, 3. The winner is friend 1.
</pre>

## My Solution
### Solution 1. Brute Force (Vector)
```cpp
class Solution {
public:
    int findTheWinner(int n, int k) {
        vector<int> friends(n, 0);
        for (int i = 0; i < n; ++i) friends[i] = i+1;
        
        int cnt = n;
        int i = 0;
        while (cnt > 1){
            int curr_n = friends.size();
            int remove_i = (i + k - 1) % curr_n;
            friends.erase(friends.begin() + remove_i);
            i = remove_i % curr_n;
            cnt--;
        }
        
        return friends[0];
    }
};
```