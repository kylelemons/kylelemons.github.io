---
layout: post
title: Bite-sized Guides to C: Life, the Universe, and Everything
categories: []
tags: []
---
Bite-sized Guides to C: Life, the Universe, and Everything
06:05 21 Jul 2010
Tags:  bite-sized, c, h2g2, programming, favorite

Kyle Lemons


* Blog Entry


This bite-sized guide to C is not precisely bite-sized, but hopefully the individual chunks are.  I think this is a good introduction to quite a few of the most basic techniques that you will require when you are writing C code.  First, have a look at the code below, and see if you can follow along.  The gritty details of what everything does follow.



  #include &#34;stdio.h&#34;
  #include &#34;math.h&#34;
  
  /**
   * \brief This function will print a number to the console in the given base.
   * \param num The number to convert
   * \param base The base in which to print (2..16)
   */
  void printNumInBase(int num, int base) {
    // Loop variable(s)
    int i;
  
    // Create a static array of the digits for use up to base 16
    char digits[] = {&#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;,
                     &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;,
                     &#39;8&#39;, &#39;9&#39;, &#39;A&#39;, &#39;B&#39;,
                     &#39;C&#39;, &#39;D&#39;, &#39;E&#39;, &#39;F&#39;};
  
    // Determine the number of digits in the final string
    int length = (int)(log(num)/log(base) &#43; 0.5);
  
    // Print out each digit
    for (i = 0; i &lt; length; &#43;&#43;i) {
      // Use the base digit calculation
      int digit = (int)(num / pow(base, length-i-1)) % base;
  
      // Print out the determined character
      printf(&#34;%c&#34;, digits[digit]);
    }
  
    // Print out a newline at the end for good measure
    printf(&#34;\n&#34;);
  }
  
  /**
   * \brief Entry point for the program
   */
  void main() {
    // Create a integral value
    int two = 2;
  
    // Create variables for later
    int six, nine;
  
    // Increment the value of two
    two&#43;&#43;;
  
    // Assign to new integer values
    six = 2 * two;
    nine = two * two;
  
    // Print values
    printf(&#34;What do you get when you multiply %d by %d?\n&#34;, six, nine);
  
    // Print the answer, decrementing two inline
    printNumInBase(six * nine, six &#43; nine - --two);
  }
  


** Data types



Four major data types are used in the program given above.  Can you find them?  I&#39;ll give you a hint: not all of them are explicitly named.  The answer, of course is this: `int`, `char`, `double`, and `char*`.  Now look back at the code and see if you can find where each one is used.



*** Chars, Integers, Floats and Doubles



C has two ways to represent numbers.  You can represent a number as either an integer, that is a number without a fractional part, or as a floating-point number that does have a fractional part.  How both are stored is a discussion for a future guide, but for now just realize that `char` and `int` are integer values (the first intended to represent ASCII characters) and that `float` and `double` are floating-point numbers (the second able to hold both bigger and smaller numbers than the first).



*** Precision and Casting


In the code above, `double`s are used implicitly.  If you have a look at [[http://www.cppreference.com/wiki/c/math/start][math.h]], you will notice that `log` and `pow` both take doubles and return doubles.  You might say to yourself, &#34;Self, how does it take a double if we are passing it an int?&#34;  Well, the answer is, of course, that a double can represent any integer value, and thus it is legal to use an integer in the place of a double without any compiler warnings or errors.  However, when you want to take the double return value and assign it to an integer variable (as we do on lines 20 and 25), you must tell the compiler explicitly that you do not care that an integer variable cannot hold all possible double values and that you wish to discard anything that won&#39;t fit.  This is called a _cast_, and has the form: `(type)expression`.  In C, chars can _promote_ implicitly to ints, ints to floats, and floats to doubles.  A promotion can skip up that chain, but a demotion (going backward) requires a cast.



*** Arrays



So, I actually lied.  I would promise that it won&#39;t happen again, but it probably will.  There aren&#39;t really any `char*`s used in the code above.  If you read my [[][pointers PDF]], this will all be explained in terribly gory detail.  You may safely ignore this for now, and pretend that the array type used on line 14 is actually a pointer to an array and ignore the compiler optimizations.




Dynamic memory allocation will come later, so just remember for now that the array on line 14 is actually a _pointer_ to a group of `char`s in memory somewhere.  Each member of the group is in a specific order.  If it helps, think of it as a row of apartment mailboxes... they are all shoved together, and each one has its own number, and they can be either all in a row, or in two dimensions.  Arrays in C can have more dimensions, but this is rare and not covered here.  The array on line 14 is one-dimensional, and (like all arrays in C) has a first element whose number is zero.  You can see how to get at an element in the array (like opening the mailbox with a given number) on line 28.



** The Code


Enough of the formalities. Let&#39;s get to the code.



`#include`&#34;stdio.h&#34;`
`#include`&#34;math.h&#34;`

First, we include the standard I/O library and the math library.  Unlike in our first bite-sized guide, we are using quotes here.  This is to demonstrate the difference between angle quotes and double quotes in an include, or rather the lack of a difference in this case.  (Wait, I lied again.  It&#39;s because syntax highlighting is broken for angle quotes.  However, double quotes are still legal for the reason I am about to share with you.)  When you use angle quotes for an include, C will search the system directories for includes and then stop.  When you use double quotes, C will search the local directory and then try the system directories before stopping.  Typically, you will see angle quotes used for system includes and double quotes used for local includes, but I may have to break from that convention here because of the syntax highligher.  Sorry, folks!


It is also important to note that there is a good chance that you will have to link with the math library on your system for this binary to work.  An example of a command (to be run in the directory in which the .c file exists, in my case called &#34;test.c&#34;):



`gcc`-o`test`test.c`-lm`

The `-lm` indicates that the C compiler (in my case gcc) should also link with libm, which happens to be the math library. This library contains optimized mathematical functions that are listed in math.h.



`void`printNumInBase(int`num,`int`base)`{...}`

This creates a function called `printNumInBase` which takes in two integers and returns nothing (`void` is a pseudo-type used in place of a data type to indicate the lack of a type or value).  The comment before the function attempts to explain what it does and what the parameters are.  The curly braces block off a section of code and associate it with the function.  In general, a function is created to perform a specific task, often requiring inputs to determine their result or to perform their task.  In this case, the task being performed is the conversion of a number into a printed string using a non-decimal base, which requires both the number to display and the base in which to display it.



`int`i;`

  char digits[] = {&#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;,
  Â Â Â Â Â Â Â Â Â Â Â Â Â Â Â Â Â &#39;8&#39;, &#39;9&#39;, &#39;A&#39;, &#39;B&#39;, &#39;C&#39;, &#39;D&#39;, &#39;E&#39;, &#39;F&#39;};

`int`length`=`(int)(log(num)/log(base)`&#43;`0.5);`

At the beginning of every function, all local variables must be declared and may be initialized.  The variable used in the loop (`i`) is declared to be an integer, but not initialized.  Be careful about not initializing variables, as this can lead to insidious bugs.  Always set a variable to a known value before you start using it.  The variable `digits` is declared to be an array (denoted by the `[]`; there could be a number between the brackets to specify the length of the array, but without it the compiler will count the initializer elements and use that) and initialized with the characters usable in base16 (hexadecimal).  I should also point out here that strings use double quotes and characters (really a number in ASCII) use single quotes.  String processing will be covered later.  I won&#39;t go over the math, but the length variable is initialized to the log of the given number in the given base, rounded up, as an integer.



`for`(i`=`0;`i`&lt;`length;`&#43;&#43;i)`{...}`

This is a &#34;for&#34; loop.  It is called such because it is typically used to run a piece of code _for_ a range of values.  The for statement always have three parts, which may have a single expression or may be blank.  The first is run before the loop, and I will call it the &#34;initializer.&#34;  The second, which I will call the &#34;conditional&#34; is run before each iteration of the loop, and if the result of the expression is true or nonzero will cause the iteration to execute.  If it is not, the loop will terminate and continue with the first statement after the loop.  The final part is the &#34;increment,&#34; and is run after each iteration of the loop before checking the conditional.



`i`=`0`

This initializes the counter to zero.  Arrays in C start at zero, so counters and such will typically follow this convention.



`i`&lt;`length`

The result of this expression is true (nonzero) if the counter is less than the value in the length variable.  When this is no longer the case, the loop terminates.



`&#43;&#43;i`

The value of the counter is incremented by one.  The value is the expression is the incremented value, though it is discarded in this case.



When the above are combined, it creates a for loop that runs once for each value from zero up to (but not including) the length calculated at the beginning of the function (e.g. the number of characters in the string-based representation of the number in the new base).



`int`digit`=`(int)(num`/`pow(base,`length-i-1))`%`base;`
`printf(&#34;%c&#34;,`digits[digit]);`

Within the loop, the digit is calculated (the first line) and printed (the second).  Notice the `&#34;%c&#34;`; the percent sign followed by a &#39;c&#39; indicates a format specifier to take the argument to `printf` and print it out as a character.  The argument is `digits[digit]`, which takes the `digit`&#39;th element (starting from the zeroth) element of the `digits` array.  The array holds characters, so the result of this expression is a character, which is what will be printed on the screen.  More information on printf can be found [[http://www.cppreference.com/wiki/c/io/printf][online]] or in the manual page on linux.



`int`two`=`2;`
`int`six,`nine;`
`two&#43;&#43;;`
`six`=`2`*`two;`
`nine`=`two`*`two;`

In the main function (explained previously), we set up some variables, increment one (if this were in an expression, the value of the variable _before_ the increment would be used; this is a postfix increment), and do some math.



`printf(&#34;What`do`you`get`when`you`multiply`%d`by`%d?\n&#34;,`six,`nine);`

We then use printf again (with decimal format specifiers) to print out the ultimate question.



`printNumInBase(six`*`nine,`six`&#43;`nine`-`--two);`

And finally, we call the function that we defined earlier with two arguments.  Notice the inline prefix decrement `--two` which will decrement the variable and use the decremented value in the calculation.  Try to see if you can figure out what the arguments to this and the printf are before you run it, and what the output will be.  If you don&#39;t want to do the base computation, there are plenty of online tools to convert between bases for you.



** Happy coding!

