package packages

import (
	"path/filepath"
)

func init() {
	Packages["path/filepath"] = map[string]interface{}{
		"Abs":          filepath.Abs,
		"Base":         filepath.Base,
		"Clean":        filepath.Clean,
		"Dir":          filepath.Dir,
		"EvalSymlinks": filepath.EvalSymlinks,
		"Ext":          filepath.Ext,
		"FromSlash":    filepath.FromSlash,
		"Glob":         filepath.Glob,
		"HasPrefix":    filepath.HasPrefix,
		"IsAbs":        filepath.IsAbs,
		"Join":         filepath.Join,
		"Match":        filepath.Match,
		"Rel":          filepath.Rel,
		"Split":        filepath.Split,
		"SplitList":    filepath.SplitList,
		"ToSlash":      filepath.ToSlash,
		"VolumeName":   filepath.VolumeName,
	}
}
