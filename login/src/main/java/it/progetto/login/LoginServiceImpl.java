
package it.progetto.login;

import io.grpc.stub.StreamObserver;
import it.progetto.login.controller.LoginController;
import it.progetto.login.controller.ProfileController;
import it.progetto.login.grpc.*;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

/*
    Implementa i servizi grpc registrati
*/

@RestController
@GrpcService()
public class LoginServiceImpl extends LoginServiceGrpc.LoginServiceImplBase {

    @Autowired
    LoginController loginController;

    @Autowired
    ProfileController profileController;

    // Esegue il Login dell'utente
    @Override
    public void login(LoginRequest request, StreamObserver<LoginResponse> responseObserver){
        System.out.println("LOGIN SERVICE + " + request.getUsername() + request.getPassword());
        String user = loginController.login(request.getUsername(),request.getPassword());
        LoginResponse response = LoginResponse.newBuilder()
                .setUser(user).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Esegue il SignIn dell'utente
    @Override
    public void signin(LoginRequest request, StreamObserver<LoginResponse> responseObserver){
        System.out.println("SIGNIN SERVICE + " + request.getUsername() + request.getPassword());
        String user = loginController.signin(request.getUsername(),request.getPassword());
        LoginResponse response = LoginResponse.newBuilder()
                .setUser(user).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Controlla se un utente è attualmente loggato
    @Override
    public void checkLogin(CheckRequest request, StreamObserver<CheckResponse> responseObserver){
        System.out.println("CHECK SERVICE + "+ request.getUserId());
        boolean isLogged = loginController.isLogged(request.getUserId());
        CheckResponse response = CheckResponse.newBuilder()
                .setIsLogged(isLogged).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Esegue il Logout dell'utente
    @Override
    public void logOut(CheckRequest request, StreamObserver<CheckResponse> responseObserver){
        System.out.println("LOG OUT "+ request.getUserId());
        loginController.logOut(request.getUserId());
        CheckResponse response = CheckResponse.newBuilder()
                .setIsLogged(true).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Recupera le informazioni del profilo di un utente
    @Override
    public void getProfile(ProfileRequest request, StreamObserver<ProfileResponse> responseObserver){
        System.out.println("GET PROFILE "+ request.getUserId());
        String profile = profileController.getUserProfile(request.getUserId());
        System.out.println("PROFILO: "+profile);
        ProfileResponse response = ProfileResponse.newBuilder()
                        .setProfile(profile).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Aggiorna le informazioni di profilo di un utente
    @Override
    public void updateProfile(ProfileRequest request, StreamObserver<ProfileResponse> responseObserver){
        System.out.println("UPDATE PROFILE "+ request.getProfile());
        String profile = profileController.updateUserProfile(request.getUserId(), request.getProfile());
        ProfileResponse response = ProfileResponse.newBuilder()
                .setProfile(profile).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Controlla se lo username inserito è presente all'interno del sistema
    @Override
    public void checkUsername(CheckUsernameRequest request, StreamObserver<CheckUsernameResponse> responseObserver){
        System.out.println("CHECK USERNAME");
        Boolean isRegistered = loginController.isRegistered(request.getUsername());
        CheckUsernameResponse response = CheckUsernameResponse.newBuilder()
                .setIsRegistered(isRegistered).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();

    }

}
