package libs

import (
	_ "github.com/lib/pq"
	"github.com/three-plus-three/modules/permissions"
	"github.com/three-plus-three/modules/types"
	"github.com/three-plus-three/modules/web_ext"
)

// Lifecycle 表示一个运行周期，它包含了所有业务相关的对象
type Lifecycle struct {
	*web_ext.Lifecycle
	Definitions *types.TableDefinitions
	DB          permissions.DB
	DataDB      permissions.DB
}
