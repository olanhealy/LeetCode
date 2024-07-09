class Solution {
    public int mySqrt(int x) {
        if (x == 0 || x == 1) {
            return x;
        }

        int start = 1;
        int end = x;
        int middle;

        while (start <= end) {
            middle = start + (end - start) / 2;

            long square = (long) middle * middle;

            if (square == x) {
                return middle;
            } else if (square < x) {
                start = middle + 1;
            } else {
                end = middle - 1;
            }
        }

        return end;
    }
}
