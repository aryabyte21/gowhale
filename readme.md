
## How to start with GoLang
- Introduction
- hi hello
- Hello this in intro

Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language. Programs are assembled by using packages, for efficient management of dependencies. This language also supports environment adopting patterns alike to dynamic languages. For eg., type inference (y := 0 is a valid declaration of a variable y of type float).

text editors for go
Text Editor: Text editor gives you a platform where you write your source code. Following are the list of text editors:  

Windows notepad
OS Edit command
Brief
Epsilon
vm or vi
Emacs
VS Code
Finding a Go Compiler: Go distribution comes as a binary installable for FreeBSD (release 8 and above), Linux, Mac OS X (Snow Leopard and above), and Windows operating systems with 32-bit (386) and 64-bit (amd64) x86 processor architectures. 

 Beginning with Go programming

There are various online IDEs such as The Go Playground, repl.it, etc. which can be used to run Go programs without installing. 
hello
First, let’s talk about C syntax. C took an unusual and clever approach to declaration syntax. Instead of describing the types with special syntax, one writes an expression involving the item being declared, and states what type that expression will have. Thus

int x;
declares x to be an int: the expression ‘x’ will have type int. In general, to figure out how to write the type of a new variable, write an expression involving that variable that evaluates to a basic type, then put the basic type on the left and the expression on the right.

Thus, the declarations
The pointer in C language is a variable which stores the address of another variable. This variable can be of type int, char, array, function, or any other pointer. The size of the pointer depends on the architecture. However, in 32-bit architecture the size of a pointer is 2 byte.

int *p;
int a[3];
In C, we can divide a large program into the basic building blocks known as function. The function contains the set of programming statements enclosed by {}. A function can be called multiple times to provide reusability and modularity to the C program. In other words, we can say that the collection of functions creates a program. The function is also known as procedureor subroutinein other programming languages.
Following is another example:  


package main
import "fmt"
 
func main() {
   fmt.Println("1 + 1 =", 1 + 1)
}
Output:  

1 + 1 = 2
Explanation of the above program: 
In this above program, the same package line, the same import line, the same function declaration and uses the same Println function as we have used in 1st GO program. This time instead of printing the string “Hello, geeksforgeeks” we print the string 1 + 1 = followed by the result of the expression 1 + 1. This expression is made up of three parts: the numeric literal 1 (which is of type int), the + operator (which represents addition) and another numeric literal 1. 

Why this “Go language”?

Because Go language is an effort to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aims to be modern, with support for networked and multicore computing. 

What excluding in Go which is present in other languages? 

Go attempts to reduce the amount of typing in both senses of the word. Throughout its design, developers tried to reduce clutter and complexity.
There are no forward declarations and no header files; everything is declared exactly once.
Stuttering is reduced by simple type derivation using the := declare-and-initialize construct.
There is no type hierarchy: types just are, they don’t have to announce their relationships.
Hardware Limitations

We have observed that in a decade, the hardware and processing configuration is changing at a very slow rate. In 2004, P4 was having the clock speed of 3.0 GHz and now in 2018, Macbook pro has the clock speed of Approx (2.3Ghz v 2.66Ghz). To speed up, the functionality we use more processors, but using more processor the cost also increases. And due to this we use limited processors and using limited processor we have a heavy programming language whose threading takes more memory and slows down the performance of our system. Hence, to overcome such problem Golang has been designed in such a way that instead of using threading it uses Goroutine, which is similar to threading but consumes very less memory. 
Like threading consumes 1MB whereas Goroutine consumes 2KB of memory, hence at the same time, we can have millions of goroutine triggered. 
So the above-discussed point makes golang a strong language that handles concurrency like C++ and Java. 
State that p is a pointer to int because ‘*p’ has type int, and that a is an array of ints because a[3] (ignoring the particular index value, which is punned to be the size of the array) has type int.

