package it.progetto.reviewmanager;

import io.grpc.stub.StreamObserver;
import it.progetto.reviewmanager.controller.ReviewController;
import it.progetto.reviewmanager.grpc.*;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

// Implementazione dei servizi grpc definiti

@RestController
@GrpcService()
public class ReviewManagerServiceImpl extends ReviewManagerServiceGrpc.ReviewManagerServiceImplBase{

    @Autowired
    ReviewController reviewController;

    // Aggiunge una nuova recensione
    @Override
    public void addReview(AddReviewRequest request, StreamObserver<AddReviewResponse> responseObserver){
        String review = reviewController.addReview(request.getReview());
        AddReviewResponse response = AddReviewResponse.newBuilder()
                .setAddedReview(review).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Restituisce tutte le recensioni di un sentiero
    @Override
    public void getReviews(GetReviewsRequest request, StreamObserver<GetReviewsResponse> responseObserver){
        String reviews = reviewController.getReviews(request.getMountainPathName());
        GetReviewsResponse response = GetReviewsResponse.newBuilder()
                .setReviewsList(reviews).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
