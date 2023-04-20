If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

  * What happens if you remove the `go-command` from the `Seek` call in the `Seek` function?

We always get the same results because it goes through the slice in the same way. Anna always sends to Bob and so on. However, if we use goroutines to create 5 concurrent processes, the matching becomes random. There is always one person left who can't send to anybody.

If we remove the go command from the `Seek` call in the `Seek` function, the `Seek` function will be executed sequentially instead of concurrently in separate goroutines.
`Seek` call will block until the previous one completes, and the program will always try to match person with the one to their left in the people array. Go runtime schedules goroutines randomly or **non-deterministically**.

  * What happens if you switch the declaration `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?

When we use `wg := new(sync.WaitGroup)`, it has a built-in pointer, so we can pass it without `&` as a parameter to the go `Seek` function. However, if we change it to `var wg sync.WaitGroup`, we have to include a `pointer` when passing it to the `Seek` function.

Each `Seek` goroutine will  run its own copy of the `WaitGroup`. => `wg.Done()` call won't count correctly and the program will end in a deadlock. The `Seek` function passes the `WaitGroup` **by value** because `sync.WaitGroup` contains `sync.noCopycopylocks`. 
"By preventing a wait group from being copied, Go ensures that there is only one wait group instance that is used by all goroutines that need to wait on it."-internet.

The program will output:
 'Eva sent a message to Dave'
 'Cody sent a message to Bob'. 
'fatal error: all goroutines are asleep - deadlock!', 

The `wg.Wait()` in the main function won't return even after all `Seek` routines are done, because the counter will still be non-zero.

  * What happens if you remove the buffer on the channel match?

An uneven number of people means that not every person will receive a message. 
A buffer allows a channel `ch` to store one element without the need for it to be immediately removed. 
However, removing the `buffer` from the channel will cause a block until someone receives the element. 
If no other routines exist to take the element from the channel, the code will end up in a `deadlock`.

For example:
 Anna sent a message to Eva 
 Cody sent a message to Bob
 fatal error: all goroutines are asleep - deadlock!

 If there are more senders than receivers, or if the channel has a limited buffer size, some messages may not be received. If  no routines exist to take the messages out of the channel, the code may end up in a `deadlock`.

* What happens if you remove the default-case from the case-statement in the `main` function?

The default statement in the main program is important in case we have an **even** number of people. 
However, if we have an odd number of people, we can remove the default statement without consequences.

If we don't have a default statement, the only case in which the program will try to take an element out of the channel is when there are an even number of people. In this case, there will be no elements left in the channel to be taken out, and the code will end up in a deadlock.

In case of odd number of people:
Cody sent a message to Eva.
Bob sent a message to Dave.
No one received Annaâ€™s message.
Nohing -> Ran as normal..

$$
\begin{array}{|c|r|}
\hline \text { Variant } & \text { Runtime (ms) } \\
\hline \text { singleworker } & 1045 \\
\hline \text { mapreduce } & 568 \\
\hline
\end{array}
$$