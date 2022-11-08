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

    private List<Integer> loggedUsers;

    @Autowired
    private CredentialDao credentialDao;


    public LoginController(){
        loggedUsers = new ArrayList<>();
    }

    public LoginController(List<Integer> loggedUsers){
        this.loggedUsers = loggedUsers;
    }

    public int login(String username, String password){
        UserCredential credential = credentialDao.findByUsername(username);
        int userId;
        if(credential == null || !Objects.equals(credential.getPassword(), password)){
            System.out.println("CREDENTIAL: "+ credential);
            userId = -1;
        }else {
            userId = credential.getUser().getId();
            loggedUsers.add(userId);
        }

        return userId;
    }

    public int signin(String username, String password){
        if(credentialDao.findByUsername(username) != null){
            return -1;
        }
        UserCredential credential = new UserCredential(username,password, new LoggedUser());
        credential = credentialDao.save(credential);
        int userId = credential.getUser().getId();
        loggedUsers.add(userId);
        return userId;
    }

    public Boolean isLogged(Integer userId){
        for(Integer id: loggedUsers){
            if(id == userId)
                return true;
        }
        return false;
    }

    public void logOut(Integer userId){
        for(Integer id: loggedUsers){
            if(id == userId){
                loggedUsers.remove(id);
                return;
            }
        }
    }

    public Boolean isRegistered(String username){
        return (credentialDao.findByUsername(username) != null);
    }
}
