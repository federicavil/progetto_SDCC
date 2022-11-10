package it.progetto.reviewmanager;

import io.grpc.stub.StreamObserver;
import it.progetto.reviewmanager.controller.ReviewController;
import it.progetto.reviewmanager.grpc.*;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

@RestController
@GrpcService()
public class ReviewManagerServiceImpl extends ReviewManagerServiceGrpc.ReviewManagerServiceImplBase{

    @Autowired
    ReviewController reviewController;

    @Override
    public void addReview(AddReviewRequest request, StreamObserver<AddReviewResponse> responseObserver){
        System.out.println("AGGIUNTA REVIEW "+request.getReview());
        String review = reviewController.addReview(request.getReview());
        AddReviewResponse response = AddReviewResponse.newBuilder()
                .setAddedReview(review).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getReviews(GetReviewsRequest request, StreamObserver<GetReviewsResponse> responseObserver){
        System.out.println("GET REVIEWS "+request.getMountainPathName());
        String reviews = reviewController.getReviews(request.getMountainPathName());
        System.out.println(reviews);
        GetReviewsResponse response = GetReviewsResponse.newBuilder()
                .setReviewsList(reviews).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}