package algorithms

func reverseStringArrayRecursive(s []byte)  {
    reverseFromIndex(s,len(s)-1)
}

func reverseFromIndex(s []byte, i int){
    if s == nil || i <= 0 {
        return 
    }
    tmp := s[i]
    reverseFromIndex(s,i-1)
    for j:=i-1;j>=0;j-- {
        s[j+1] = s[j]
    }
    s[0] = tmp
}

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