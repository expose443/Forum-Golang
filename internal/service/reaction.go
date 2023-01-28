package service

import "log"

// 400 - http status Bad request
// 500 - http status Internal server error
// 200 - http status Ok

func (p *postService) LikePost(postId, userId int) int {
	post, err := p.repository.GetPostById(int64(postId))
	if err != nil {
		log.Println(err)
		return 400
	}
	like := p.repository.GetLikeStatus(postId, userId)
	dislike := p.repository.GetDisLikeStatus(postId, userId)
	if like == 0 && dislike == 0 {
		err = p.repository.LikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like++
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return 500
		}
	} else if like == 0 && dislike == 1 {
		err = p.repository.DeletePostDisLike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		err = p.repository.LikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like++
		post.Dislike++
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return 500
		}
	} else {
		err = p.repository.DeletePostLike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like--
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			return 500
		}
	}
	return 200
}

func (p *postService) DislikePost(postId, userId int) int {
	post, err := p.repository.GetPostById(int64(postId))
	if err != nil {
		log.Println(err)
		return 400
	}
	like := p.repository.GetLikeStatus(postId, userId)
	dislike := p.repository.GetDisLikeStatus(postId, userId)
	if dislike == 0 && like == 0 {
		err = p.repository.DisLikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Dislike--
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return 500
		}
	} else if dislike == 0 && like == 1 {
		err = p.repository.DeletePostLike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		err = p.repository.DisLikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like--
		post.Dislike--
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return 500
		}
	} else {
		err = p.repository.DeletePostDisLike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Dislike++
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			return 500
		}
	}
	return 200
}
