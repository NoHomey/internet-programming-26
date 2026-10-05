# Task 1

Add a function for comparing if two BST(Binary Search Tree)s are equivalent to `lecture-3/equiv-bsts/main.go`.

# Task 2

Try to fix the select stament from `lecture-3/select/main.go`.

Hint: Reading from closed channel is always a ready channel operation and returns `<zero value>, true` (in this case `0, true`).
Which means that such case in a select statement can always be selected.
So we want to disable cases of already consumed (fully read and closed) channels.
If all three channels are closed we obviously want to exit the `for` loop and return the `s` the sum.
