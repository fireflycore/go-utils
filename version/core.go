package version

import (
	"fmt"
	"strconv"
	"strings"
)

// Version 表示一个语义化版本号
// 遵循 SemVer 规范：Major.Minor.Patch
type Version struct {
	Major int // 主版本号：不兼容的 API 修改
	Minor int // 次版本号：向下兼容的功能性新增
	Patch int // 修订号：向下兼容的问题修正
}

// NewVersion 解析版本字符串并返回 Version 结构体
// verStr: 版本字符串（如 "v1.0.1" 或 "1.0.1"）
// 返回: Version 指针，如果格式错误则返回 error
func NewVersion(verStr string) (*Version, error) {
	verStr = strings.TrimPrefix(verStr, "v") // 去除前缀 'v'
	parts := strings.Split(verStr, ".")      // 按点分割
	if len(parts) != 3 {
		return nil, fmt.Errorf("版本格式不正确: %s", verStr)
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}
	return &Version{Major: major, Minor: minor, Patch: patch}, nil
}

// String 格式化输出版本号
// 格式: "v{Major}.{Minor}.{Patch}"
func (v *Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Increment 对版本进行累加，修订号满 100 则次版本号加 1，次版本号满 100 则主版本号加 1。
func (v *Version) Increment() {
	v.Patch++
	if v.Patch == 100 {
		v.Patch = 0
		v.Minor++
		if v.Minor == 100 {
			v.Minor = 0
			v.Major++
		}
	}
}

// Compare 比较两个版本号的大小
// other: 待比较的另一个版本
// 返回:
//   - 1: 当前版本 > other
//   - 0: 当前版本 == other
//   - -1: 当前版本 < other
func (v *Version) Compare(other *Version) int {
	if v.Major > other.Major {
		return 1
	} else if v.Major < other.Major {
		return -1
	}
	if v.Minor > other.Minor {
		return 1
	} else if v.Minor < other.Minor {
		return -1
	}
	if v.Patch > other.Patch {
		return 1
	} else if v.Patch < other.Patch {
		return -1
	}
	return 0
}
