package algorithms 

func getRow(i int) []int {
    
    var ret []int
    if i == 0 {
        return append(ret,1)
    }
    for j:=0;j<=i;j++{
        ret = append(ret,pascalElement(i,j))
    }
    
    return ret
}

func pascalElement(i,j int) int {
    if j == i || j==0 || i ==0 {
        return 1
    } else{
        return pascalElement(i-1,j-1)+pascalElement(i-1,j)
    }
    
    
}