What about functions?
In C, we can divide a large program into the basic building blocks known as function. The function contains the set of programming statements enclosed by {}. A function can be called multiple times to provide reusability and modularity to the C program. In other words, we can say that the collection of functions creates a program. The function is also known as procedureor subroutinein other programming languages.

In c, we can divide a large program into the basic building blocks known as function. The function contains the set of programming statements enclosed by {}. A function can be called multiple times to provide reusability and modularity to the C program. In other words, we can say that the collection of functions creates a program. The function is also known as procedureor subroutinein other programming languages.

 


Again, we see that main is a function because the expression main(argc, argv) returns an int. In modern notation we’d write

int main(int argc, char *argv[]) { /* ... */ }
but the basic structure is the same.

This is a clever syntactic idea that works well for simple types but can get confusing fast. The famous example is declaring a function pointer. Follow the rules and you get this:

int (*fp)(int a, int b);
Here, fp is a pointer to a function because if you write the expression (*fp)(a, b) you’ll call a function that returns int. What if one of fp’s arguments is itself a function?

int (*fp)(int (*ff)(int x, int y), int b)
That’s starting to get hard to read.

Of course, we can leave out the name of the parameters when we declare a function, so main can be declared

int main(int, char *[])
Recall that argv is declared like this,

char *argv[]
so you drop the name from the middle of its declaration to construct its type. It’s not obvious, though, that you declare something of type char *[] by putting its name in the middle.

And look what happens to fp’s declaration if you don’t name the parameters:

int (*fp)(int (*)(int, int), int)
Not only is it not obvious where to put the name inside

go is a good language
int (*)(int, int)
it’s not exactly clear that it’s a function pointer declaration at all. And what if the return type is a function pointer?

int (*(*fp)(int (*)(int, int), int))(int, int)
It’s hard even to see that this declaration is about fp.

You can construct more elaborate examples but these should illustrate some of the difficulties that C’s declaration syntax can introduce.

There’s one more point that needs to be made, though. Because type and declaration syntax are the same, it can be difficult to parse expressions with types in the middle. This is why, for instance, C casts always parenthesize the type, as in

(int)M_PI
Go syntax
Languages outside the C family usually use a distinct type syntax in declarations. Although it’s a separate point, the name usually comes first, often followed by a colon. Thus our examples above become something like (in a fictional but illustrative language)

x: int
p: pointer to int
a: array[3] of int
These declarations are clear, if verbose - you just read them left to right. Go takes its cue from here, but in the interests of brevity it drops the colon and removes some of the keywords:

x int
p *int
a [3]int
There is no direct correspondence between the look of [3]int and how to use a in an expression. (We’ll come back to pointers in the next section.) You gain clarity at the cost of a separate syntax.

Now consider functions. Let’s transcribe the declaration for main as it would read in Go, although the real main function in Go takes no arguments:

func main(argc int, argv []string) int
Superficially that’s not much different from C, other than the change from char arrays to strings, but it reads well from left to right:

function main takes an int and a slice of strings and returns an int.

Drop the parameter names and it’s just as clear - they’re always first so there’s no confusion.

func main(int, []string) int
One merit of this left-to-right style is how well it works as the types become more complex. Here’s a declaration of a function variable (analogous to a function pointer in C):



f func(func(int,int) int, int) func(int, int) int
It still reads clearly, from left to right, and it’s always obvious which name is being declared - the name comes first.

The distinction between type and expression syntax makes it easy to write and invoke closures in Go:

sum := func(a, b int) int { return a+b } (3, 4)
Pointers
Pointers are the exception that proves the rule. Notice that in arrays and slices, for instance, Go’s type syntax puts the brackets on the left of the type but the expression syntax puts them on the right of the expression:

var a []int
x = a[1]
For familiarity, Go’s pointers use the * notation from C, but we could not bring ourselves to make a similar reversal for pointer types. Thus pointers work like this

var p *int
x = *p
We couldn’t say

