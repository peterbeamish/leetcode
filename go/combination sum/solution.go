type BT struct {
    Results [][]int
    Comb []int
    Candidates []int
}

func combinationSum(candidates []int, target int) [][]int {
        
    
    bt := &BT {
        Results: make([][]int,0),
        Comb: make([]int,0),
        Candidates: candidates,
    }
    
    bt.backtrack(target,0)
    
   // fmt.Println(bt.Results)
    
    return bt.Results
}

func (bt *BT) backtrack(remain int, index int) {
    
    // We found an answer, record
    if remain == 0 { 
        copyComb := make([]int,len(bt.Comb))
        copy(copyComb,bt.Comb)
        
        bt.Results = append(bt.Results, copyComb)
        //fmt.Println("good combo")
        //fmt.Println(bt.Comb)
    } else if  remain < 0 {
        // bad branch, return without append
        //fmt.Println("bad combo")
        //fmt.Println(bt.Comb)
        return
    }
    
    for i := index; i < len(bt.Candidates); i++ {
        
        bt.Comb = append(bt.Comb,bt.Candidates[i])        
        
        // recurse
        newRemains := remain - bt.Candidates[i]
        ///fmt.Println(bt.Comb)
        bt.backtrack(newRemains,i)
        
        // remove last
        bt.Comb = bt.Comb[:len(bt.Comb)-1]
        
    }
    
}