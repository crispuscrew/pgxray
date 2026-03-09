package db                                                    
				
import (
	"context"
	"testing"
	"strings"
)

func TestLoadDatabaseInfo(t *testing.T) {
	dbInfo, err := testConn.LoadDatabaseInfo(context.Background())
	if err != nil { t.Fatalf("LoadDatabaseInfo failed: %v", err) }

	if dbInfo.Name != testConn.dbName { t.Errorf("Expected database name %q, got %q", testConn.dbName, dbInfo.Name) }
	if !dbInfo.ChildrenLoaded { t.Errorf("Expected ChildrenLoaded to be true, got false") }

	if len(dbInfo.Schemas) == 0 { t.Errorf("Expected at least one schema, got 0") }

	foundPublic := false
	for _, schema := range dbInfo.Schemas {
		if strings.HasPrefix(schema.Name, "pg_") || schema.Name == "information_schema" {
			t.Errorf("Unexpected system schema %q included in results", schema.Name)
		} else if schema.Name == "public" {
			foundPublic = true
		}
	}
	
	if !foundPublic { t.Errorf("Expected to find 'public' schema, but it was not found") }
}

func TestLoadSchemaInfo(t *testing.T) {
	schemaInfo, err := testConn.LoadSchemaInfo(context.Background(), "public")
	if err != nil { t.Fatalf("LoadSchemaInfo failed: %v", err) }

	if schemaInfo.Name != "public" { t.Errorf("Expected schema name 'public', got %q", schemaInfo.Name) }
	if !schemaInfo.ChildrenLoaded { t.Errorf("Expected ChildrenLoaded to be true, got false") }

	if len(schemaInfo.Tables) == 0 { t.Errorf("Expected at least one table, got 0") }

	foundUsers := false
	for _, table := range schemaInfo.Tables {
		if strings.HasPrefix(table.Name, "pg_") || table.Name == "information_schema" {
			t.Errorf("Unexpected system table %q included in results", table.Name)
		} else if table.Name == "users" {
			foundUsers = true
		}
	}
	
	if !foundUsers { t.Errorf("Expected to find 'users' table, but it was not found") }
}