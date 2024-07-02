class Solution {
    public static boolean isAnagram(String s, String t) {

        char sNew[] = s.toCharArray();
        char tNew[] = t.toCharArray();

        if (sNew.length != tNew.length) {
            return false;
        }
        Arrays.sort(tNew);
        Arrays.sort(sNew);


        return Arrays.equals(sNew, tNew);
    }
}
