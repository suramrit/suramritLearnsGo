package algorithms


func reverseString(s []byte)  {
    reverseFromIndex(s,0,len(s)-1)
}

func reverseFromIndex(s []byte, start,end int){
    if(start >= end || s==nil) {
        return 
    }
    //swap
    tmp := s[start]
    s[start] = s[end]
    s[end] = tmp 
    reverseFromIndex(s,start+1, end-1)
}