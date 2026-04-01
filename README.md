## hmusa

I implemented the function `lowN()` that handles the transformation of a specified number of words in a string to lowercase

How this was done:
First, i found the marker by identifying the index of the opening and the closing parentheses

Secondly, i put the words before the marker into a slice so they'll be easier to manipulate through indexing

Thirdly, i looped inside the marker to find the number, then created another loop starting at the lenght of the slice minus that number, then transform everything from that point onwards

Lastly, i appended the remaining of the text after the marker to that slice and return it as a string through strings.Join.
---

## kinonoja

When we started; I was given the job of handling the conversion of hexadecimal to decimal. 
* HOW THIS WAS DONE
First, I had to use "strconv" because we were converting a string to an integer

Then i decided to declare a variable that err := strconv.ParseInt cause ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.

After that i used if statement that if the error is not eqaul to it should return 0 and print "invalid hex string:" then i returned decimal and the nil figure

Then i had to wriite another function that will bring the input and return everything in the output figure as presented using regexp.Must.Compile

Then after i that, i joined in writing the read file and input to make sure the file brings out the output in everything

