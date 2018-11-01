# Övning 2 pallinda18

This work should be completed before the exercise on **Thursday/Friday 15th/16th November**.

Samtliga uppgifter på kursen ska lämnas in på
[organisationen dd1396-ht18 på KTH GitHub](https://gits-15.sys.kth.se/dd1396-ht18).
Repos för de tre övningarna heter `username-ovnx`.
Användaren `nisse` hittar alltså sitt repo för övning 2 på
[https://gits-15.sys.kth.se/pallinda18/nisse-ovn2](https://gits-15.sys.kth.se/dd1396-ht18/nisse-ovn2).

Gör (minst) en fil per uppgift. Utgå från mallarna i
[/pallinda18/ovn0/](https://github.com/yourbasic/pallinda18/tree/master/ovn0).
Lösningar skall vara inlämnade **innan övningen börjar**.

- Vid övningen ska du vara beredd att muntligt presentera och diskutera dina lösningar och din programkod.
- Uppgifter märkta med HANDIN ska ockå lämnas in skriftligt innan övningens start.

### Homework
Study the following course literature:

- Read the following from the [Step-by-step guide to concurrency](http://yourbasic.org/golang/concurrent-programming/)
  - [Channels](http://yourbasic.org/golang/channels-explained/)
  - [Select](http://yourbasic.org/golang/select-explained/)
  - [Data races](http://yourbasic.org/golang/data-races-explained/)
  - [Detecting data races](http://yourbasic.org/golang/detect-data-races/)
  - [Deadlock](http://yourbasic.org/golang/detect-deadlock/)

### Task 1 - Debugging Concurrent Programs (HANDIN)

Explain what is wrong in the code below, and then fix the code so that all data really passes
through the channel and gets printed.

#### Bug 1
```Go
package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
    ch := make(chan string)
    ch <- "Hello world!"
    fmt.Println(<-ch)
}
```
See: [bug01.go](code/bug01.go) for source code to modify.

#### Bug 2
```Go
package main

import "fmt"

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
    ch := make(chan int)
    go Print(ch)
    for i := 1; i <= 11; i++ {
        ch <- i
    }
    close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
    for n := range ch { // reads from channel until it's closed
        fmt.Println(n)
    }
}
```
See: [bug02.go](code/bug02.go) for source code to modify.

### Task 2 - Many Senders; Many Receivers (HANDIN)

The program [many2many.go](code/many2many.go) contains four producers that together send 32 strings over a channel.
At the other end there are two consumers that receive the strings.
Describe what happens, and explain why it happens, if you make the following changes in the program.
Try first to reason your way through, and then test your hypothesis by changing and running the program.

* What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?
* What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?
* What happens if you remove the statement `close(ch)` completely?
* What happens if you increase the number of consumers from 2 to 4?
* Can you be sure that all strings are printed before the program stops?

Finally, modify the code by adding a new WaitGroup that waits for all consumers to finish.

### Task 3 - Pythia, the Oracle of Delphi (HANDIN)

The code in [oracle.go](code/oracle.go) contains the outline for a program that will answer 'questions'.
Complete the `Oracle` function.  You should not modify the `main` function or other function signatures.
Note that answers should not appear immediately; instead there should be a delay or **pause for thought**. 
Also, the Oracle will still print **helpful predictions** even if there are not any questions.
You may structure your solution into multiple functions.

Your program should contain two channels: One channel for questions, and one for answers and predictions.
In the `Oracle` function you should start three indefinite go-routines.

* A go-routine that receives all questions, and for each incoming question, creates a separate go-routine that answers that question
* A go-routine that generates predictions
* A go-routine that receives all answers and predictions, and prints then to stdout

Whilst the `Oracle` function is the most important of the assignment, you may also want to improve the answer-algorithm.

* The [strings](https://golang.org/pkg/strings/) and [regex](https://golang.org/pkg/regexp/) packages may be of some help
* The program can seem more human if the Oracle prints it answers one character at a time
* Take a look at the story of [ELIZA](https://en.wikipedia.org/wiki/ELIZA)
