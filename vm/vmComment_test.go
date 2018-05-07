package vm

import (
	"os"
	"testing"
)

func TestComment(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `# 1`},
		{script: `# 1;`},
		{script: `# 1 // 2`},
		{script: `# 1 \n 2`},
		{script: `# 1 # 2`},

		{script: `1# 1`, runOutput: int64(1)},
		{script: `1# 1;`, runOutput: int64(1)},
		{script: `1# 1 // 2`, runOutput: int64(1)},
		{script: `1# 1 \n 2`, runOutput: int64(1)},
		{script: `1# 1 # 2`, runOutput: int64(1)},

		{script: `1
# 1`, runOutput: int64(1)},
		{script: `1
# 1;`, runOutput: int64(1)},
		{script: `1
# 1 // 2`, runOutput: int64(1)},
		{script: `1
# 1 \n 2`, runOutput: int64(1)},
		{script: `1
# 1 # 2`, runOutput: int64(1)},

		{script: `// 1`},
		{script: `// 1;`},
		{script: `// 1 // 2`},
		{script: `// 1 \n 2`},
		{script: `// 1 # 2`},

		{script: `1// 1`, runOutput: int64(1)},
		{script: `1// 1;`, runOutput: int64(1)},
		{script: `1// 1 // 2`, runOutput: int64(1)},
		{script: `1// 1 \n 2`, runOutput: int64(1)},
		{script: `1// 1 # 2`, runOutput: int64(1)},

		{script: `1
// 1`, runOutput: int64(1)},
		{script: `1
// 1;`, runOutput: int64(1)},
		{script: `1
// 1 // 2`, runOutput: int64(1)},
		{script: `1
// 1 \n 2`, runOutput: int64(1)},
		{script: `1
// 1 # 2`, runOutput: int64(1)},

		{script: `/* 1 */`},
		{script: `/* * 1 */`},
		{script: `/* 1 * */`},
		{script: `/** 1 */`},
		{script: `/*** 1 */`},
		{script: `/**** 1 */`},
		{script: `/* 1 **/`},
		{script: `/* 1 ***/`},
		{script: `/* 1 ****/`},
		{script: `/** 1 ****/`},
		{script: `/*** 1 ****/`},
		{script: `/**** 1 ****/`},

		{script: `1/* 1 */`, runOutput: int64(1)},
		{script: `1/* * 1 */`, runOutput: int64(1)},
		{script: `1/* 1 * */`, runOutput: int64(1)},
		{script: `1/** 1 */`, runOutput: int64(1)},
		{script: `1/*** 1 */`, runOutput: int64(1)},
		{script: `1/**** 1 */`, runOutput: int64(1)},
		{script: `1/* 1 **/`, runOutput: int64(1)},
		{script: `1/* 1 ***/`, runOutput: int64(1)},
		{script: `1/* 1 ****/`, runOutput: int64(1)},
		{script: `1/** 1 ****/`, runOutput: int64(1)},
		{script: `1/*** 1 ****/`, runOutput: int64(1)},
		{script: `1/**** 1 ****/`, runOutput: int64(1)},

		{script: `/* 1 */1`, runOutput: int64(1)},
		{script: `/* * 1 */1`, runOutput: int64(1)},
		{script: `/* 1 * */1`, runOutput: int64(1)},
		{script: `/** 1 */1`, runOutput: int64(1)},
		{script: `/*** 1 */1`, runOutput: int64(1)},
		{script: `/**** 1 */1`, runOutput: int64(1)},
		{script: `/* 1 **/1`, runOutput: int64(1)},
		{script: `/* 1 ***/1`, runOutput: int64(1)},
		{script: `/* 1 ****/1`, runOutput: int64(1)},
		{script: `/** 1 ****/1`, runOutput: int64(1)},
		{script: `/*** 1 ****/1`, runOutput: int64(1)},
		{script: `/**** 1 ****/1`, runOutput: int64(1)},

		{script: `1
/* 1 */`, runOutput: int64(1)},
		{script: `1
/* * 1 */`, runOutput: int64(1)},
		{script: `1
/* 1 * */`, runOutput: int64(1)},
		{script: `1
/** 1 */`, runOutput: int64(1)},
		{script: `1
/*** 1 */`, runOutput: int64(1)},
		{script: `1
/**** 1 */`, runOutput: int64(1)},
		{script: `1
/* 1 **/`, runOutput: int64(1)},
		{script: `1
/* 1 ***/`, runOutput: int64(1)},
		{script: `1
/* 1 ****/`, runOutput: int64(1)},
		{script: `1
/** 1 ****/`, runOutput: int64(1)},
		{script: `1
/*** 1 ****/`, runOutput: int64(1)},
		{script: `1
/**** 1 ****/`, runOutput: int64(1)},
	}
	runTests(t, tests)
}
