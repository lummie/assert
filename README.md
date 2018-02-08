# assert
A simple lightweight assert extension for use with testing.T


# Example usage

```
func TestSimpleTransitions(t *testing.T) {
	tr := fsm.NewTransitions()
	assert.NotNil(t, tr, "Transitions should have been created")

    assert.Equal(t, 12,12, "Should match")
    assert.NotEqual(t, "expected", "actual", "Should not match")

	tr.Add(Red, RedAmber, nil, nil, "")

	ok, g, a := tr.Get(Red, Green)
	assert.False(t, ok, "Red, Green not expected")
	assert.Nil(t, g)
	assert.Nil(t, a)

	ok, g, a = tr.Get(Red, RedAmber)
	assert.True(t, ok, "Expected Red, RedAmber")
	assert.Nil(t, g)
	assert.Nil(t, a)
}
```

Currently there are the following comparison functions

* Equal | NotEqual
* Len
* Nil | NotNil
* True | False
* Error | NotError


Generally each methods is supplied with:

1. a ```*testing.T``` as the first parameter.
2. An **expected** then **actual** value (e.g. for Equal & Len) or just an **actual** value
3. *Optional* Messages that will be output on a failure


If you construct a test that asserts in a descendant function, then the line reported in the assert failure will be the 
line in that descendant function. You can alter this so the line in the calling function is identified by adding a call 
to `CallStackAdjust(1)` at the start of the descendant and `CallStackReset()` at the end.

The parameter to CallStackAdjust is the number of stack levels UP to report, usually 1 will suffice, but if you have 
several levels of nesting before the assert then this can be adjusted as necessary. 


