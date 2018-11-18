# Övning 3 pallinda18

This work should be completed before the exercise on **Friday 30th November**.

Samtliga uppgifter på kursen ska lämnas in på
[organisationen dd1396-ht18 på KTH GitHub](https://gits-15.sys.kth.se/dd1396-ht18).
Repos för de tre övningarna heter `username-ovnx`.
Användaren `nisse` hittar alltså sitt repo för övning 3 på
[https://gits-15.sys.kth.se/pallinda18/nisse-ovn3](https://gits-15.sys.kth.se/dd1396-ht18/nisse-ovn3).

Gör (minst) en fil per uppgift. Utgå från mallarna i
[/pallinda18/ovn0/](https://github.com/yourbasic/pallinda18/tree/master/ovn0).
Lösningar skall vara inlämnade **innan övningen börjar**.

- Vid övningen ska du vara beredd att muntligt presentera och diskutera dina lösningar och din programkod.
- Uppgifter märkta med HANDIN ska ockå lämnas in skriftligt innan övningens start.

### Homework
Study the following course literature:

- Read the following from the [Step-by-step guide to concurrency](http://yourbasic.org/golang/concurrent-programming/)
  - [Mutual exclusion](http://yourbasic.org/golang/mutex-explained/)
  - [Efficient parallel computation](http://yourbasic.org/golang/efficient-parallel-computation/)

### Task 1 - Matching Behaviour (HANDIN)

Take a look at the program [matching.go](code/matching.go). Explain what happens and why it happens if you make the following changes. Try first to reason about it, and then test your hypothesis by changing and running the program.

  * What happens if you remove the `go-command` from the `Seek` call in the `main` function?
  * What happens if you switch the declaration `wg := new(sync.WaitGroup`) to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?
  * What happens if you remove the buffer on the channel match?
  * What happens if you remove the default-case from the case-statement in the `main` function?

> **Hint:** Think about the order of the instructions and what happens with arrays of different lengths.

### Task 2 - Fractal Images (HANDIN)

The file [julia.go](code/julia.go) contains a program that creates images and writes them to file. The program is pretty slow. Your assignment is to divide the computations so that they run in parallel on all available CPUs. Use the ideas from the example in the [efficient parallel computation](http://yourbasic.org/golang/efficient-parallel-computation/) section of the course literature.

You can also make changes to the program, such as using different functions and other colourings.

How many CPUs does you program use? How much faster is your parallel version?

> **Hint:** In more recent versions of Golang (since 1.5), the runtime will default to use as many operating system threads as it is allowed. To see differences in behaviour, refer to the [GOMAXPROCS](https://golang.org/pkg/runtime/#GOMAXPROCS) function and vary the value.

### Task 3 - Weather station (HANDIN)

The file [server.go](code/server.go) contains a program that simulates three independent weather stations that show the temperature at KTH. The results are published at the following addresses once the serves are operational:

  * `http://localhost:8080`
  * `http://localhost:8081`
  * `http://localhost:8082`

Start the program and try to visit the three addresses in your browser. You'll soon find that the three services aren't very reliable; they're pretty slow and sometimes you get no response at all. You might also get the error message "Service unavailable".

Your assignment is to write a client that simultaneously asks all servers and terminates the search as soon as one has responded with a correct temperature. The request should also terminate if no-one has answered within a given time. The file [client.go](code/client.go) contains a template from which you should build on.

  * Read through the code and start the client whilst the weather stations are operational
  * Implement the function `MultiGet`
