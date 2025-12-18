/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
    let n = new ListNode(0);
    let m = n;
    let a = 0;
    while (l1 || l2 || a !== 0) {
        let s = a;
        if (l1) {
            s = s + l1.val;
            l1 = l1.next;
        }
        if (l2) {
            s = s + l2.val;
            l2 = l2.next;
        }
        a = Math.floor(s / 10);
        m.next = new ListNode(s % 10);
        m = m.next;
    }
    return n.next;
};
