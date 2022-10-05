func maxArea(height []int) int {
    max := 0
    leftPos := 0
    rightPos := len(height)-1
        
    for leftPos != rightPos {
        area := (rightPos -leftPos) * int(math.Min(float64(height[leftPos]),float64(height[rightPos])))
        if area > max {
            max = area
        }
        
        if height[leftPos] < height[rightPos] {
            leftPos++
        } else {
            rightPos--
        }
    }
    
    return max
}