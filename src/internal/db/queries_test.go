package db

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestLoadDatabaseInfo(t *testing.T) {
	dbInfo, err := testConn.LoadDatabaseInfo(context.Background())
	if err != nil { t.Fatalf("LoadDatabaseInfo failed: %v", err) }

	if dbInfo.Name != testConn.dbName { t.Errorf("Expected database name %q, got %q", testConn.dbName, dbInfo.Name) }
	if !dbInfo.IsLoaded { t.Errorf("Expected IsLoaded to be true, got false") }

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
	if !schemaInfo.IsLoaded { t.Errorf("Expected IsLoaded to be true, got false") }

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

func TestLoadTableInfo(t *testing.T) {
	tableInfo, err := testConn.LoadTableInfo(context.Background(), "public", "users")
	if err != nil { t.Fatalf("LoadTableInfo failed: %v", err) }

	if tableInfo.Name != "users" { t.Errorf("Expected table name 'users', got %q", tableInfo.Name) }
	if !tableInfo.IsLoaded { t.Errorf("Expected IsLoaded to be true") }

	// Columns — seed has: id, name, email, created_at
	if len(tableInfo.Columns) != 4 { t.Errorf("Expected 4 columns, got %d", len(tableInfo.Columns)) }

	colByName := make(map[string]Column)
	for _, col := range tableInfo.Columns { colByName[col.Name] = col }

	if _, ok := colByName["id"]; !ok { t.Errorf("Expected column 'id'") }
	if _, ok := colByName["email"]; !ok { t.Errorf("Expected column 'email'") }

	if colByName["id"].notNullable != true  { t.Errorf("Column 'id' should be not nullable") }
	if colByName["name"].notNullable != true { t.Errorf("Column 'name' should be not nullable") }
	if colByName["email"].notNullable != false { t.Errorf("Column 'email' should be nullable") }

	if colByName["id"].Default == ""          { t.Errorf("Column 'id' should have a default (nextval)") }
	if colByName["created_at"].Default == ""  { t.Errorf("Column 'created_at' should have a default (now())") }
	if colByName["name"].Default != ""        { t.Errorf("Column 'name' should have no default") }

	// Indexes — seed has: users_pkey, idx_users_email
	if len(tableInfo.Indexes) < 2 { t.Errorf("Expected at least 2 indexes, got %d", len(tableInfo.Indexes)) }

	idxByName := make(map[string]Index)
	for _, idx := range tableInfo.Indexes { idxByName[idx.Name] = idx }

	if _, ok := idxByName["users_pkey"]; !ok { t.Errorf("Expected index 'users_pkey'") }
	if _, ok := idxByName["idx_users_email"]; !ok { t.Errorf("Expected index 'idx_users_email'") }

	if idxByName["users_pkey"].Kind != "btree"     { t.Errorf("Expected users_pkey to be btree") }
	if idxByName["users_pkey"].IsUnique != true     { t.Errorf("Expected users_pkey to be unique") }
	if idxByName["idx_users_email"].IsUnique != false { t.Errorf("Expected idx_users_email to be non-unique") }
	if idxByName["idx_users_email"].Partial != ""   { t.Errorf("Expected idx_users_email to be non-partial") }

	// Constraints — seed has: users_pkey (PRIMARY KEY), users_email_key (UNIQUE)
	if len(tableInfo.Constraints) < 2 { t.Errorf("Expected at least 2 constraints, got %d", len(tableInfo.Constraints)) }

	conByName := make(map[string]Constraint)
	for _, con := range tableInfo.Constraints { conByName[con.Name] = con }

	if _, ok := conByName["users_pkey"]; !ok { t.Errorf("Expected constraint 'users_pkey'") }
	if conByName["users_pkey"].Kind != "PRIMARY KEY" { t.Errorf("Expected users_pkey to be PRIMARY KEY") }

	if _, ok := conByName["users_email_key"]; !ok { t.Errorf("Expected constraint 'users_email_key'") }
	if conByName["users_email_key"].Kind != "UNIQUE" { t.Errorf("Expected users_email_key to be UNIQUE") }
}

func TestLoadTableInfo_PartialIndex(t *testing.T) {
	tableInfo, err := testConn.LoadTableInfo(context.Background(), "public", "orders")
	if err != nil { t.Fatalf("LoadTableInfo failed: %v", err) }

	idxByName := make(map[string]Index)
	for _, idx := range tableInfo.Indexes { idxByName[idx.Name] = idx }

	if _, ok := idxByName["idx_orders_pending"]; !ok { t.Fatalf("Expected index 'idx_orders_pending'") }
	if idxByName["idx_orders_pending"].Partial == "" { t.Errorf("Expected idx_orders_pending to be partial") }
}

func TestLoadTableInfo_ForeignKey(t *testing.T) {
	tableInfo, err := testConn.LoadTableInfo(context.Background(), "public", "orders")
	if err != nil { t.Fatalf("LoadTableInfo failed: %v", err) }

	conByName := make(map[string]Constraint)
	for _, con := range tableInfo.Constraints { conByName[con.Name] = con }

	if _, ok := conByName["orders_user_id_fkey"]; !ok { t.Fatalf("Expected constraint 'orders_user_id_fkey'") }
	if conByName["orders_user_id_fkey"].Kind != "FOREIGN KEY" { t.Errorf("Expected orders_user_id_fkey to be FOREIGN KEY") }
}

func TestLoadQuery(t *testing.T) {
	result, err := testConn.LoadQuery(context.Background(), "SELECT id, name FROM users ORDER BY id")
	if err != nil { t.Fatalf("LoadQuery failed: %v", err) }

	if len(result.ColumnName) != 2 { t.Errorf("Expected 2 columns, got %d", len(result.ColumnName)) }
	if result.ColumnName[0] != "id"   { t.Errorf("Expected column 0 to be 'id', got %q", result.ColumnName[0]) }
	if result.ColumnName[1] != "name" { t.Errorf("Expected column 1 to be 'name', got %q", result.ColumnName[1]) }

	if result.ColumnType[0] != "int4" { t.Errorf("Expected id type 'int4', got %q", result.ColumnType[0]) }
	if result.ColumnType[1] != "text" { t.Errorf("Expected name type 'text', got %q", result.ColumnType[1]) }

	if len(result.Rows) != 2 { t.Errorf("Expected 2 rows (Alice, Bob), got %d", len(result.Rows)) }
}

func TestLoadQuery_RejectsWrite(t *testing.T) {
	_, err := testConn.LoadQuery(context.Background(), "INSERT INTO users (name) VALUES ('hacker')")
	if err == nil { t.Errorf("Expected error for write query in read-only transaction, got nil") }
}

func TestLoadTableRows(t *testing.T) {
	result, err := testConn.LoadTableRows(context.Background(), "public", "users", 1, 10, 0)
	if err != nil { t.Fatalf("LoadTableRows failed: %v", err) }

	if len(result.ColumnName) == 0 { t.Errorf("Expected columns, got 0") }
	if len(result.Rows) != 2 { t.Errorf("Expected 2 rows, got %d", len(result.Rows)) }
}

func TestLoadTableRows_Pagination(t *testing.T) {
	page1, err := testConn.LoadTableRows(context.Background(), "public", "users", 1, 1, 0)
	if err != nil { t.Fatalf("LoadTableRows page1 failed: %v", err) }

	page2, err := testConn.LoadTableRows(context.Background(), "public", "users", 1, 1, 1)
	if err != nil { t.Fatalf("LoadTableRows page2 failed: %v", err) }

	if len(page1.Rows) != 1 { t.Errorf("Expected 1 row on page1, got %d", len(page1.Rows)) }
	if len(page2.Rows) != 1 { t.Errorf("Expected 1 row on page2, got %d", len(page2.Rows)) }

	if fmt.Sprintf("%v", page1.Rows[0]) == fmt.Sprintf("%v", page2.Rows[0]) {
		t.Errorf("Expected page1 and page2 to return different rows")
	}
}