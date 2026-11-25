How to declare a variable in Go?
First what is a variable? A variable is a storage location in memory that has a name and a type. It can hold different values during the execution of a program.

var name type = value

var is a keyword used to declare a variable. The name is the identifier for the variable, and the type specifies what kind of data the variable can hold. The value is the initial value assigned to the variable.

name: The name of the variable. It should be a valid identifier in Go, which means it can contain letters, digits, and underscores, but cannot start with a digit.

type: The type of the variable. Go is a statically typed language, which means you must specify the type of the variable when you declare it. Common types include int, float64, string, bool, etc.

There are two ways to declare a variable in Go:

1. Using the var keyword:
var variableName type = value

2. Using the shorthand declaration:
variableName := value
Note: Use shorthand declarations only inside a function. Go infers the variable's type from the value you assign.
