package repository

func (p *postQuery) GetDisLikeStatus(postId, userId int) int {
	query := `SELECT status FROM dislikes WHERE post_id = ? AND user_id = ?`
	var dislikeStatus int
	p.db.QueryRow(query, postId, userId).Scan(&dislikeStatus)
	return dislikeStatus
}

func (p *postQuery) DeletePostDisLike(post_id, user_id int) error {
	query := `DELETE FROM dislikes WHERE post_id = ? AND user_id = ?`
	_, err := p.db.Exec(query, post_id, user_id)
	return err
}

func (p *postQuery) DisLikePost(post_id, user_id, status int) error {
	query := `INSERT INTO dislikes(post_id, user_id, status) VALUES(?,?,?)`
	_, err := p.db.Exec(query, post_id, user_id, status)
	return err
}
