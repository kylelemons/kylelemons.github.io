---
layout: post
title: Bite-sized Guides to C: Pointers - Not Just Your Neighbor&#39;s Dog
categories: []
tags: []
---
Bite-sized Guides to C: Pointers - Not Just Your Neighbor&#39;s Dog
17:43 14 Aug 2010
Tags:  c, dogs, pointers

Kyle Lemons


* Blog Entry

This guide is going to be slightly different than the previous ones.  It is significantly more self-driven, and you will be asked to do some experimentation yourself.  Pointers are not a subject that can be spoon-fed; their use-cases cannot be itemized nicely and memorized.  In order to understand pointers, you must understand the fundamentals of what they are.  You learn new applications for them over time; this can be through observing other people&#39;s code or simply through experience, having internalized the concepts enough to recognize where they are applicable.

In the code below, many things are done two different ways: one way using array notation, which you should be familiar with from any other language, and the other using pointer notation.  Read over the code and make sure you understand what the array notation is doing, then look at the corresponding piece that is using pointers.  See if you can spot the similarities.  Once you have done that, read the explanation of pointers below the code, and then look over the code once more.  When you&#39;re satisfied you understand what is going on, move on to the exercises at the bottom of this guide.

  #include &#34;stdio.h&#34;
  #include &#34;string.h&#34;
  
  int main() {
    // Local variables
    char *p, **w; int *o; int i, j, k;   
    char str1[] = &#34;My neighbor just got a new dog!&#34;;
    char *str2 = &#34;That dog is a pointer.&#34;;
    char buf[256] = &#34;&#34;; int index[32]; char* word[32];
    int pick1[] = {0, 1, 2, 3, 4, 6}; int pick2[] = {0, 2, 3, 4};
  
    // Get the lengths and sizes of our strings and arrays
    int str1size = sizeof(str1);   int str1len = strlen(str1);
    int str2size = sizeof(str2);   int str2len = strlen(str2);
    int bufsize = sizeof(buf);     int buflen = strlen(buf);
    int indexsize = sizeof(index); int indexlen = sizeof(index)/sizeof(int);
    int wordsize = sizeof(word);   int wordlen = sizeof(word)/sizeof(char*);
    int pick1size = sizeof(pick1); int pick1len = sizeof(pick1)/sizeof(int);
    int pick2size = sizeof(pick2); int pick2len = sizeof(pick2)/sizeof(int);
  
    // Print out the sizes
    printf(&#34;------ Sizes ------\n&#34;);
    printf(&#34;String 1: size=%-3d length=%-3d\n&#34;, str1size,  str1len);
    printf(&#34;String 2: size=%-3d length=%-3d\n&#34;, str2size,  str2len);
    printf(&#34;Buffer:   size=%-3d length=%-3d\n&#34;, bufsize,   buflen);
    printf(&#34;Indices:  size=%-3d length=%-3d\n&#34;, indexsize, indexlen);
    printf(&#34;Words:    size=%-3d length=%-3d\n&#34;, wordsize,  wordlen);
    printf(&#34;Pick 1:   size=%-3d length=%-3d\n&#34;, pick1size, pick1len);
    printf(&#34;Pick 2:   size=%-3d length=%-3d\n&#34;, pick2size, pick2len);
  
    // Modify string 1
    str1[str1len-1] = &#39; &#39;;
  
    // Find the words from str1
    j = 0;
    for (i = 0; str1len &gt; i; &#43;&#43;i) // some browsers don&#39;t like lessthan
      if (i == 0 || str1[i-1] == &#39; &#39;) index[j&#43;&#43;] = i;
    index[j] = -1;
  
    // Find the words from str2
    for (w = word, p = str2; *p; &#43;&#43;p)
      if (p == str2 || p[-1] == &#39; &#39;) {
        *(w&#43;&#43;) = p;
        *w = NULL;
      }
  
    // Build the final string
    k = 0;
    for (i = 0; pick1len &gt; i; &#43;&#43;i) {
      int wordnum = pick1[i];
      int offset = index[wordnum];
      for (j = offset; str1[j] &amp;&amp; str1[j] != &#39; &#39;; &#43;&#43;j, &#43;&#43;k)
        buf[k] = str1[j];
      buf[k&#43;&#43;] = &#39; &#39;;
    }
    if (k) --k;
    for (p = buf&#43;k, o = pick2; j = *o, pick2len &gt; o-pick2; &#43;&#43;o) {
      int len = strlen(word[j]);
      if (word[j&#43;1]) len = word[j&#43;1] - word[j] - 1;
      *(p&#43;&#43;) = &#39; &#39;;
      strncpy(p, word[j], len);
      if (o == pick2) *p &#43;= &#39;a&#39; - &#39;A&#39;;
      p &#43;= len;
    }
  
    printf(&#34;------ Final ------\n%s\n&#34;, buf);
    return 0;
  }
  


** Pointer variables


The first thing that is unclear to most people about pointers are variable declarations and pointer types.  This is actually pretty straight-forward, as long as you remember a few really simple rules:



- *A*pointer*is*a*number*.  If you are in my recitation, you will learn this well.  A pointer is simply a number, and as such is not magical and cannot follow values around in memory.  This also implies that there is a certain amount of math that is possible with them.
- *All*pointers*have*an*associated*type*. In addition to being a pointer (which is its own data type), all pointers have an associated type.  With the exception of `void`, this type is a standard data type that could be assigned to/from another variable, including another pointer type.
- *Never*dereference*NULL*.  I will explain NULL later, but suffice it to say here that you need to always protect yourself from the possibility of dereferencing NULL, as this will cause your program to crash, and it will be very difficult to debug.


The general format of a pointer declaration is `[type]*[variable]`.  The C compiler doesn&#39;t care about white space as long as it can tell what you mean, so often you will see it written with a space either before or after the asterisk.  Just remember that in a variable declaration, the asterisk is always between the type and the (possibly empty) variable name.  Look at the following examples and identify the `[type]` and `[variable]`: (highlight to see answers)



- `int`*a` The type is &#34;int&#34; and the variable is &#34;a&#34;.
- `char*`str` The type is &#34;char&#34; and the variable is &#34;str&#34;.
- `float`*` The type is &#34;float&#34; and there is no variable.
- `char`**words` The type is &#34;char*&#34; and the variable is &#34;words&#34;.


** Addressing and Dereferencing



As I mentioned above, pointers are simply a number.  By convention, this number is one of two things: the address of a value in memory, or NULL, which is zero on most systems.  The special value NULL is used to indicate the absence of a value, e.g. as the return value of a function.  This does not mean that pointers cannot hold other numbers, though using them for other purposes will probably make your code much less readable.  Because of this convention, the C compiler requires that you choose a value type for the pointer, and it uses this to interpret the memory at the location specified by the pointer when you *dereference* it.  It also uses this type to give you warnings when you *address* another variable and assign it to the pointer, should the types be incompatible.



Dereferencing is arguably the most common operation you will see performed on a pointer.  The operation takes the number stored in the pointer (because a pointer is a number) and looks that location up in memory.  It then takes the value type of the pointer (e.g. char for a char*) and interprets the memory at that location as if it were that data type.  This does NOT have to be the actual data type of any variables which happen to be stored at that memory location, though if it isn&#39;t you should probably really know what you&#39;re doing or the results could surprise you.  The syntax for dereferencing a pointer is simply `*var` where var is any pointer variable.  One trick to figuring out what value a dereferencing expression has is to look at its data type and &#34;cancel&#34; out an asterisk.  For example, for `char`*x`, the expression `*x` has one fewer asterisk, and thus is a char.  In the same way, for `int`***p`, the expression `**p` is int*.  (Try to figure out the answers for yourself before you highlight them.)  When you are reading code aloud, the dereferencing operator is often read &#34;the value at,&#34; so the code `12`&#43;`*p` would be read &#34;twelve plus the value at p.&#34;



Addressing is the opposite of dereferencing.  If dereferencing takes the number stored in the pointer and makes it a value, then it only makes sense that an operation exists to take a value in memory and get the number representing its address.  In C, such an operation is only valid on a variable, and only on a variable stored on the stack or in the heap.  The syntax is simply `&amp;var` where var is any variable.  The trick to figuring out the resulting type of an addressing expression is to add an asterisk to the variable&#39;s type (the counterpart to canceling one out in the case of the dereference option).  For example, with `int`y`, the expression `&amp;y` has type `int*`.  To combine the two, with `float`**x`, the expression `***(&amp;x)` has type int.


** Arithmetic and Indexing



I will cover pointer arithmetic briefly here, as a means of explaining arrays in C, but for a full explanation you will probably want to go have a look at my [[http://kylelemons.net/downloads/?did=1][pointers PDF]].  When you add a number to a pointer, you are not adding directly to the base number stored in the pointer.  What you are adding are called &#34;strides&#34; of the pointer&#39;s value type, which correspond to the `sizeof` that type.  So, if you add 10 to a `int*` you are actually adding `10*sizeof(int)`.  On a 64-bit system, the size of an integer is 8 bytes, and so you are actually adding 80 to the pointer value.  Another way of thinking about this is that you are going on to the nUnknown atom &#34;sup&#34; []

th value type after the one pointed to by the pointer, where the value stored at the address in the pointer is the zeroth.  So, if you add 1 to a `char*` you are referencing the character immediately after the one stored in memory at the location referenced in the pointer.



If you think all that talk of &#34;the nUnknown atom &#34;sup&#34; []

th value after the one referenced by the pointer&#34; sounds like arrays, you&#39;d be correct.  Because an array in C is essentially a pointer to the first element in that array, we can write `*arr` where arr is some suitable array variable.  Of course, from basic math class, nobody should be surprised that this is equivalent to `*(arr&#43;0)`.  Now, from the discussion above, it should not be a surprise to you that `*(arr&#43;1)` will now be the second (or the 1th) value in the array.  Because this syntax is cumbersome, C has given us some syntactic sugar such that we can say `a[b]` and it will interpret this as `*(a&#43;b)`.  This makes it possible to do some really interesting things, but you will have to go read my pointers PDF for that.  For now, I should simply mention that because of this property of indexing arrays in C, you need to be very careful because there is absolutely no bounds checking done by the compiler or the computer, and thus if you index too far off your array into no-man&#39;s land, the operating system will give your program a swift kick, causing it to report &#34;Segmentation Fault&#34; or, affectionately, SIGSEGV.


** The Output



If you&#39;ve read through everything above, I challenge you to spend some quality time tracing through the code above before you continue on and see the output.  If you can&#39;t figure some piece out, go read my pointers PDF and then take a crack at it.  If that still doesn&#39;t work, then have a look at the output and try to work backward to figure out what that particular piece of dark pointer magic was doing.  It is also important to compare the array notation and the pointer notation, as I purposefully did nearly identical tasks using both to help demonstrate the dual syntaxes for pointers.


  ------ Sizes ------
  String 1: size=32  length=31 
  String 2: size=4   length=22 
  Buffer:   size=256 length=0  
  Indices:  size=128 length=32 
  Words:    size=128 length=32 
  Pick 1:   size=24  length=6  
  Pick 2:   size=16  length=4  
  ------ Final ------
  My neighbor just got a dog that is a pointer.
