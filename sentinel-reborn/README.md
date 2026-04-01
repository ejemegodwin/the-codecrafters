README.md must include:

- What the program does
- How to run it with examples
- Each transformation listed and explained
- What your personal contribution was
- One thing you found hardest today
- One thing you understand now that
  you did not understand this morning


## hmusa

I implemented the function `lowN()` that handles the transformation of a specified number of words in a string to lowercase

How this was done:
First, i found the marker by identifying the index of the opening and the closing parentheses

Secondly, i put the words before the marker into a slice so they'll be easier to manipulate through indexing

Thirdly, i looped inside the marker to find the number, then created another loop starting at the lenght of the slice minus that number, then transform everything from that point onwards

Lastly, i appended the remaining of the text after the marker to that slice and return it as a string through strings.Join.
---
