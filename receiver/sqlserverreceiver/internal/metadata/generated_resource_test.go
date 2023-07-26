// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetSqlserverComputerName("sqlserver.computer.name-val")
			rb.SetSqlserverDatabaseName("sqlserver.database.name-val")
			rb.SetSqlserverInstanceName("sqlserver.instance.name-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return 0

			switch test {
			case "default":
				assert.Equal(t, 1, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 3, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("sqlserver.computer.name")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "sqlserver.computer.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("sqlserver.database.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "sqlserver.database.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("sqlserver.instance.name")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "sqlserver.instance.name-val", val.Str())
			}
		})
	}
}