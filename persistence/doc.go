package persistence

// Package persistence - persistence package for the application.
// it is responsible for handling all persistence operations
// Main Operations:
// 1. Write to Disk
// 2. Read from Disk
// Timing:
// 1. On Startup
//   1.1. Read from Disk
//   1.2. Load into memory
// 2. Upon Change
//   2.1. Write to Memory
//   2.2. Write to Disk
// 3. Every n seconds
//   3.1. Compare Memory and Disk
//   3.2. If different, write to disk
// 4. On Shutdown
//   4.1. Compare Memory and Disk
//   4.2. If different, write to disk
