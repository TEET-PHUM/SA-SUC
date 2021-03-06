// Code generated by entc, DO NOT EDIT.

package foodmenu

const (
	// Label holds the string label denoting the foodmenu type in the database.
	Label = "foodmenu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFoodmenuName holds the string denoting the foodmenu_name field in the database.
	FieldFoodmenuName = "foodmenu_name"
	// FieldFoodmenuType holds the string denoting the foodmenu_type field in the database.
	FieldFoodmenuType = "foodmenu_type"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeEatinghistory holds the string denoting the eatinghistory edge name in mutations.
	EdgeEatinghistory = "eatinghistory"

	// Table holds the table name of the foodmenu in the database.
	Table = "foodmenus"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "foodmenus"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// EatinghistoryTable is the table the holds the eatinghistory relation/edge.
	EatinghistoryTable = "eatinghistories"
	// EatinghistoryInverseTable is the table name for the Eatinghistory entity.
	// It exists in this package in order to avoid circular dependency with the "eatinghistory" package.
	EatinghistoryInverseTable = "eatinghistories"
	// EatinghistoryColumn is the table column denoting the eatinghistory relation/edge.
	EatinghistoryColumn = "foodmenu_id"
)

// Columns holds all SQL columns for foodmenu fields.
var Columns = []string{
	FieldID,
	FieldFoodmenuName,
	FieldFoodmenuType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Foodmenu type.
var ForeignKeys = []string{
	"owner_id",
}

var (
	// FoodmenuNameValidator is a validator for the "foodmenu_name" field. It is called by the builders before save.
	FoodmenuNameValidator func(string) error
	// FoodmenuTypeValidator is a validator for the "foodmenu_type" field. It is called by the builders before save.
	FoodmenuTypeValidator func(string) error
)
