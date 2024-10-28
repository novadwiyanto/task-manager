package database

func Query() {
	db := DatabaseConnection()
	db.Exec("CREATE TYPE status_enum AS ENUM ('not_started', 'in_progress', 'completed')")
	db.Exec("CREATE TYPE priority_enum AS ENUM ('low', 'medium', 'high')")
	db.Exec("CREATE TYPE role_enum AS ENUM ('admin', 'member', 'guest')")
	db.Exec("CREATE TYPE notification_type_enum AS ENUM ('assigned', 'updated', 'deleted')")
	

	createTasksTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id bigserial PRIMARY KEY,
		status status_enum DEFAULT 'not_started',
		priority priority_enum DEFAULT 'low'
	);`
	db.Exec(createTasksTableSQL)

	cerateTeamMembersTableSQL := `
	CREATE TABLE IF NOT EXISTS team_members (
		id bigserial PRIMARY KEY,
		role role_enum DEFAULT 'guest'
	);`
	db.Exec(cerateTeamMembersTableSQL)

	createtaskHistoriesTableSQL := `
	CREATE TABLE IF NOT EXISTS task_histories (
		id bigserial PRIMARY KEY,
		status status_enum DEFAULT 'not_started',
		priority priority_enum DEFAULT 'low'
	);`
	db.Exec(createtaskHistoriesTableSQL)

	createNotificationsTableSQL := `
	CREATE TABLE IF NOT EXISTS notifications (
		id bigserial PRIMARY KEY,
		type notification_type_enum DEFAULT 'assigned'
	);`
	db.Exec(createNotificationsTableSQL)
}
