
## How to start with GoLang
- Introduction
C syntax
First, let’s talk about C syntax. C took an unusual and clever approach to declaration syntax. Instead of describing the types with special syntax, one writes an expression involving the item being declared, and states what type that expression will have. Thus

int x;
declares x to be an int: the expression ‘x’ will have type int. In general, to figure out how to write the type of a new variable, write an expression involving that variable that evaluates to a basic type, then put the basic type on the left and the expression on the right.

Thus, the declarations
The pointer in C language is a variable which stores the address of another variable. This variable can be of type int, char, array, function, or any other pointer. The size of the pointer depends on the architecture. However, in 32-bit architecture the size of a pointer is 2 byte.

int *p;
int a[3];
In C, we can divide a large program into the basic building blocks known as function. The function contains the set of programming statements enclosed by {}. A function can be called multiple times to provide reusability and modularity to the C program. In other words, we can say that the collection of functions creates a program. The function is also known as procedureor subroutinein other programming languages.

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

f func(func(int,int) int, int) int
Or if f returns a function:

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
Package http provides HTTP client and server implementations.

Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:

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
log.Fatal(s.ListenAndServe())
Starting with Go 1.6, the http package has transparent support for the HTTP/2 protocol when using HTTPS. Programs that must disable HTTP/2 can do so by setting Transport.TLSNextProto (for clients) or Server.TLSNextProto (for servers) to a non-nil, empty map. Alternatively, the following GODEBUG environment variables are currently supported:

GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps


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
Fuzzing has a few limitations as well. In your unit test, you could predict the expected output of the Reverse function, and verify that the actual output met those expectations.
There are a few different ways you could debug this error. If you are using VS Code as your text editor, you can set up your debugger to investigate.

In this tutorial, we will log useful debugging info to your terminal.

First, consider the docs for utf8.ValidString.

ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
The current Reverse function reverses the string byte-by-byte, and therein lies our problem. In order to preserve the UTF-8-encoded runes of the original string, we must instead reverse the string rune-by-rune.

To examine why the input (in this case, the Chinese character 泃) is causing Reverse to produce an invalid string when reversed, you can inspect the number of runes in the reversed string.

Like before, there are several ways you could debug this failure. In this case, using a debugger would be a great approach.

we will log useful debugging info in the Reverse function.

Look closely at the reversed string to spot the error. In Go, a string is a read only slice of bytes, and can contain bytes that aren’t valid UTF-8. The original string is a byte slice with one byte, '\x91'. When the input string is set to []rune, Go encodes the byte slice to UTF-8, and replaces the byte with the UTF-8 character �. When we compare the replacement UTF-8 character to the input byte slice, they are clearly not equal.

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

Open a command prompt and change to your home directory.

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

From the command prompt, create a directory for your code called generics.

$ mkdir generics
$ cd generics
Create a module to hold your code.

Run the go mod init command, giving it your new code’s module path.

$ go mod init example/generics
go: creating new go.mod: module example/generics
Note: For production code, you’d specify a module path that’s more specific to your own needs. For more, be sure to see Managing dependencies.

Next, you’ll add some simple code to work with maps.

Add non-generic functions
In this step, you’ll add two functions that each add together the values of a map and return the total.

You’re declaring two functions instead of one because you’re working with two different types of maps: one that stores int64 values, and one that stores float64 values.

Write the code
Using your text editor, create a file called main.go in the generics directory. You’ll write your Go code in this file.

Into main.go, at the top of the file, paste the following package declaration.

package main
A standalone program (as opposed to a library) is always in package main.

Beneath the package declaration, paste the following declaration of an album struct. You’ll use this to store album data in memory.
