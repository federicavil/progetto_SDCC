package it.progetto.login.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.google.gson.Gson;
import it.progetto.login.dao.CredentialDao;
import it.progetto.login.dao.UserDao;
import it.progetto.login.model.LoggedUser;
import it.progetto.login.model.UserCredential;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

  /*
       Implementa la logica di gestione del profilo utente
  */

@Service
public class ProfileController {

    @Autowired
    private UserDao userDao;

    @Autowired
    private CredentialDao credentialDao;


    public ProfileController(){}

    /*
    * Recupera all'interno del database le informazioni sul profilo dell'utente
    * @param {username}: username dell'utente di cui recuperare le informazioni
    * @return json contentente le informazioni del profilo
     */
    public String getUserProfile(String username){
        UserCredential credential = credentialDao.findByUsername(username);
        LoggedUser user = userDao.findLoggedUserById(credential.getUser().getId());
        return "{"+user.toString()+"}";
    }
    /*
    * Aggiorna i dati all'interno del database relativi al profilo utente
    * @param {username}: username dell'utente di cui aggiornare le informazioni
    * @param {userProfile}: stringa json contentente le informazioni aggiornate
    * @return json contentente le informazioni del profilo
     */
    public String updateUserProfile(String username, String userProfile){
        Integer userId = credentialDao.findByUsername(username).getUser().getId();
        userProfile = "{\"id\": " + userId +","+userProfile.substring(1);
        System.out.println(userProfile);
        Gson gson = new Gson();
        LoggedUser user = gson.fromJson(userProfile, LoggedUser.class);
        LoggedUser updateUser = userDao.save(user);
        String json = updateUser.toString();
        System.out.println(json);
        return json;
    }
}
