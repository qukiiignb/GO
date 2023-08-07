# go (goland) 
`last update jul 26 2023`
## welcome
you are still here? great so les't go!

## documentation
wiki,oficial web, <a href="https://cs.opensource.google/go/go">code open source</a>.
## Instalation 
<a href="https://go.dev/dl/">go.dev</a>
## challenges
###	1 tutorial (theGoProgramingLenguage)
   - /a  Get started with Hello, World. With the follow <a href="https://go.dev/doc/tutorial/getting-started#code">doc</a> create a programm that print "hello world" in the terminal then make the testing for this.
   - /b  Call code in an external package. With the follow <a href="https://go.dev/doc/tutorial/getting-started#code">doc</a>
   - /c Create a Go module with the follow <a href="https://go.dev/doc/tutorial/create-module">documentaion</a> create two modules.
   - /d Test, with this documentation <a href="https://pkg.go.dev/testing#hdr-Examples">testing</a> make a funcion that must retur a positive num use Abs function provided by the math package. and test the number that this func return you must make two file num.go and num_test.go
   - e/ Start a module that other can use.
   - f/ Make collection of distinguishable entities such that for any entity it is determined without ambiguity whether it belongs to the collection or not ( set, objects).
   - g/ How you can determine if x belong to X or does not belong to x write a programme that do this for you.
   - h/ Prints its command-line arguments on a single line use os.Args[]
   - i/ Modify the os.Arg[] program to print the index and value of each of its arguments,one per line.
   - j/ Reads input lines, counts their occurrences, and prints lines that appear more than once.
   - k/ Read from the standard input or handle a list of file names,using os.Open to open each one
   - l/ Introduce the function ReadFile in the program k (from the io/ioutil package), which reads the entire contents of a named file, and strings.Split, which splits a string into a slice of substrings.
   - m/ Demonstrates basic usage of Go’s standard image packages, which we’ll use to create a sequence of bit-mapped images and then encode the sequence as a GIF animation.
   - n/ Called fetch that fetches the content of each specified URL and prints it as uninter preted text.
   - o/ Use io.Copy(dst, src) reads from src and writes to dst. Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
   - p/ Fetching url concurrently in parallel and reports their times and sizes.
   - q/ Find a web site that produces a large amount of data. Investigate caching byrunning fetchall twice in succession to see whether the reported time changes much. do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.
   - r/ Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond?
   - s/ Write a web server that responds to client requests, if the request is for http://localhost:8000/hello, the response will be URL.Path = "/hello".
   - t/ Counts the number of requests; a request to the URL /count returns the count so far, excluding /count requests themselves:
   - u/ Can report on the headers and form data that it receives, making the server useful for inspecting and debugging requests
   - v/ Each time you load the page, you’ll see a new animation (lissajous)
   - w/ Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of the default 5
   - x/ 
   - y
   - z
### 2 (theGoProgramingLenguage)
   - a/	Prints the boiling point of water.
   - b/	Print two Fahrenheit to Celsius conversions
   - c/	A set of variables can also be initialized by calling a function that returns multiple values.
   - d/ Pointers are key to the flag package, which uses a program’s command-line arguments to set the values of certain variables distributed throughout the program. To illustrate, this variation on the earlier echo command takes two optional flags: -n causes echo to omit the trailing newline that would normally be printed, and -s sep causes it to separate the output arguments by the contents of the string sep instead of the default single space.
   - e/ Each variable that escapes requires an extra memory allocation
   - f/ Tuple : the n-th Fibonacci number iteratively
   - g/ Computing the greatest common divisor of two intergers
   - h/ Illustrate type declarations performs Celsius and Fahrenheit temperature computations. package tempconv.
   - i/ Stored in two files to show how declarations in separate files of a package are accessed
   - j/ how to precompute a table of values, which is often a useful programming technique.
   - k/ what are the lexical block, scope and rules? 
   - l/ how bitwise operations can be used?
   - m/ Use the built-in len function returns a signed int, as in this loop which announces prize medals in reverse order
   - o/ How you can use the lib math for prints the powers of e with three decimal digits of precision?
   - p/ Surface computes an SVG rendering of a 3-D surface function
   - q/
   - r/
   - s/
   - t/
   - u/
   - v/
   - w/
   - x/
   - y/
   - z/
### 3.3/ Basic Data Types
3. -./3.3 Complex Numbers
   - /a how can you use build-in function complex adn build-in function real, imag extract components?
   - /b how you can create a program that uses complex128 arithmetic to generate a Mandelbrot set?
   - /c : Implementafull-color Mandelbro t set using the function image.NewRGBA and the typ e color.RGBA or color.YCbCr.
   - /d Supersamp ling is a technique to reduce the effec t of pixe lat ion by computing the color value at several points wit hin each pixel and tak ing the average.The simplest met hod is to div ide each pixel int o four ‘‘su bpixe ls.’’ Implement it.
   - /e  Anot her simple frac tal uses Newton’s met hod to find complex solut ion s to a func tion such as z4−1=0. Shade each starting point by the number of iterat ions required to get clos e to one of the four roots. Color each point by the root it appro ach es
   - /f  Render ing frac tals at hig h zo om level s demands gre at arithmetic pre cision. Implement the same frac tal using four dif ferent represent ation s of numbers: complex64, complex128, big.Float, and big.Rat. (The latter two typ es are found in the math/big package.Float uses arbitrar y but bounded-precision floating-point; Rat uses unbounded-precision rational numbers.) How do the y compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?
   - /g  Wr ite a web ser ver that renders frac tals and writes the image dat a to the client. Al low the client to specif y the x, y, and zoom values as parameters to the HTTP request.
### 3.4/ Booleans
   - /a how you whiting a operation of conversion function that return 0 if is true and 1 if false
   - /b how you will be a reports wether i is non-zero?
### 3.5/ String
   - /a how can you use build-in len function return the number  of bytes (not runes) in a string and the index operation s[i] retrievas the i-th byte  of string s.
   - /b can you do A range lo op decodes a UTF-8-enco ded str ing?
   - /c The bytes package provides the Buffer type for efficient manipu lat ion of byte slices. A
Buffer st arts out emp ty but grows as dat a of typ es li ke string, byte, and []byte are writt en to it
   - /d
### 3.6/ iota
   - /a use iota for create a secuense of related values without spelling out each one explicitly
   - /b
### 4.1 array
   - /a The bui lt-in function len.
### 4.2 append
   - /a app end s it ems to slices.
   - /b Each call to appendYouFun must check whether the slice has suficient capacity to hold the new elements in the existing array
### 4.4 Struct enbedding and anonymous field
   - /a GO's enbedding struct
### 4.6 text and HTML templates
   - /a
### 5.1 function
