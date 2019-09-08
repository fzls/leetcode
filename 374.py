# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num):

class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        low = 1
        high = n
        while low <= high:
            i = low + (high - low) // 2
            res = guess(i)
            if res == 0:
                return i
            elif res == -1:
                high = i - 1
            else:
                low = i + 1
