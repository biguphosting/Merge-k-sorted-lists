/**
 * Definition for singly-linked list.
  *type ListNode struct {
  *    Val int
  *    Next *ListNode
  *}**/
 
// The basic approach is to disarm the whole given structure, get advantage of the sorting methods and then rearm this.
// I'll be putting all the values into a []int, sort them and then create a new *ListNode that will satisfy the answer
// I had to ask for help to walk the tree and @Mi-Br showed me how to do the walk on a tree.  
// Check the solution here: https://github.com/Mi-Br/Data-Structures/blob/main/GO/BST/bst/bst.go

// Base function provided by leetcode. Can't change this
  func mergeKLists(lists []*ListNode) *ListNode {
    // This is the []int I'll be using to store the results and aux is a *ListNode I'll use to move around the given values
    var res []int
    var aux *ListNode
     
    for i:=0;i<len(lists);i++{
        as we advance on every array inside the []*ListNode , I save each value into aux.
        // Just to make sure and explain something here. When you fmt.Println(lists)  you get (in the given example) 
        // three arrays. But each array has a single ListNode showing you the first value and the address of the next one. 
        // So this is like a hybrid walk on the struc.  First as an array, then as a binary tree.
        aux=lists[i]

        // I've always feared this kind of loops but this is the way to go
        for ;;{
            // First, check if aux has something in it. This was added later when submitting the solution 
            // to leetcode because they were tricky on the testcases providing you of [[]] on the second test case and of a [] on the third case.
            if aux!=nil{
            
            // Add the values only to the []int
            res=append(res,aux.Val)
            // And check if the node has the address of the Next node, move there and repeat.
            // If there's no address (it's nil) then break the loop and go back to the for loop
            if aux.Next!=nil{      
                aux=aux.Next
            } else {
                break
            }
            }else{
                break
            }
        }
    }

    // Once out, sort the []int
    sort.Ints(res)

    // Here, I have to check if the arrays given have a single value because it's Next address will be nil
    // So, if it has a single value, just assign the ListNode to a var called last, return it.  And return an empty last variable
    // if the given values are empty (remember the provided [] and [[]]).
    // To walk the []int, I'll be running it backwards so I don't have to worry for the nil address of the last one.
    if len(res) > 1{
        // assign value to variable last and assign the address of last item to a var called temp
        last:=ListNode{Val:res[len(res)-1],Next:nil}
        temp:=&last
        for i:=len(res)-2;i>=0;i--{
            // Here I'll be running backwards, creating a newNode with it's value and assigning the temp address to the Next of this Node.
            // After that, then capture this new node's address
            newNode:=ListNode{res[i],temp}
            temp=&newNode
        }
        // After that, return temp, which is already a *ListNode - The End.
    return temp
    } else if len(res)==1 {
        last:=ListNode{Val:res[0], Next:nil}
        return &last
    }else{
        var last *ListNode
        return last
        
    }
}