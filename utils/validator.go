package utils

// I could use any go validator package but tried to create those validations by myself
type Validator interface{ IsValid() bool }
