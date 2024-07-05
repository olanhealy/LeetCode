# Notes
Had one issue wehre I said i > 0 so if the input was one element, it would also add a 1 to start whcih is incorrect

# Complexity
> O(n) n being the size of the int array

# Code
```Java
class Solution {
    public int[] plusOne(int[] digits) {

        for ( int i = digits.length - 1; i >= 0; i--) {
            if (digits[i] < 9) {
                digits [i]++;
                return digits;
            }
            digits[i] = 0;
        }
        digits = new int[digits.length + 1];
        digits[0] = 1;
        return digits;

    }
}
```
