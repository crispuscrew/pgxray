package db

type DBInstance struct {
	Name 				string
	IsLoaded		bool
}

type Database struct {
	DBInstance
	Schemas 			[]Schema
}

type Schema struct {
	DBInstance
	Tables 				[]Table
	Views 				[]View
	Sequences 			[]Sequence
}

type Table struct {
	DBInstance
	Columns 			[]Column
	Indexes 			[]Index
	Constraints 		[]Constraint
}

type Column struct {
	DBInstance
	DataType 			string
	notNullable			bool
	Default      		string
	Comment      		string
}

type Index struct {
	DBInstance
	Columns   			[]string
	Kind      			string  // btree, hash, gin...
	IsUnique  			bool
	Partial   			string  // WHERE condition if exists
}

type Constraint struct {
	DBInstance
	Kind       			string  // PRIMARY KEY, FOREIGN KEY, CHECK, UNIQUE
	Definition			string
}

type View struct {
	DBInstance
	Definition 			string
}

type Sequence struct {
	DBInstance
	CurrentValue 		*int64
	MinValue     		int64
	MaxValue     		int64
	Increment    		int64
}