var p *int
x = p*
because that postfix * would conflate with multiplication. We could have used the Pascal ^, for example:

var p ^int
x = p^
and perhaps we should have (and chosen another operator for xor), because the prefix asterisk on both types and expressions complicates things in a number of ways. For instance, although one can write

[]int("hi")
as a conversion, one must parenthesize the type if it starts with a *:

(*int)(nil)
Had we been willing to give up * as pointer syntax, those parentheses would be unnecessary.

So Go’s pointer syntax is tied to the familiar C form, but those ties mean that we cannot break completely from using parentheses to disambiguate types and expressions in the grammar.

Overall, though, we believe Go’s type syntax is easier to understand than C’s, especially when things get complicated.

Notes
Go’s declarations read left to right. It’s been pointed out that C’s read in a spiral! See The “Clockwise/Spiral Rule” by David Anderson.


## Getting started with servers and GoLang
The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.
Package http provides HTTP client and server implementations.

Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:
dwedwdcdwcdfd
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
The client must close the response body when finished with it:

resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
// ...
For control over HTTP client headers, redirect policy, and other settings, create a Client:

client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:

tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.

ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
More control over the server's behavior is available by creating a custom Server:

s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
Postfix increment
let x = 3;
const y = x++;

// x = 4
// y = 3
Copy to Clipboard
Prefix increment
let x = 3;
const y = ++x;

// x = 4
// y = 4
log.Fatal(s.ListenAndServe())
Starting with Go 1.6, the http package has transparent support for the HTTP/2 protocol when using HTTPS. Programs that must disable HTTP/2 can do so by setting Transport.TLSNextProto (for clients) or Server.TLSNextProto (for servers) to a non-nil, empty map. Alternatively, the following GODEBUG environment variables are currently supported:

GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
You can use the pkg.go.dev site to find published modules whose packages have functions you can use in your own code. Packages are published in modules -- like rsc.io/quote -- where others can use them. Modules are improved with new versions over time, and you can upgrade your code to use the improved versions.

Add a fuzz test
The unit test has limitations, namely that each input must be added to the test by the developer. One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}


Developing modules
Developing and publishing modules
You can collect related packages into modules, then publish the modules for other developers to use. This topic gives an overview of developing and publishing modules.

Module release and versioning workflow
When you develop modules for use by other developers, you can follow a workflow that helps ensure a reliable, consistent experience for developers using the module. This topic describes the high-level steps in that workflow.

Managing module source
When you're developing modules to publish for others to use, you can help ensure that your modules are easier for other developers to use by following the repository conventions described in this topic.

Developing a major version update
A major version update can be very disruptive to your module's users because it includes breaking changes and represents a new module. Learn more in this topic.

Publishing a module
When you want to make a module available for other developers, you publish it so that it's visible to Go tools. Once you've published the module, developers importing its packages will be able to resolve a dependency on the module by running commands such as go get.
If used postfix, with operator after operand (for example, x++), the increment operator increments and returns the value before incrementing.

If used prefix, with operator before operand (for example, ++x), the increment operator increments and returns the value after incrementing.
Module version numbering
A module's developer uses each part of a module's version number to signal the version’s stability and backward compatibility. For each new release, a module's release version number specifically reflects the nature of the module's changes since the preceding release.

Fuzzing has a few limitations as well. In your unit test, you could predict the expected output of the Reverse function, and verify that the actual output met those expectations.
There are a few different ways you could debug this error. If you are using VS Code as your text editor, you can set up your debugger to investigate.

In this tutorial, we will log useful debugging info to your terminal.

First, consider the docs for utf8.ValidString.

ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
The current Reverse function reverses the string byte-by-byte, and therein lies our problem. In order to preserve the UTF-8-encoded runes of the original string, we must instead reverse the string rune-by-rune.

To examine why the input (in this case, the Chinese character 泃) is causing Reverse to produce an invalid string when reversed, you can inspect the number of runes in the reversed string.

