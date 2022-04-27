package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	// (1) If error
	if err != nil {
		// (1) Rollback transaction
		errorRollback := tx.Rollback()
		// (2) Handle error from transactional
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
