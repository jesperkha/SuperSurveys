## Code Formatting Guidelines

<br>

### Variables

- All variables should be self-descriptive unless they are temporary or the context makes it obvious what they are used for.

<br>

### Functions

- Functions names should be descriptive and a user should know what the function does based on the function and parameter names. Additional info can be added in a function descriptor if necessary.
- All functions that could create an error should return said error and exit gracefully.
- Errors returned by functions should always be handled.
- Functions should have a 2 line gap between eachother.

<br>

### Packages

- Package names should be self-descriptive and one word.
- All packages must contain a main file with the same name as the package. The main file should declare all global variables and types. All other files should just contain functions.

<br>

### Other

- Your code should use minimal indentation to prioritize readability. Avoid 3+ line indentations where possible.
- Do your best at all times to keep code readable to the point where comments are only needed in rare cases.
- Keep functions short and consise and, if possible, under 50 lines. If your functions exceeds this try to break it down into smaller components.
- Use line breaks to create smaller blocks of code that are easy to digest. Avoid longwinded code where possible.