Like before, there are several ways you could debug this failure. In this case, using a debugger would be a great approach.

we will log useful debugging info in the Reverse function.

Following is another example:  


package main
import "fmt"
 
func main() {
   fmt.Println("1 + 1 =", 1 + 1)
}
Output:  

1 + 1 = 2

Look closely at the reversed string to spot the error. In Go, a string is a read only slice of bytes, and can contain bytes that aren’t valid UTF-8. The original string is a byte slice with one byte, '\x91'. When the input string is set to []rune, Go encodes the byte slice to UTF-8, and replaces the byte with the UTF-8 character �. When we compare the replacement UTF-8 character to the input byte slice, they are clearly not equal.
The major Cloud providers (GCP, AWS, Azure) have Go APIs for their services, and popular open source libraries provide support for API tooling (Swagger), transport (protocol buffers, gRPC), monitoring (OpenCensus), Object-Relational Mapping (gORM), and authentication (JWT). The open source community has also provided several service frameworks, including Go Kit, Go Micro, and Gizmo, which can be a great way to get started quickly.
f.Fuzz(func(t *testing.T, orig string) {
    rev := Reverse(orig)
    doubleRev := Reverse(rev)
    t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
    if orig != doubleRev {
        t.Errorf("Before: %q, after: %q", orig, doubleRev)
    }
    if utf8.ValidString(orig) && !utf8.ValidString(rev) {
        t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
    }
})
Notice that your code calls the Go function, printing a clever message about communication.
## further general information
Godoc extracts and generates documentation for Go programs.

It runs as a web server and presents the documentation as a web page.

godoc -http=:6060
Usage:

godoc [flag]
The flags are:

-v
	verbose mode
-timestamps=true
	show timestamps with directory listings
-index
	enable identifier and full text search index
	(no search box is shown if -index is not set)
-index_files=""
	glob pattern specifying index files; if not empty,
	the index is read from these files in sorted order
-index_throttle=0.75
	index throttle value; a value of 0 means no time is allocated
	to the indexer (the indexer will never finish), a value of 1.0
	means that index creation is running at full throttle (other
	goroutines may get no time while the index is built)
-index_interval=0
	interval of indexing; a value of 0 sets it to 5 minutes, a
	negative value indexes only once at startup
-play=false
	enable playground
-links=true
	link identifiers to their declarations
-write_index=false
	write index to a file; the file name must be specified with
	-index_files
-maxresults=10000
	maximum number of full text search results shown
	(no full text index is built if maxresults <= 0)
-notes="BUG"
	regular expression matching note markers to show
	(e.g., "BUG|TODO", ".*")
-goroot=$GOROOT
	Go root directory
-http=addr
	HTTP service address (e.g., '127.0.0.1:6060' or just ':6060')
-templates=""
	directory containing alternate template files; if set,
	the directory may provide alternative template files
	for the files in $GOROOT/lib/godoc
-url=path
	print to standard output the data that would be served by
	an HTTP request for path
-zip=""
	zip file providing the file system to serve; disabled if empty
By default, godoc looks at the packages it finds via $GOROOT and $GOPATH (if set). This behavior can be altered by providing an alternative $GOROOT with the -goroot flag.

When the -index flag is set, a search index is maintained. The index is created at startup.

The index contains both identifier and full text search information (searchable via regular expressions). The maximum number of full text search results shown can be set with the -maxresults flag; if set to 0, no full text results are shown, and only an identifier index but no full text search index is created.


Disadvantages:  

It has no support for generics, even if there are many discussions about it.
The packages distributed with this programming language is quite useful but Go is not so object-oriented in the conventional sense.
There is absence of some libraries especially a UI tool kit.
Some popular Applications developed in Go Language

Docker: a set of tools for deploying linux containers
Openshift: a cloud computing platform as a service by Red Hat.
Kubernetes: The future of seamlessly automated deployment processes
Dropbox: migrated some of their critical components from Python to Go.
Netflix: for two part of their server architecture.
InfluxDB: is an open-source time series database developed by InfluxData.
Golang: The language itself was written in Go.



