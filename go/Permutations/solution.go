func permute(nums []int) [][]int {
    results := make([][]int,0)
    
      
    backtrack(len(nums), nums, &results,0)
    
    return results
    
}

func backtrack(size int, nums []int,results *[][]int,first int) {
    
    if  first == size {
        r := deepCopySlice(nums)
        *results = append(*results,r)
    }
    
    for i :=first; i<size; i++ {
        
        nums[i], nums[first]= nums[first], nums[i]
        
        backtrack(size, nums, results, first + 1)
        
        // backtrack
        nums[i], nums[first]= nums[first], nums[i]
        
    }
    
}

func deepCopySlice(nums []int) []int{
    r := make([]int,len(nums))
    copy(r,nums)
    return r
}