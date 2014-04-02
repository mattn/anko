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

# vim: set ft=anko:

