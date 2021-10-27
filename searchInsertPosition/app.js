var searchInsert = function(nums, target) {

    if (target > nums[nums.length - 1]) {
        return nums.length
    } 
    const tree = new BinarySearchTree();
    for (let i = 0; i < nums.length; i++) {
        tree.insert(nums[i])
    }
    return tree.lookup(target)
};

class Node {
    constructor(value) {
        this.left = null
        this.right = null
        this.value = value
    }    
}

class BinarySearchTree {
    constructor() {
        this.root = null
    }
    insert(value){
        const newNode = new Node(value)
        if (!this.root) {
            this.root = newNode
            return this
        } 
        let cursorNode = this.root;
        while (true) {
            if (value < cursorNode.value) {
                 if (!cursorNode.left) {
                    cursorNode.left = newNode
                    return this
                 }
                 cursorNode = cursorNode.left               
            } else {
                if (!cursorNode.right) {
                    cursorNode.right = newNode
                    return this
                }
                cursorNode = cursorNode.right
            }
        }
    }
    lookup(value) {
        if (!this.root) {
            return false;
        }
        let cursorNode = this.root
        let index = 0
        while(cursorNode) {
            if (value < cursorNode.value) {
                cursorNode = cursorNode.left
            } else if (value > cursorNode.value) {
                cursorNode = cursorNode.right
            } else {
                return index
            }
            if (!cursorNode) {
                return index
            }             
            index++           
        }
    }    
}

const test1 = [1,3,5,6]

const result = searchInsert(test1,5)

console.log(result)