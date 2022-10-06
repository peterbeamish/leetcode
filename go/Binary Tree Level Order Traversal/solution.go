/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type LevelPair struct {
    Node *TreeNode
    Level int
}


func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var levels [][]int = make([][]int,0)
    
    q := make([]*LevelPair,0)
    
    // prime the q by putting the root in the q
    q = append(q, &LevelPair{
        Node: root,
        Level: 0,
    })
    
    currentLevel := 0
    
    // while q has elem, dequeue. Add child to collection. Add child back to q for next level
    for len(q) != 0 {
        
        item := q[0]
        q = q[1:]
        currentLevel = item.Level
        
        if len(levels)-1 < currentLevel{
            levels = append(levels,[]int{})
        }
        
        levels[currentLevel] = append(levels[currentLevel],item.Node.Val)
        
        leftPair := &LevelPair{
            Node: item.Node.Left,
            Level: currentLevel+1,
        }
        
        rightPair := &LevelPair{
            Node: item.Node.Right,
            Level: currentLevel+1,
        }
        
        if leftPair.Node != nil {
            q = append(q,leftPair)    
        }
        
        if rightPair.Node != nil {
            q = append(q,rightPair)    
        }
        
    }
    return levels
}