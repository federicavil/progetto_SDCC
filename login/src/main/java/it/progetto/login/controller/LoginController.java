package it.progetto.login.controller;

import it.progetto.login.dao.CredentialDao;
import it.progetto.login.model.UserCredential;
import it.progetto.login.model.LoggedUser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/*
    implementa la logina di login del microservizio
*/

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


    /*
    * Effettua operazione di login dell'utente
    * @param {username}: username dell'utente che deve effettuare il login
    * @param {password}: password dell'utente che deve effettuare il login
    * @return username dell'utente se login avvenuto con successo, -1 altrimenti
     */
    public String login(String username, String password){
        UserCredential credential = credentialDao.findByUsername(username);
        if(credential == null || !Objects.equals(credential.getPassword(), password)){
            username = String.valueOf(-1);
        }else {
            username = credential.getUsername();
            loggedUsers.add(username);
        }

        return username;
    }

    /*
        * Effettua operazione di signin dell'utente
        * @param {username}: username dell'utente che deve effettuare il login
        * @param {password}: password dell'utente che deve effettuare il login
        * @return username dell'utente appena registrato
         */

    public String signin(String username, String password){
        if(credentialDao.findByUsername(username) != null){
            return String.valueOf(-1);
        }
        UserCredential credential = new UserCredential(username,password, new LoggedUser("Insert your name","Insert your surname","Insert a description"));
        credentialDao.save(credential);
        loggedUsers.add(username);
        return username;
    }

    /*
        * Controlla se l'utente indicato è attualmente loggato nell'applicazione
        * @param {userId}: cookie dell'utente
        * @return true se l'utente è loggato, false altrimenti
         */
    public Boolean isLogged(String userId){
        for(String id: loggedUsers){
            if(Objects.equals(id, userId))
                return true;
        }
        return false;
    }
      /*
        * Effettua operazione di logOut dell'utente
        * @param {userId}: username dell'utente che deve effettuare il logout
              */
    public void logOut(String userId){
        for(String id: loggedUsers){
            if(Objects.equals(id, userId)){
                loggedUsers.remove(id);
                return;
            }
        }
    }

  /*
        * Controlla se un dato username è presente all'interno del sistema
        * @param {username}: username dell'utente da verificare
        * @return true se lo username inserito appartiene ad un utente registrato all'interno del sistema, false altrimenti
         */
    public Boolean isRegistered(String username){
        return (credentialDao.findByUsername(username) != null);
    }
}
