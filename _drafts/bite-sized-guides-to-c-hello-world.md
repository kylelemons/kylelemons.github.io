---
layout: post
title: Bite-sized Guides to C: Hello World!
created: 2015-03-28 02:14:15.3 -0700 PDT
categories: []
tags: []
---
Bite-sized Guides to C: Hello World!
15:00 18 Aug 2009
Tags:  bite-sized, programming, favorite

Kyle Lemons


* Blog Entry

In this, the first, bite-sized guide to C we will be going over all of the gory details of your first C program, Hello, World!.  This code (without line numbers) goes in a file by itself, preferably with a trailing newline (to avoid compiler warnings), with a &#34;.c&#34; extension (e.g. &#34;hello.c&#34;).
Unknown atom &#34;blockquote&#34; []

01 #include &lt;stdio.h&gt;
02 int main(int argc, char **argv)
03 {
04 Â  printf(&#34;Hello, World!&#34;);
05 Â  return 0;
06 }
That&#39;s about as short of a C program as you can get while still printing something out to the screen and not cutting any boilerplate corners.  Let&#39;s go through it line by line (and, sometimes, character by character.

01 #include &lt;stdio.h&gt;
Just to be clear, the numbers at the beginning of the line are line numbers.  Everything after that is actually code that needs to be in your C file to work.  The first lines of any C file are usually one of two things: comments detailing what is in the file and #include statements.  We&#39;ll cover comments later, the first line in our example is an include statement.

When the compiler scans through a C file, it treats lines that start with a pound sign (#) specially.  These are called &#34;preprocessor definitions&#34; and are processed before the compiler actually looks at the code.  In the case of #include, the compiler rips out the line and replaces it with whatever is in the specified file.  Now, if you&#39;ll notice, the compiler will be looking for a file called &#34;stdio.h&#34; which we did not provide for you.  This is called a standard library and is stored in a special set of directories in your system, and which provides the same functionality on all standard C installations on any platform.  In our case, this file (called a header) contains standard input and output function definitions, which we need in order to be able to print to the screen.  As you learn more and more functions, you will need more header files in order to be able to use them, and you will learn which function calls require which headers.

02 int main(int argc, char **argv)
This line requires a fairly significant amount of explanation if you are new to C.  For now, we will suffice it to say that this is the header for a function.  You will learn about function calls, return values, function parameters, and arguments later.  For now, understand that this line is a necessary invocation that interfaces your program with the operating system.  When you double click on your application or run it from the command line, the operating system does two things that are important to us: It tries to give information to the program as a list, and it expects the program to tell it whether it ran successfully or failed.  The list of information is encapsulated in the argument count (argc) and the argument values (argv) input arguments, and the success or failure is sent back to the operating system as an integral number (int).  There are a few other ways in which this line can be written, but this is the most common and the one which will be the most flexible in terms of making use of this operating system integration when you need to.  For now, just understand that this line is required so that the OS can run your program.

03 {
06 }
The basic code structuring of C programs is done in blocks.  A block of code is surrounded by curly braces.  Blocks of code contain lines of code, which can have their own blocks of code.  Blocks can also be nested for scoping purposes.  In this example, the curly braces on lines three and six surround the code that is within the main function.  This is the only block of code and contains only two lines of code.  When you write code yourself, you probably want to devise a coding style that allows you to readily see which curly brace matches up with what, as this will be very beneficial when you need to debug your code.

04 printf(&#34;Hello, World!&#34;);
As mentioned above, we need a function to print something out to the screen.  The &#34;printf&#34; function does exactly that.  It is a much more flexible function than we are making use of here, but this is the simplest way to take a string of characters and print them out to the screen.  If you have used any other programming languages, you need to note that the double quoted strings in C are a string and you should remember that single quotes surround a character literal.  These are treated differently by the compiler, and are not interchangeable.

05 return 0;
This exits the main function of your program with an exit status of zero (which indicates success; nonzero would indicate failure).  As mentioned above, this status code is returned to the operating system, and could be used to notify the user or some script that your program completed successfully.

Compiling
The compiler is the program which translates the above code into an executable.  If you are using the GNU Compiler Collection&#39;s C Compiler (gcc), you might run a command like this to compile it:
`gcc`-Wall`-Werror`-ansi`-pedantic`-o`hello`hello.c`
The arguments here are simply to turn on all warnings and enforce the strictest ANSI C.  I recommend that you do this with whatever compiler you use, as it will make your code as cross-platform as possible, and it will give you warnings when you do something which might be fishy.

There are three stages to compiling.  The above command does all three in one.  First, the compiler preprocesses the input files, processing #includes, #defines, etc.  Then, the compiler compiles the source into an object file.  The object file contains pieces of machine code which correspond to all of the functions and variables in the program with a large table of where it put them.  Then, the linker takes all of the objects and connects all references between functions and will let the programmer know if there are undefined references (e.g. your code calls a function that the linker can&#39;t find).  Once this is done, the linker produces a binary executable and the compiling process concludes.

If you have questions about what is contained in here, email me or leave a comment.  If you find a mistake, error, omission, whatever, please let me know as soon as possible so that I can get it corrected.