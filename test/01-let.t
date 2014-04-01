a = nil
is(nil, a, "let nil")
a = 1
is(1, a, "let int")
a = "foo"
is("foo", a, "let string")
is("1", toString(1), "toString(int)")
is("1.2", toString(1.2), "toString(float)")
