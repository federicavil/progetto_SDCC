
package it.progetto.login;

import io.grpc.stub.StreamObserver;
import it.progetto.login.controller.LoginController;
import it.progetto.login.controller.ProfileController;
import it.progetto.login.grpc.*;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

@RestController
@GrpcService()
public class LoginServiceImpl extends LoginServiceGrpc.LoginServiceImplBase {

    @Autowired
    LoginController loginController;

    @Autowired
    ProfileController profileController;

    @Override
    public void login(LoginRequest request, StreamObserver<LoginResponse> responseObserver){
        System.out.println("LOGIN SERVICE + " + request.getUsername() + request.getPassword());
        int userId = loginController.login(request.getUsername(),request.getPassword());
        LoginResponse response = LoginResponse.newBuilder()
                .setUser(String.valueOf(userId)).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void signin(LoginRequest request, StreamObserver<LoginResponse> responseObserver){
        System.out.println("SIGNIN SERVICE + " + request.getUsername() + request.getPassword());
        int userId = loginController.signin(request.getUsername(),request.getPassword());
        LoginResponse response = LoginResponse.newBuilder()
                .setUser(String.valueOf(userId)).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void checkLogin(CheckRequest request, StreamObserver<CheckResponse> responseObserver){
        System.out.println("CHECK SERVICE + "+ request.getUserId());
        boolean isLogged = loginController.isLogged(Integer.valueOf(request.getUserId()));
        CheckResponse response = CheckResponse.newBuilder()
                .setIsLogged(isLogged).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void logOut(CheckRequest request, StreamObserver<CheckResponse> responseObserver){
        System.out.println("LOG OUT "+ request.getUserId());
        loginController.logOut(Integer.valueOf(request.getUserId()));
        CheckResponse response = CheckResponse.newBuilder()
                .setIsLogged(true).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getProfile(ProfileRequest request, StreamObserver<ProfileResponse> responseObserver){
        System.out.println("GET PROFILE "+ request.getUserId());
        String profile = profileController.getUserProfile(Integer.valueOf(request.getUserId()));
        System.out.println("PROFILO: "+profile);
        ProfileResponse response = ProfileResponse.newBuilder()
                        .setProfile(profile).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void updateProfile(ProfileRequest request, StreamObserver<ProfileResponse> responseObserver){
        System.out.println("UPDATE PROFILE "+ request.getProfile());
        String profile = profileController.updateUserProfile(Integer.valueOf(request.getUserId()),request.getProfile());
        ProfileResponse response = ProfileResponse.newBuilder()
                .setProfile(profile).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

}
