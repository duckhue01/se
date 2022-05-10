## tips and to for writing lean and maintainable go code
### Optimizing function implementations for readability
- "if a function implementation does not fit on a single screen, then it must be split up into smaller function"
- it follow the SRP and the function name talk all about what it do
- if you find yourself navigating through a lengthy function that contains deeply nested if/else blocks, repeated block of code, or its implementation tackles several seemingly unrelated concerns, it would be a great opportunity to apply some drive-by refactoring and extract any potential self-contained blocks of logic into separate functions.
- everything have drawback, splitting the business logic across function sometimes takes a toll on performance (pushing arguments to stack, popping things off the stack when the function call returns, checking stack have enough space or need to spread up)
- when the end goal is to produce high-performance code that contains tight inner loops or code that is expected to be called with high frequency, it may be a good idea to keep the implementation neatly tucked within a function.

### Variable naming conventions
### Using go interfaces effectively
- " accept interface, return structs"
### zero values are tour friends
- go channel, nil channels indefinitely block go routines attempting to read off them
- the zero value for go slice, this is an empty slice that this can be appended to
- the sync.Mutex type, whose zero value indicates that the mutex is unlocked
- the bytes.Buffer type, whose zero value indicates an empty buffer
  
### using tools to analyze and manipulate go programs