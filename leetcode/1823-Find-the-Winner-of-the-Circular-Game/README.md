# [1823. Find the Winner of the Circular Game](https://leetcode.com/problems/find-the-winner-of-the-circular-game/)

There are `n` friends that are playing a game. The friends are sitting in a circle and are numbered from 1 to n in **clockwise order**. More formally, moving clockwise from the ith friend brings you to the <code>(i+1)<sup>th</sup></code> friend for `1 <= i < n`, and moving clockwise from the nth friend brings you to the 1st friend.

The rules of the game are as follows:

1. **Start** at the <code>1<sup>st</sup></code> friend.
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

![example1](./example1.png)

<p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> n = 6, k = 5
<strong>Output:</strong> 1
<strong>Explanation:</strong> The friends leave in this order: 5, 4, 6, 2, 3. The winner is friend 1.
</pre>

## My Solutions
### Solution 1. Brute Force (Vector)
- Time complexity: O(N)
- Space complexity: O(N)

#### C++
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

### Solution 2. Brute Force (Circular Linked List)
- Time complexity: O(N)
- Space complexity: O(N)

#### C++
```C++
struct MyListNode {
    int val;
    MyListNode *next;
    MyListNode() : val(0), next(NULL) {}
    MyListNode(int x) : val(x), next(NULL) {}
};


class CircularList {
    int n;
    int k;
    MyListNode* head;
    MyListNode* tail;
    
public:
    CircularList() : n(0), k(0), head(NULL) {}
    ~CircularList() {
        tail->next = NULL;
        tail = NULL;
        while (head){
            MyListNode* curr = head;
            head = head->next;
            delete curr;
        }
    }
    
    CircularList(int N, int K) : n(N), k(K) {
        init();
    }
    
    void init() {
        MyListNode* res = new MyListNode();
        MyListNode* curr = res;
        for (int i = 1; i <= n; ++i){
            curr->next = new MyListNode(i);
            curr = curr->next;
        }
        head = res->next;
        curr->next = res->next;
        tail = curr;
    }
    
    MyListNode* getHead(){ return head; }
    MyListNode* getTail(){ return tail; }
    
    void kill(){
        int cnt = n;
        MyListNode* start = tail;
        while (cnt > 1){
            MyListNode* curr = start;
            for (int i = 1; i < k; ++i) curr = curr->next;
            
            MyListNode* dead = curr->next;
            
            curr->next = curr->next->next;
            
            delete dead;
            cnt--;
            start = curr;
        }
        head = start;
        tail = start;
    }
};


class Solution {
public:
    int findTheWinner(int n, int k) {
        CircularList cl(n, k);
        MyListNode* curr = cl.getHead();
        
        cl.kill();
        return cl.getHead()->val;
    }
};
```

### Solution 3. Josephus problem
- Time complexity: O(N)
- Space complexity: O(1)

#### C++ (recursive solution)
- General term: `W(n, k) = ( W(n-1, k) + k ) mod n` (index base: 0)
- Base term: `W(1, k) = 0` (index base: 0)
- To find the surviver in the <code>n<sup>th</sup></code> game can simply transform into the question of finding the next wordsman in the <code>n-1<sup>th</sup></code> game.

```cpp
class Solution {
public:
    int findTheWinner(int n, int k) {
        return n == 1 ? 1 : (findTheWinner(n-1, k) + k-1) % n + 1;
    }
};
```