Using and understanding Go
Effective Go
A document that gives tips for writing clear, idiomatic Go code. A must read for any new Go programmer. It augments the tour and the language specification, both of which should be read first.

Editor plugins and IDEs
A document that summarizes commonly used editor plugins and IDEs with Go support.

Diagnostics
Summarizes tools and methodologies to diagnose problems in Go programs.

A Guide to the Go Garbage Collector
A document that describes how Go manages memory, and how to make the most of it.

Managing dependencies
When your code uses external packages, those packages (distributed as modules) become dependencies.

Fuzzing
Main documentation page for Go fuzzing.

Accessing databases
Tutorial: Accessing a relational database
Introduces the basics of accessing a relational database using Go and the database/sql package in the standard library.

Accessing relational databases
An overview of Go's data access features.

Opening a database handle
You use the Go database handle to execute database operations. Once you open a handle with database connection properties, the handle represents a connection pool it manages on your behalf.

Executing SQL statements that don't return data
For SQL operations that might change the database, including SQL INSERT, UPDATE, and DELETE, you use Exec methods.

Querying for data
For SELECT statements that return data from a query, using the Query or QueryRow method.

Using prepared statements
Defining a prepared statement for repeated use can help your code run a bit faster by avoiding the overhead of re-creating the statement each time your code performs the database operation.

Executing transactions
sql.Tx exports methods representing transaction-specific semantics, including Commit and Rollback, as well as methods you use to perform common database operations.

Canceling in-progress database operations
Using context.Context, you can have your application's function calls and services stop working early and return an error when their processing is no longer needed.

Managing connections
For some advanced programs, you might need to tune connection pool parameters or work with connections explicitly.

Avoiding SQL injection risk
You can avoid an SQL injection risk by providing SQL parameter values as sql package function arguments
By default, godoc uses the system's GOOS/GOARCH. You can provide the URL parameters "GOOS" and "GOARCH" to set the output on the web page for the target system.

The presentation mode of web pages served by godoc can be controlled with the "m" URL parameter; it accepts a comma-separated list of flag names as value:

all	show documentation for all declarations, not just the exported ones
methods	show all embedded methods, not just those of unexported anonymous fields
src	show the original source code rather than the extracted documentation
flat	present flat (not indented) directory listings using full paths
For instance, https://golang.org/pkg/math/big/?m=all shows the documentation for all (not just the exported) declarations of package big.

By default, godoc serves files from the file system of the underlying OS. Instead, a .zip file may be provided via the -zip flag, which contains the file system to serve. The file paths stored in the .zip file must use slash ('/') as path separator; and they must be unrooted. $GOROOT (or -goroot) must be set to the .zip file directory path containing the Go root directory. For instance, for a .zip file created by the command:

zip -r go.zip $HOME/go
one may run godoc as follows:

godoc -http=:6060 -zip=go.zip -goroot=$HOME/go

benefits in cloud and network security

Key Benefits

Address tradeoff between development cycle time and server performance

Go was created to address exactly these concurrency needs for scaled applications, microservices, and cloud development. In fact, over 75 percent of projects in the Cloud Native Computing Foundation are written in Go.

Go helps reduce the need to make this tradeoff, with its fast build times that enable iterative development, lower memory and CPU utilization. Servers built with Go experience instant start up times and are cheaper to run in pay-as-you-go and serverless deployments.

Address challenges with the modern cloud, delivering standard idiomatic APIs

Go addresses many challenges developers face with the modern cloud, delivering standard idiomatic APIs, and built in concurrency to take advantage of multicore processors. Go’s low-latency and “no knob” tuning make Go a great balance between performance and productivity - granting engineering teams the power to choose and the power to move.

Use Case

Use Go for Cloud Computing

