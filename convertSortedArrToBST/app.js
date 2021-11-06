 function TreeNode(val, left, right) {
      this.val = (val===undefined ? 0 : val)
      this.left = (left===undefined ? null : left)
      this.right = (right===undefined ? null : right)
  }

  /**
 * @param {number[]} nums
 * @return {TreeNode}
 */
 var sortedArrayToBST = function(nums) {
     return getMiddle(nums)
 };

 function getMiddle(nums) {
    if (nums.length === 0) {
        return 
    }
    var length = nums.length
    if (length % 2 != 0) {
        length--
    } 
    var index =  length/2
    var half1 = nums.slice(0, index)
    var half2 = nums.slice(index + 1)
    var node = new TreeNode(nums[index], getMiddle(half1), getMiddle(half2))
    return node    
 }

var input = [-10,-3,0,5,9]

console.log(sortedArrayToBST(input))