package it.progetto.reviewmanager.dao;

import it.progetto.reviewmanager.model.Review;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface ReviewDao extends JpaRepository<Review,Long> {

    List<Review> findReviewByMountainPathName(String mountainPathName);

    Review findReviewById(Long id);
}