Go’s strengths shine when it comes to building services. Its speed and built-in support for concurrency results in fast and efficient services, while static typing, robust tooling, and emphasis on simplicity and readability help build reliable and maintainable code.

Go has a strong ecosystem supporting service development. The standard library includes packages for common needs like HTTP servers and clients, JSON/XML parsing, SQL databases, and a range of security/encryption functionality, while the Go runtime includes tools for race detection, benchmarking/profiling, code generation, and static code analysis.

The major Cloud providers (GCP, AWS, Azure) have Go APIs for their services, and popular open source libraries provide support for API tooling (Swagger), transport (protocol buffers, gRPC), monitoring (OpenCensus), Object-Relational Mapping (gORM), and authentication (JWT). The open source community has also provided several service frameworks, including Go Kit, Go Micro, and Gizmo, which can be a great way to get started quickly.




Benefits of go in web development!

Key Benefits

Deploy across platforms in record speed

For enterprises, Go is preferred for providing rapid cross-platform deployment. With its goroutines, native compilation, and the URI-based package namespacing, Go code compiles to a single, small binary—with zero dependencies—making it very fast.

Leverage Go’s out-of-the-box performance to scale with ease

Tigran Bayburtsyan, Co-Founder and CTO at Hexact Inc., summarizes five key reasons his company switched to Go:

Compiles into a single binary — “Using static linking, Go actually combining all dependency libraries and modules into one single binary file based on OS type and architecture.”

Static type system — “Type system is really important for large scale applications.”

Performance — “Go performed better because of its concurrency model and CPU scalability. Whenever we need to process some internal request, we are doing it with separate Goroutines which are 10x cheaper in resources than Python Threads.”

No need for a web framework — “In most of the cases you really don’t need any third-party library.”

Great IDE support and debugging — “After rewriting all projects to Go, we got 64 percent less code than we had earlier.”
First line of code in go (try it yourselve!)


// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

Go was named Programming Language of the Year by the TIOBE Programming Community Index in its first year, 2009, for having a larger 12-month increase in popularity (in only 2 months after its introduction in November) than any other language that year, and reached 13th place by January 2010, surpassing established languages like Pascal. By June 2015, its ranking had dropped to below 50th in the index, placing it lower than COBOL and Fortran.But as of January 2017, its ranking had surged to 13th, indicating significant growth in popularity and adoption. Go was awarded TIOBE programming language of the year 2016.
Advantages:  

Flexible- It is concise, simple and easy to read.
Concurrency- It allows multiple process running simultaneously and effectively.
Quick Outcome- Its compilation time is very fast.
Library- It provides a rich standard library.
Garbage collection- It is a key feature of go. Go excels in giving a lot of control over memory allocation and has dramatically reduced latency in the most recent versions of the garbage collector.
It validates for the interface and type embedding
this programme will print Hello
congrats! you have written your first Go programme
command line interface
Overview

CLI developers prefer Go for portability, performance, and ease of creation

Command line interfaces (CLIs), unlike graphical user interfaces (GUIs), are text-only. Cloud and infrastructure applications are primarily CLI-based due to their easy automation and remote capabilities.

Key benefits

Leverage fast compile times to build programs that start quickly and run on any system

Developers of CLIs find Go to be ideal for designing their applications. Go compiles very quickly into a single binary, works across platforms with a consistent style, and brings a strong development community. From a single Windows or Mac laptop, developers can build a Go program for every one of the dozens of architectures and operating systems Go supports in a matter of seconds, no complicated build farms are needed. No other compiled language can be built as portably or quickly. Go applications are built into a single self contained binary making installing Go applications trivial.

Specifically, programs written in Go run on any system without requiring any existing libraries, runtimes, or dependencies. And programs written in Go have an immediate startup time—similar to C or C++ but unobtainable with other programming languages.
devops and site reliability engineer
Overview

Go helps enterprises automate and scale

