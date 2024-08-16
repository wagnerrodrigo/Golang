package db

type UptimeResult struct {
	URL       string `json:"url"`
	Status    bool   `json:"status"`
	CheckedAt string `json:"checked_at"`
}

func GetAllResults() ([]UptimeResult, error) {
	rows, err := db.Query("SELECT url,status,checked_at FROM uptime_results")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []UptimeResult
	for rows.Next() {
		var result UptimeResult
		err := rows.Scan(&result.URL, &result.Status, &result.CheckedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
