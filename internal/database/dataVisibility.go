package database

func hideCredentials(database DatabaseConfig) DatabaseConfig {
	database.User = HIDDEN_PASSWORD
	database.Pass = HIDDEN_PASSWORD
	return database
}

func hideSensitiveFields(database DatabaseConfig) DatabaseConfig {
	database.Pass = HIDDEN_PASSWORD
	return database
}

func handleDataVisibility(AreCredsHidden, isDecoded bool, database DatabaseConfig) DatabaseConfig {
	if AreCredsHidden {
		return hideCredentials(database)
	}

	if isDecoded {
		if isBase64Encoded(database.Pass) {
			database.Pass = decodePassword(database.Pass)
		}
		return database
	}

	return hideSensitiveFields(database)
}