Development Operations (DevOps) teams help engineering organizations automate tasks and improve their continuous integration and continuous delivery and deployment (CI/CD) process. DevOps can topple developmental silos and implement tooling and automation to enhance software development, deployment, and support.

Site Reliability Engineering (SRE) was born at Google to make the company’s “large-scale sites more reliable, efficient, and scalable,” writes Silvia Fressard, an independent DevOps consultant. “And the practices they developed responded so well to Google’s needs that other big tech companies, such as Amazon and Netflix, also adopted them.” SRE requires a mix of development and operations skills, and “empowers software developers to own the ongoing daily operation of their applications in production.”

Go serves both siblings, DevOps and SRE, from its fast build times and lean syntax to its security and reliability support. Go’s concurrency and networking features also make it ideal for tools that manage cloud deployment—readily supporting automation while scaling for speed and code maintainability as development infrastructure grows over time.

DevOps/SRE teams write software ranging from small scripts, to command-line interfaces (CLI), to complex automation and services, and Go’s feature set has benefits for every situation.

Key Benefits

Easily build small scripts with Go’s robust standard library and static typing

Go’s fast build and startup times. Go’s extensive standard library—including packages for common needs like HTTP, file I/O, time, regular expressions, exec, and JSON/CSV formats—lets DevOps/SREs get right into their business logic. Plus, Go’s static type system and explicit error handling make even small scripts more robust.

Quickly deploy CLIs with Go’s fast build times

Every site reliability engineer has written “one-time use” scripts that turned into CLIs used by dozens of other engineers every day. And small deployment automation scripts turn into rollout management services. With Go, DevOps/SREs are in a great position to be successful when software scope inevitably creeps. Starting with Go puts you in a great position to be successful when that happens.

Scale and maintain larger applications with Go’s low memory footprint and doc generator

Go’s garbage collector means DevOps/SRE teams don’t have to worry about memory management. And Go’s automatic documentation generator (godoc) makes code self-documenting–lowering maintenance overhead and establishing best practices from the get-go.

Getting started in generics!

This tutorial introduces the basics of generics in Go. With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

In this tutorial, you’ll declare two simple non-generic functions, then capture the same logic in a single generic function.

You’ll progress through the following sections:

Create a folder for your code.
Add non-generic functions.
Add a generic function to handle multiple types.
Remove type arguments when calling the generic function.
Declare a type constraint.
Note: For other tutorials, see Tutorials.

Note: If you prefer, you can use the Go playground in “Go dev branch” mode to edit and run your program instead.

Prerequisites
An installation of Go 1.18 or later. For installation instructions, see Installing Go.
A tool to edit your code. Any text editor you have will work fine.
A command terminal. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.
=======
how to do this!!


Developing and publishing modules
You can collect related packages into modules, then publish the modules for other developers to use. This topic gives an overview of developing and publishing modules.

To support developing, publishing, and using modules, you use:

A workflow through which you develop and publish modules, revising them with new versions over time. See Workflow for developing and publishing modules.
Design practices that help a module’s users understand it and upgrade to new versions in a stable way. See Design and development.
A decentralized system for publishing modules and retrieving their code. You make your module available for other developers to use from your own repository and publish with a version number. See Decentralized publishing.
A package search engine and documentation browser (pkg.go.dev) at which developers can find your module. See Package discovery.
A module version numbering convention to communicate expectations of stability and backward compatibility to developers using your module. See Versioning.
Go tools that make it easier for other developers to manage dependencies, including getting your module’s source, upgrading, and so on. See Managing dependencies.

Create a folder for your code
To begin, create a project for the code you’ll write.
This tutorial introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework (Gin).

You’ll get the most out of this tutorial if you have a basic familiarity with Go and its tooling. If this is your first exposure to Go, please see Tutorial: Get started with Go for a quick introduction.

Gin simplifies many coding tasks associated with building web applications, including web services. In this tutorial, you’ll use Gin to route requests, retrieve request details, and marshal JSON for responses.

