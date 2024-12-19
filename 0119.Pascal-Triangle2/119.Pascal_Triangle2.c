/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

int getRowNum(int rowIndex, int index) {
    if (rowIndex == 0) {
        return 1;
    }
    if (index == 0 || index == rowIndex) {
        return 1;
    }
    return getRowNum(rowIndex - 1, index) + getRowNum(rowIndex - 1, index - 1);
}

int* getRow(int rowIndex, int* returnSize) {
    int* ret = malloc(sizeof(int) * (rowIndex + 1));
    *returnSize = rowIndex + 1;
    for (int i = 0; i < (rowIndex + 1); i++) {
        if (i == 0 || i == rowIndex) {
            ret[i] = 1;
        } else {
            ret[i] = getRowNum(rowIndex, i);
        }
    }
    return ret;
}