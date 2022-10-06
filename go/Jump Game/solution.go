/*
 *      root
 *    3      1
 *  1 1 4     1
 * 1  4   e    1 
 *              4
 */

// Bottom up greedy

func canJump(nums []int) bool {
    lastIndex := len(nums) -1
    lastPos := lastIndex
    
    for i:=lastIndex-1; i>=0;i-- {
        if nums[i]+i >= lastPos {
            lastPos = i
        }
    }
    
    return lastPos == 0
}
