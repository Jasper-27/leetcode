var isPalindrome = function(x) {

    sx = x.toString()

    rev =  sx.split("").reverse().join("");

    if (rev == sx){
        return true
    }else{
        return false
    }
    
};


console.log(isPalindrome(102))
console.log(isPalindrome(202))
console.log(isPalindrome(-1001))