

ok(1 > 0, "1 > 0")
ok(1 != 1.0, "1 != 1.0")
ok(1 != "1", "1 != \"1\"")
ok(1 == 1, "1 == 1")
ok(1.1 == 1.1, "1.1 == 1.1")
ok("1" == "1", "\"1\" == \"1\"")

ok(false != "1", "false != \"1\"")
ok(false != true, "false != true")
ok(false == false, "false == false")
ok(true == true, "true == true")
ok(true == true, "true == true")

ok(1 <= 1, "1 <= 1")
ok(1.0 <= 1.0, "1.0 <= 1.0")

is(true, 1 <= 2 ? true : false, "1 == 1 ? true : false")

a = 1; a += 1
is(2, a, "+=")

a = 2; a -= 1
is(1, a, "-=")

a = 2; a *= 2
is(4, a, "*=")

a = 3; a /= 2
is(1.5, a, "/=")

a = 2; a++
is(3, a, "++")

a = 2; a--
is(1, a, "--")

a = 2**3
is(8, a, "**")

a = 1
a &= 2
is(0, a, "&=")

a = 1
a |= 2
is(3, a, "|=")

# vim: set ft=anko:

