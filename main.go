/**
 * Definition for singly-linked list.
  *type ListNode struct {
  *    Val int
  *    Next *ListNode
  *}**/
 


// Base function provided by leetcode. Can't change this
  func mergeKLists(lists []*ListNode) *ListNode {
    var res []int
    var aux *ListNode
     
    for i:=0;i<len(lists);i++{
        aux=lists[i]
        for ;;{
            if aux!=nil{
            res=append(res,aux.Val)
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
    sort.Ints(res)
    if len(res) > 1{
        last:=ListNode{Val:res[len(res)-1],Next:nil}
        temp:=&last
        for i:=len(res)-2;i>=0;i--{
            newNode:=ListNode{res[i],temp}
            temp=&newNode
        }
    return temp
    } else if len(res)==1 {
        last:=ListNode{Val:res[0], Next:nil}
        return &last
    }else{
        var last *ListNode
        return last
        
    }
}