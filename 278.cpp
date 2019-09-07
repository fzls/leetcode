// 本题可选语言中没有go，就用cpp来实现吧
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
  public:
    int firstBadVersion(int n) {
      int low = 1;
      int high = n;
      while (low <= high) {
        int i = low + (high - low) / 2;
        if (isBadVersion(i)) {
          high = i - 1;
        } else {
          low = i + 1;
        }
      }

      return high - 1;
    }
};