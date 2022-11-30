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

// Implementa la logica del microservizio di gestione delle recensioni

@Service
public class ReviewController {

    @Autowired
    ReviewDao reviewDao;
    /*
    * Aggiunge una nuova recensione all'interno del sistema relativamente ad un sentiero
    * @Param {reviewJson}: stringa json contenente la nuova recenisone da aggiungere
    * @returns {String}: stringa json contenente la nuova recensione aggiunta
     */
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
    /*
    * Recupera tutte le recensioni relative ad un sentiero
    * @Param {mountainPathName}: nome del sentiero di cui prendere tutte le recensioni presenti
    * @returns {String}: stringa json contenente le recensioni recuperate dal db
     */
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
