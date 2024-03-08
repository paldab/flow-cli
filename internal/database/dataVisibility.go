package database

func hideSensitiveFields(database DatabaseConfig) DatabaseConfig {
	database.Pass = HIDDEN_PASSWORD
	return database
}

func handleDataVisibility(isDecoded bool, database DatabaseConfig) DatabaseConfig {
	if isDecoded {
		if isBase64Encoded(database.Pass) {
			database.Pass = decodePassword(database.Pass)
		}
		return database
	}

	return hideSensitiveFields(database)
}
