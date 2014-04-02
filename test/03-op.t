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

# vim: set ft=anko:

