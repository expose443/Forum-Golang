package repository

func (p *postQuery) GetLikedPostIdByUserId(userId int) ([]int64, error) {
	var postId []int64
	query := `SELECT post_id FROM likes WHERE user_id = ?`
	rows, err := p.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		postId = append(postId, id)
	}
	return postId, nil
}

func (p *postQuery) GetLikeStatus(postId, userId int) int {
	query := `SELECT status FROM likes WHERE post_id = ? AND user_id = ?`
	var likeStatus int
	p.db.QueryRow(query, postId, userId).Scan(&likeStatus)
	return likeStatus
}

func (p *postQuery) LikePost(post_id, user_id, status int) error {
	query := `INSERT INTO likes(post_id, user_id, status) VALUES(?,?,?)`
	_, err := p.db.Exec(query, post_id, user_id)
	return err
}

func (p *postQuery) UpdatePostLikeDislike(post_id, like, dislike int) error {
	query := `UPDATE posts SET like = ?, dislike = ? WHERE post_id = ?`
	_, err := p.db.Exec(query, like, dislike, post_id)
	return err
}

func (p *postQuery) DeletePostLike(post_id, user_id int) error {
	query := `DELETE FROM likes WHERE post_id = ? AND user_id = ?`
	_, err := p.db.Exec(query, post_id, user_id)
	return err
}
