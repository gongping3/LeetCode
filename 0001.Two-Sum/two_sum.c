class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        std::unordered_map<int, int> hashmap;
        
        for (int i = 0; i < nums.size(); ++i) {
            if (hashmap.find(target - nums[i]) != hashmap.end()) {
                return vector<int>{hashmap[target - nums[i]], i};
            } else {
                hashmap[nums[i]] = i;
            }
        }
        return {};
    }
};