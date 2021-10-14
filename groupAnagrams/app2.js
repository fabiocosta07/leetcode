function findSmallestRepeatedSubsequence(primary, secondary) {
    // Write your code here
    if(secondary.search(primary) === -1 ) {
        return - (secondary.length)
    } else {
        var min = primary.length
        loop1:
        for (var i = primary.length; i > 1 ;i--) {            
            if (secondary.search(primary.substr(0,i)) === -1) {
                continue
            }
            const splitArr = secondary.split(primary.substr(0,i))
            for (var j = 0; j < splitArr.length; j++){
                if (splitArr[j].length > 0) {
                    continue loop1
                }
            }            
            min = primary.substr(0,i).length    
        }
        return min
    }

}

console.log(findSmallestRepeatedSubsequence('CATGCATGCATGCATG', 'CATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATGCATG'))