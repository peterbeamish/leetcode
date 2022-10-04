
type state int

const (
    WHITE state = 0
    GRAY state = 1
    BLACK state = 2
)


func eventualSafeNodes(graph [][]int) []int {
    
    // Allocate a collection sized to the number of nodes in the graph
    stateList := make([]state, len(graph))
  
    var ret []int
    
    // foreach node in the graph, apply dfs
    for i := 0; i<len(graph); i++ {
        // dfs returns bool. True if all branches lead to leaf. False in case of loop.
        if dfs(i, graph, stateList) {
            ret = append(ret,i)
        }
    }

    return ret        
}

func dfs(root int, graph [][]int, stateList[]state) bool {    
    // during dfs, each node node in the state collection will have its state updated
    // possible states include: unvisited, possible-safe, looped
        
    if stateList[root] != WHITE {
        return stateList[root] == BLACK
    }

    children := graph[root]

    stateList[root] = GRAY
    
    for _, child := range children {
        
        // optimize visited nodes
        if stateList[child] == BLACK {
            continue
        }
        
        // recurse
        if stateList[child] == GRAY || !dfs(child,graph,stateList) {
            return false
        }
    }
    
    stateList[root] = BLACK
    
    return true
}