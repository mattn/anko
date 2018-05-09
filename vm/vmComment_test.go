package vm

import (
	"os"
	"testing"

	"github.com/mattn/anko/internal/testlib"
)

func TestComment(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
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
	testlib.RunTests(t, tests, nil)
}
