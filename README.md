# assert
A simple lightweight assert extension for use with testing.T


# Example usage

```
	Assert(t, 1, EqualInt, 1, "Expected both to be equal")
	Assert(t, someArray, Len, 3, "Expected array with a length of 3")
```

Currently there are the following comparison functions

* EqualDeep
* EqualString
* EqualInt
* EqualFloat
* NotEqualDeep
* NotEqualString
* NotEqualInt
* NotEqualFloat
* Len

Feel free to contribute more comparisons that are useful.
