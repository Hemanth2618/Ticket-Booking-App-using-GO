# Ticket-Booking-App-using-GO

I developed a CLI based Ticket booking application using GO programming language. This design illustrates some key concepts/data structures of GO language such as Maps, Structs, Go routines for concurrency, WaitGroup for synchronizing goroutines.
Channels 

Achieving concurrency is very easy in Go when compared to other languages like Java. Go creates light weight threads called "Go routines". These are cheaper and lightweight. Hundreds of thousands of of goroutines can be run without affecting the performance.

The Go runtime scheduler employs an algorithm called work stealing to efficiently distribute the execution of goroutines among a limited number of operating system threads. It ensures that all available processor cores are utilized effectively.
By multiplexing goroutines onto a smaller set of threads, Go can achieve efficient concurrency and parallelism while keeping the memory footprint and overhead associated with operating system threads relatively low. Goroutines are lightweight and have a small stack size compared to traditional threads, which allows for the creation of thousands or even millions of goroutines without significant resource consumption.
