import (
    "strings"
)

func capitalizeTitle(title string) string {
    
    output := ""
    
    words := strings.Split(title," ")
    
    for _,w := range words {
        lower := strings.ToLower(w)
        
        if len(lower)>2{
            // Capitalize the first letter
            firstLetter := lower[0:1]
            firstLetter = strings.ToUpper(firstLetter)
            
            rest := lower[1:len(lower)]
            
            lower = firstLetter+rest
            
        }
        
        output+=lower+" "
        
    }
    
    output = output[0:len(output)-1]
    
    return output
    
}