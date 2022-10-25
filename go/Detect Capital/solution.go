func detectCapitalUse(word string) bool {

    chars:=strings.Split(word,"")
    
    loopEntered := false
    isLower := false
    
    for i,c := range chars{
        if strings.ToLower(c) == c{
            if !loopEntered {
                isLower = true
            } else {
                if !isLower{
                    return false
                }
            } 
        } else {
            if i == 0 {
                continue
            }
            
            if loopEntered && isLower{
                return false
            }
            
        }
        
        if !loopEntered {
            loopEntered = true
        }
    }
    
    return true
}