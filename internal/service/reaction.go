package service

import "log"

func (p *postService) LikePost(postId, userId int) error {
	post, err := p.repository.GetPostById(int64(postId))
	if err != nil {
		log.Println(err)
		return err
	}
	like := p.repository.GetLikeStatus(postId, userId)
	dislike := p.repository.GetDisLikeStatus(postId, userId)
	if like == 0 && dislike == 0 {
		err = p.repository.LikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		post.Like++
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return err
		}
	} else if like == 0 && dislike == 1 {
		err = p.repository.DeletePostDisLike(postId, userId)
		if err != nil {
			log.Println(err)
			return err
		}
		err = p.repository.LikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		post.Like++
		post.Dislike--
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		err = p.repository.DeletePostLike(postId, userId)
		if err != nil {
			log.Println(err)
			return err
		}
		post.Like--
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
	}
	return nil
}

func (p *postService) DislikePost(postId, userId int) error {
	post, err := p.repository.GetPostById(int64(postId))
	if err != nil {
		log.Println(err)
		return err
	}
	like := p.repository.GetLikeStatus(postId, userId)
	dislike := p.repository.GetDisLikeStatus(postId, userId)
	if dislike == 0 && like == 0 {
		err = p.repository.DisLikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		post.Dislike++
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return err
		}
	} else if dislike == 0 && like == 0 {
		err = p.repository.DeletePostLike(postId, userId)
		if err != nil {
			log.Println(err)
			return err
		}
		err = p.repository.DisLikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		post.Like--
		post.Dislike++
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		err = p.repository.DeletePostDisLike(postId, userId)
		if err != nil {
			log.Println(err)
			return err
		}
		post.Dislike--
		err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
	}
	return nil
}
