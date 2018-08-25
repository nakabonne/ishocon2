package main

// Vote Model
type Vote struct {
	ID          int
	UserID      int
	CandidateID int
	Keyword     string
}

func getVoteCountByCandidateID(candidateID int) (count int) {
	row := db.QueryRow("SELECT COUNT(*) AS count FROM votes WHERE candidate_id = ?", candidateID)
	row.Scan(&count)
	return
}

func getUserVotedCount(userID int) (count int) {
	row := db.QueryRow("SELECT COUNT(*) AS count FROM votes WHERE user_id =  ?", userID)
	row.Scan(&count)
	return
}

func createVote(userID int, candidateID int, keyword string) {
	db.Exec("INSERT INTO votes (user_id, candidate_id, keyword) VALUES (?, ?, ?)",
		userID, candidateID, keyword)
}

func getVoiceOfSupporter(candidateIDs []int) (voices []string) {
	rows, err := db.Query(`
    SELECT keyword
    FROM votes
    WHERE candidate_id IN (?)
    GROUP BY keyword
    ORDER BY COUNT(*) DESC
    LIMIT 10`)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			panic(err.Error())
		}
		voices = append(voices, keyword)
	}
	return
}

func getVoteCountByParty(party string) (count int) {
	rows, err := db.Query(`
    select count(*) as count
    from votes v
    inner join candidates c on c.id = v.candidate_id
    where c.political_party = ?;
`, party)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}
