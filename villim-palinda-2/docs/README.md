If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

#### Task 1

##### Buggy Code 1
1. What is wrong: I want this program to print "Hello world!", but it doesn't work. -> Deadlock -> Once we put something into the ch it will wait until another routine takes it out.
In the case of this bug the same go routine tries to send and recive. Since there are no other gorutines we get -> Deadlock.
2. How it was fixed: Created a separate goroutine. When a goroutine sends a message to a channel, it waits for another goroutine to receive it. I tried Closure to get a separate goroutine and it seems to work.

##### Buggy Code 2
1. What is wrong: This program should go to 11, but it only prints 1 to 10. Main finishes before print functions gets to print all values. Once the main is done it closes all others and we dont get the last number printed.
2. How it was fixed: via waitgroup -> we get main to wait for 1 extra go routine to finnish. Then before main returns it calls wait_group.Wait() to wait for every goroutine to finish. I added defer wait_group.Done() to Print function to indicate when the Print function is done.

#### Task 2

| Question | What I expected | What happened | Why I believe this happened |
|---|---|---|---|
| What happens if you do X? | Program would still work as before | Program ended up in a deadlock | Because of reasons 🤷 |
| What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function? | WaitGroup is used to ensure that all processes finish before we close the channel. I expect an error to happen if we change the order | Error -> panic: send on closed channel | Main closing the channel before producer had enough time to send. Since ch is closed, no more data can be sent on the channel. When they try to send data to a closed channel, we get an error.|
| What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce? | It will crash because the ch is closed prematurely on another thread | panic: send on closed channel -> It runs for a while and prints before we get an error message | If each Produce function closes the ch -> the first producer that finishes will close the ch that all the others use! When any other processes try to send data to a closed channel, it will crash.|
| What happens if you remove the statement close(ch) completely? | The code will run as normal | Nothing happens | Go has a garbage collector, so it will automatically close the channel. We don't always have to close the channel manually, but it is good practice to do so. |
| What happens if you increase the number of consumers from 2 to 4? | The code should run as normal. I expect no errors, and it could run faster since we now have a balance of 4 producers to 4 consumers | Code ran without errors and it did run faster! | More consumers means that we can divide workload more efficiently between goroutines. |
| Can you be sure that all strings are printed before the program stops? | I think we can NOT be sure | Tried to count... | Since Print sleeps for 10 milliseconds, we have this delay between printing. When main sends all elements into ch, it closes the ch. Because we have a delay in print, this means we can't be sure we printed everything in time before ch was closed. |