In this tutorial, you will build a RESTful API server with two endpoints. Your example project will be a repository of data about vintage jazz records.

The tutorial includes the following sections:

Design API endpoints.
Create a folder for your code.
Create the data.
Write a handler to return all items.
Write a handler to add a new item.
Write a handler to return a specific item.
Note: For other tutorials, see Tutorials.

To try this as an interactive tutorial you complete in Google Cloud Shell, click the button below.

Open in Cloud Shell

Prerequisites
An installation of Go 1.16 or later. For installation instructions, see Installing Go.
A tool to edit your code. Any text editor you have will work fine.
A command terminal. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.
The curl tool. On Linux and Mac, this should already be installed. On Windows, it’s included on Windows 10 Insider build 17063 and later. For earlier Windows versions, you might need to install it. For more, see Tar and Curl Come to Windows.
Design API endpoints
You’ll build an API that provides access to a store selling vintage recordings on vinyl. So you’ll need to provide endpoints through which a client can get and add albums for users.

When developing an API, you typically begin by designing the endpoints. Your API’s users will have more success if the endpoints are easy to understand.

Here are the endpoints you’ll create in this tutorial.

/albums

GET – Get a list of all albums, returned as JSON.
POST – Add a new album from request data sent as JSON.
/albums/:id

GET – Get an album by its ID, returning the album data as JSON.

Creating a go module

Tutorial: Create a Go module
Table of Contents
Prerequisites
Start a module that others can use
This is the first part of a tutorial that introduces a few fundamental features of the Go language. If you're just getting started with Go, be sure to take a look at Tutorial: Get started with Go, which introduces the go command, Go modules, and very simple Go code.

In this tutorial you'll create two modules. The first is a library which is intended to be imported by other libraries or applications. The second is a caller application which will use the first.

This tutorial's sequence includes seven brief topics that each illustrate a different part of the language.

Create a module -- Write a small module with functions you can call from another module.
Call your code from another module -- Import and use your new module.
Return and handle an error -- Add simple error handling.
Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).
Return greetings for multiple people -- Store key/value pairs in a map.
Add a test -- Use Go's built-in unit testing features to test your code.
Compile and install the application -- Compile and install your code locally.
Create a folder for your code
To begin, create a folder for the code you’ll write.


On Linux or Mac:

$ cd
On Windows:

C:\> cd %HOMEPATH%
Using the command prompt, create a directory for your code called web-service-gin.

$ mkdir web-service-gin
$ cd web-service-gin
Create a module in which you can manage dependencies.

Run the go mod init command, giving it the path of the module your code will be in.

$ go mod init example/web-service-gin
go: creating new go.mod: module example/web-service-gin
This command creates a go.mod file in which dependencies you add will be listed for tracking. For more about naming a module with a module path, see Managing dependencies.

Next, you’ll design data structures for handling data.

Create the data
To keep things simple for the tutorial, you’ll store data in memory. A more typical API would interact with a database.

Note that storing data in memory means that the set of albums will be lost each time you stop the server, then recreated when you start it.

Using your text editor, create a file called main.go in the web-service directory. You’ll write your Go code in this file.

The rest of the tutorial will show a $ as the prompt. The commands you use will work on Windows too.

Run your code to see the message generated by the function you're calling.
$ go run .
Don't communicate by sharing memory, share memory by communicating.

From the command prompt, create a directory for your code called generics.

$ mkdir generics
$ cd generics
Create a module to hold your code.

Run the go mod init command, giving it your new code’s module path.
Go language is just like Java language as it support platform independency. Due to its modular design and modularity i.e., the code is compiled and is converted into binary form which is as small as possible and hence, it requires no dependency. Its code can be compiled in any platform or any server and application you work on.

Go merges modern day developer workflow of working with Open Source projects and includes that in the way it manages external packages. Support is provided directly in the tooling to get external packages and publish your own packages in a set of easy commands.