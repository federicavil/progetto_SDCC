package it.progetto.reviewmanager.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.google.gson.Gson;
import it.progetto.reviewmanager.dao.ReviewDao;
import it.progetto.reviewmanager.model.Review;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ReviewController {

    @Autowired
    ReviewDao reviewDao;

    public String addReview(String reviewJson){
        Gson gson = new Gson();
        Review newReview = gson.fromJson(reviewJson,Review.class);
        newReview = reviewDao.save(newReview);

        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        try {
            reviewJson = ow.writeValueAsString(newReview);
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
        return reviewJson;
    }

    public String getReviews(String mountainPathName){
        List<Review> reviews = reviewDao.findReviewByMountainPathName(mountainPathName);
        String reviewsJson = "";
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        try {
            reviewsJson = ow.writeValueAsString(reviews);
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
        return reviewsJson;
    }
}
