# Type Casting

Type casting happens when we assign a value of one data type to another data type. Statically typed languages like C++ and Java provide the support for Implicit Type Conversion. Still, the Go language is different, as it doesn’t support an Automatic Type Conversion or Implicit Type Conversion even if the data types are compatible.

## Types of Typecasting
1. Implicit Typecasting
2. Explicit Typecasting

## Syntax
If you want to convert a value from one data type to another, you must follow the following syntax.
```
newDataTypeVariable = T(oldDataTypeVariable)
```
Where T is the new data type.
## Example
```
var badboys int = 1921
// explicit type conversion
var badboys2 float64 = float64(badboys)
```

## It conversion not casting!
In other programming languages we call it casting, but we don't do that here :P, rather we call it conversion.