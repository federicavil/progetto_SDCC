package it.progetto.login.controller;

import it.progetto.login.dao.CredentialDao;
import it.progetto.login.model.UserCredential;
import it.progetto.login.model.LoggedUser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

@Service
public class LoginController {

    private List<String> loggedUsers;

    @Autowired
    private CredentialDao credentialDao;


    public LoginController(){
        loggedUsers = new ArrayList<>();
    }

    public LoginController(List<String> loggedUsers){
        this.loggedUsers = loggedUsers;
    }

    public String login(String username, String password){
        UserCredential credential = credentialDao.findByUsername(username);
        if(credential == null || !Objects.equals(credential.getPassword(), password)){
            System.out.println("CREDENTIAL: "+ credential);
            username = String.valueOf(-1);
        }else {
            username = credential.getUsername();
            loggedUsers.add(username);
        }

        return username;
    }

    public String signin(String username, String password){
        if(credentialDao.findByUsername(username) != null){
            return String.valueOf(-1);
        }
        UserCredential credential = new UserCredential(username,password, new LoggedUser());
        credentialDao.save(credential);
        loggedUsers.add(username);
        return username;
    }

    public Boolean isLogged(String userId){
        System.out.println(userId);
        System.out.println(loggedUsers);
        for(String id: loggedUsers){
            if(Objects.equals(id, userId))
                return true;
        }
        return false;
    }

    public void logOut(String userId){
        for(String id: loggedUsers){
            if(Objects.equals(id, userId)){
                loggedUsers.remove(id);
                return;
            }
        }
    }

    public Boolean isRegistered(String username){
        return (credentialDao.findByUsername(username) != null);
    }
}
