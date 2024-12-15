/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
    for (int i = digitsSize - 1; i >= 0; i--)
    {
        if (digits[i] != 9)
        {
            digits[i] += 1;
            *returnSize = digitsSize;
            return digits;
        }
        digits[i] = 0;
    }
    int *res = (int*)malloc(sizeof(int) * (digitsSize + 1));
    memset(res, 0, sizeof(int) * (digitsSize + 1));
    res[0] = 1;
    *returnSize = digitsSize + 1;
    return res;
}