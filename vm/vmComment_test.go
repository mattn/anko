package vm

import (
	"os"
	"testing"
)

func TestComment(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `# 1`},
		{Script: `# 1;`},
		{Script: `# 1 // 2`},
		{Script: `# 1 \n 2`},
		{Script: `# 1 # 2`},

		{Script: `1# 1`, RunOutput: int64(1)},
		{Script: `1# 1;`, RunOutput: int64(1)},
		{Script: `1# 1 // 2`, RunOutput: int64(1)},
		{Script: `1# 1 \n 2`, RunOutput: int64(1)},
		{Script: `1# 1 # 2`, RunOutput: int64(1)},

		{Script: `1
# 1`, RunOutput: int64(1)},
		{Script: `1
# 1;`, RunOutput: int64(1)},
		{Script: `1
# 1 // 2`, RunOutput: int64(1)},
		{Script: `1
# 1 \n 2`, RunOutput: int64(1)},
		{Script: `1
# 1 # 2`, RunOutput: int64(1)},

		{Script: `// 1`},
		{Script: `// 1;`},
		{Script: `// 1 // 2`},
		{Script: `// 1 \n 2`},
		{Script: `// 1 # 2`},

		{Script: `1// 1`, RunOutput: int64(1)},
		{Script: `1// 1;`, RunOutput: int64(1)},
		{Script: `1// 1 // 2`, RunOutput: int64(1)},
		{Script: `1// 1 \n 2`, RunOutput: int64(1)},
		{Script: `1// 1 # 2`, RunOutput: int64(1)},

		{Script: `1
// 1`, RunOutput: int64(1)},
		{Script: `1
// 1;`, RunOutput: int64(1)},
		{Script: `1
// 1 // 2`, RunOutput: int64(1)},
		{Script: `1
// 1 \n 2`, RunOutput: int64(1)},
		{Script: `1
// 1 # 2`, RunOutput: int64(1)},

		{Script: `/* 1 */`},
		{Script: `/* * 1 */`},
		{Script: `/* 1 * */`},
		{Script: `/** 1 */`},
		{Script: `/*** 1 */`},
		{Script: `/**** 1 */`},
		{Script: `/* 1 **/`},
		{Script: `/* 1 ***/`},
		{Script: `/* 1 ****/`},
		{Script: `/** 1 ****/`},
		{Script: `/*** 1 ****/`},
		{Script: `/**** 1 ****/`},

		{Script: `1/* 1 */`, RunOutput: int64(1)},
		{Script: `1/* * 1 */`, RunOutput: int64(1)},
		{Script: `1/* 1 * */`, RunOutput: int64(1)},
		{Script: `1/** 1 */`, RunOutput: int64(1)},
		{Script: `1/*** 1 */`, RunOutput: int64(1)},
		{Script: `1/**** 1 */`, RunOutput: int64(1)},
		{Script: `1/* 1 **/`, RunOutput: int64(1)},
		{Script: `1/* 1 ***/`, RunOutput: int64(1)},
		{Script: `1/* 1 ****/`, RunOutput: int64(1)},
		{Script: `1/** 1 ****/`, RunOutput: int64(1)},
		{Script: `1/*** 1 ****/`, RunOutput: int64(1)},
		{Script: `1/**** 1 ****/`, RunOutput: int64(1)},

		{Script: `/* 1 */1`, RunOutput: int64(1)},
		{Script: `/* * 1 */1`, RunOutput: int64(1)},
		{Script: `/* 1 * */1`, RunOutput: int64(1)},
		{Script: `/** 1 */1`, RunOutput: int64(1)},
		{Script: `/*** 1 */1`, RunOutput: int64(1)},
		{Script: `/**** 1 */1`, RunOutput: int64(1)},
		{Script: `/* 1 **/1`, RunOutput: int64(1)},
		{Script: `/* 1 ***/1`, RunOutput: int64(1)},
		{Script: `/* 1 ****/1`, RunOutput: int64(1)},
		{Script: `/** 1 ****/1`, RunOutput: int64(1)},
		{Script: `/*** 1 ****/1`, RunOutput: int64(1)},
		{Script: `/**** 1 ****/1`, RunOutput: int64(1)},

		{Script: `1
/* 1 */`, RunOutput: int64(1)},
		{Script: `1
/* * 1 */`, RunOutput: int64(1)},
		{Script: `1
/* 1 * */`, RunOutput: int64(1)},
		{Script: `1
/** 1 */`, RunOutput: int64(1)},
		{Script: `1
/*** 1 */`, RunOutput: int64(1)},
		{Script: `1
/**** 1 */`, RunOutput: int64(1)},
		{Script: `1
/* 1 **/`, RunOutput: int64(1)},
		{Script: `1
/* 1 ***/`, RunOutput: int64(1)},
		{Script: `1
/* 1 ****/`, RunOutput: int64(1)},
		{Script: `1
/** 1 ****/`, RunOutput: int64(1)},
		{Script: `1
/*** 1 ****/`, RunOutput: int64(1)},
		{Script: `1
/**** 1 ****/`, RunOutput: int64(1)},
	}
	RunTests(t, tests, nil)
}
