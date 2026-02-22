class Solution:
    def binaryGap(self, n: int) -> int:
        list1 = [i for i, bit in enumerate(format(n, 'b')) if bit == '1']
        return max([list1[i]-list1[i-1] for i in range(1, len(list1))], default=0)


        