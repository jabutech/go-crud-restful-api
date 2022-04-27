package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	// Use recover for handle panic error.
	err := recover()
	// (1) If error
	if err != nil {
		// (1) Rollback transaction if Error
		errorRollback := tx.Rollback()
		// (2) Handle error from transaction rollback
		PanicErr(errorRollback)
		// (2) Handle error from recover
		panic(err)
	} else {
		// (1) If not error, commit transaction
		errorCommit := tx.Commit()
		// (2) Handle error from transaction commit
		PanicErr(errorCommit)
	}
}
