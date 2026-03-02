package db

type Database struct {
	Name 				string
	Schemas 			[]Schema
}

type Schema struct {
	Name 				string
	Tables 				[]Table
	Views 				[]View
	Sequences 			[]Sequence
}

type Table struct {
	Name 				string
	Columns 			[]Column
	Indexes 			[]Index
	Constraints 		[]Constraint
}

type Column struct {
	Name 				string
	DataType 			string
	IsNullable 			bool
	Default      		string
	Comment      		string
}

type Index struct {
	Name      			string
	Columns   			[]string
	Kind      			string  // btree, hash, gin...
	IsUnique  			bool
	Partial   			string  // WHERE condition if exists
}

type Constraint struct {
	Name       			string
	Kind       			string  // PRIMARY KEY, FOREIGN KEY, CHECK, UNIQUE
	Columns    			[]string
	References 			string  // for FK: "table(column)"
	Check      			string  // for CHECK: condition
}

type View struct {
	Name       			string
	Definition 			string
}

type Sequence struct {
	Name         		string
	CurrentValue 		int64
	MinValue     		int64
	MaxValue     		int64
	Increment    		int64
}