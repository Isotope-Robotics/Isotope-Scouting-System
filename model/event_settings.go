package model

type EventSettings struct {
	Id                 int `db:"id"`
	Name               string
	TbaDownloadEnabled bool
	TbaEventCode       string
	TbaSecretId        string
	AdminPassword      string
}

func (database *Database) GetEventSettings() (*EventSettings, error) {
	allEventSettings, err := database.eventSettingsTable.getAll()
	if err != nil {
		return nil, err
	}
	if len(allEventSettings) == 1 {
		return &allEventSettings[0], nil
	}

	// Database record doesn't exist yet; create it now.
	eventSettings := EventSettings{
		Name:               "Untitled Event",
		TbaDownloadEnabled: true,
	}

	if err := database.eventSettingsTable.create(&eventSettings); err != nil {
		return nil, err
	}
	return &eventSettings, nil
}

func (database *Database) UpdateEventSettings(eventSettings *EventSettings) error {
	return database.eventSettingsTable.update(eventSettings)
}
