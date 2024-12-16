void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n) {
    int index2, index1;
    int index1_begin = 0;
    int be_inserted_num;
    int insert_index;

    for (index2 = 0; index2 < nums2Size; index2++) {
        be_inserted_num = nums2[index2];

        for (index1 = index1_begin; index1 < m + index2; index1++) {
            if (be_inserted_num < nums1[index1]) {
                // nums1中的插入位置索引就是j
                insert_index = index1;
                nums1[m + index2] = nums1[insert_index];
                nums1[insert_index] = be_inserted_num;
                index1_begin = insert_index + 1;
                break;
            }
        }
        // 循环没有提前退出情况处理
        if (index1 == (m + index2)) {
            nums1[m + index2] = be_inserted_num;
            index1_begin = index1 + 1;
        }
    }
}