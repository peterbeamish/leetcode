func minPathSum(grid [][]int) int {
    memo := [200][200]int{}
   
    for i := len(grid)-1;i>=0;i--{
        for j := len(grid[0])-1; j>=0; j-- {
            
            if i == len(grid)-1 && j < len(grid[0])-1 {
                // max right
                memo[i][j] = memo[i][j+1] + grid[i][j]
            } else if j == len(grid[0])-1 && i < len(grid)-1 {
                // max bottom
                memo[i][j] = memo[i+1][j] + grid[i][j]
            } else if i != len(grid)-1 && j != len(grid[0]) -1 {
                // middle
                memo[i][j] = grid[i][j] + int(math.Min(float64(memo[i+1][j]),float64(memo[i][j+1])))
            } else {
                // first block
                memo[i][j] = grid[i][j]
            }
        }
    }
    
    
    return memo[0][0]
}