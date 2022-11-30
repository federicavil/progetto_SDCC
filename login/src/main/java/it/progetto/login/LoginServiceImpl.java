
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
        String user = loginController.login(request.getUsername(),request.getPassword());
        LoginResponse response = LoginResponse.newBuilder()
                .setUser(user).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Esegue il SignIn dell'utente
    @Override
    public void signin(LoginRequest request, StreamObserver<LoginResponse> responseObserver){
        String user = loginController.signin(request.getUsername(),request.getPassword());
        LoginResponse response = LoginResponse.newBuilder()
                .setUser(user).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Controlla se un utente è attualmente loggato
    @Override
    public void checkLogin(CheckRequest request, StreamObserver<CheckResponse> responseObserver){
        boolean isLogged = loginController.isLogged(request.getUserId());
        CheckResponse response = CheckResponse.newBuilder()
                .setIsLogged(isLogged).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Esegue il Logout dell'utente
    @Override
    public void logOut(CheckRequest request, StreamObserver<CheckResponse> responseObserver){
        loginController.logOut(request.getUserId());
        CheckResponse response = CheckResponse.newBuilder()
                .setIsLogged(true).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Recupera le informazioni del profilo di un utente
    @Override
    public void getProfile(ProfileRequest request, StreamObserver<ProfileResponse> responseObserver){
        String profile = profileController.getUserProfile(request.getUserId());
        ProfileResponse response = ProfileResponse.newBuilder()
                        .setProfile(profile).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Aggiorna le informazioni di profilo di un utente
    @Override
    public void updateProfile(ProfileRequest request, StreamObserver<ProfileResponse> responseObserver){
        String profile = profileController.updateUserProfile(request.getUserId(), request.getProfile());
        ProfileResponse response = ProfileResponse.newBuilder()
                .setProfile(profile).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    // Controlla se lo username inserito è presente all'interno del sistema
    @Override
    public void checkUsername(CheckUsernameRequest request, StreamObserver<CheckUsernameResponse> responseObserver){
        Boolean isRegistered = loginController.isRegistered(request.getUsername());
        CheckUsernameResponse response = CheckUsernameResponse.newBuilder()
                .setIsRegistered(isRegistered).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();

    }

}
