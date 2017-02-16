// Package UserManager contains The User Management.
// The Data is stored in memory for speed for Read operations. On startup, a databasefile gets loaded into memory and
// write operations get written back to this database file. This ensures fast reads and persistent writes.
package UserManager
